package bratatat

func Do(f func()) {
	for i := 0; i < 3; i++ {
		f()
	}
}
