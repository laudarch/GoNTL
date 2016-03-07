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
```

## Functions

Pollard Rho

```
Pollard_Rho(problem *big.Int, generator *big.Int, modulus *big.Int, order *big.Int, a *big.Int, b *big.Int) (discrete_log *big.Int)
```

Chinese Remainder Theorem:

```
CRT2(a *big.Int, b *big.Int, moda *big.Int, modb *big.Int) (*big.Int)
```

