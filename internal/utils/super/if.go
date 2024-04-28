package super

func If[T any](expression bool, t T, f T) T {
	if expression {
		return t
	}
	return f
}
