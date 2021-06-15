package main

import (
	"fmt"
	"os"

	"github.com/charoleizer/system-agent/core"
	"github.com/charoleizer/system-agent/utils"
)

const (
	DISK   = "disk"
	MEMORY = "memory"
	CPU    = "cpu"
)

func main() {

	if utils.Contains(os.Args, DISK) {
		fmt.Println(string(core.GetDiskUsage()))
	}
}
