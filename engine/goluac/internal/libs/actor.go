package libs

import (
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/vivid"
	lua "github.com/yuin/gopher-lua"
	"net/url"
	"time"
)

func init() {
	injectModuleFunction("actor", "tell", tell)
	injectModuleFunction("actor", "ask", ask)
	injectModuleFunction("actor", "future_ask", futureAsk)
}

func tell(ctx vivid.ActorContext) lua.LGFunction {
	return func(state *lua.LState) int {
		refUrl := state.ToString(1)
		message, err := readToJson(state, 2)
		if err != nil {
			return pushError(state, err)
		}
		u, err := url.Parse(refUrl)
		if err != nil {
			return pushError(state, err)
		}
		ctx.Tell(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message})
		return 0
	}
}

func ask(ctx vivid.ActorContext) lua.LGFunction {
	return func(state *lua.LState) int {
		refUrl := state.ToString(1)
		message, err := readToJson(state, 2)
		if err != nil {
			return pushError(state, err)
		}
		u, err := url.Parse(refUrl)
		if err != nil {
			return pushError(state, err)
		}
		ctx.Ask(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message})
		return 0
	}
}

func futureAsk(ctx vivid.ActorContext) lua.LGFunction {
	return func(state *lua.LState) int {
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
			f = ctx.FutureAsk(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message}, time.Duration(timeout)*time.Millisecond)
		} else {
			f = ctx.FutureAsk(vivid.NewActorRef(u.Host, u.Path), &LuaMessage{Data: message})
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
}
