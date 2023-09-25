package calc

import "math/big"

func BigModulus(base int, exponent int, divisor int) int {
	rem := new(big.Int).Exp(big.NewInt(int64(base)), big.NewInt(int64(exponent)), big.NewInt(int64(divisor)))
	rem_int := int(rem.Int64())
	return rem_int
}
