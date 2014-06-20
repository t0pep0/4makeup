package controllers

import (
	"bitbucket.org/skill-im/skillim/views"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func RenderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html")
	ur := r.URL.Path
	var tmplPath string
	var flag bool
	if len(ur) > 5 {
		flag = ur[len(ur)-len(".gold"):] == ".gold"
		tmplPath = ur[:len(ur)-len(".gold")]
	} else {
		flag = false
	}
	if flag {
		views.Render(w, tmplPath, nil)
	} else {
		info, _ := os.Stat("./views" + ur)
		html := `<pre><a href='../'>../</a><br>`
		if info.IsDir() {
			files, _ := ioutil.ReadDir("./views" + ur)
			for _, file := range files {
				fileName := file.Name()
				if file.IsDir() {
					fileName += "/"
				}
				html += "<a href='" + fileName + "'>" + fileName + "</a><br>"
			}
			html += "</pre>"
		} else {
		}
		fmt.Fprint(w, html)
	}
}
