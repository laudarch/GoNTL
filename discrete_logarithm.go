package gontl

import(
	"math/big"
	"log"
)

////////////////////////////////////////////////////////////////////////
// Multiplicative groups
////////////////////////////////////////////////////////////////////////

func iteration(x *big.Int, a *big.Int, b *big.Int, alpha *big.Int, beta *big.Int, modulus *big.Int, order *big.Int) {

	test := new(big.Int)
	three := SetScalar(3)
	test.Mod(x, three)

	if CmpScalar(test, 1) == 0 {
		x.Mul(x, beta)
		x.Mod(x, modulus)

		b.Add(b, Big1)
		b.Mod(b, order)

	} else if CmpScalar(test, 0) == 0 {
		x.Mul(x, x)
		x.Mod(x, modulus)

		a.Mul(a, Big2)
		a.Mod(a, order)
		b.Mul(b, Big2)
		b.Mod(b, order)

	} else {
		x.Mul(x, alpha)
		x.Mod(x, modulus)

		a.Add(a, Big1)
		a.Mod(a, order)
	}
}

/*
 * This is the implementation of the algorithm introduced in
 * the Handbook of Applied Cryptography chapter 3.6.3
 * It is definitely not the latest improvement, 
 * nor the parallelizable version.
 * You need to know the order of the generator to use it
 */
func Pollard_Rho(problem *big.Int, generator *big.Int, modulus *big.Int, order *big.Int, a *big.Int, b *big.Int) (discrete_log *big.Int) {
	// init
	alpha := new(big.Int)
	alpha.Set(generator)

	beta := new(big.Int)
	beta.Set(problem)

	// x
	x := new(big.Int)

	if a.Cmp(Big0) != 0 || b.Cmp(Big0) != 0{
		x.Exp(alpha, a, modulus)
		y:= new(big.Int)
		y.Exp(beta, b, modulus)
		x.Mul(x, y)
		x.Mod(x, modulus)
	} else {
		x.SetInt64(int64(1))
	}

	x1 := new(big.Int)
	x2 := new(big.Int)
	x1.Set(x)
	x2.Set(x)

	a1 := new(big.Int)
	a2 := new(big.Int)
	a1.Set(a)
	a2.Set(a)
	
	b1 := new(big.Int)
	b2 := new(big.Int)
	b1.Set(b)
	b2.Set(b)

	// loop
	for {
		// iterate
		iteration(x1, a1, b1, alpha, beta, modulus, order)

		iteration(x2, a2, b2, alpha, beta, modulus, order)
		iteration(x2, a2, b2, alpha, beta, modulus, order)

		// detect collision
		if x1.Cmp(x2) == 0 {
			r := new(big.Int)
			r.Sub(b1, b2)
			r.Mod(r, order)
			if r.Cmp(Big0) == 0 {
				break // failure
			} else {
				r.ModInverse(r, order)
				a1.Sub(a2, a1)
				a1.Mod(a1, order)
				r.Mul(r, a1)
				return r.Mod(r, order)
			}
						
		}
	}
	log.Println("failure")
	// failure
	return Big0
}

////////////////////////////////////////////////////////////////////////
// Finite Fields
////////////////////////////////////////////////////////////////////////


////////////////////////////////////////////////////////////////////////
// Elliptic Curves
////////////////////////////////////////////////////////////////////////

