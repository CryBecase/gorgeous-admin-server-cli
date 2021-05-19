package models

var tmplStr = `
// generated by gorgeous-admin-server-cli
package {{ .Models }}

{{ $importsNum := len .Imports }}
{{ $table := .Table }}
{{ $jsonTag := .LongestJsonTagLength }}

{{ if gt $importsNum 0 }}
import (
	{{ range .Imports }} "{{ . }}" {{ end }}
)
{{ end }}

// {{ Mapper $table.Name }} db model
type {{ Mapper $table.Name }} struct {
{{ range $table.ColumnsSeq }} {{ $col := $table.GetColumn . }} {{ Mapper $col.Name }} {{ Type $col }} {{ Tag $table $col $jsonTag }}
{{ end }}
}

func (t *{{ Mapper $table.Name }}) TableName() string {
	return "{{ $table.Name }}"
}

func (t *{{ Mapper $table.Name }}) TableComment() string {
	return "{{ $table.Comment }}"
}

`