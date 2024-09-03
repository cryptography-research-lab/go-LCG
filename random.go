package go_LCG

const (
	// TODO 设置合适的默认值
	DefaultA = 69069
	DefaultC = 1
	DefaultM = 4294967296
)

// LCG 线性同余发生器
type LCG struct {

	// 乘数
	A int

	// 增量
	C int

	// 模数
	M int

	Seed int
}

type LinearCongruentialGenerator LCG

func NewLCG() *LCG {
	return &LCG{
		A:    DefaultA,
		C:    DefaultC,
		M:    DefaultM,
		Seed: 1,
	}
}

func (x *LCG) Next() int {
	// R0=（A*seed+C）mod M
	x.Seed = (x.A*x.Seed + x.C) % x.M
	return x.Seed
}
