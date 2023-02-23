package taggedunion

func (ns25519 *A) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{Val uint8})(ns25519)

	// ZZ: (uint8)(ns25519.Val)
	ns25520 += 1

	return ns25520
}

func (ns25521 *A) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (struct{Val uint8})(ns25521)

	// ZZ: (uint8)(ns25521.Val)
	dst[ns25522+0] = byte(ns25521.Val >> 0)
	ns25522++

	return ns25522
}

func (ns25523 *A) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Val uint8})(ns25523)

	// ZZ: (uint8)(ns25523.Val)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	ns25523.Val = uint8(
		uint8(src[offset+0]) << 0)
	offset += 1

	ok = true
	return
}

func (ns25524 *B) SizeGOBE() uint64 {
	var ns25525 uint64

	// ZZ: (struct{Val uint16})(ns25524)

	// ZZ: (uint16)(ns25524.Val)
	ns25525 += 2

	return ns25525
}

func (ns25526 *B) MarshalGOBE(dst []byte) uint64 {
	var ns25527 uint64

	// ZZ: (struct{Val uint16})(ns25526)

	// ZZ: (uint16)(ns25526.Val)
	_ = dst[ns25527+1]
	dst[ns25527+0] = byte(ns25526.Val >> 0)
	dst[ns25527+1] = byte(ns25526.Val >> 8)
	ns25527 += 2

	return ns25527
}

func (ns25528 *B) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Val uint16})(ns25528)

	// ZZ: (uint16)(ns25528.Val)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25528.Val = uint16(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}

func (ns25529 *C) SizeGOBE() uint64 {
	var ns25530 uint64

	// ZZ: (struct{Val uint32})(ns25529)

	// ZZ: (uint32)(ns25529.Val)
	ns25530 += 4

	return ns25530
}

func (ns25531 *C) MarshalGOBE(dst []byte) uint64 {
	var ns25532 uint64

	// ZZ: (struct{Val uint32})(ns25531)

	// ZZ: (uint32)(ns25531.Val)
	_ = dst[ns25532+3]
	dst[ns25532+0] = byte(ns25531.Val >> 0)
	dst[ns25532+1] = byte(ns25531.Val >> 8)
	dst[ns25532+2] = byte(ns25531.Val >> 16)
	dst[ns25532+3] = byte(ns25531.Val >> 24)
	ns25532 += 4

	return ns25532
}

func (ns25533 *C) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Val uint32})(ns25533)

	// ZZ: (uint32)(ns25533.Val)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	ns25533.Val = uint32(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	ok = true
	return
}

func (ns25534 *MyEnum) SizeGOBE() uint64 {
	var ns25535 uint64

	// ZZ: (struct{Type ./example/tagged_union.Type; A *./example/tagged_union.A "gobe_enum:\"Type=AType\""; B *./example/tagged_union.B "gobe_enum:\"Type=BType\""; C *./example/tagged_union.C "gobe_enum:\"Type=CType\""})(ns25534)

	// ZZ: (./example/tagged_union.Type)(ns25534.Type)

	// ZZ: (uint16)(ns25534.Type)
	ns25535 += 2
	// ENUM: KEY=<Type>
	switch ns25534.Type {
	case AType: // Type == AType

		// ZZ: (*./example/tagged_union.A)(ns25534.A)
		ns25535 += 1
		if ns25534.A != nil {

			// ZZ: (./example/tagged_union.A)((*ns25534.A))
			ns25535 += (*ns25534.A).SizeGOBE()
		}
	case BType: // Type == BType

		// ZZ: (*./example/tagged_union.B)(ns25534.B)
		ns25535 += 1
		if ns25534.B != nil {

			// ZZ: (./example/tagged_union.B)((*ns25534.B))
			ns25535 += (*ns25534.B).SizeGOBE()
		}
	case CType: // Type == CType

		// ZZ: (*./example/tagged_union.C)(ns25534.C)
		ns25535 += 1
		if ns25534.C != nil {

			// ZZ: (./example/tagged_union.C)((*ns25534.C))
			ns25535 += (*ns25534.C).SizeGOBE()
		}
	}

	return ns25535
}

func (ns25536 *MyEnum) MarshalGOBE(dst []byte) uint64 {
	var ns25537 uint64

	// ZZ: (struct{Type ./example/tagged_union.Type; A *./example/tagged_union.A "gobe_enum:\"Type=AType\""; B *./example/tagged_union.B "gobe_enum:\"Type=BType\""; C *./example/tagged_union.C "gobe_enum:\"Type=CType\""})(ns25536)

	// ZZ: (./example/tagged_union.Type)(ns25536.Type)

	// ZZ: (uint16)(ns25536.Type)
	_ = dst[ns25537+1]
	dst[ns25537+0] = byte(ns25536.Type >> 0)
	dst[ns25537+1] = byte(ns25536.Type >> 8)
	ns25537 += 2
	// ENUM: KEY=<Type>
	switch ns25536.Type {
	case AType: // Type == AType

		// ZZ: (*./example/tagged_union.A)(ns25536.A)
		if ns25536.A != nil {
			dst[ns25537] = 1
			ns25537++

			// ZZ: (./example/tagged_union.A)((*ns25536.A))
			ns25537 += (*ns25536.A).MarshalGOBE(dst[ns25537:])
		} else {
			dst[ns25537] = 0
			ns25537++
		}

	case BType: // Type == BType

		// ZZ: (*./example/tagged_union.B)(ns25536.B)
		if ns25536.B != nil {
			dst[ns25537] = 1
			ns25537++

			// ZZ: (./example/tagged_union.B)((*ns25536.B))
			ns25537 += (*ns25536.B).MarshalGOBE(dst[ns25537:])
		} else {
			dst[ns25537] = 0
			ns25537++
		}

	case CType: // Type == CType

		// ZZ: (*./example/tagged_union.C)(ns25536.C)
		if ns25536.C != nil {
			dst[ns25537] = 1
			ns25537++

			// ZZ: (./example/tagged_union.C)((*ns25536.C))
			ns25537 += (*ns25536.C).MarshalGOBE(dst[ns25537:])
		} else {
			dst[ns25537] = 0
			ns25537++
		}

	}

	return ns25537
}

func (ns25538 *MyEnum) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Type ./example/tagged_union.Type; A *./example/tagged_union.A "gobe_enum:\"Type=AType\""; B *./example/tagged_union.B "gobe_enum:\"Type=BType\""; C *./example/tagged_union.C "gobe_enum:\"Type=CType\""})(ns25538)

	// ZZ: (./example/tagged_union.Type)(ns25538.Type)

	// ZZ: (uint16)(ns25538.Type)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25538.Type = Type(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2
	// ENUM: KEY=<Type>
	switch ns25538.Type {
	case AType: // Type == AType

		// ZZ: (*./example/tagged_union.A)(ns25538.A)
		if uint64(len(src)) < offset+1 {
			return
		}
		if src[offset] == 1 {
			offset++
			if ns25538.A == nil {
				ns25538.A = new(A)
			}

			// ZZ: (./example/tagged_union.A)((*ns25538.A))
			ns25539, ns25540 := (*ns25538.A).UnmarshalGOBE(src[offset:])
			offset += ns25539
			if !ns25540 {
				return
			}
		} else {
			offset++
			ns25538.A = nil
		}
	case BType: // Type == BType

		// ZZ: (*./example/tagged_union.B)(ns25538.B)
		if uint64(len(src)) < offset+1 {
			return
		}
		if src[offset] == 1 {
			offset++
			if ns25538.B == nil {
				ns25538.B = new(B)
			}

			// ZZ: (./example/tagged_union.B)((*ns25538.B))
			ns25541, ns25542 := (*ns25538.B).UnmarshalGOBE(src[offset:])
			offset += ns25541
			if !ns25542 {
				return
			}
		} else {
			offset++
			ns25538.B = nil
		}
	case CType: // Type == CType

		// ZZ: (*./example/tagged_union.C)(ns25538.C)
		if uint64(len(src)) < offset+1 {
			return
		}
		if src[offset] == 1 {
			offset++
			if ns25538.C == nil {
				ns25538.C = new(C)
			}

			// ZZ: (./example/tagged_union.C)((*ns25538.C))
			ns25543, ns25544 := (*ns25538.C).UnmarshalGOBE(src[offset:])
			offset += ns25543
			if !ns25544 {
				return
			}
		} else {
			offset++
			ns25538.C = nil
		}
	}

	ok = true
	return
}

func (ns25545 *Type) SizeGOBE() uint64 {
	var ns25546 uint64

	// ZZ: (./example/tagged_union.Type)(*ns25545)

	// ZZ: (uint16)(*ns25545)
	ns25546 += 2

	return ns25546
}

func (ns25547 *Type) MarshalGOBE(dst []byte) uint64 {
	var ns25548 uint64

	// ZZ: (./example/tagged_union.Type)(*ns25547)

	// ZZ: (uint16)(*ns25547)
	_ = dst[ns25548+1]
	dst[ns25548+0] = byte(*ns25547 >> 0)
	dst[ns25548+1] = byte(*ns25547 >> 8)
	ns25548 += 2

	return ns25548
}

func (ns25549 *Type) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (./example/tagged_union.Type)(*ns25549)

	// ZZ: (uint16)(*ns25549)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	*ns25549 = Type(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}
