package models

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"path"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/xormplus/xorm/names"
	"github.com/xormplus/xorm/schemas"

	"gorgeous-admin-server-cli/internal/file"
	"gorgeous-admin-server-cli/internal/log"
	"gorgeous-admin-server-cli/pkg/task"
	"gorgeous-admin-server-cli/pkg/util"
	"gorgeous-admin-server-cli/pkg/xorm"
)

type XormTmpl struct {
	Table                *schemas.Table
	Imports              map[string]string
	Models               string
	LongestJsonTagLength int
}

func execute(cmd *cobra.Command, args []string) {
	// init config from config.json
	initConfig()
	initDB()

	genDir := _config.OutPath

	err := util.MakeNotExistDir(genDir)
	if err != nil {
		log.Error(err.Error())
		return
	}

	// snake name
	mapper := &names.SnakeMapper{}

	// init template
	t := template.New("t")
	t.Funcs(template.FuncMap{
		"Mapper": mapper.Table2Obj, // a_b_c => ABC
		"Type":   xorm.TypeTransform,
		"Tag":    tag,
	})
	t, err = t.Parse(tmplStr)
	if err != nil {
		log.Error(err.Error())
		return
	}

	tables, err := db.DBMetas()
	if err != nil {
		log.Error(err.Error())
		return
	}

	// 需要生成的表
	tableMap := make(map[string]struct{})
	for _, table := range _config.Tables {
		tableMap[table.Name] = struct{}{}
	}

	tasks := make([]task.Task, 0)
	rule := file.NewRule()
	for _, table := range tables {
		if _, ok := tableMap[table.Name]; !ok { // 不需要生成的表跳过
			continue
		}

		outPath := path.Join(genDir, table.Name+".go")

		// 统计最长的 json tag
		longestJsonTagLength := 0
		for _, col := range table.ColumnsSeq() {
			if len(col) > longestJsonTagLength {
				longestJsonTagLength = len(col)
			}
		}

		tmpl := &XormTmpl{
			Table:                table,
			Imports:              xorm.GenerateImports(table),
			Models:               "model",
			LongestJsonTagLength: longestJsonTagLength,
		}

		// execute template
		bufStr := bytes.NewBufferString("")
		err = t.Execute(bufStr, tmpl)
		if err != nil {
			log.Error(err.Error())
			return
		}

		tmplContent, err := ioutil.ReadAll(bufStr)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// format
		tmplContent, err = format.Source(tmplContent)
		if err != nil {
			log.Error("go format error: " + err.Error())
			return
		}

		diskFile := file.NewDiskFile(outPath, string(tmplContent), 0666, false, nil)
		tasks = append(tasks, file.NewDiskFileTask(diskFile, rule))
	}

	taskGroup := task.NewGroup("Generating models", tasks)
	err = taskGroup.Exec()
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func tag(table *schemas.Table, col *schemas.Column, longestJsonTag int) string {
	return xorm.GenerateTag(table, col, func(generator *xorm.TagGenerator) {
		generator.LongestJsonTagLength = longestJsonTag
	})
}
