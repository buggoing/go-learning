package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Mon
	Tue
	Wen
	Thu
	Fri
	Sat
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBd
	FlagLp
	FlagPP
	FlagMc
)

const (
	iB = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	fmt.Printf("%05b\n", FlagUp)
	fmt.Printf("%05b\n", FlagBd)
	fmt.Printf("%05b\n", FlagLp)
	fmt.Printf("%05b\n", FlagPP)
	fmt.Printf("%05b\n", FlagMc)

	st := fmt.Sprintf("%05b", FlagUp)
	// fmt.Printf("%T", YiB)
	fmt.Println(st)
}
