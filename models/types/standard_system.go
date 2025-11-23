package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// StandardSystem 表示运动等级标准体系
type StandardSystem int8

const (
	StandardSystemUnknown StandardSystem = 0
	StandardSystemPE      StandardSystem = 1 // 学生体测（Physical Education）
	StandardSystemChina   StandardSystem = 2 // 中国田协/国家体育总局标准
	StandardSystemWorld   StandardSystem = 3 // 国际田联（World Athletics）等国际标准
	StandardSystemSelf    StandardSystem = 4 // 自定义等级
)

func (s StandardSystem) String() string {
	switch s {
	case StandardSystemPE:
		return "学生体测"
	case StandardSystemChina:
		return "中国田协"
	case StandardSystemWorld:
		return "国际田联"
	case StandardSystemSelf:
		return "自定义等级"
	default:
		return "未知"
	}
}

func ParseStandardSystem(str string) (StandardSystem, error) {
	switch str {
	case "体测", "pe", "PE", "school":
		return StandardSystemPE, nil
	case "中国", "china", "national":
		return StandardSystemChina, nil
	case "国际", "world", "international", "iaaf":
		return StandardSystemWorld, nil
	case "自定义等级", "自定义", "self":
		return StandardSystemSelf, nil
	case "", "unknown":
		return StandardSystemUnknown, nil
	default:
		return StandardSystemUnknown, fmt.Errorf("invalid standard system: %s", str)
	}
}

func (s StandardSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *StandardSystem) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case float64:
		*s = StandardSystem(val)
	case string:
		parsed, err := ParseStandardSystem(val)
		if err != nil {
			return err
		}
		*s = parsed
	default:
		return fmt.Errorf("invalid standard system type: %T", v)
	}
	return nil
}

func (s StandardSystem) Value() (driver.Value, error) {
	return int8(s), nil
}

func (s *StandardSystem) Scan(value interface{}) error {
	if value == nil {
		*s = StandardSystemUnknown
		return nil
	}
	switch v := value.(type) {
	case int64:
		*s = StandardSystem(v)
	case int32:
		*s = StandardSystem(v)
	case int:
		*s = StandardSystem(v)
	case []byte:
		*s = StandardSystem(v[0])
	default:
		return fmt.Errorf("cannot scan %T into StandardSystem", value)
	}
	return nil
}

func (s StandardSystem) Valid() bool {
	return s > StandardSystemUnknown && s < StandardSystemSelf
}
