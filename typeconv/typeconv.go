package typeconv

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"

	"github.com/lemon-mint/gobe/typeinfo"
)

func getCustomTypeIface() *types.Interface {
	fset := token.NewFileSet()
	const code = `
	package typeconv

	type GOBE_CUSTOM_TYPE interface {
		MarshalGOBE(dst []byte) error
		UnmarshalGOBE(src []byte) error
		SizeGOBE() uint64
	}
	`
	f, err := parser.ParseFile(fset, "customtype.go", code, 0)
	if err != nil {
		panic(err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("typeconv", fset, []*ast.File{f}, nil)
	if err != nil {
		panic(err)
	}

	return pkg.Scope().Lookup("GOBE_CUSTOM_TYPE").Type().Underlying().(*types.Interface).Complete()
}

var ctypeIface = getCustomTypeIface()

type TypeConvCtx struct {
	typemap map[types.Type]typeinfo.Type
}

func TypeConv(ctx *TypeConvCtx, tt types.Type) typeinfo.Type {
TAIL_CALL:
	if ctx.typemap == nil {
		ctx.typemap = make(map[types.Type]typeinfo.Type)
	}

	if v, ok := ctx.typemap[tt]; ok {
		return v
	}

	// Check if tt implements GOBE_CUSTOM_TYPE interface.
	if t, ok := tt.(*types.Named); ok {
		if types.AssignableTo(t, ctypeIface) || types.AssignableTo(types.NewPointer(t), ctypeIface) {
			vv := &typeinfo.CustomType{
				Underlying: tt,
			}
			return vv
		} else {
		}
	}

	switch t := tt.(type) {
	case *types.Basic:
		vv := &typeinfo.BasicType{}
		ctx.typemap[t] = vv

		switch t.Kind() {
		case types.Bool:
			vv.Kind = "bool"
		case types.Int:
			vv.Kind = "int"
		case types.Int8:
			vv.Kind = "int8"
		case types.Int16:
			vv.Kind = "int16"
		case types.Int32:
			vv.Kind = "int32"
		case types.Int64:
			vv.Kind = "int64"
		case types.Uint:
			vv.Kind = "uint"
		case types.Uint8:
			vv.Kind = "uint8"
		case types.Uint16:
			vv.Kind = "uint16"
		case types.Uint32:
			vv.Kind = "uint32"
		case types.Uint64:
			vv.Kind = "uint64"
		case types.Uintptr:
			vv.Kind = "uintptr"
		case types.Float32:
			vv.Kind = "float32"
		case types.Float64:
			vv.Kind = "float64"
		case types.Complex64:
			vv.Kind = "complex64"
		case types.Complex128:
			vv.Kind = "complex128"
		case types.String:
			vv.Kind = "string"
		case types.UnsafePointer:
			vv.Kind = "unsafe.Pointer"
		}
		return vv
	case *types.Named:
		tt = tt.Underlying()
		goto TAIL_CALL
	case *types.Pointer:
		vv := &typeinfo.PointerType{}
		ctx.typemap[t] = vv
		vv.Elem = TypeConv(ctx, t.Elem())
		return vv
	case *types.Slice:
		vv := &typeinfo.SliceType{}
		ctx.typemap[t] = vv
		vv.Elem = TypeConv(ctx, t.Elem())
		return vv
	case *types.Array:
		vv := &typeinfo.ArrayType{}
		ctx.typemap[t] = vv
		vv.Elem = TypeConv(ctx, t.Elem())
		vv.Len = t.Len()
		return vv
	case *types.Struct:
		vv := &typeinfo.StructType{}
		ctx.typemap[t] = vv
		vv.Fields = make([]typeinfo.StructField, t.NumFields())
		for i := 0; i < t.NumFields(); i++ {
			f := t.Field(i)
			vv.Fields[i].Name = f.Name()
			vv.Fields[i].Anonymous = f.Anonymous()
			vv.Fields[i].Type = TypeConv(ctx, f.Type())
			vv.Fields[i].Tag = reflect.StructTag(t.Tag(i))
		}
		return vv
	case *types.Map:
		vv := &typeinfo.MapType{}
		ctx.typemap[t] = vv
		vv.Key = TypeConv(ctx, t.Key())
		vv.Elem = TypeConv(ctx, t.Elem())
		return vv
	default:
		// Ignore other types
	}

	return typeinfo.InvalidType
}
