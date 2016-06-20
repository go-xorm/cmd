{{if not .IsAppend}}
package {{.Model}}
{{end}}

import (
	{{range .Imports}}"{{.}}"{{end}}
)

{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .Columns}}	{{Mapper .Name}}	{{Type .}}
{{end}}
}

{{end}}
