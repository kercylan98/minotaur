package geometry

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"net/http"
	"strings"
	"sync"
	"time"
)

// NewPreviewBoard2D 创建一个 2D 预览画板
func NewPreviewBoard2D(fps, width, height int) *PreviewBoard2D {
	ctx, cancel := context.WithCancel(context.Background())
	return &PreviewBoard2D{
		ctx:    ctx,
		cancel: cancel,
		fps:    fps,
		width:  width,
		height: height,
		cond:   sync.NewCond(&sync.Mutex{}),
	}
}

// PreviewBoard2D 用于调试等行为的预览画板，该画板生成一个简单的 HTML 页面，该页面包含一个 canvas 元素，可以在该画板上绘制图形。
// 另外也会根据 EventStream 的事件流来试试更新画板上的图形
type PreviewBoard2D struct {
	ctx    context.Context
	cancel context.CancelFunc
	fps    int
	width  int
	height int

	currFrameLock sync.RWMutex
	currFrame     []string
	frames        [][]string
	cond          *sync.Cond
}

func (b *PreviewBoard2D) Update(vec2 ...Vector) {
	var builder strings.Builder
	builder.WriteString("[")
	for i, v := range vec2 {
		builder.WriteString("{\"x\": " + convert.Float64ToString(v[0]) + ", \"y\": " + convert.Float64ToString(v[1]) + "}")
		if i != len(vec2)-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")
	b.currFrameLock.Lock()
	defer b.currFrameLock.Unlock()
	b.currFrame = append(b.currFrame, builder.String())
}

func (b *PreviewBoard2D) Stop() {
	b.cancel()
	b.cond.Broadcast()
}

func (b *PreviewBoard2D) Start(addr string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var builder strings.Builder
		builder.WriteString("<!DOCTYPE html><html><head><title>Preview Board</title></head><body>")
		builder.WriteString("<canvas id=\"canvas\" width=\"" + convert.IntToString(b.width) + "\" height=\"" + convert.IntToString(b.height) + "\"></canvas>")
		builder.WriteString("<script>")
		builder.WriteString("var canvas = document.getElementById('canvas');")
		builder.WriteString("canvas.style.backgroundColor = 'black';")
		builder.WriteString("var ctx = canvas.getContext('2d');")
		builder.WriteString("var eventSource = new EventSource('/event');")
		builder.WriteString("eventSource.onmessage = function(event) {")
		builder.WriteString("var data = JSON.parse(event.data);")
		builder.WriteString("ctx.clearRect(0, 0, canvas.width, canvas.height);")
		builder.WriteString("for (var i = 0; i < data.length; i++) {")
		builder.WriteString("ctx.beginPath();")
		builder.WriteString("ctx.arc(data[i].x, data[i].y, 5, 0, 2 * Math.PI);")
		builder.WriteString("ctx.fillStyle = 'blue';")
		builder.WriteString("ctx.fill();")
		builder.WriteString("}")
		builder.WriteString("};")

		builder.WriteString("</script>")
		builder.WriteString("</body></html>")
		_, _ = writer.Write([]byte(builder.String()))
	})

	http.HandleFunc("/event", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/event-stream")
		writer.Header().Set("Cache-Control", "no-cache")
		writer.Header().Set("Connection", "keep-alive")

		var frame = 0
		for {
			b.cond.L.Lock()
			b.cond.Wait()
			b.cond.L.Unlock()
			select {
			case <-b.ctx.Done():
				return
			default:
				b.currFrameLock.RLock()

				if frame < len(b.frames) {
					for _, frameData := range b.frames[frame] {
						_, _ = writer.Write([]byte("data: " + frameData + "\n\n"))
					}
					frame++
				}

				b.currFrameLock.RUnlock()
			}
		}
	})

	go func() {
		var ticker = time.NewTicker(time.Second / time.Duration(b.fps))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				b.currFrameLock.Lock()
				b.frames = append(b.frames, b.currFrame)
				b.currFrame = nil
				b.currFrameLock.Unlock()
				b.cond.Broadcast()
			case <-b.ctx.Done():
				return
			}
		}
	}()
	return http.ListenAndServe(addr, nil)
}
