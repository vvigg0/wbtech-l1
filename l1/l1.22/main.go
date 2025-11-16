package main

import (
	"errors"
	"fmt"
	"math/big"
)

type BigVars struct {
	A *big.Int
	B *big.Int
}

func loadBigVars(a, b string) (*BigVars, error) {
	bigA, ok := new(big.Int).SetString(a, 10)
	if !ok {
		return nil, errors.New("can't parse variable a")
	}
	bigB, ok := new(big.Int).SetString(b, 10)
	if !ok {
		return nil, errors.New("can't parse variable b")
	}
	return &BigVars{A: bigA, B: bigB}, nil
}

func (bV *BigVars) add() *big.Int {
	return new(big.Int).Add(bV.A, bV.B)
}

func (bV *BigVars) sub() *big.Int {
	return new(big.Int).Sub(bV.A, bV.B)
}
func (bV *BigVars) mul() *big.Int {
	return new(big.Int).Mul(bV.A, bV.B)
}
func (bV *BigVars) div() (*big.Int, error) {
	if bV.B.Sign() == 0 {
		return new(big.Int), errors.New("division by zero")
	}
	return new(big.Int).Div(bV.A, bV.B), nil
}
func main() {
	a := "98765432109876543210"
	b := "0"

	result, err := loadBigVars(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.add().String())
	fmt.Println(result.sub().String())
	fmt.Println(result.mul().String())
	x, err := result.div()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x.String())
	}
}
