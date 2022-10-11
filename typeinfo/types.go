package typeinfo

import (
	"reflect"
	"strconv"
)

type Type interface {
	String() string
}

type Kind uint16

//go:generate go run golang.org/x/tools/cmd/stringer -type=Kind

const Version = uint64(9988560860994558250)

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	String
	UnsafePointer

	Struct
	Pointer
	Slice
	Array
	Map
)

func TypeKind(t Type) Kind {
	switch t := t.(type) {
	case *BasicType:
		switch t.Kind {
		case "bool":
			return Bool
		case "int":
			return Int
		case "int8":
			return Int8
		case "int16":
			return Int16
		case "int32":
			return Int32
		case "int64":
			return Int64
		case "uint":
			return Uint
		case "uint8":
			return Uint8
		case "uint16":
			return Uint16
		case "uint32":
			return Uint32
		case "uint64":
			return Uint64
		case "uintptr":
			return Uintptr
		case "float32":
			return Float32
		case "float64":
			return Float64
		case "complex64":
			return Complex64
		case "complex128":
			return Complex128
		case "string":
			return String
		case "unsafe.Pointer":
			return UnsafePointer
		}
	case *StructType:
		return Struct
	case *PointerType:
		return Pointer
	case *SliceType:
		return Slice
	case *ArrayType:
		return Array
	case *MapType:
		return Map
	case *invalidType:
		return Invalid
	}

	return Invalid
}

type BasicType struct {
	Kind string
}

func (b BasicType) String() string {
	return b.Kind
}

type StructType struct {
	Fields []StructField
}

type StructTag = reflect.StructTag

type StructField struct {
	Name      string
	Type      Type
	Tag       StructTag
	Anonymous bool
}

func (s StructType) String() string {
	return "struct { ... }"
}

type PointerType struct {
	Elem Type
}

func (p PointerType) String() string {
	return "*" + p.Elem.String()
}

type SliceType struct {
	Elem Type
}

func (s SliceType) String() string {
	return "[]" + s.Elem.String()
}

type ArrayType struct {
	Elem Type
	Len  int64
}

func (a ArrayType) String() string {
	return "[" + strconv.FormatInt(a.Len, 10) + "]" + a.Elem.String()
}

type MapType struct {
	Key, Elem Type
}

func (m MapType) String() string {
	return "map[" + m.Key.String() + "]" + m.Elem.String()
}

type invalidType struct{}

func (invalidType) String() string {
	return "invalid"
}

var InvalidType = &invalidType{}
