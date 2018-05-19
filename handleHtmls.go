package main

import (
	"net/http"
)

type params struct {
	Title  string
	Header string
	imgS   []string
}

func handleHtmls() {
	http.HandleFunc("/home", homeHndl)
}

func homeHndl(w http.ResponseWriter, r *http.Request) {
	imgs := []string{"photos/home.png"}
	pars := params{"Home - ImageWeb", "Home", imgs}
	//fmt.Println(pars.Title)
	//fmt.Println(pars.Header)
	//fmt.Println(pars.imgS[0])
	htmlTpl.ExecuteTemplate(w, "home.html", pars)
}
