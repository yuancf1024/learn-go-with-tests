package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x) 
	// ValueOf返回一个新值，初始化为存储在接口i中的具体值。ValueOf(nil)返回零值。 
	field := val.Field(0)
	// Field返回结构v的第i个字段。如果v的类型不是struct或i不在范围内，它会panics。
	fn(field.String())
}