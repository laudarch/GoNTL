# GoNTL

Sets of useful functions for cryptanalysis/math

Don't use in production

```go
package main

import(
    "github.com/mimoo/GoNTL"
)

new_big_int := gontl.SetScalar(5)

fmt.Println(gontl.CmpScalar(new_big_int, 5)) // -> 0

fmt.Println(new_big_int.Mul(new_big_int, gontl.Big0)) // -> 0

// Pollard Rho for discrete logarithm
problem := gontl.SetScalar(549862755487)
modulus := gontl.SetScalar(1127698306139)
order := gontl.SetScalar(563849153069)
secret := gontl.SetScalar(150672668021)

fmt.Println(gontl.Pollard_Rho(problem, gontl.Big2, modulus, order, gontl.Big0, gontl.Big0)) // -> secret
```
