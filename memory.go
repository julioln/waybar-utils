package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)

		meminfo, err := os.ReadFile("/proc/meminfo")
		if err != nil {
			continue
		}

		var memtotal int
		var memavailable int
		var swaptotal int
		var swapfree int

		for _, mem_line := range strings.Split(string(meminfo), "\n") {
			if strings.Contains(mem_line, "MemTotal:") {
				memtotal, err = strconv.Atoi(strings.Fields(mem_line)[1])
			}
			if strings.Contains(mem_line, "MemAvailable:") {
				memavailable, err = strconv.Atoi(strings.Fields(mem_line)[1])
			}
			if strings.Contains(mem_line, "SwapTotal:") {
				swaptotal, err = strconv.Atoi(strings.Fields(mem_line)[1])
			}
			if strings.Contains(mem_line, "SwapFree:") {
				swapfree, err = strconv.Atoi(strings.Fields(mem_line)[1])
			}
		}

		mem_used := (float32((memtotal - memavailable)) / float32(memtotal)) * 100
		swap_used := (float32((swaptotal - swapfree)) / float32(swaptotal)) * 100

		fmt.Printf("%.0f%% +%.0f%%\n", mem_used, swap_used)
	}
}
