package transport

type (
	FiberLoadedHook  func(app *FiberApp) error
	FiberMountedHook func(app *FiberApp) error
	FiberLaunchHook  func(app *FiberApp) error
)

func newFiberHooks(app *FiberApp) *FiberHooks {
	return &FiberHooks{
		app: app,
	}
}

type FiberHooks struct {
	app          *FiberApp
	loadedHooks  []FiberLoadedHook
	mountedHooks []FiberMountedHook
	launchHooks  []FiberLaunchHook
}

func (h *FiberHooks) BindLoadedHook(hooks ...FiberLoadedHook) *FiberHooks {
	h.loadedHooks = append(h.loadedHooks, hooks...)
	return h
}

func (h *FiberHooks) onLoadedHook() error {
	for _, hook := range h.loadedHooks {
		if err := hook(h.app); err != nil {
			return err
		}
	}

	return nil
}

func (h *FiberHooks) BindMountedHook(hooks ...FiberMountedHook) *FiberHooks {
	h.mountedHooks = append(h.mountedHooks, hooks...)
	return h
}

func (h *FiberHooks) onMountedHook() error {
	for _, hook := range h.mountedHooks {
		if err := hook(h.app); err != nil {
			return err
		}
	}

	return nil
}
func (h *FiberHooks) BindLaunchHook(hooks ...FiberLaunchHook) *FiberHooks {
	h.launchHooks = append(h.launchHooks, hooks...)
	return h
}

func (h *FiberHooks) onLaunchHook() error {
	for _, hook := range h.launchHooks {
		if err := hook(h.app); err != nil {
			return err
		}
	}

	return nil
}
