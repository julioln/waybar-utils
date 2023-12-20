package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)

		gpu_busy, err := os.ReadFile("/sys/class/drm/card1/device/gpu_busy_percent")
		if err != nil {
			continue
		}

		fmt.Printf("%s%%\n", strings.Fields(string(gpu_busy))[0])
	}
}
