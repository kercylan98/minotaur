package goluac

import (
	"embed"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
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
	// 已经加载到内存中的静态脚本（它具备优先级，其中 0：模块名称 1：编译信息）
	loadedScripts [][2]any
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
	temp := lua.NewState()
	defer temp.Close()

	for _, entry := range loadScripts() {
		if entry.IsDir() {
			continue
		}
		name := strings.SplitN(entry.Name(), ".", 2)[0]
		fileBytes, err := libraryScriptEmbedFS.ReadFile(internalLuaScriptDirName + "/" + entry.Name())
		if err != nil {
			panic(fmt.Errorf("read lib %s error: %v", entry.Name(), err))
		}

		// 编译
		code := string(fileBytes)
		reader := strings.NewReader(code)
		chunk, err := parse.Parse(reader, code)
		if err != nil {
			panic(err)
		}
		proto, err := lua.Compile(chunk, code)
		if err != nil {
			panic(err)
		}

		loadedScripts = append(loadedScripts, [2]any{name, temp.NewFunctionFromProto(proto)})
	}
}

func applyModuleGoFuncInject(ctx *actorContext, state *lua.LState, name string, code *lua.LFunction, currError error, errorHandler func(err error)) int {
	if currError != nil {
		return 0
	}

	state.Push(code)
	if err := state.PCall(0, lua.MultRet, nil); err != nil {
		errorHandler(newDoLuaScriptError(name, err))
		return 0
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
