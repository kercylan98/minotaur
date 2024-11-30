package goluac

import (
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/toolkit/log"
	lua "github.com/yuin/gopher-lua"
	"net/url"
	"sync"
	"time"
)

var (
	// 用于确保只检查一次组件是否注册的标记
	onceComponentCheck sync.Once
)

// AttachActorContext 通过一段 Lua 入口代码创建一个 Lua 虚拟机并绑定到传入的 ActorContext 上下文
func AttachActorContext(ctx vivid.ActorContext, luaCode string) (err error) {
	onceComponentCheck.Do(func() {
		if !vivid.HasComponent[*component](ctx.System()) {
			err = ErrorNotLoadGoluacComponent
		}
	})
	if err != nil {
		return
	}

	ac := &actorContext{
		ActorContext: ctx,
		lua:          lua.NewState(),
	}
	ac.initLuaActorContext()
	ctx.SetValue(actorContextKey, ac)

	// 标准库注入
	defer func() {
		if err != nil {
			ac.lua.Close()
		}
	}()

	for _, script := range loadedScripts {
		name, code := script[0], script[1]
		ac.lua.PreloadModule(name, func(state *lua.LState) int {
			return applyModuleGoFuncInject(ac, state, name, code, err, func(injectErr error) {
				err = injectErr
			})
		})
	}
	if err != nil {
		return
	}

	if err = ac.lua.DoString(luaCode); err != nil {
		return err
	}

	// 模块加载
	ac.luaActorModule = ac.lua.GetGlobal("actor").(*lua.LTable)
	ac.luaOnReceiveHandler = ac.luaActorModule.RawGetString("on_receive").(*lua.LFunction)

	return
}

type actorContext struct {
	vivid.ActorContext                 // Actor 上下文
	lua                 *lua.LState    // Lua 状态
	luaCtx              *lua.LTable    // Actor 上下文表
	luaActorModule      *lua.LTable    // Actor 模块
	luaOnReceiveHandler *lua.LFunction // actor.on_receive 函数缓存
	messageCache        lua.LValue     // 解码过的消息缓存
}

func (c *actorContext) initLuaActorContext() {
	c.luaCtx = c.lua.NewTable()
	c.luaCtx.RawSet(lua.LString("ref"), c.lua.NewFunction(c.ref))
	c.luaCtx.RawSet(lua.LString("sender"), c.lua.NewFunction(c.sender))
	c.luaCtx.RawSet(lua.LString("parent"), c.lua.NewFunction(c.parent))
	c.luaCtx.RawSet(lua.LString("logical_address"), c.lua.NewFunction(c.logicalAddress))
	c.luaCtx.RawSet(lua.LString("physical_address"), c.lua.NewFunction(c.physicalAddress))
	c.luaCtx.RawSet(lua.LString("message"), c.lua.NewFunction(c.message))
	c.luaCtx.RawSet(lua.LString("tell"), c.lua.NewFunction(c.tell))
	c.luaCtx.RawSet(lua.LString("ask"), c.lua.NewFunction(c.ask))
	c.luaCtx.RawSet(lua.LString("future_ask"), c.lua.NewFunction(c.futureAsk))
	c.luaCtx.RawSet(lua.LString("reply"), c.lua.NewFunction(c.reply))
}

func (c *actorContext) OnReceive(m *LuaMessage, resetCache bool) {
	if resetCache {
		c.messageCache = nil
	}
	if err := c.lua.CallByParam(lua.P{
		Fn:      c.luaOnReceiveHandler,
		NRet:    1,
		Protect: true,
	}, c.luaCtx, lua.LString(m.Data)); err != nil {
		c.System().Logger().Error("LuaActor", log.Any("on_receive", err))
	}
}

func (c *actorContext) tell(state *lua.LState) int {
	refUrl := state.ToString(1)
	message, err := readToJson(state, 2)
	if err != nil {
		return pushError(state, err)
	}
	u, err := url.Parse(refUrl)
	if err != nil {
		return pushError(state, err)
	}
	c.Tell(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message})
	return 0
}

func (c *actorContext) ask(state *lua.LState) int {
	refUrl := state.ToString(1)
	message, err := readToJson(state, 2)
	if err != nil {
		return pushError(state, err)
	}
	u, err := url.Parse(refUrl)
	if err != nil {
		return pushError(state, err)
	}
	c.Ask(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message})
	return 0
}

func (c *actorContext) futureAsk(state *lua.LState) int {
	refUrl := state.ToString(1)
	message, err := readToJson(state, 2)
	if err != nil {
		return pushError(state, err)
	}
	timeout, hasTimeout := state.Get(3).(lua.LNumber)
	u, err := url.Parse(refUrl)
	if err != nil {
		return pushError(state, err)
	}

	// 获取 msgpack module
	var f future.Future[vivid.Message]
	if hasTimeout {
		f = c.FutureAsk(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message}, time.Duration(timeout)*time.Millisecond)
	} else {
		f = c.FutureAsk(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message})
	}

	futureTable := state.NewTable()
	futureTable.RawSet(lua.LString("result"), state.NewFunction(func(state *lua.LState) int {
		result, err := f.Result()
		if err != nil {
			return pushError(state, err)
		}
		state.Push(lua.LString(result.(*LuaMessage).Data))
		return 1
	}))

	futureTable.RawSet(lua.LString("wait"), state.NewFunction(func(state *lua.LState) int {
		if err = f.Wait(); err != nil {
			return pushError(state, err)
		}
		state.Push(lua.LNil)
		return 1
	}))

	state.Push(futureTable)
	return 1
}

func (c *actorContext) createMessageCache(message *LuaMessage) error {
	value, err := decodeFromJson(c.lua, message.Data)
	if err != nil {
		return err
	}
	c.messageCache = value
	return nil
}

func (c *actorContext) message(state *lua.LState) int {
	cache := c.messageCache
	if cache != nil {
		state.Push(cache)
		return 1
	}

	if luaMessage, ok := c.Message().(*LuaMessage); ok {
		if err := c.createMessageCache(luaMessage); err != nil {
			return pushError(state, err)
		}
		state.Push(c.messageCache)
		return 1
	}
	c.messageCache = lua.LNil
	state.Push(lua.LNil)
	return 1
}

func (c *actorContext) ref(state *lua.LState) int {
	state.Push(lua.LString(c.Ref().URL().String()))
	return 1
}

func (c *actorContext) sender(state *lua.LState) int {
	if c.Sender() == nil {
		state.Push(lua.LNil)
	} else {
		state.Push(lua.LString(c.Sender().URL().String()))
	}
	return 1
}

func (c *actorContext) parent(state *lua.LState) int {
	if c.Parent() == nil {
		state.Push(lua.LNil)
	} else {
		state.Push(lua.LString(c.Parent().URL().String()))
	}
	return 1
}

func (c *actorContext) logicalAddress(state *lua.LState) int {
	state.Push(lua.LString(c.LogicalAddress()))
	return 1
}

func (c *actorContext) physicalAddress(state *lua.LState) int {
	state.Push(lua.LString(c.PhysicalAddress()))
	return 1
}

func (c *actorContext) reply(state *lua.LState) int {
	message, err := readToJson(state, 1)
	if err != nil {
		return pushError(state, err)
	}
	c.Reply(&LuaMessage{Data: message})
	return 0
}
