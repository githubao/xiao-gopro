// 模板格式化
// author: baoqiang
// time: 2019/3/4 下午8:32
package ch04

import (
	"time"
	"os"
	"log"
	"code.byted.org/gopkg/htmlutil/gosrc/template"
)

const templ = `{{.TotalCount}} Issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title|printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

//19223 Issues:
//----------------------------------------
//Number: 2
//User: golangfm
//Title: golang-packages
//Age: 19 days
//----------------------------------------
//Number: 38735
//User: turbo
//Title: Golang panic causes Dockerd to crash
//Age: 17 days
//----------------------------------------

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func RunTextTemplate() {
	report := template.Must(template.New("issueList").
		Funcs(template.FuncMap{"daysAgo":daysAgo}).
		Parse(templ))

	result, err := SearchIssues([]string{"golang", "docker"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
