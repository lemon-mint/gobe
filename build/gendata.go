package build

import (
	"unicode"
	"unicode/utf8"

	"github.com/lemon-mint/gobe/typeinfo"
)

type GenerateData interface {
	GenerateData()
}

type StructFieldGenerateData struct {
	Name string
	Type typeinfo.StructField

	Data GenerateData
}

type StructGenerateData struct {
	Type *typeinfo.StructType

	Fields []StructFieldGenerateData
}

func (s *StructGenerateData) GenerateData() {}

type PointerGenerateData struct {
	Type *typeinfo.PointerType

	Data GenerateData
}

func (p *PointerGenerateData) GenerateData() {}

type SliceGenerateData struct {
	Type *typeinfo.SliceType

	Data GenerateData
}

func (s *SliceGenerateData) GenerateData() {}

type ArrayGenerateData struct {
	Type *typeinfo.ArrayType

	Data GenerateData
}

func (a *ArrayGenerateData) GenerateData() {}

type MapGenerateData struct {
	Type *typeinfo.MapType

	KeyData  GenerateData
	ElemData GenerateData
}

func (m *MapGenerateData) GenerateData() {}

type CustomGenerateData struct {
	Type *typeinfo.CustomType
}

func (c *CustomGenerateData) GenerateData() {}

type BytesGenerateData struct {
	Type *typeinfo.SliceType
}

func (b *BytesGenerateData) GenerateData() {}

type StringGenerateData struct {
	Type *typeinfo.BasicType
}

func (s *StringGenerateData) GenerateData() {}

type U8GenerateData struct {
	Type *typeinfo.BasicType
}

func (u *U8GenerateData) GenerateData() {}

type BoolGenerateData struct {
	Type *typeinfo.BasicType
}

func (b *BoolGenerateData) GenerateData() {}

type VarUintGenerateData struct {
	Type *typeinfo.BasicType
}

func (v *VarUintGenerateData) GenerateData() {}

type Float32GenerateData struct {
	Type *typeinfo.BasicType
}

func (f *Float32GenerateData) GenerateData() {}

type Float64GenerateData struct {
	Type *typeinfo.BasicType
}

func (f *Float64GenerateData) GenerateData() {}

type Complex64GenerateData struct {
	Type *typeinfo.BasicType
}

func (c *Complex64GenerateData) GenerateData() {}

type Complex128GenerateData struct {
	Type *typeinfo.BasicType
}

func (c *Complex128GenerateData) GenerateData() {}

type ConvContext struct {
	TypeInfoGendataMap map[typeinfo.Type]GenerateData
}

func Conv(t typeinfo.Type, ctx *ConvContext) GenerateData {
	if ctx == nil {
		ctx = &ConvContext{
			TypeInfoGendataMap: make(map[typeinfo.Type]GenerateData),
		}
	}

	if v, ok := ctx.TypeInfoGendataMap[t]; ok {
		return v
	}

	switch t := t.(type) {
	case *typeinfo.StructType:
		g := &StructGenerateData{
			Type: t,
		}
		ctx.TypeInfoGendataMap[t] = g

		for _, f := range t.Fields {
			// Ignore unexported fields
			ch, _ := utf8.DecodeRuneInString(f.Name)
			if !unicode.IsUpper(ch) {
				continue
			}

			g.Fields = append(g.Fields, StructFieldGenerateData{
				Name: f.Name,
				Type: f,
				Data: Conv(f.Type, ctx),
			})
		}
		return g
	case *typeinfo.PointerType:
		g := &PointerGenerateData{
			Type: t,
		}
		ctx.TypeInfoGendataMap[t] = g
		g.Data = Conv(t.Elem, ctx)
		return g
	case *typeinfo.SliceType:
		// []byte, []int8 optimization
		switch typeinfo.TypeKind(t) {
		case typeinfo.Uint8:
			return &BytesGenerateData{
				Type: t,
			}
		case typeinfo.Int8:
			return &BytesGenerateData{
				Type: t,
			}
		}

		g := &SliceGenerateData{
			Type: t,
		}
		ctx.TypeInfoGendataMap[t] = g
		g.Data = Conv(t.Elem, ctx)
		return g
	case *typeinfo.ArrayType:
		g := &ArrayGenerateData{
			Type: t,
		}
		ctx.TypeInfoGendataMap[t] = g
		g.Data = Conv(t.Elem, ctx)
		return g
	case *typeinfo.MapType:
		g := &MapGenerateData{
			Type: t,
		}
		ctx.TypeInfoGendataMap[t] = g
		g.KeyData = Conv(t.Key, ctx)
		g.ElemData = Conv(t.Elem, ctx)
		return g
	case *typeinfo.CustomType:
		return &CustomGenerateData{
			Type: t,
		}
	case *typeinfo.BasicType:
		switch typeinfo.TypeKind(t) {
		case typeinfo.String:
			return &StringGenerateData{
				Type: t,
			}
		case typeinfo.Uint8, typeinfo.Int8:
			return &U8GenerateData{
				Type: t,
			}
		case typeinfo.Bool:
			return &BoolGenerateData{
				Type: t,
			}
		case typeinfo.Uint, typeinfo.Uint16, typeinfo.Uint32, typeinfo.Uint64, typeinfo.Uintptr,
			typeinfo.Int, typeinfo.Int16, typeinfo.Int32, typeinfo.Int64:
			return &VarUintGenerateData{
				Type: t,
			}
		case typeinfo.Complex64:
			return &Complex64GenerateData{
				Type: t,
			}
		case typeinfo.Complex128:
			return &Complex128GenerateData{
				Type: t,
			}
		case typeinfo.Float32:
			return &Float32GenerateData{
				Type: t,
			}
		case typeinfo.Float64:
			return &Float64GenerateData{
				Type: t,
			}
		}
	}
	return nil
}

type NamedInfo struct {
	Name string

	Data GenerateData
}
