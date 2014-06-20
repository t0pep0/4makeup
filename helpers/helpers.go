package helpers

import (
	"github.com/shaoshing/train"
	"html/template"
)

var Helpers template.FuncMap

func init() {
	Helpers = train.HelperFuncs
}
