package compsortingalgo

func swap(a, b *int32) {
	tmp := *a
	*a = *b
	*b = tmp
}
