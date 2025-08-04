package exercises

import "fmt"

type Printable interface {
	~int | ~float64 | ~string | any
	fmt.Stringer
}

func GenericPrint[T Printable](val T) {
	fmt.Println(val)
}
