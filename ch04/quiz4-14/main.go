package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"./github"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>Title</th>
	<th>State</th>
  <th>User</th>
	<th>MileStone</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.Title}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	{{if .MileStone}}
<td><a href='{{.MileStone.HTMLURL}}'>{{.MileStone.Title}}</a></td>
{{else}}
<td>none</td>
{{end}}
</tr>
{{end}}
</table>
`))

func FormatHTML(w io.Writer, res *github.IssuesSearchResult) {
	if err := issueList.Execute(w, res); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	sen := "localhost:8000/isuues?q=[key1]+[key2]+[key3]...\nPlease like this input url bar\n\nex)http://localhost:8000/issues?q=repo:golang/go+is:open+http"
	fmt.Fprintf(w, sen)
}

func IssuesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	r.ParseForm()
	res := r.Form.Get("q")
	keys := strings.Split(res, " ")

	result, err := github.SearchIssues(keys)
	if err != nil {
		log.Print(err)
		w.WriteHeader(400) // Bad Request
		body := "<h1>Error! Bad Request</h1>"
		w.Write([]byte(body))
	} else {
		FormatHTML(w, result)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/issues", IssuesHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
