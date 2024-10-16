package cores

import (
	"math/big"
	"math/rand"
)

type Bool = bool
type Int = int
type Uint = uint
type Int16 = int16
type Uint16 = uint16
type Int32 = int32
type Uint32 = uint32
type Int64 = int64
type Uint64 = uint64
type Float32 = float32
type Float64 = float64
type String = string

type Char = int8
type Byte = uint8
type Short = int16
type Ushort = uint16
type Wchar = uint16
type Word = uint16
type Dword = uint
type Dword32 = uint32
type Dword64 = uint64

type Long = int
type Ulong = uint

type Long64 = int64
type Ulong64 = uint64

type UintPtr = uintptr

type Float = float32
type Double = float64

type UnionTypeBool interface {
	~bool | Bool
}

type UnionTypeInt interface {
	~int | Int | Long
}

type UnionTypeUint interface {
	~uint | Uint | Ulong
}

type UnionTypeInt16 interface {
	~int16 | Int16 | Short
}

type UnionTypeUint16 interface {
	~uint16 | Uint16 | Ushort
}

type UnionTypeInt32 interface {
	~int32 | Int32
}

type UnionTypeUint32 interface {
	~uint32 | Uint32
}

type UnionTypeInt64 interface {
	~int64 | Int64 | Long64
}

type UnionTypeUint64 interface {
	~uint64 | Uint64 | Ulong64
}

type UnionTypeFloat32 interface {
	~float32 | Float32 | Float
}

type UnionTypeFloat64 interface {
	~float64 | Float64 | Double
}

type UnionTypeString interface {
	~string | String
}

type UnionTypeChar interface {
	~int8 | Char
}

type UnionTypeByte interface {
	~uint8 | Byte
}

type UnionTypeShort interface {
	~int16 | Short
}

type UnionTypeUshort interface {
	~uint16 | Ushort
}

type UnionTypeWchar interface {
	~uint16 | Wchar | Word | Ushort
}

type UnionTypeWord interface {
	~uint16 | Wchar | Word | Ushort
}

type UnionTypeDword interface {
	~uint | Dword | Uint | Ulong
}

type UnionTypeDword32 interface {
	~uint32 | Dword32 | Uint32
}

type UnionTypeDword64 interface {
	~uint64 | Dword64 | Uint64 | Ulong64
}

type UnionTypeLong interface {
	~int | Long | Int
}

type UnionTypeUlong interface {
	~uint | Ulong | Uint
}

type UnionTypeLong64 interface {
	~int64 | Long64 | Int64
}

type UnionTypeUlong64 interface {
	~uint64 | Ulong64 | Uint64
}

type UnionTypeUintPtr interface {
	~uintptr | UintPtr | Uint64 | Ulong64
}

type UnionTypeFloat interface {
	~float32 | Float | Float32
}

type UnionTypeDouble interface {
	~float64 | Double | Float64
}

type Ordering int

var (
	OrderingLess    Ordering = -1
	OrderingEqual   Ordering = 0
	OrderingGreater Ordering = 1
)

type ComparableImpl[T any] interface {
	CompareTo(other T) Ordering
	Equals(other T) bool
}

type EquatableImpl[T any] interface {
	Equals(other T) bool
}

type StringableImpl interface {
	ToString() string
}

type CountableImpl interface {
	Len() int
}

type HashableImpl interface {
	HashCode() int
}

type JsonableImpl interface {
	ToJson() string
}

type YamlableImpl interface {
	ToYaml() string
}

type TomlableImpl interface {
	ToToml() string
}

type XamlableImpl interface {
	ToXaml() string
}

type BigFloatImpl interface {
	SetPrec(prec uint) *big.Float
	SetMode(mode big.RoundingMode) *big.Float
	Prec() uint
	MinPrec() uint
	Mode() big.RoundingMode
	Acc() big.Accuracy
	Sign() int
	MantExp(mant *big.Float) (exp int)
	SetMantExp(mant *big.Float, exp int) *big.Float
	Signbit() bool
	IsInf() bool
	IsInt() bool
	SetUint64(x uint64) *big.Float
	SetInt64(x int64) *big.Float
	SetFloat64(x float64) *big.Float
	SetInt(x *big.Int) *big.Float
	SetRat(x *big.Rat) *big.Float
	SetInf(signbit bool) *big.Float
	Set(x *big.Float) *big.Float
	Copy(x *big.Float) *big.Float
	Uint64() (uint64, big.Accuracy)
	Int64() (int64, big.Accuracy)
	Float32() (float32, big.Accuracy)
	Float64() (float64, big.Accuracy)
	Int(z *big.Int) (*big.Int, big.Accuracy)
	Rat(z *big.Rat) (*big.Rat, big.Accuracy)
	Abs(x *big.Float) *big.Float
	Neg(x *big.Float) *big.Float
	Add(x, y *big.Float) *big.Float
	Sub(x, y *big.Float) *big.Float
	Mul(x, y *big.Float) *big.Float
	Quo(x, y *big.Float) *big.Float
	Cmp(y *big.Float) int
}

type BigIntImpl interface {
	Sign() int
	SetInt64(x int64) *big.Int
	SetUint64(x uint64) *big.Int
	Set(x *big.Int) *big.Int
	Bits() []big.Word
	SetBits(abs []big.Word) *big.Int
	Abs(x *big.Int) *big.Int
	Neg(x *big.Int) *big.Int
	Add(x, y *big.Int) *big.Int
	Sub(x, y *big.Int) *big.Int
	Mul(x, y *big.Int) *big.Int
	MulRange(a, b int64) *big.Int
	Binomial(n, k int64) *big.Int
	Quo(x, y *big.Int) *big.Int
	Rem(x, y *big.Int) *big.Int
	QuoRem(x, y, r *big.Int) (*big.Int, *big.Int)
	Div(x, y *big.Int) *big.Int
	Mod(x, y *big.Int) *big.Int
	DivMod(x, y, m *big.Int) (*big.Int, *big.Int)
	Cmp(y *big.Int) (r int)
	CmpAbs(y *big.Int) int
	Int64() int64
	Uint64() uint64
	IsInt64() bool
	IsUint64() bool
	Float64() (float64, big.Accuracy)
	SetString(s string, base int) (*big.Int, bool)
	SetBytes(buf []byte) *big.Int
	Bytes() []byte
	FillBytes(buf []byte) []byte
	BitLen() int
	TrailingZeroBits() uint
	Exp(x, y, m *big.Int) *big.Int
	GCD(x, y, a, b *big.Int) *big.Int
	Rand(rnd *rand.Rand, n *big.Int) *big.Int
	ModInverse(g, n *big.Int) *big.Int
	ModSqrt(x, p *big.Int) *big.Int
	Lsh(x *big.Int, n uint) *big.Int
	Rsh(x *big.Int, n uint) *big.Int
	Bit(i int) uint
	SetBit(x *big.Int, i int, b uint) *big.Int
	And(x, y *big.Int) *big.Int
	AndNot(x, y *big.Int) *big.Int
	Or(x, y *big.Int) *big.Int
	Xor(x, y *big.Int) *big.Int
	Not(x *big.Int) *big.Int
	Sqrt(x *big.Int) *big.Int
}

type RandomImpl interface {
	Seed(seed int64)
	Int63() int64
	Uint32() uint32
	Uint64() uint64
	Int31() int32
	Int() int
	Int63n(n int64) int64
	Int31n(n int32) int32
	Intn(n int) int
	Float64() float64
	Float32() float32
	Perm(n int) []int
	Shuffle(n int, swap func(i, j int))
	Read(p []byte) (n int, err error)
}
