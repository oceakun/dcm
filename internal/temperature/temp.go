package temperature

import (
	"fmt"
	"log"
	"strings"

	"github.com/ssimunic/gosensors"
)


func GetTemperatures() map[string]float64 {
	temperatures := make(map[string]float64)
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		log.Printf("failed to initialize sensors: %v", err)
		return temperatures
	}

	for _, readings := range sensors.Chips {
		// Iterate over sensor readings for this chip
		for label, reading := range readings {
			// Clean up and format the label and reading
			cleanLabel := strings.TrimSpace(label)
			cleanReading := strings.TrimSpace(reading)
			var temp float64
			_, err := fmt.Sscanf(cleanReading, "%f", &temp)
			if err == nil {
				temperatures[cleanLabel] = temp
			}
		}
	}

	return temperatures
}