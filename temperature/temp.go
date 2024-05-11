package temperature

import (
	"fmt"
	"log"
	"strings"

	"github.com/ssimunic/gosensors"
)

func GetTemperatures() {
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		log.Fatalf("failed to initialize sensors: %v", err)
	}

	// Iterate over chips
	for chip, readings := range sensors.Chips {
		// Chip name
		fmt.Printf("Chip: %s\n", chip)

		// Iterate over sensor readings for this chip
		for label, reading := range readings {
			// Clean up and format the label and reading
			cleanLabel := strings.TrimSpace(label)
			cleanReading := strings.TrimSpace(reading)
			fmt.Printf("  %s: %s\n", cleanLabel, cleanReading)
		}
		fmt.Println()
	}
}
