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
			content = str[:index] + "<script type=\"text/javascript\" src=\"../scripts/coders.js\"></script>" + str[index:]
		} else if preferences == "design" {
			content = str[:index] + "<script type=\"text/javascript\" src=\"scripts/designers.js\"></script>" + str[index:]
		}

		w.Write([]byte(content))
	})

	router.HandleFunc("/styles/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/styles/styles.css")
	}).Methods(http.MethodGet)

	router.HandleFunc("/images/coder.jpeg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/images/coder.jpeg")
	}).Methods(http.MethodGet)

	router.HandleFunc("/images/designer.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/images/designer.jpg")
	}).Methods(http.MethodGet)

	router.HandleFunc("/fonts/Roboto-Regular.ttf", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/fonts/Roboto-Regular.ttf")
	}).Methods(http.MethodGet)

	router.HandleFunc("/scripts/initScript.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/scripts/initScript.js")
	}).Methods(http.MethodGet)

	router.HandleFunc("/scripts/coders.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/scripts/coders.js")
	}).Methods(http.MethodGet)

	router.HandleFunc("/test/testData.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(testData.GetCodersData())
	}).Methods(http.MethodGet)
}
