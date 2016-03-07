package gontl

import(
	"testing"
	"math/big"
	"log"
)

func TestPollardRho(t *testing.T){
	// test Pollard Rho
	problem := SetScalar(549862755487)
	modulus := SetScalar(1127698306139)
	order := SetScalar(563849153069)
	secret := SetScalar(150672668021)

	//fmt.Println("testing Pollard Rho, secret to find is", secret)

	result := Pollard_Rho(problem, Big2, modulus, order, Big0, Big0)

	//fmt.Println("secret found:", result)

	if result.Cmp(secret) != 0 {
		t.Error("Expected", secret, ", got ", result)
	}
}

func Test2CRT(t *testing.T){

	x := SetScalar(39248302948940234)
	p := SetScalar(2706403677983)
	q := SetScalar(2402895179297)
	n := new(big.Int)
	n.Mul(p, q)

	xp := new(big.Int)
	xp.Mod(x, p)
	xq := new(big.Int)
	xq.Mod(x, q)

	result := CRT2(xp, xq, p, q)
	if result.Cmp(x) != 0 {
		t.Error("Expected", x, ", got ", result)
	}
}
