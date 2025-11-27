package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Level 表示运动等级
type Level int8

const (
	LevelNone               Level = 0 // 未达标
	LevelInternationalElite Level = 1 // 国际级运动健将
	LevelElite              Level = 2 // 运动健将
	LevelFirst              Level = 3 // 一级运动员
	LevelSecond             Level = 4 // 二级运动员
	LevelThird              Level = 5 // 三级运动员
	LevelParticipate        Level = 6 // 参与级
)

func (l Level) String() string {
	switch l {
	case LevelNone:
		return "未达标"
	case LevelInternationalElite:
		return "国际级运动健将"
	case LevelElite:
		return "运动健将"
	case LevelFirst:
		return "一级运动员"
	case LevelSecond:
		return "二级运动员"
	case LevelThird:
		return "三级运动员"
	case LevelParticipate:
		return "参与级"
	default:
		return "未知等级"
	}
}

func ParseLevel(s string) (Level, error) {
	switch s {
	case "健将", "elite":
		return LevelElite, nil
	case "一级", "first":
		return LevelFirst, nil
	case "二级", "second":
		return LevelSecond, nil
	case "三级", "third":
		return LevelThird, nil
	case "参与级", "participate":
		return LevelParticipate, nil
	case "未达标", "none", "":
		return LevelNone, nil
	default:
		return LevelNone, fmt.Errorf("invalid level: %s", s)
	}
}

func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(int8(l))
}

func (l *Level) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case float64:
		*l = Level(val)
	case string:
		parsed, err := ParseLevel(val)
		if err != nil {
			return err
		}
		*l = parsed
	default:
		return fmt.Errorf("invalid level type: %T", v)
	}
	return nil
}

func (l Level) Value() (driver.Value, error) {
	return int8(l), nil
}

func (l *Level) Scan(value interface{}) error {
	if value == nil {
		*l = LevelNone
		return nil
	}
	if v, ok := value.(int64); ok {
		*l = Level(v)
	} else if v, ok := value.(int32); ok {
		*l = Level(v)
	} else {
		return fmt.Errorf("cannot scan %T into Level", value)
	}
	return nil
}

func (l Level) Valid() bool {
	return l >= LevelNone && l <= LevelParticipate
}
