package utils

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
