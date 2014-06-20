package views

import (
	"github.com/shaoshing/train"
	"github.com/t0pep0/4makeup/helpers"
	"github.com/t0pep0/gold"
	"log"
	"net/http"
)

const viewsDir = "./views/"
const goldExt = ".gold"

var g = gold.NewGenerator(train.IsInProduction()) //Use cache for views if production

func Render(w http.ResponseWriter, name string, TmplParam interface{}) {
	g.SetHelpers(helpers.Helpers)
	tpl, err := g.ParseFile(viewsDir + name + goldExt)
	if err != nil {
		log.Println("Template:", name)
		log.Println(err)
		return
	}
	err = tpl.Execute(w, TmplParam)
	if err != nil {
		log.Println("Template:", name)
		log.Println(err)
		return
	}
}
