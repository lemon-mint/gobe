package build

import "github.com/lemon-mint/gobe/typeinfo"

type HasLoopContext struct {
	LoopedTypes map[typeinfo.Type]bool
}

func HasLoop(t typeinfo.Type, ctx *HasLoopContext) bool {
	if ctx == nil {
		ctx = &HasLoopContext{LoopedTypes: map[typeinfo.Type]bool{}}
	}
	if ctx.LoopedTypes[t] {
		return true
	}
	ctx.LoopedTypes[t] = true
	switch t := t.(type) {
	case *typeinfo.StructType:
		for _, f := range t.Fields {
			if HasLoop(f.Type, ctx) {
				return true
			}
		}
	case *typeinfo.PointerType:
		return HasLoop(t.Elem, ctx)
	case *typeinfo.SliceType:
		return HasLoop(t.Elem, ctx)
	case *typeinfo.ArrayType:
		return HasLoop(t.Elem, ctx)
	case *typeinfo.MapType:
		return HasLoop(t.Key, ctx) || HasLoop(t.Elem, ctx)
	}
	return false
}
