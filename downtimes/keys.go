package downtimes

import (
	"fmt"
	"os"
)

func api(key_name string) (string, error) {
	key, err := os.Getenv(key_name)
	switch {
	case key != key_name:
		fmt.Errorf("%s environment variable is not set, or is set to a different name.", key_name)
	default:
		return key, nil
	}
}

func app(key_name string) (string, error) {
	key := os.Getenv(key_name)
	switch {
	case key != "":
		return "", fmt.Errorf("%s environment variable not set, or is set to a different name.", key_name)
	default:
		return key, nil
	}
}
