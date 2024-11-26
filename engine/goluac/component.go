package goluac

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/vivid"
	lua "github.com/yuin/gopher-lua"
)

type (
	luaStateKeyStruct struct{}
	Script            struct {
		Filepath  string // lua 脚本文件路径
		LuaScript string // lua 脚本代码
	}
)

var (
	_           LuaComponent = (*luaComponent)(nil)
	luaStateKey              = new(luaStateKeyStruct)
)

// New 创建一个可用于绑定到 vivid.ActorSystem 的 LuaComponent 组件
func New() LuaComponent {
	return new(luaComponent)
}

// BindLuaScript 绑定一个 Lua 脚本
func BindLuaScript(ctx vivid.ActorContext, script string) {
	ctx.Tell(ctx.Ref(), &Script{LuaScript: script})
}

// BindLuaFile 绑定一个 Lua 脚本文件
func BindLuaFile(ctx vivid.ActorContext, filepath string) {
	ctx.Tell(ctx.Ref(), &Script{Filepath: filepath})
}

// LuaComponent 组件可以为每一个 vivid.Actor 绑定一个 Lua 脚本状态，仅需要对其投递一个 Script 消息即可
type LuaComponent interface {
	vivid.Component
	vivid.ActorReceiveMessageCaptureComponent
}

type luaComponent struct {
}

func (l *luaComponent) OnInitialize(actorSystem *vivid.ActorSystem) error {
	return nil
}

func (l *luaComponent) OnActorReceiveMessageCapture(ctx vivid.ActorContext) (abort bool) {
	switch m := ctx.Message().(type) {
	case *Script:
		return l.onBindLuaState(ctx, m)
	}

	return false
}

func (l *luaComponent) onBindLuaState(ctx vivid.ActorContext, m *Script) bool {
	// 初始化状态
	if ctx.HasValue(luaStateKey) {
		panic(fmt.Errorf("actor %s already has a lua state", ctx.Ref().URL()))
	}
	var err error
	var state = lua.NewState()
	ctx.SetValue(luaStateKey, state)

	// 导入模块
	// ...

	// 加载脚本
	if m.Filepath != "" {
		err = state.DoFile(m.Filepath)
	} else {
		err = state.DoString(m.LuaScript)
	}
	if err != nil {
		panic(err)
	}
	return true
}
