package modular

// Dimension 维度接口
//   - 维度与服务的区别在于，维度是对非全局性的服务进行抽象，例如：依赖特定游戏房间的局内玩家管理服务
type Dimension[Owner any] interface {
	// OnInit 服务初始化阶段，该阶段不应该依赖其他任何服务
	OnInit(owner Owner) error

	// OnPreload 预加载阶段，在进入该阶段时，所有服务已经初始化完成，可在该阶段注入其他服务的依赖
	OnPreload() error

	// OnMount 挂载阶段，该阶段所有服务本身及依赖的服务都已经初始化完成，可在该阶段进行服务功能的定义
	OnMount() error
}

// RunDimensions 运行维度
func RunDimensions[Owner any](owner Owner, dimensions ...Dimension[Owner]) error {
	// OnInit
	for _, dimension := range dimensions {
		if err := dimension.OnInit(owner); err != nil {
			return err
		}
	}

	// OnPreload
	for _, dimension := range dimensions {
		if err := dimension.OnPreload(); err != nil {
			return err
		}
	}

	// OnMount
	for _, dimension := range dimensions {
		if err := dimension.OnMount(); err != nil {
			return err
		}
	}

	return nil
}
