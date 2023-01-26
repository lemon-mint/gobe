package main

import (
	ns25529 "unsafe"
)

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
	var ns25530 []byte = src[offset : offset+ns25528]
	ns25527.Name = *(*string)(ns25529.Pointer(&ns25530))
	offset += ns25528

	// ZZ: (./example.CustomUint8)(ns25527.Health)
	ns25531, ns25532 := ns25527.Health.ZZUnmarshalGOBE(src[offset:])
	offset += ns25531
	if !ns25532 {
		return
	}

	// ZZ: ([]./example.Weapon)(ns25527.Weapons)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25533 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(cap(ns25527.Weapons)) < ns25533 {
		if ns25533*uint64(18) <= 1<<15 {
			ns25527.Weapons = make([]Weapon, ns25533)
			for ns25534 := uint64(0); ns25534 < ns25533; ns25534++ {

				// ZZ: (./example.Weapon)(ns25527.Weapons[ns25534])
				ns25536, ns25537 := ns25527.Weapons[ns25534].UnmarshalGOBE(src[offset:])
				offset += ns25536
				if !ns25537 {
					return
				}
			}
		} else {
			// Slice too large, Using Append
			ns25527.Weapons = ns25527.Weapons[:0]
			for ns25534 := uint64(0); ns25534 < ns25533; ns25534++ {
				var ns25535 Weapon

				// ZZ: (./example.Weapon)(ns25535)
				ns25538, ns25539 := ns25535.UnmarshalGOBE(src[offset:])
				offset += ns25538
				if !ns25539 {
					return
				}
				ns25527.Weapons = append(ns25527.Weapons, ns25535)
			}
		}
	} else {
		ns25527.Weapons = ns25527.Weapons[:ns25533]
		for ns25534 := uint64(0); ns25534 < ns25533; ns25534++ {

			// ZZ: (./example.Weapon)(ns25527.Weapons[ns25534])
			ns25540, ns25541 := ns25527.Weapons[ns25534].UnmarshalGOBE(src[offset:])
			offset += ns25540
			if !ns25541 {
				return
			}
		}
	}

	// ZZ: (./example.Position)(ns25527.Position)
	ns25542, ns25543 := ns25527.Position.UnmarshalGOBE(src[offset:])
	offset += ns25542
	if !ns25543 {
		return
	}

	ok = true
	return
}

func (ns25544 *Position) SizeGOBE() uint64 {
	var ns25545 uint64

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25544)

	// ZZ: (uint64)(ns25544.X)
	ns25545 += 8

	// ZZ: (uint64)(ns25544.Y)
	ns25545 += 8

	// ZZ: (uint64)(ns25544.Z)
	ns25545 += 8

	return ns25545
}

func (ns25546 *Position) MarshalGOBE(dst []byte) uint64 {
	var ns25547 uint64

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25546)

	// ZZ: (uint64)(ns25546.X)
	_ = dst[ns25547+7]
	dst[ns25547+0] = byte(ns25546.X >> 0)
	dst[ns25547+1] = byte(ns25546.X >> 8)
	dst[ns25547+2] = byte(ns25546.X >> 16)
	dst[ns25547+3] = byte(ns25546.X >> 24)
	dst[ns25547+4] = byte(ns25546.X >> 32)
	dst[ns25547+5] = byte(ns25546.X >> 40)
	dst[ns25547+6] = byte(ns25546.X >> 48)
	dst[ns25547+7] = byte(ns25546.X >> 56)
	ns25547 += 8

	// ZZ: (uint64)(ns25546.Y)
	_ = dst[ns25547+7]
	dst[ns25547+0] = byte(ns25546.Y >> 0)
	dst[ns25547+1] = byte(ns25546.Y >> 8)
	dst[ns25547+2] = byte(ns25546.Y >> 16)
	dst[ns25547+3] = byte(ns25546.Y >> 24)
	dst[ns25547+4] = byte(ns25546.Y >> 32)
	dst[ns25547+5] = byte(ns25546.Y >> 40)
	dst[ns25547+6] = byte(ns25546.Y >> 48)
	dst[ns25547+7] = byte(ns25546.Y >> 56)
	ns25547 += 8

	// ZZ: (uint64)(ns25546.Z)
	_ = dst[ns25547+7]
	dst[ns25547+0] = byte(ns25546.Z >> 0)
	dst[ns25547+1] = byte(ns25546.Z >> 8)
	dst[ns25547+2] = byte(ns25546.Z >> 16)
	dst[ns25547+3] = byte(ns25546.Z >> 24)
	dst[ns25547+4] = byte(ns25546.Z >> 32)
	dst[ns25547+5] = byte(ns25546.Z >> 40)
	dst[ns25547+6] = byte(ns25546.Z >> 48)
	dst[ns25547+7] = byte(ns25546.Z >> 56)
	ns25547 += 8

	return ns25547
}

func (ns25548 *Position) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{X uint64; Y uint64; Z uint64})(ns25548)

	// ZZ: (uint64)(ns25548.X)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25548.X = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25548.Y)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25548.Y = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25548.Z)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25548.Z = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25549 *Weapon) SizeGOBE() uint64 {
	var ns25550 uint64

	// ZZ: (struct{Name string; Damage uint16})(ns25549)

	// ZZ: (string)(ns25549.Name)
	ns25550 += 8 + uint64(len(ns25549.Name))

	// ZZ: (uint16)(ns25549.Damage)
	ns25550 += 2

	return ns25550
}

func (ns25551 *Weapon) MarshalGOBE(dst []byte) uint64 {
	var ns25552 uint64

	// ZZ: (struct{Name string; Damage uint16})(ns25551)

	// ZZ: (string)(ns25551.Name)
	var ns25553 uint64 = uint64(len(ns25551.Name))
	_ = dst[ns25552+7]
	dst[ns25552+0] = byte(ns25553 >> 0)
	dst[ns25552+1] = byte(ns25553 >> 8)
	dst[ns25552+2] = byte(ns25553 >> 16)
	dst[ns25552+3] = byte(ns25553 >> 24)
	dst[ns25552+4] = byte(ns25553 >> 32)
	dst[ns25552+5] = byte(ns25553 >> 40)
	dst[ns25552+6] = byte(ns25553 >> 48)
	dst[ns25552+7] = byte(ns25553 >> 56)
	copy(dst[ns25552+8:], ns25551.Name)
	ns25552 += 8 + ns25553

	// ZZ: (uint16)(ns25551.Damage)
	_ = dst[ns25552+1]
	dst[ns25552+0] = byte(ns25551.Damage >> 0)
	dst[ns25552+1] = byte(ns25551.Damage >> 8)
	ns25552 += 2

	return ns25552
}

func (ns25554 *Weapon) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Name string; Damage uint16})(ns25554)

	// ZZ: (string)(ns25554.Name)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25555 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25555 {
		return
	}
	var ns25556 []byte = src[offset : offset+ns25555]
	ns25554.Name = *(*string)(ns25529.Pointer(&ns25556))
	offset += ns25555

	// ZZ: (uint16)(ns25554.Damage)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25554.Damage = uint16(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}
