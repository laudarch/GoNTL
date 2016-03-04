package gontl

import(
	"testing"
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
