package main

import (
	"fmt"

	"github.com/charoleizer/system-agent/core"
)

func main() {
	fmt.Println(string(core.GetDiskUsage()))
}
