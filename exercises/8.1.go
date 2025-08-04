package exercises

type Number interface {
	~int | ~float32 | ~float64
}

func Double[T Number](val T) T {
	return val * 2
}
