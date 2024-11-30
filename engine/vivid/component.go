package vivid

// Component 是以钩子的形式进行实现的用于对 ActorSystem 进行扩展的接口，同时也是各项对 ActorSystem 扩展功能实现的基础
type Component interface {
	// OnInitialize 是在 ActorSystem 初始化完毕并正常运行后执行的方法，它被用于组件生命周期的初始化
	//  - 在该阶段可以生成组件内部的 Actor 或其他工作内容
	OnInitialize(actorSystem *ActorSystem) error
}

// FunctionalComponent 是 Component 的函数式组件
type FunctionalComponent func(actorSystem *ActorSystem) error

func (f FunctionalComponent) OnInitialize(actorSystem *ActorSystem) error {
	return f(actorSystem)
}

// ShutdownComponent 是用于清理组件自身的扩展接口，实现该接口的组件将在 ActorSystem 彻底关闭后执行清理工作
type ShutdownComponent interface {
	Component

	// OnShutdown 在 ActorSystem 彻底关闭后被调用，用于清理组件自身资源
	OnShutdown(actorSystem *ActorSystem) error
}

// FunctionalShutdownComponent 是 ShutdownComponent 的函数式组件，但是它无法在 Component.OnInitialize 中执行任何工作
type FunctionalShutdownComponent func(actorSystem *ActorSystem) error

func (f FunctionalShutdownComponent) OnInitialize(actorSystem *ActorSystem) error {
	return nil
}

func (f FunctionalShutdownComponent) OnShutdown(actorSystem *ActorSystem) error {
	return f(actorSystem)
}

// ActorDefineCaptureComponent 是用于对 Actor 定义进行捕获的扩展接口，实现该接口的组件可以捕获 Actor 定义，并对 Actor 进行一些必要的处理
type ActorDefineCaptureComponent interface {
	Component

	// OnActorDefineCapture 在 Actor 还未生成前被调用，此刻对于 ActorDescriptor 和 ActorProvider 的修改都是有效的，并且会在 Actor 重启时保留影响
	OnActorDefineCapture(actorSystem *ActorSystem, provider ActorProvider, descriptor *ActorDescriptor)
}

// FunctionalActorDefineCaptureComponent 是 ActorDefineCaptureComponent 的函数式组件，但是它无法在 Component.OnInitialize 中执行任何工作
type FunctionalActorDefineCaptureComponent func(actorSystem *ActorSystem, provider ActorProvider, descriptor *ActorDescriptor)

func (f FunctionalActorDefineCaptureComponent) OnInitialize(actorSystem *ActorSystem) error {
	return nil
}

func (f FunctionalActorDefineCaptureComponent) OnActorDefineCapture(actorSystem *ActorSystem, provider ActorProvider, descriptor *ActorDescriptor) {
	f(actorSystem, provider, descriptor)
}

// ActorContextCaptureComponent 是用于在 Actor 创建完成且正常运行后对其 ActorContext 进行捕获的扩展接口
//   - 该接口仅捕获首次次创建的 ActorContext，后续的重启等状态不会被触发
type ActorContextCaptureComponent interface {
	Component

	// OnActorContextCapture 在 Actor 创建完成且正常运行后被调用
	//
	// 特殊标注：
	//  - MarkNonStrictConcurrencySafety
	OnActorContextCapture(ctx ActorContext)
}

// FunctionalActorContextCaptureComponent 是 ActorContextCaptureComponent 的函数式组件，但是它无法在 Component.OnInitialize 中执行任何工作
type FunctionalActorContextCaptureComponent func(ctx ActorContext)

func (f FunctionalActorContextCaptureComponent) OnInitialize(actorSystem *ActorSystem) error {
	return nil
}
func (f FunctionalActorContextCaptureComponent) OnActorContextCapture(ctx ActorContext) {
	f(ctx)
}

// ActorReceiveMessageCaptureComponent 是用于在 Actor 收到消息时进行捕获的扩展接口
//   - 该接口在收到消息但还未处理时被调用
type ActorReceiveMessageCaptureComponent interface {
	Component

	// OnActorReceiveMessageCapture 在 Actor 收到消息但还未处理时被调用
	OnActorReceiveMessageCapture(ctx ActorContext) (abort bool)
}

// FunctionalActorReceiveMessageCaptureComponent 是 ActorReceiveMessageCaptureComponent 的函数式组件，但是它无法在 Component.OnInitialize 中执行任何工作
type FunctionalActorReceiveMessageCaptureComponent func(context ActorContext) (abort bool)

func (f FunctionalActorReceiveMessageCaptureComponent) OnInitialize(actorSystem *ActorSystem) error {
	return nil
}

func (f FunctionalActorReceiveMessageCaptureComponent) OnActorReceiveMessageCapture(ctx ActorContext) (abort bool) {
	return f(ctx)
}

// ActorReplyCaptureComponent 是用于在 Actor 发送回复时进行捕获的扩展接口
type ActorReplyCaptureComponent interface {
	Component

	// OnActorReplyCapture 在 Actor 发送回复时被调用
	OnActorReplyCapture(ctx ActorContext, message Message)
}

// FunctionalActorReplyCaptureComponent 是 ActorReplyCaptureComponent 的函数式组件，但是它无法在 Component.OnInitialize 中执行任何工作
type FunctionalActorReplyCaptureComponent func(ctx ActorContext, message Message)

func (f FunctionalActorReplyCaptureComponent) OnInitialize(actorSystem *ActorSystem) error {
	return nil
}

func (f FunctionalActorReplyCaptureComponent) OnActorReplyCapture(ctx ActorContext, message Message) {
	f(ctx, message)
}
