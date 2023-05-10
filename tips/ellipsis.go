package tips

type Opt struct {
	Name string
	Num  int
	IsOk bool
}

type Option func(*Opt)

func WithName(name string) Option {
	return func(opt *Opt) {
		opt.Name = name
	}
}

func WithNum(num int) Option {
	return func(opt *Opt) {
		opt.Num = num
	}
}

func WithIsOk(isOk bool) Option {
	return func(opt *Opt) {
		opt.IsOk = isOk
	}
}

// VariadicFunc using ellipsis, return the length of args
func VariadicFunc(args ...Option) int {
	return len(args)
}
