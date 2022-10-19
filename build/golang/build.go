package golang

import (
	"fmt"
	"go/format"

	"github.com/lemon-mint/gobe/build"
)

func go_format(b []byte) []byte {
	v, err := format.Source(b)
	if err != nil {
		return b
	}
	return v
}

func appendSize(b []byte, d build.GenerateData, name string) []byte {
	switch d := d.(type) {
	case *build.U8GenerateData, *build.BoolGenerateData:
		b = append(b, "__size += 1\n"...)
	case *build.VarUintGenerateData:
		b = fmt.Appendf(
			b,
			"__size += (bits.Len64(uint64(%s|1))+6)/7\n",
			name,
		)
	case *build.Float32GenerateData:
		b = append(b, "__size += 4\n"...)
	case *build.Float64GenerateData:
		b = append(b, "__size += 8\n"...)
	case *build.BytesGenerateData:
		b = fmt.Appendf(
			b,
			"__size += (bits.Len64(uint64(len(%s)|1))+6)/7) + uint64(len(%s))\n",
			name,
			name,
		)
	case *build.StringGenerateData:
		b = fmt.Appendf(
			b,
			"__size += (bits.Len64(uint64(len(%s)|1))+6)/7) + uint64(len(%s))\n",
			name,
			name,
		)
	case *build.Complex64GenerateData:
		b = append(b, "__size += 8\n"...)
	case *build.Complex128GenerateData:
		b = append(b, "__size += 16\n"...)
	case *build.StructGenerateData:
		for _, f := range d.Fields {
			b = appendSize(b, f.Data, name+"."+f.Name)
		}
	case *build.ArrayGenerateData:
		b = fmt.Appendf(
			b,
			"for __i := range %s {\n",
			name,
		)
		b = appendSize(b, d.Data, name+"[__i]")
		b = append(b, '}', '\n')
	case *build.SliceGenerateData:
		b = fmt.Appendf(
			b,
			"__size += (bits.Len64(uint64(len(%s)|1))+6)/7\n",
			name,
		)
		b = fmt.Appendf(
			b,
			"for __i := range %s {\n",
			name,
		)
		b = appendSize(b, d.Data, name+"[__i]")
		b = append(b, '}', '\n')
	case *build.CustomGenerateData:
		b = fmt.Appendf(
			b,
			"__size += %s.SizeGOBE()\n",
			name,
		)
	case *build.MapGenerateData:
		// TODO: implement
	case *build.PointerGenerateData:
		b = fmt.Appendf(
			b,
			"if %s != nil {\n"+
				"__ptr += 1\n"+
				"__size += (bits.Len64(uint64(__ptr|1))+6)/7\n",
			name,
		)
		b = appendSize(b, d.Data, "(*"+name+")")
		b = append(b, "} else { __size += 1 }\n"...) // Existence flag
	}
	return b
}
func Generate(n build.NamedInfo) []byte {
	var b []byte

	switch v := n.Data.(type) {
	case *build.StructGenerateData:
		b = fmt.Appendf(
			b,
			"func (v *%s) SizeGOBE() uint64 {\n"+
				"var __ptr uint64\n"+
				"var __size uint64\n"+
				"_ = __ptr\n",
			n.Name,
		)
		b = appendSize(b, v, "v")
		b = append(b, "return __size\n}\n\n"...)
	}
	return go_format(b)
}
