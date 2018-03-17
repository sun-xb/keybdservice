package main

import "github.com/sun-xb/keybdservice/btsdp"

func main() {
	sdp := btsdp.New()
	sdp.RegisterService()
}
