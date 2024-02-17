package alternative

func FooPublic() int {
	zmienna := fooPrivate()
	return zmienna
}

func fooPrivate() int {
	return 123
}

func Closure() func() int {
	var authMiddleware int
	return func() int {
		authMiddleware++
		return authMiddleware
	}
}

func SumAndReturn(a, b, c, d int, e, f float64) (int, float64) {
	return a + b + c + d, e + f
}
