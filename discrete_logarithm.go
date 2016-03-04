package GoNTL

import(
	"math/big"
)

////////////////////////////////////////////////////////////////////////
// Multiplicative groups
////////////////////////////////////////////////////////////////////////

func iteration_function(x big.Int) (y big.Int) {

}

/*
 * This is the implementation of the algorithm introduced in
 * the Handbook of Applied Cryptography chapter 3.6.3
 * It is definitely not the latest improvement, 
 * nor the parallelizable version.
 * You need to know the order of the generator to use it
 */
func Pollard_Rho(problem bit.Int, generator big.Int, modulus big.Int, order big.int, a big.Int, b big.Int) (discrete_log bit.Int) {
	// init
	if generator.CmpScalar(0) {
		alpha := SetScalar(2)
	}
	big.Set(beta, problem)

	if a.CmpScalar(0) || b.CmpScalar(0) {
		
	}
}

////////////////////////////////////////////////////////////////////////
// Finite Fields
////////////////////////////////////////////////////////////////////////


////////////////////////////////////////////////////////////////////////
// Elliptic Curves
////////////////////////////////////////////////////////////////////////

