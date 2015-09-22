package main

import (
	"fmt"
	govirt "libvirt-go"
)

func main() {
	virConn, err := govirt.NewVirConnectionReadOnly("qemu:///system")
	if err != nil {
		fmt.Println(err)
		return
	}

	dom, err := virConn.LookupDomainByName("instance-000000aa")
	if err != nil {
		fmt.Println(err)
		return
	}

	stat, err := dom.MemoryStats()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stat)
}
