package GoNTL

import(
	"math/big"
)

// Set a big int to a scalar
func SetScalar(scalar int) (*big.Int) {
	a := new(big.Int)
	a.SetInt64(int64(scalar))
	return a
}

var big0 *big.Int = SetScalar(0)
var big1 *big.Int = SetScalar(1)
var big2 *big.Int = SetScalar(2)

// Compare a big int to a scalar
func CmpScalar(a *big.Int, scalar int) (int) {
	b := SetScalar(scalar)
	return a.Cmp(b)
}

func AddScalar(a *big.Int, scalar int) (*big.Int) {
	b := SetScalar(scalar)
	result := new(big.Int)
	return result.Add(a, b)
}

func MulScalar(a *big.Int, scalar int) (*big.Int) {
	b := SetScalar(scalar)
	result := new(big.Int)
	return result.Mul(a, b)
}

