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
