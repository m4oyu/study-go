package main

import (
	"flag"
	"fmt"
)

var nFlag = flag.Int("nflag", 1234, "helpmessage")
var flagvar int

func main() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message")
	flag.Parse()
	fmt.Println("flagvar has value ", flagvar)
	fmt.Println("nFlag has value ", *nFlag)
}
