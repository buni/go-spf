package example

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/buni/go-spf"
)

// SPFPageRender a sample httphander implementing both normal template rendering and SPF response
func SPFPageRender(w http.ResponseWriter, r *http.Request) {

	spfq := r.URL.Query().Get("spf")

	if spfq == "navigate" || spfq == "load" {
		w.Header().Set("Content-Type", "application/json")

		buf := new(bytes.Buffer)
		t, _ := template.ParseFiles("views/index.tmpl")
		t.ExecuteTemplate(buf, "content", "")

		sp := spf.New()
		sp.SetBody("content", buf.String())
		sp.SetTitle("Index page")
		sp.SetAttribute("load", "value", "done")
		spfjson, _ := sp.EncodeJSON()

		w.Write(spfjson.Bytes())

		return
	}

	t, _ := template.ParseFiles("views/layout.tmpl", "views/index.tmpl")

	w.Header().Set("Content-Type", "text/html")

	t.ExecuteTemplate(w, "layout", "")
}
