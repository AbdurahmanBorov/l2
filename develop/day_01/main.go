package main

import (
	"day_01/ntp"
)

func main() {
	if err := ntp.GetTime(); err != nil {
		panic(err)
	}
}
