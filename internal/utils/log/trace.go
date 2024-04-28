package log

import (
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"runtime"
	"strings"
)

type beautyTrace struct {
	trace []string
}

func formatTraceError(err error, beauty bool) slog.Value {
	var groupValues []slog.Attr

	if msg := err.Error(); msg != "" {
		groupValues = append(groupValues, slog.Any("msg", err))
	}

	type StackTracer interface {
		StackTrace() errors.StackTrace
	}

	var st StackTracer
	for err := err; err != nil; err = errors.Unwrap(err) {
		if x, ok := err.(StackTracer); ok {
			st = x
		}
	}

	if st == nil && err != nil {
		st = errors.WithStack(err).(StackTracer)
	}

	if st != nil {
		if beauty {
			groupValues = append(groupValues,
				slog.Any("trace", &beautyTrace{traceLines(st.StackTrace())}),
			)
		} else {
			groupValues = append(groupValues,
				slog.Any("trace", traceLines(st.StackTrace())),
			)
		}
	}

	return slog.GroupValue(groupValues...)
}

func traceLines(frames errors.StackTrace) []string {
	traceLines := make([]string, len(frames))

	var skipped int
	skipping := true
	for i := len(frames) - 1; i >= 0; i-- {
		pc := uintptr(frames[i]) - 1
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			traceLines[i] = "unknown"
			skipping = false
			continue
		}

		name := fn.Name()

		if skipping && strings.HasPrefix(name, "runtime.") {
			skipped++
			continue
		} else {
			skipping = false
		}

		filename, lineNr := fn.FileLine(pc)
		traceLines[i] = fmt.Sprintf("%s %s:%d", name, filename, lineNr)
	}

	return traceLines[:len(traceLines)-skipped]
}
