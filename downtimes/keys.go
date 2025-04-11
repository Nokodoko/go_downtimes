package downtimes

import (
	"fmt"
	"os"
)

func api(key_name string) (string, error) {
	err, ok := os.LookupEnv(key_name)
	switch {
	case !ok:
		return "", fmt.Errorf("%v, %s environment variable not set, or is set to a different name.", err, key_name)
	default:
		key := os.Getenv(key_name)
		return key, nil
	}
}

func app(key_name string) (string, error) {
	err, ok := os.LookupEnv(key_name)
	switch {
	case !ok:
		return "", fmt.Errorf("%v, %s environment variable not set, or is set to a different name.", key_name, err)
	default:
		key := os.Getenv(key_name)
		return key, nil
	}
}
