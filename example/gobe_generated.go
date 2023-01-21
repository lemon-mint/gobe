package main

func (ns25519 *Player) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{Name string; Health uint16; Weapons []./example.Weapon; ./example.Position})(ns25519)

	// ZZ: (string)(ns25519.Name)
	ns25520 += 8 + uint64(len(ns25519.Name))

	// ZZ: (uint16)(ns25519.Health)
	ns25520 += 2

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

	// ZZ: (struct{Name string; Health uint16; Weapons []./example.Weapon; ./example.Position})(ns25522)

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

	// ZZ: (uint16)(ns25522.Health)
	_ = dst[ns25523+1]
	dst[ns25523+0] = byte(ns25522.Health >> 0)
	dst[ns25523+1] = byte(ns25522.Health >> 8)
	ns25523 += 2

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

	// ZZ: (struct{Name string; Health uint16; Weapons []./example.Weapon; ./example.Position})(ns25527)

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

	// ZZ: (uint16)(ns25527.Health)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25527.Health = uint16(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	// ZZ: ([]./example.Weapon)(ns25527.Weapons)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25529 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(cap(ns25527.Weapons)) < ns25529 {
		if ns25529 <= 1<<15 {
			ns25527.Weapons = make([]Weapon, ns25529)
		} else {
			return
		}
	} else {
		ns25527.Weapons = ns25527.Weapons[:ns25529]
	}
	for ns25530 := uint64(0); ns25530 < ns25529; ns25530++ {

		// ZZ: (./example.Weapon)(ns25527.Weapons[ns25530])
		ns25531, ns25532 := ns25527.Weapons[ns25530].UnmarshalGOBE(src[offset:])
		offset += ns25531
		if !ns25532 {
			return
		}
	}

	// ZZ: (./example.Position)(ns25527.Position)
	ns25533, ns25534 := ns25527.Position.UnmarshalGOBE(src[offset:])
	offset += ns25533
	if !ns25534 {
		return
	}

	ok = true
	return
}

func (ns25535 *Position) SizeGOBE() uint64 {
	var ns25536 uint64

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25535)

	// ZZ: (uint64)(ns25535.X)
	ns25536 += 8

	// ZZ: (uint64)(ns25535.Y)
	ns25536 += 8

	// ZZ: (uint64)(ns25535.Z)
	ns25536 += 8

	return ns25536
}

func (ns25537 *Position) MarshalGOBE(dst []byte) uint64 {
	var ns25538 uint64

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25537)

	// ZZ: (uint64)(ns25537.X)
	_ = dst[ns25538+7]
	dst[ns25538+0] = byte(ns25537.X >> 0)
	dst[ns25538+1] = byte(ns25537.X >> 8)
	dst[ns25538+2] = byte(ns25537.X >> 16)
	dst[ns25538+3] = byte(ns25537.X >> 24)
	dst[ns25538+4] = byte(ns25537.X >> 32)
	dst[ns25538+5] = byte(ns25537.X >> 40)
	dst[ns25538+6] = byte(ns25537.X >> 48)
	dst[ns25538+7] = byte(ns25537.X >> 56)
	ns25538 += 8

	// ZZ: (uint64)(ns25537.Y)
	_ = dst[ns25538+7]
	dst[ns25538+0] = byte(ns25537.Y >> 0)
	dst[ns25538+1] = byte(ns25537.Y >> 8)
	dst[ns25538+2] = byte(ns25537.Y >> 16)
	dst[ns25538+3] = byte(ns25537.Y >> 24)
	dst[ns25538+4] = byte(ns25537.Y >> 32)
	dst[ns25538+5] = byte(ns25537.Y >> 40)
	dst[ns25538+6] = byte(ns25537.Y >> 48)
	dst[ns25538+7] = byte(ns25537.Y >> 56)
	ns25538 += 8

	// ZZ: (uint64)(ns25537.Z)
	_ = dst[ns25538+7]
	dst[ns25538+0] = byte(ns25537.Z >> 0)
	dst[ns25538+1] = byte(ns25537.Z >> 8)
	dst[ns25538+2] = byte(ns25537.Z >> 16)
	dst[ns25538+3] = byte(ns25537.Z >> 24)
	dst[ns25538+4] = byte(ns25537.Z >> 32)
	dst[ns25538+5] = byte(ns25537.Z >> 40)
	dst[ns25538+6] = byte(ns25537.Z >> 48)
	dst[ns25538+7] = byte(ns25537.Z >> 56)
	ns25538 += 8

	return ns25538
}

func (ns25539 *Position) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25539)

	// ZZ: (uint64)(ns25539.X)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25539.X = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25539.Y)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25539.Y = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25539.Z)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25539.Z = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25540 *Weapon) SizeGOBE() uint64 {
	var ns25541 uint64

	// ZZ: (struct{Name string; Damage uint16})(ns25540)

	// ZZ: (string)(ns25540.Name)
	ns25541 += 8 + uint64(len(ns25540.Name))

	// ZZ: (uint16)(ns25540.Damage)
	ns25541 += 2

	return ns25541
}

func (ns25542 *Weapon) MarshalGOBE(dst []byte) uint64 {
	var ns25543 uint64

	// ZZ: (struct{Name string; Damage uint16})(ns25542)

	// ZZ: (string)(ns25542.Name)
	var ns25544 uint64 = uint64(len(ns25542.Name))
	_ = dst[ns25543+7]
	dst[ns25543+0] = byte(ns25544 >> 0)
	dst[ns25543+1] = byte(ns25544 >> 8)
	dst[ns25543+2] = byte(ns25544 >> 16)
	dst[ns25543+3] = byte(ns25544 >> 24)
	dst[ns25543+4] = byte(ns25544 >> 32)
	dst[ns25543+5] = byte(ns25544 >> 40)
	dst[ns25543+6] = byte(ns25544 >> 48)
	dst[ns25543+7] = byte(ns25544 >> 56)
	copy(dst[ns25543+8:], ns25542.Name)
	ns25543 += 8 + ns25544

	// ZZ: (uint16)(ns25542.Damage)
	_ = dst[ns25543+1]
	dst[ns25543+0] = byte(ns25542.Damage >> 0)
	dst[ns25543+1] = byte(ns25542.Damage >> 8)
	ns25543 += 2

	return ns25543
}

func (ns25545 *Weapon) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Name string; Damage uint16})(ns25545)

	// ZZ: (string)(ns25545.Name)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25546 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25546 {
		return
	}
	ns25545.Name = string(src[offset : offset+ns25546])
	offset += ns25546

	// ZZ: (uint16)(ns25545.Damage)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25545.Damage = uint16(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}
