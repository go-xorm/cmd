package main

import (
	"github.com/go-xorm/core"
	"strings"
)

var Custom custom

type custom struct{}

//将数据表字段变成go model的字段名
func (*custom) Table2Obj(name string) string {
	//调用包内的某个方法,改不了
	name = core.SnakeMapper{}.Table2Obj(name)
	return strings.Replace(name, "Id", "ID", -1)
}

//将数据表字段变成ID的字段名
func (*custom) IdToID(name string) string {
	//调用包内的某个方法,改不了
	return strings.Replace(name, "Id", "ID", -1)
}
