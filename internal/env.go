package env

import (
	"os"
	"strconv"
)

func GetEnv(key string, callback_value any) any {
	value, exists := os.LookupEnv(key)

	if !exists {
		return callback_value
	}

	switch v := callback_value.(type) {
	case string:
		return value
	case int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return v // return default value if conversion fails
		}
		return intValue
	case bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return v // return default value if conversion fails
		}
		return boolValue
	default:
		return v // return default value for unsupported types
	}
}
