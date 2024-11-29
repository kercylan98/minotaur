package vivid

import (
	"github.com/kercylan98/minotaur/toolkit/log"
	"reflect"
)

func newComponents(actorSystem *ActorSystem) *components {
	var cs = &components{actorSystem: actorSystem, componentRecord: make(map[reflect.Type]struct{})}
	for _, comp := range actorSystem.config.components {
		cs.componentRecord[reflect.TypeOf(comp)] = struct{}{}

		tryBindComponent[ShutdownComponent](cs, &cs.shutdown, comp)
		tryBindComponent[ActorDefineCaptureComponent](cs, &cs.actorDefineCapture, comp)
		tryBindComponent[ActorContextCaptureComponent](cs, &cs.actorContextCapture, comp)
		tryBindComponent[ActorReceiveMessageCaptureComponent](cs, &cs.actorReceiveMessageCapture, comp)
	}

	cs.onInitialize()
	return cs
}

type components struct {
	actorSystem                *ActorSystem
	componentRecord            map[reflect.Type]struct{}
	shutdown                   []ShutdownComponent
	actorDefineCapture         []ActorDefineCaptureComponent
	actorContextCapture        []ActorContextCaptureComponent
	actorReceiveMessageCapture []ActorReceiveMessageCaptureComponent
}

// HasComponent 判断指定的 ActorSystem 是否绑定了指定的组件
func HasComponent[C Component](actorSystem *ActorSystem) bool {
	_, ok := actorSystem.components.componentRecord[reflect.TypeOf((*C)(nil)).Elem()]
	return ok
}

func tryBindComponent[C Component](components *components, slice *[]C, comp Component) {
	if c, ok := comp.(C); ok {
		*slice = append(*slice, c)
		components.actorSystem.Logger().Info("component", log.String("register", reflect.TypeOf(new(C)).Elem().Name()), log.String("handler", reflect.TypeOf(comp).Elem().Name()))
	}
}

func (cs *components) onActorReceiveMessageCapture(context ActorContext) bool {
	if cs == nil {
		return false
	}
	for _, c := range cs.actorReceiveMessageCapture {
		if c.OnActorReceiveMessageCapture(context) {
			return true
		}
	}
	return false
}

func (cs *components) onActorContextCapture(actorContext ActorContext) {
	if cs == nil {
		return
	}
	for _, c := range cs.actorContextCapture {
		actorContext.ExecLocalFunc(actorContext.Ref(), func(ctx ActorContext) {
			c.OnActorContextCapture(ctx)
		})
	}
}

func (cs *components) onActorDefineCapture(provider ActorProvider, descriptor *ActorDescriptor) {
	if cs == nil {
		return
	}
	for _, c := range cs.actorDefineCapture {
		c.OnActorDefineCapture(cs.actorSystem, provider, descriptor)
	}
}

func (cs *components) onInitialize() {
	if cs == nil {
		return
	}
	for _, c := range cs.actorSystem.config.components {
		if err := c.OnInitialize(cs.actorSystem); err != nil {
			panic(err)
		}
	}
}

func (cs *components) onShutdown() {
	if cs == nil {
		return
	}
	for _, c := range cs.shutdown {
		if err := c.OnShutdown(cs.actorSystem); err != nil {
			cs.actorSystem.Logger().Error("component", log.String("name", reflect.TypeOf(c).Name()), log.String("status", "shutdown failed"), log.Err(err))
		}
	}
}
