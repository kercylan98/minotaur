package libs

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	lua "github.com/yuin/gopher-lua"
	"io/fs"
	"math"
	"sort"
	"strings"
)

const libraryScriptDirName = "lua-scripts"

// 脚本加载优先级
var libraryScriptPriority = []string{
	"errors.lua",
	"json.lua",
	"actor.lua",
}

var (
	//go:embed lua-scripts/*.lua
	libraryScriptEmbedFS embed.FS
)

func NewLibrary(state *lua.LState, ctx vivid.ActorContext) *Library {
	libs := &Library{
		state: state,
	}

	libs.initLibrary(ctx)
	return libs
}

type Library struct {
	state *lua.LState
}

func loadScripts() []fs.DirEntry {
	// 加载脚本
	entries, err := libraryScriptEmbedFS.ReadDir(libraryScriptDirName)
	if err != nil {
		panic(err)
	}

	// 优先级排序
	priority := make(map[string]int)
	for i, name := range libraryScriptPriority {
		priority[name] = i
	}
	sort.Slice(entries, func(i, j int) bool {
		ip, ie := priority[entries[i].Name()]
		jp, je := priority[entries[j].Name()]
		if !ie {
			ip = math.MaxInt32
		}
		if !je {
			jp = math.MaxInt32
		}

		return ip < jp
	})

	return entries
}

func (l *Library) initLibrary(ctx vivid.ActorContext) {
	// 加载脚本
	entries := loadScripts()

	// 注入脚本
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.SplitN(entry.Name(), ".", 2)[0]
		if strings.HasSuffix(name, "_doc") {
			continue
		}

		var luaBytes []byte
		var err error
		l.state.PreloadModule(name, func(state *lua.LState) (num int) {
			if luaBytes, err = libraryScriptEmbedFS.ReadFile(libraryScriptDirName + "/" + entry.Name()); err != nil {
				panic(fmt.Errorf("read lib %s error: %v", entry.Name(), err))
			}

			if err = state.DoString(string(luaBytes)); err != nil {
				panic(fmt.Errorf("load lib %s error: %v", entry.Name(), err))
			}

			mod := state.Get(-1)
			defer state.Pop(1)

			// 模块函数注入
			functions := moduleInjector[name]
			if len(functions) > 0 {
				var gf = make(map[string]lua.LGFunction)
				for k, v := range functions {
					gf[k] = v(ctx)
				}
				mod = state.SetFuncs(mod.(*lua.LTable), gf)
			}

			num = 1
			state.Push(mod)
			return
		})
	}
}

func (l *Library) Close() {
	l.state.Close()
}
