package xorm

import (
	"fmt"
	"sort"
	"strings"

	"github.com/xormplus/xorm/schemas"
)

type TagGenerator struct {
	TypeTransformRule    map[string]string // 类型转换规则 例如：int => int64
	FieldTransformRule   map[string]string // 字段类型转换规则（强制转换） 例如：created_at => time.Time
	ColHandlerChain      []ColumnHandler   // 字段处理
	LongestJsonTagLength int               // 最长的 json tag 的长度，用来美化 tag 间距
}

func GenerateTag(table *schemas.Table, col *schemas.Column, options ...func(*TagGenerator)) string {
	generator := defaultTagGenerator()
	for _, f := range options {
		f(generator)
	}

	col.TableName = table.Name
	table.Comment = strings.ReplaceAll(table.Comment, "\n", " ")
	col.Comment = strings.ReplaceAll(col.Comment, "\n", " ")

	var xormTag []string
	for _, handler := range generator.ColHandlerChain {
		xormTag = handler(table, col, xormTag)
	}

	jsonTagStr := []byte(fmt.Sprintf(`json:"%s"`, col.Name))
	xormTagStr := fmt.Sprintf(`xorm:"%s"`, strings.Join(xormTag, " "))

	for i := generator.LongestJsonTagLength - len(col.Name); generator.LongestJsonTagLength > 0 && i > 0; i-- {
		jsonTagStr = append(jsonTagStr, ' ')
	}

	tags := []string{string(jsonTagStr)}
	if len(xormTag) > 0 {
		tags = append(tags, xormTagStr)
	}

	return fmt.Sprintf("`%s`", strings.Join(tags, " "))
}

func defaultTagGenerator() *TagGenerator {
	typeTransformRule := make(map[string]string)
	typeTransformRule["int"] = "int64"

	fieldTransformRule := make(map[string]string)
	fieldTransformRule["created_at"] = "time.Time"
	fieldTransformRule["updated_at"] = "time.Time"
	fieldTransformRule["deleted_at"] = "time.Time"

	chain := make([]ColumnHandler, 0, 9)
	chain = append(chain, handleAutoTime, handleComment, handleType, handleIndex, handleAutoincr, handleDefault, handlePK, handleNullable)

	return &TagGenerator{
		TypeTransformRule:    typeTransformRule,
		FieldTransformRule:   fieldTransformRule,
		ColHandlerChain:      chain,
		LongestJsonTagLength: 0,
	}
}

type ColumnHandler func(*schemas.Table, *schemas.Column, []string) []string

var (
	handleAutoTime ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		switch col.Name {
		case "created_at":
			res = append(res, "created")
		case "updated_at":
			res = append(res, "updated")
		case "deleted_at":
			res = append(res, "deleted")
		}
		return res
	}
	handleComment ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		if col.Comment != "" {
			res = append(res, fmt.Sprintf("comment('%s')", col.Comment))
		}
		return res
	}
	handleType ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		nstr := col.SQLType.Name
		if col.Length != 0 {
			if col.Length2 != 0 {
				nstr += fmt.Sprintf("(%v,%v)", col.Length, col.Length2)
			} else {
				nstr += fmt.Sprintf("(%v)", col.Length)
			}
		} else if len(col.EnumOptions) > 0 { //enum
			nstr += "("
			opts := ""

			enumOptions := make([]string, 0, len(col.EnumOptions))
			for enumOption := range col.EnumOptions {
				enumOptions = append(enumOptions, enumOption)
			}
			sort.Strings(enumOptions)

			for _, v := range enumOptions {
				opts += fmt.Sprintf(",'%v'", v)
			}
			nstr += strings.TrimLeft(opts, ",")
			nstr += ")"
		} else if len(col.SetOptions) > 0 { //enum
			nstr += "("
			opts := ""

			setOptions := make([]string, 0)
			for setOption := range col.SetOptions {
				setOptions = append(setOptions, setOption)
			}
			sort.Strings(setOptions)

			for _, v := range setOptions {
				opts += fmt.Sprintf(",'%v'", v)
			}
			nstr += strings.TrimLeft(opts, ",")
			nstr += ")"
		}
		res = append(res, nstr)
		return res
	}
	handleIndex ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		indices := make([]string, 0)
		for idx := range col.Indexes {
			indices = append(indices, idx)
		}
		sort.Strings(indices)

		for _, idx := range indices {
			index := table.Indexes[idx]
			var idxStr string
			if index.Type == schemas.UniqueType {
				idxStr = "unique"
			} else if index.Type == schemas.IndexType {
				idxStr = "index"
			}
			if len(index.Cols) > 1 {
				idxStr += "(" + index.Name + ")"
			}
			res = append(res, idxStr)
		}
		return res
	}
	handleAutoincr ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		if col.IsAutoIncrement {
			res = append(res, "autoincr")
		}
		return res
	}
	handleDefault ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		if col.Default != "" {
			res = append(res, "default "+col.Default)
		}
		return res
	}
	handlePK ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		if col.IsPrimaryKey {
			res = append(res, "pk")
		}
		return res
	}
	handleNullable ColumnHandler = func(table *schemas.Table, col *schemas.Column, res []string) []string {
		if !col.Nullable {
			if !col.IsPrimaryKey {
				res = append(res, "notnull")
			}
		}
		return res
	}
)
