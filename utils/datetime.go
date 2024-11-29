package utils

import (
	"database/sql/driver"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type DateTime time.Time

var TimeFormats = []string{"2006-01-02 15:04:05", "20060102150405"}

// Scan GORM Scanner 接口, 从数据库读取到类型
func (t *DateTime) Scan(value any) error {

	if v, ok := value.(time.Time); !ok {
		return errors.Errorf("failed to unmarshal CustomTime value: %v", value)
	} else {
		*t = DateTime(v)
		return nil
	}
}

// Value GORM Valuer 接口, 保存到数据库
func (t DateTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return time.Time(t), nil
}

// UnmarshalJSON JSON序列号
func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	fmt.Println(string(data))
	// 空值不进行解析
	if len(data) == 2 {
		*t = DateTime(time.Time{})
		return
	}
	var now time.Time
	for _, format := range TimeFormats {
		// 指定解析的格式
		if now, err = time.ParseInLocation(format, string(data), time.Local); err == nil {
			*t = DateTime(now)
			return
		}
		// 指定解析的格式
		if now, err = time.ParseInLocation(`"`+format+`"`, string(data), time.Local); err == nil {
			*t = DateTime(now)
			return
		}
	}

	return
}

// MarshalJSON JSON序列号
func (t DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormats[0])+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormats[0])
	b = append(b, '"')
	return b, nil
}
func (t DateTime) String() string {
	return time.Time(t).Format(TimeFormats[0])
}

func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}
