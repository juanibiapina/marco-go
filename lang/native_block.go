package lang

type nativeFunc func(*environment, *environment) Expr

type nativeBlock struct {
	f nativeFunc
}

func MakeNativeBlock(f nativeFunc) *nativeBlock {
	return &nativeBlock{f}
}

func (b *nativeBlock) Invoke(closure *environment, dynamic *environment) Expr {
	return b.f(closure, dynamic)
}

func (b *nativeBlock) String() string {
	return "NativeBlock"
}
