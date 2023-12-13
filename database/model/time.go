package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"strings"
	"time"
)

type Carbon struct {
	carbon.Carbon
}

func (t Carbon) Value() (driver.Value, error) {
	if !t.IsValid() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Carbon) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Carbon{carbon.CreateFromStdTime(value)}
		return nil
	}
	carbon.CreateFromStdTime(value)
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t Carbon) MarshalJSON() ([]byte, error) {
	if !t.IsValid() {
		return json.Marshal(nil)
	}
	return []byte(t.ToDateTimeString()), nil
}

func (t *Carbon) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	*t = Carbon{carbon.Parse(timeStr)}
	return nil
}
