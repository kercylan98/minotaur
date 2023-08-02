package combination

// NewValidator 创建一个新的校验器
func NewValidator[T Item](options ...ValidatorOption[T]) *Validator[T] {
	var validator = &Validator[T]{}
	for _, option := range options {
		option(validator)
	}
	return validator
}

// Validator 用于对组合进行验证的校验器
type Validator[T Item] struct {
	vh []func(items []T) bool // 用于对组合进行验证的函数
}

// Validate 校验组合是否符合要求
func (slf *Validator[T]) Validate(items []T) bool {
	for _, handle := range slf.vh {
		if !handle(items) {
			return false
		}
	}
	return true
}
