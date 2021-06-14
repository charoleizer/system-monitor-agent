package core

import (
	"encoding/json"
	"fmt"
	"syscall"

	"github.com/charoleizer/system-agent/utils"
)

type DiskStatus struct {
	Total             float64 `json:"total"`
	Used              float64 `json:"used"`
	Available         float64 `json:"available"`
	Used_Percent      float64 `json:"used-percent"`
	Available_Percent float64 `json:"available-percent"`
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.Total = float64(fs.Blocks) * float64(fs.Bsize)
	disk.Available = float64(fs.Bfree) * float64(fs.Bsize)
	disk.Used = disk.Total - disk.Available

	disk.Used_Percent = float64(disk.Used) * MAXPERCENT / float64(disk.Total)
	disk.Available_Percent = float64(disk.Available) * MAXPERCENT / float64(disk.Total)

	return
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB

	MAXPERCENT = 100
)

func GetDiskUsage() {
	disk := DiskUsage("/")

	disk.Total = utils.RoundTo(disk.Total/GB, 2)
	disk.Used = utils.RoundTo(disk.Used/GB, 2)
	disk.Available = utils.RoundTo(disk.Available/GB, 2)
	disk.Used_Percent = utils.RoundTo(disk.Used_Percent, 2)
	disk.Available_Percent = utils.RoundTo(disk.Available_Percent, 2)

	response, _ := json.Marshal(disk)
	fmt.Println(string(response))
}
