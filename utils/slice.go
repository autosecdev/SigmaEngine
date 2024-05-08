package utils

import "fmt"

// ToStringSlice converts a slice of interface{} to a slice of string.
func ToStringSlice(slice []interface{}) ([]string, error) {
	var result []string
	for _, v := range slice {
		str, ok := v.(string)
		if ok {
			result = append(result, str)
		} else {

			fmt.Printf("Warning: Non-string value '%v' of type %T skipped.\n", v, v)
		}
	}
	return result, nil
}
