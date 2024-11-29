package goluac

import (
	"embed"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"io/fs"
	"math"
	"sort"
	"strings"
)

// internalLuaScriptDirName 内部静态 Lua 脚本所在目录名称
const internalLuaScriptDirName = "lua-scripts"

var (
	// 内部静态脚本路径，用于加载到二进制静态资源
	//go:embed lua-scripts/*.lua
	libraryScriptEmbedFS embed.FS
	// 已经加载到内存中的静态脚本（它具备优先级，其中 0：模块名称 1：脚本代码）
	loadedScripts [][2]string
	// 固定静态脚本加载优先级（当未命中名称时将最末加载）
	internalLuaScriptLoadPriority = []string{
		"errors.lua",
		"json.lua",
		"actor.lua",
	}
)

type (
	moduleName           = string
	functionName         = string
	moduleGoFuncInjector = func(ctx *actorContext) lua.LGFunction
)

var (
	// 模块函数注入器列表，它将在初始化时向模块添加额外的 Go 函数
	moduleGoFuncInjectorList = map[moduleName]map[functionName]moduleGoFuncInjector{}
)

func init() {
	for _, entry := range loadScripts() {
		if entry.IsDir() {
			continue
		}
		moduleName := strings.SplitN(entry.Name(), ".", 2)[0]
		fileBytes, err := libraryScriptEmbedFS.ReadFile(internalLuaScriptDirName + "/" + entry.Name())
		if err != nil {
			panic(fmt.Errorf("read lib %s error: %v", entry.Name(), err))
		}
		loadedScripts = append(loadedScripts, [2]string{moduleName, string(fileBytes)})
	}
}

func applyModuleGoFuncInject(ctx *actorContext, state *lua.LState, name, code string) int {
	if err := state.DoString(code); err != nil {
		panic(fmt.Errorf("load lib %s error: %v", name, err))
	}

	mod := state.Get(-1)
	defer state.Pop(1)

	// 模块函数注入
	functions := moduleGoFuncInjectorList[name]
	if len(functions) > 0 {
		var gf = make(map[string]lua.LGFunction)
		for k, v := range functions {
			gf[k] = v(ctx)
		}
		mod = state.SetFuncs(mod.(*lua.LTable), gf)
	}

	state.Push(mod)
	state.SetGlobal(name, mod)
	return 1
}

func moduleGoFuncInject(moduleName moduleName, funcName functionName, function moduleGoFuncInjector) {
	functions, ok := moduleGoFuncInjectorList[moduleName]
	if !ok {
		functions = make(map[functionName]moduleGoFuncInjector)
		moduleGoFuncInjectorList[moduleName] = functions
	}
	functions[funcName] = function
}

func loadScripts() []fs.DirEntry {
	// 加载脚本
	entries, err := libraryScriptEmbedFS.ReadDir(internalLuaScriptDirName)
	if err != nil {
		panic(err)
	}

	// 优先级排序
	priority := make(map[string]int)
	for i, name := range internalLuaScriptLoadPriority {
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
