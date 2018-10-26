package foobarbaz

type Bar struct {
	X int
}

func ProvideBar(foo Foo) Bar {
	return Bar{-foo.X}
}
