package expose

var LoginExpose Login

type Login interface {
	Name() string
}
