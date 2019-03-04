// 自动转义
// author: baoqiang
// time: 2019/3/4 下午8:34
package ch04

import (
	"html/template"
	"os"
	"github.com/labstack/gommon/log"
)

// <p>A: &lt;b&gt;Hello!&lt;/b&gt;</p><p>B: <b>Hello!</b></p>
func RunAutoEscape() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`

	t := template.Must(template.New("escape").Parse(templ))

	var data struct {
		A string
		B template.HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

}
