package field

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func SetField(cfg interface{}, key, value string) error {
	//fmt.Printf("key:%s, value:%s\n", key, value)

	stype := reflect.ValueOf(cfg)
	fs := strings.Split(key, ".")
	for i := 0; i < len(fs); i++ {
		//stype = stype.FieldByName(fs[i])
		member := fs[i]
		r := regexp.MustCompile(`^([a-zA-Z]+)\[(\d+)\]$`)
		matches := r.FindStringSubmatch(member)

		if len(matches) == 3 {
			name := matches[1]
			index, _ := strconv.Atoi(matches[2])
			stype = reflect.Indirect(stype).FieldByName(name)
			if !stype.IsValid() {
				return fmt.Errorf("Field [%s] is not invalid\n", name)
			} else {
				//fmt.Printf("stype: %v, kind: %d\n", stype, stype.Kind())
				if stype.Kind() == reflect.Array || stype.Kind() == reflect.Slice {
					if index < 0 || index >= stype.Len() {
						return fmt.Errorf("Field [%s] index [%d] exceed!\n", name, index)
					} else {
						stype = stype.Index(index)
					}
				}
			}
		} else {
			stype = reflect.Indirect(stype).FieldByName(member)
			if !stype.IsValid() {
				return fmt.Errorf("Field [%s] not found\n", member)
			}
		}
	}

	if stype.CanSet() {
		switch stype.Kind() {
		case reflect.String:
			stype.SetString(value)
		case reflect.Int:
			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("Field Int set failed: %s\n", value)
			} else {
				stype.SetInt(num)
			}
		case reflect.Bool:
			flag, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("Field Bool set failed: %s\n", value)
			} else {
				stype.SetBool(flag)
			}
		}
		//fmt.Printf("stype: %v\n", stype)
		//fmt.Printf("cfg: %v\n", cfg)
	} else {
		return fmt.Errorf("Field [%s] can not set\n", key)
	}

	return nil
}

func GetField(stype reflect.Type) {
	for i := 0; i < stype.NumField(); i++ {
		fieldType := stype.Field(i)
		fmt.Printf("属性名: %v, 字段是否可导出: %v, tag: %v, struct中第 %v 位\n",
			fieldType.Name, fieldType.IsExported(), fieldType.Tag, fieldType.Index)
		if fieldType.Type.Kind() == reflect.Struct {
			GetField(fieldType.Type)
		} else {
			fmt.Printf("属性名: %v, 字段是否可导出: %v, tag: %v, struct中第 %v 位\n",
				fieldType.Name, fieldType.IsExported(), fieldType.Tag, fieldType.Index)
		}
	}
}
