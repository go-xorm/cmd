// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	//"fmt"
	"strings"
	"text/template"

	"github.com/go-xorm/core"
)

var (
	ObjcTmpl LangTmpl = LangTmpl{
		template.FuncMap{"Mapper": mapper.Table2Obj,
			"Type":    objcTypeStr,
			"UnTitle": unTitle,
		},
		nil,
		genCPlusImports,
	}
)

func objcTypeStr(col *core.Column) string {
	tp := col.SQLType
	name := strings.ToUpper(tp.Name)
	switch name {
	case core.Bit, core.TinyInt, core.SmallInt, core.MediumInt, core.Int, core.Integer, core.Serial:
		return "int"
	case core.BigInt, core.BigSerial:
		return "long"
	case core.Char, core.Varchar, core.TinyText, core.Text, core.MediumText, core.LongText:
		return "NSString*"
	case core.Date, core.DateTime, core.Time, core.TimeStamp:
		return "NSString*"
	case core.Decimal, core.Numeric:
		return "NSString*"
	case core.Real, core.Float:
		return "float"
	case core.Double:
		return "double"
	case core.TinyBlob, core.Blob, core.MediumBlob, core.LongBlob, core.Bytea:
		return "NSString*"
	case core.Bool:
		return "BOOL"
	default:
		return "NSString*"
	}
	return ""
}

func genObjcImports(tables []*core.Table) map[string]string {
	imports := make(map[string]string)

	for _, table := range tables {
		for _, col := range table.Columns() {
			switch objcTypeStr(col) {
			case "time_t":
				imports[`<time.h>`] = `<time.h>`
			case "tstring":
				imports["<string>"] = "<string>"
				//case "__int64":
				//    imports[""] = ""
			}
		}
	}
	return imports
}
