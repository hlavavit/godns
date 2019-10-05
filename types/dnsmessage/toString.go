package dnsmessage

import "fmt"

func formatLabelValue(label interface{}, value interface{}) string {
	return fmt.Sprintf("%-21v : %v", label, value)
}
