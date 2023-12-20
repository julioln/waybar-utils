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

		file_content, err := os.ReadFile("/sys/class/drm/card1/device/mem_info_vram_total")
		if err != nil {
			continue
		}
		memory_total, err := strconv.Atoi(strings.Fields(string(file_content))[0])
		if err != nil {
			continue
		}

		file_content, err = os.ReadFile("/sys/class/drm/card1/device/mem_info_vram_used")
		if err != nil {
			continue
		}
		memory_used, err := strconv.Atoi(strings.Fields(string(file_content))[0])
		if err != nil {
			continue
		}

		file_content, err = os.ReadFile("/sys/class/drm/card1/device/mem_info_gtt_total")
		if err != nil {
			continue
		}
		gtt_total, err := strconv.Atoi(strings.Fields(string(file_content))[0])
		if err != nil {
			continue
		}

		file_content, err = os.ReadFile("/sys/class/drm/card1/device/mem_info_gtt_used")
		if err != nil {
			continue
		}
		gtt_used, err := strconv.Atoi(strings.Fields(string(file_content))[0])
		if err != nil {
			continue
		}

		memory_percent := (float32(memory_used) / float32(memory_total)) * 100
		gtt_percent := (float32(gtt_used) / float32(gtt_total)) * 100

		fmt.Printf("%.0f%% +%.0f%%\n", memory_percent, gtt_percent)
	}
}
