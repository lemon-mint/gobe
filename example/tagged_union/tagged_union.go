package taggedunion

type MyEnum struct {
	Type Type
	A    *A `gobe_enum:"Type=AType"`
	B    *B `gobe_enum:"Type=BType"`
	C    *C `gobe_enum:"Type=CType"`
}

type Type uint16

const (
	AType Type = iota
	BType
	CType
)

type A struct {
	Val uint8
}

type B struct {
	Val uint16
}

type C struct {
	Val uint32
}
