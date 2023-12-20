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

		thermal, err := os.ReadFile("/proc/acpi/ibm/thermal")
		if err != nil {
			continue
		}
		fan, err := os.ReadFile("/proc/acpi/ibm/fan")
		if err != nil {
			continue
		}

		temperature := strings.Fields(string(thermal))[1]
		fan_speed := "0"

		for _, fan_line := range strings.Split(string(fan), "\n") {
			if strings.Contains(fan_line, "speed") {
				fan_speed = strings.Fields(fan_line)[1]
			}
		}

		fan_c, err := strconv.Atoi(fan_speed)

		if fan_c == 0 {
			fmt.Printf("%s°C\n", temperature)
		} else {
			fmt.Printf("%s°C (%.1fK)\n", temperature, float32(fan_c)/1000)
		}
	}
}
