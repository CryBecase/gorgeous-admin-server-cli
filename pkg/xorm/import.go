package xorm

import "github.com/xormplus/xorm/schemas"

type ImportGenerator struct {
	ImportRule map[string]string // 导入规则 例如：time.Time 类型需要导入 time 包
}

func defaultImportGenerator() *ImportGenerator {
	rule := make(map[string]string)
	rule["time.Time"] = "time"

	return &ImportGenerator{ImportRule: rule}
}

func GenerateImports(table *schemas.Table, options ...func(*ImportGenerator)) map[string]string {
	generator := defaultImportGenerator()
	for _, f := range options {
		f(generator)
	}

	res := make(map[string]string)
	for _, col := range table.Columns() {
		typ := TypeTransform(col)

		if v, ok := generator.ImportRule[typ]; ok {
			res[typ] = v
		}
	}
	return res
}
