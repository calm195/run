package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Gender 表示性别
type Gender int8

const (
	GenderUnknown Gender = 0
	GenderMale    Gender = 1
	GenderFemale  Gender = 2
)

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "男"
	case GenderFemale:
		return "女"
	default:
		return "未知"
	}
}

func ParseGender(s string) (Gender, error) {
	switch s {
	case "男", "male", "Male", "M":
		return GenderMale, nil
	case "女", "female", "Female", "F":
		return GenderFemale, nil
	case "未知", "unknown", "":
		return GenderUnknown, nil
	default:
		return GenderUnknown, fmt.Errorf("invalid gender: %s", s)
	}
}

// MarshalJSON 输出为可读字符串
func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(int8(g))
}

// UnmarshalJSON 支持从数字或字符串反序列化
func (g *Gender) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case float64: // JSON 数字解析为 float64
		*g = Gender(val)
	case string:
		parsed, err := ParseGender(val)
		if err != nil {
			return err
		}
		*g = parsed
	default:
		return fmt.Errorf("invalid gender type: %T", v)
	}
	return nil
}

func (g Gender) Value() (driver.Value, error) {
	return int8(g), nil
}

func (g *Gender) Scan(value interface{}) error {
	if value == nil {
		*g = GenderUnknown
		return nil
	}
	if v, ok := value.(int64); ok {
		*g = Gender(v)
	} else if v, ok := value.(int32); ok {
		*g = Gender(v)
	} else {
		return fmt.Errorf("cannot scan %T into Gender", value)
	}
	return nil
}

func (g Gender) Valid() bool {
	return g == GenderMale || g == GenderFemale
}
