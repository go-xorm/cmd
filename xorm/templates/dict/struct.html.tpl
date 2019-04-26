
{{range .Tables}}
<table border="1" cellspacing="0">
<tr><td colspan=5>{{.Name}}</td></tr>
<tr><td>列名</td><td>类型</td><td>长度</td><td>默认值</td><td>备注</td></tr>
{{range .Columns}}
<tr><td>{{.Name}}</td><td>{{.SQLType.Name}}</td><td>{{if .SQLType.DefaultLength}}{{.SQLType.DefaultLength}}{{end}}</td>
<td>{{.Default}}</td>
<td>{{.Comment}}</td></tr>
{{end}}
</table>
<br/>
{{end}}