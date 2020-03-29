package config

import "html/template"

var TPL *template.Template

func init() {
	TPL = template.Must(TPL.ParseGlob("./templates/*.html"))
}
