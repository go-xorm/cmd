package {{.Models}}

{{$ilen := len .Imports}}
import (
	"amiba.fun/amiba/go-component/orm"
	{{if gt $ilen 0}}
	{{range .Imports}}"{{.}}"{{end}}
{{end}}
)


{{range .Tables}}
type {{Mapper .Name}} struct {
orm.Model `xorm:"-"`
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}} 	{{Type $col}}  {{Tag $table $col }}
{{end}}

}
// DatabaseAlias 数据库别名
func (*{{Mapper .Name}}) DatabaseAlias() string {
	return "database" //TODO 请自行输入数据库别名!
}

// TableName 表名
func (*{{Mapper .Name}}) TableName() string {
	return "{{.Name}}"
}
{{end}}


