package GoNTL

import(
	"math/big"
)

func power_mod(base big.Int, exponent big.Int, modulus big.Int) (result big.Int) {
	
	if CmpScalar(modulus, 1) == 0 {
		return SetScalar(0)
	}
	
	result := SetScalar(1)

	//
	
	if modulus = 1 then return 0
	Assert :: (modulus - 1) * (modulus - 1) does not overflow base
	result := 1
	base := base mod modulus
	while exponent > 0
	if (exponent mod 2 == 1):
	result := (result * base) mod modulus
	exponent := exponent >> 1
	base := (base * base) mod modulus
	    return result
}
