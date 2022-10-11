package structs

import "math/big"

type aa[T any] struct {
	UnsigedInt uint
	Big        big.Int
	bb         bb
	V          T
}

func (a aa[T]) Foo() {}

func ss() {

}

type C aa[int64]

type bb struct {
	VV *bb
}

func init() {
	V := aa[int]{}
	_ = V
}
