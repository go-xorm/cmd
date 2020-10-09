package {{.Models}}

{{$ilen := len .Imports}}
import (
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
	return "xxxx"
}

// TableName 表名
func (*{{Mapper .Name}}) TableName() string {
	return "{{.Name}}"
}
func (t *{{Mapper .Name}}) GetId() uint64 {
	return t.Id
}
{{end}}


