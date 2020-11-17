package must

// NilErr panics if given err is not nil.
func NilErr(err error) {
	if err != nil {
		panic(err)
	}
}

// True panics if given arg is not truthy.
func True(truthy bool) {
	if !truthy {
		panic("not true")
	}
}
