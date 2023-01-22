package main

func (ns25519 *Player) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{Name string; Health ./example.CustomUint8; Weapons []./example.Weapon; ./example.Position})(ns25519)

	// ZZ: (string)(ns25519.Name)
	ns25520 += 8 + uint64(len(ns25519.Name))

	// ZZ: (./example.CustomUint8)(ns25519.Health)
	ns25520 += ns25519.Health.ZZSizeGOBE()

	// ZZ: ([]./example.Weapon)(ns25519.Weapons)
	ns25520 += 8
	for ns25521 := 0; ns25521 < len(ns25519.Weapons); ns25521++ {

		// ZZ: (./example.Weapon)(ns25519.Weapons[ns25521])
		ns25520 += ns25519.Weapons[ns25521].SizeGOBE()
	}

	// ZZ: (./example.Position)(ns25519.Position)
	ns25520 += ns25519.Position.SizeGOBE()

	return ns25520
}

func (ns25522 *Player) MarshalGOBE(dst []byte) uint64 {
	var ns25523 uint64

	// ZZ: (struct{Name string; Health ./example.CustomUint8; Weapons []./example.Weapon; ./example.Position})(ns25522)

	// ZZ: (string)(ns25522.Name)
	var ns25524 uint64 = uint64(len(ns25522.Name))
	_ = dst[ns25523+7]
	dst[ns25523+0] = byte(ns25524 >> 0)
	dst[ns25523+1] = byte(ns25524 >> 8)
	dst[ns25523+2] = byte(ns25524 >> 16)
	dst[ns25523+3] = byte(ns25524 >> 24)
	dst[ns25523+4] = byte(ns25524 >> 32)
	dst[ns25523+5] = byte(ns25524 >> 40)
	dst[ns25523+6] = byte(ns25524 >> 48)
	dst[ns25523+7] = byte(ns25524 >> 56)
	copy(dst[ns25523+8:], ns25522.Name)
	ns25523 += 8 + ns25524

	// ZZ: (./example.CustomUint8)(ns25522.Health)
	ns25523 += ns25522.Health.ZZMarshalGOBE(dst[ns25523:])

	// ZZ: ([]./example.Weapon)(ns25522.Weapons)
	var ns25525 uint64 = uint64(len(ns25522.Weapons))
	_ = dst[ns25523+7]
	dst[ns25523+0] = byte(ns25525 >> 0)
	dst[ns25523+1] = byte(ns25525 >> 8)
	dst[ns25523+2] = byte(ns25525 >> 16)
	dst[ns25523+3] = byte(ns25525 >> 24)
	dst[ns25523+4] = byte(ns25525 >> 32)
	dst[ns25523+5] = byte(ns25525 >> 40)
	dst[ns25523+6] = byte(ns25525 >> 48)
	dst[ns25523+7] = byte(ns25525 >> 56)
	ns25523 += 8

	for ns25526 := 0; ns25526 < len(ns25522.Weapons); ns25526++ {

		// ZZ: (./example.Weapon)(ns25522.Weapons[ns25526])
		ns25523 += ns25522.Weapons[ns25526].MarshalGOBE(dst[ns25523:])
	}

	// ZZ: (./example.Position)(ns25522.Position)
	ns25523 += ns25522.Position.MarshalGOBE(dst[ns25523:])

	return ns25523
}

func (ns25527 *Player) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Name string; Health ./example.CustomUint8; Weapons []./example.Weapon; ./example.Position})(ns25527)

	// ZZ: (string)(ns25527.Name)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25528 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25528 {
		return
	}
	ns25527.Name = string(src[offset : offset+ns25528])
	offset += ns25528

	// ZZ: (./example.CustomUint8)(ns25527.Health)
	ns25529, ns25530 := ns25527.Health.ZZUnmarshalGOBE(src[offset:])
	offset += ns25529
	if !ns25530 {
		return
	}

	// ZZ: ([]./example.Weapon)(ns25527.Weapons)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25531 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(cap(ns25527.Weapons)) < ns25531 {
		if ns25531 <= 1<<15 {
			ns25527.Weapons = make([]Weapon, ns25531)
		} else {
			return
		}
	} else {
		ns25527.Weapons = ns25527.Weapons[:ns25531]
	}
	for ns25532 := uint64(0); ns25532 < ns25531; ns25532++ {

		// ZZ: (./example.Weapon)(ns25527.Weapons[ns25532])
		ns25533, ns25534 := ns25527.Weapons[ns25532].UnmarshalGOBE(src[offset:])
		offset += ns25533
		if !ns25534 {
			return
		}
	}

	// ZZ: (./example.Position)(ns25527.Position)
	ns25535, ns25536 := ns25527.Position.UnmarshalGOBE(src[offset:])
	offset += ns25535
	if !ns25536 {
		return
	}

	ok = true
	return
}

func (ns25537 *Position) SizeGOBE() uint64 {
	var ns25538 uint64

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25537)

	// ZZ: (uint64)(ns25537.X)
	ns25538 += 8

	// ZZ: (uint64)(ns25537.Y)
	ns25538 += 8

	// ZZ: (uint64)(ns25537.Z)
	ns25538 += 8

	return ns25538
}

func (ns25539 *Position) MarshalGOBE(dst []byte) uint64 {
	var ns25540 uint64

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25539)

	// ZZ: (uint64)(ns25539.X)
	_ = dst[ns25540+7]
	dst[ns25540+0] = byte(ns25539.X >> 0)
	dst[ns25540+1] = byte(ns25539.X >> 8)
	dst[ns25540+2] = byte(ns25539.X >> 16)
	dst[ns25540+3] = byte(ns25539.X >> 24)
	dst[ns25540+4] = byte(ns25539.X >> 32)
	dst[ns25540+5] = byte(ns25539.X >> 40)
	dst[ns25540+6] = byte(ns25539.X >> 48)
	dst[ns25540+7] = byte(ns25539.X >> 56)
	ns25540 += 8

	// ZZ: (uint64)(ns25539.Y)
	_ = dst[ns25540+7]
	dst[ns25540+0] = byte(ns25539.Y >> 0)
	dst[ns25540+1] = byte(ns25539.Y >> 8)
	dst[ns25540+2] = byte(ns25539.Y >> 16)
	dst[ns25540+3] = byte(ns25539.Y >> 24)
	dst[ns25540+4] = byte(ns25539.Y >> 32)
	dst[ns25540+5] = byte(ns25539.Y >> 40)
	dst[ns25540+6] = byte(ns25539.Y >> 48)
	dst[ns25540+7] = byte(ns25539.Y >> 56)
	ns25540 += 8

	// ZZ: (uint64)(ns25539.Z)
	_ = dst[ns25540+7]
	dst[ns25540+0] = byte(ns25539.Z >> 0)
	dst[ns25540+1] = byte(ns25539.Z >> 8)
	dst[ns25540+2] = byte(ns25539.Z >> 16)
	dst[ns25540+3] = byte(ns25539.Z >> 24)
	dst[ns25540+4] = byte(ns25539.Z >> 32)
	dst[ns25540+5] = byte(ns25539.Z >> 40)
	dst[ns25540+6] = byte(ns25539.Z >> 48)
	dst[ns25540+7] = byte(ns25539.Z >> 56)
	ns25540 += 8

	return ns25540
}

func (ns25541 *Position) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25541)

	// ZZ: (uint64)(ns25541.X)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25541.X = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25541.Y)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25541.Y = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25541.Z)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25541.Z = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25542 *Weapon) SizeGOBE() uint64 {
	var ns25543 uint64

	// ZZ: (struct{Name string; Damage uint16})(ns25542)

	// ZZ: (string)(ns25542.Name)
	ns25543 += 8 + uint64(len(ns25542.Name))

	// ZZ: (uint16)(ns25542.Damage)
	ns25543 += 2

	return ns25543
}

func (ns25544 *Weapon) MarshalGOBE(dst []byte) uint64 {
	var ns25545 uint64

	// ZZ: (struct{Name string; Damage uint16})(ns25544)

	// ZZ: (string)(ns25544.Name)
	var ns25546 uint64 = uint64(len(ns25544.Name))
	_ = dst[ns25545+7]
	dst[ns25545+0] = byte(ns25546 >> 0)
	dst[ns25545+1] = byte(ns25546 >> 8)
	dst[ns25545+2] = byte(ns25546 >> 16)
	dst[ns25545+3] = byte(ns25546 >> 24)
	dst[ns25545+4] = byte(ns25546 >> 32)
	dst[ns25545+5] = byte(ns25546 >> 40)
	dst[ns25545+6] = byte(ns25546 >> 48)
	dst[ns25545+7] = byte(ns25546 >> 56)
	copy(dst[ns25545+8:], ns25544.Name)
	ns25545 += 8 + ns25546

	// ZZ: (uint16)(ns25544.Damage)
	_ = dst[ns25545+1]
	dst[ns25545+0] = byte(ns25544.Damage >> 0)
	dst[ns25545+1] = byte(ns25544.Damage >> 8)
	ns25545 += 2

	return ns25545
}

func (ns25547 *Weapon) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Name string; Damage uint16})(ns25547)

	// ZZ: (string)(ns25547.Name)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25548 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25548 {
		return
	}
	ns25547.Name = string(src[offset : offset+ns25548])
	offset += ns25548

	// ZZ: (uint16)(ns25547.Damage)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25547.Damage = uint16(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}
