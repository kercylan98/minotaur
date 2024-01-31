package expose

var AttackExpose Attack

type Attack interface {
	Name() string
}
