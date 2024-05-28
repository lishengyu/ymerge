package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

func readYAMLFile(filename string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	data, err := readYAMLFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading file1.yaml: %v", err)
	}

	iter := reflect.ValueOf(data).MapRange()
	for iter.Next() {
		fmt.Printf("key:%v\n\tvalue:%v\n", iter.Key(), iter.Value())
	}
	/*
		for i := 0; i < sType.NumField(); i++ {
			fieldType := sType.Field(i)
			fmt.Printf("属性名: %v, 字段是否可导出: %v, tag: %v, struct中第 %v 位\n", fieldType.Name, fieldType.IsExported(), fieldType.Tag, fieldType.Index)
		}
	*/

	/*
		for i, v := range data {
			fmt.Printf("index:%s\n", i)
			fmt.Printf("\t%v\n", v)
		}
	*/

	return
}
