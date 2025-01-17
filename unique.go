package main

import "unique"

type addrDetail struct {
	isV6   bool
	zoneV6 string
}

func IsUnique() {
	h1 := unique.Make(addrDetail{isV6: true, zoneV6: "2001:0db8:0001:0000:0000:0ab9:C0A8:0102"})

	// This addrDetail won't be allocated as it already exists in the underlying map
	h2 := unique.Make(addrDetail{isV6: true, zoneV6: "2001:0db8:0001:0000:0000:0ab9:C0A8:0102"})

	if h1 == h2 {
		println("addresses are equal")
	}

	// Value() returns a copy of the underlying value (i.e., different memory address)
	println(h1.Value().zoneV6)
}
