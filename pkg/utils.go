package utils

import "fmt"

func MapToKeysAndValues(table map[string]string) ([]string, []string) {
	var keys []string
	var values []string

	for key, value := range table {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

func SurroundWithQuotes(value string) string {
	return fmt.Sprintf("%q", value)
}
