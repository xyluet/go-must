package must

// NilErr panics if given err is not nil.
func NilErr(err error) {
	if err != nil {
		panic(err)
	}
}
