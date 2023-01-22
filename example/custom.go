package main

type GOBE_CUSTOM_TYPE interface {
	ZZMarshalGOBE(dst []byte) uint64
	ZZUnmarshalGOBE(src []byte) (offset uint64, ok bool)
	ZZSizeGOBE() uint64
}

type CustomUint8 uint8

var _ GOBE_CUSTOM_TYPE = (*CustomUint8)(nil)

func (c *CustomUint8) ZZMarshalGOBE(dst []byte) uint64 {
	dst[0] = byte(*c)
	return 1
}

func (c *CustomUint8) ZZUnmarshalGOBE(src []byte) (offset uint64, ok bool) {
	*c = CustomUint8(src[0])
	return 1, true
}

func (c *CustomUint8) ZZSizeGOBE() uint64 {
	return 1
}
