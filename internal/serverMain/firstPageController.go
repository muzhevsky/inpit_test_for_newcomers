package serverMain

import (
	"MyServer/testData"
	"MyServer/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handle(server *Server) {
	router := server.GetRouter()
	router.HandleFunc("/page1", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		forms := request.Form

		preferences := forms.Get("preferences")

		if preferences == "" {
			util.SendHtml(writer, "static/page1.html")
			return
		}

		if preferences == "test1" {
			util.SendHtml(writer, "static/testForCoders.html")
		} else {
			util.SendHtml(writer, "static/testForCoders.html")
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/styles/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/styles/styles.css")
	}).Methods(http.MethodGet)

	router.HandleFunc("/scripts/coders.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/scripts/coderTestScript.js")
	}).Methods(http.MethodGet)

	router.HandleFunc("/coderTestData.json", func(w http.ResponseWriter, r *http.Request) {
		response, _ := json.Marshal(testData.NewCoderQuestion())
		fmt.Println("test")
		w.Write(response)
	}).Methods(http.MethodGet)
}
