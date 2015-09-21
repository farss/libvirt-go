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

	istat, err := dom.InterfaceStats("tapa9948def-52")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", istat)

	dstat, err := dom.BlockStats("vda")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", dstat)

	stat, err := dom.MemoryStats(8, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", stat)

	vstat, err := dom.GetVcpus(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", vstat)
}
