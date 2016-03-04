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

// Compare a big int to a scalar
func CmpScalar(a *big.Int, scalar int) (int) {
	b := SetScalar(scalar)
	return a.Cmp(b)
}

