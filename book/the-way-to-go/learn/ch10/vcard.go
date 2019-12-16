package main

import (
	"fmt"
	"time"
)

type Address struct {
	addr string
}

type VCard struct {
	name     string
	address  map[string]*Address
	birthday time.Time
	photo    byte
}

func main() {
	vcard := new(VCard)
	fmt.Printf("name: %s \n", vcard.name)
	fmt.Printf("address: %v \n", vcard.address)
	fmt.Printf("birthday: %s \n", vcard.birthday)
	fmt.Printf("photo: %s \n", vcard.photo)
}
