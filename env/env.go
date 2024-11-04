package env

import (
	"fmt"
	"os"
	"strconv"
)

// BasicType is a constraint that allows only basic types.
type BasicType interface {
	int | int8 | int32 | int64 | uint | uint8 | uint32 | uint64 | float32 | float64 | bool | string
}

// Get 获取环境变量
func Get[V BasicType](key string) V {
	return GetWithDefault(key, *new(V))
}

// GetWithDefault 获取环境变量，如果不存在则返回默认值
func GetWithDefault[V BasicType](key string, defaultValue V) V {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	var result V
	if err := convert(value, &result); err != nil {
		return defaultValue
	}
	return result
}

// convert 将字符串转换为指定类型
func convert[V BasicType](value string, result *V) error {
	switch any(*result).(type) {
	case int:
		parsedValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		*result = any(parsedValue).(V)
	case int8:
		parsedValue, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return err
		}
		*result = any(int8(parsedValue)).(V)
	case int32:
		parsedValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		*result = any(int32(parsedValue)).(V)
	case int64:
		parsedValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		*result = any(parsedValue).(V)
	case uint:
		parsedValue, err := strconv.ParseUint(value, 10, 0)
		if err != nil {
			return err
		}
		*result = any(uint(parsedValue)).(V)
	case uint8:
		parsedValue, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return err
		}
		*result = any(uint8(parsedValue)).(V)
	case uint32:
		parsedValue, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		*result = any(uint32(parsedValue)).(V)
	case uint64:
		parsedValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		*result = any(parsedValue).(V)
	case float32:
		parsedValue, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		*result = any(float32(parsedValue)).(V)
	case float64:
		parsedValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		*result = any(parsedValue).(V)
	case bool:
		parsedValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		*result = any(parsedValue).(V)
	case string:
		*result = any(value).(V)
	default:
		return fmt.Errorf("unsupported type")
	}
	return nil
}
