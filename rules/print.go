package rules

import (
	"fmt"

	"github.com/fatih/color"
)

func printReport(report map[string][]string) {
	fmt.Print("\n")

	if len(report) == 0 {
		color.Green("✔ No Border Patrol Violations")
	} else {
		fmt.Print("Border Violations:\n")

		for module, violations := range report {
			color.Red("\n" + module + ":")
			for _, violation := range violations {
				color.Red("  ✘ " + violation)
			}
		}
	}

	fmt.Print("\n")
}
