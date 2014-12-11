package lang

type nativeFunc func(*Environment, *Environment) Expr

type nativeBlock struct {
	f nativeFunc
}

func MakeNativeBlock(f nativeFunc) *nativeBlock {
	return &nativeBlock{f}
}

func (b *nativeBlock) Invoke(closure *Environment, dynamic *Environment) Expr {
	return b.f(closure, dynamic)
}

func (b *nativeBlock) String() string {
	return "NativeBlock"
}
