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
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		util.SendHtml(writer, "static/init.html")
	}).Methods(http.MethodGet)

	router.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		forms := r.Form

		preferences := forms.Get("preferences")

		if preferences == "" {
			util.SendHtml(w, "static/init.html")
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

		if preferences == "coding" {
			content = str[:index] + "<script type=\"text/javascript\" src=\"../static/scripts/coders.js\"></script>" + str[index:]
		} else if preferences == "design" {
			content = str[:index] + "<script type=\"text/javascript\" src=\"../static/scripts/designers.js\"></script>" + str[index:]
		}

		w.Write([]byte(content))
	})

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/test/codersData.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(testData.GetCodersData())
	}).Methods(http.MethodGet)

	router.HandleFunc("/test/designersData.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(testData.GetDesignersData())
	}).Methods(http.MethodGet)
}
