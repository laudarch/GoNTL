# GoNTL

Sets of useful functions for cryptanalysis/math

```go
new_big_int := GoNTL.SetScalar(5)

fmt.Println(GoNTL.CmpScalar(new_big_int, 5)) // -> 0

fmt.Println(new_big_int.Mul(new_big_int, GoNTL.big0)) // -> 0
```
