package main

import (
	"fmt"

	"github.com/Nitecon/ib"
)

func main() {
	ib := &ib.IB{BackupFile: "../TestMyStuffs.dat", CreateBackup: true}
	c := ib.Start("localhost:7497")
	c.CreateBackup = true
	for {
		select {
		case portfolio := <-c.PortFolio:
			fmt.Printf("System: %#v\n", portfolio)
		}
	}
}
