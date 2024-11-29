package libs

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	lua "github.com/yuin/gopher-lua"
)

var moduleInjector = map[string]map[string]moduleFunctionInjector{}

type moduleFunctionInjector func(ctx vivid.ActorContext) lua.LGFunction

func injectModuleFunction(moduleName string, functionName string, function moduleFunctionInjector) {
	if _, ok := moduleInjector[moduleName]; !ok {
		moduleInjector[moduleName] = make(map[string]moduleFunctionInjector)
	}
	moduleInjector[moduleName][functionName] = function
}
