package scope

import (
	"github.com/samber/lo"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
)

func GetFields(model interface{}) []string {
	s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		panic("failed to parse schema")
	}

	fields := make([]string, 0)
	for _, field := range s.Fields {
		if field.DBName != "" {
			fields = append(fields, field.DBName)
		}
	}
	return fields
}

func GetFieldStrings(model interface{}, except []string, prefix ...string) string {
	s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		panic("failed to parse schema")
	}

	fields := make([]string, 0)
	for _, field := range s.Fields {
		if !lo.Contains(except, field.DBName) {
			if field.DBName != "" {
				if len(prefix) == 0 {
					fields = append(fields, "`"+field.DBName+"`")
				} else if len(prefix) == 1 {
					fields = append(fields, prefix[0]+"`"+field.DBName+"`")
				} else if len(prefix) == 2 {
					fields = append(fields, prefix[0]+"`"+field.DBName+"` AS "+prefix[1]+field.DBName)
				}
			}
		}
	}
	return strings.Join(fields, ",")
}
