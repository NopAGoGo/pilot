package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Product struct {
	SN        string `json:"编号" property:"编号,-"`
	WrapperSN int    `json:"包装编号" property:"包装编号,-"`
	Length    string `json:"-" property:"长度,1"`
	Width     string `json:"-" property:"宽度,2"`
	Thickness string `json:"-" property:"厚度,3"`
	Color     string `json:"颜色" property:"颜色,4"`
	Remark    string `json:"备注" property:"备注,-"`
	Num       int    `json:"数量" property:"数量,-"`
}

type WrapperProp struct {
	ID    int    `json:"-"`
	SN    int    `json:"sn"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	product := Product{
		"201906261254",
		1,
		"3000",
		"1000",
		"100",
		"白",
		"无",
		2,
	}

	fmt.Printf("%+v\n", splitProductProperty(product))
}

// 反射取tag
func splitProductProperty(product Product) (wrapperProps []WrapperProp) {
	reflectProduct := reflect.ValueOf(product)
	wrapperProp := WrapperProp{}
	wrapperProp.SN = product.WrapperSN
	for i := 0; i < reflectProduct.NumField(); i++ {
		sf := reflectProduct.Type().Field(i)
		if sf.PkgPath != "" && !sf.Anonymous {
			continue
		}
		tag := sf.Tag.Get("property")
		if tag == "-" {
			continue
		}
		key, opts := parseTag(tag)
		if opts == "-" {
			continue
		}
		index, _ := strconv.Atoi(opts)
		wrapperProp.ID = index
		wrapperProp.Key = key
		switch reflectProduct.Field(i).Type().String() {
		case "string":
			wrapperProp.Value = reflectProduct.Field(i).String()
		case "float64":
			wrapperProp.Value = strconv.FormatFloat(reflectProduct.Field(i).Interface().(float64), 'f', 6, 64)
		case "int":
			wrapperProp.Value = strconv.Itoa(reflectProduct.Field(i).Interface().(int))
		}
		wrapperProps = append(wrapperProps, wrapperProp)
	}
	return
}

func parseTag(tag string) (string, string) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tag[idx+1:]
	}
	return tag, ""
}
