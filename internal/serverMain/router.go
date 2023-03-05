package serverMain

import (
	"MyServer/testData"
	"MyServer/util"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (server *Server) ConfigureRouter() {
	Handle(server)
}

func Handle(server *Server) {
	router := server.GetRouter()
	router.HandleFunc("/init", func(writer http.ResponseWriter, request *http.Request) {

		request.ParseForm()
		forms := request.Form

		preferences := forms.Get("preferences")

		if preferences == "" {
			util.SendHtml(writer, "static/init.html")
			return
		}

		page, err := os.ReadFile("static/test.html")

		if err != nil {
			fmt.Println(err)
			return
		}
		str := string(page)
		index := strings.Index(str, "</head>")
		var content string

		if preferences == "test0" {
			content = str[:index] + "<script type=\"text/javascript\" src=\"scripts/coders.js\"></script>" + str[index:]
		} else if preferences == "test1" {
			content = str[:index] + "<script type=\"text/javascript\" src=\"scripts/designers.js\"></script>" + str[index:]
		}

		writer.Write([]byte(content))

	}).Methods(http.MethodGet)

	router.HandleFunc("/styles/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/styles/styles.css")
	}).Methods(http.MethodGet)

	router.HandleFunc("/scripts/coders.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/scripts/coders.js")
	}).Methods(http.MethodGet)

	router.HandleFunc("/testData.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(testData.GetCodersData())
	}).Methods(http.MethodGet)
}
