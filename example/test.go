package main

import (
	"fmt"
	"os/exec"
)

func Sn() []byte {
	var sn []byte
	var err error
	if sn, err = exec.Command("CMD", "/C", "WMIC BIOS GET SERIALNUMBER").CombinedOutput(); err != nil {
		fmt.Println(err)
	}
	return sn
}

func main() {
	sn := Sn()
	fmt.Println(sn)
}
