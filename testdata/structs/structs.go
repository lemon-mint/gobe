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

type CustomType struct {
	AnyJSON interface{}
}

func (c *CustomType) MarshalGOBE(dst []byte) error {
	return nil
}

func (c *CustomType) UnmarshalGOBE(src []byte) error {
	return nil
}

func (c *CustomType) SizeGOBE() uint64 {
	return 0
}

func init() {
	V := aa[int]{}
	_ = V
}
