package main

type A struct { a int }

type B struct { a int }

type C B

type D = B

func main() {
	var ( a A; b B; c C; d D )
	a = b	// nope
	b = c	// also nope
	d = b	// new in Go 1.9!
}
