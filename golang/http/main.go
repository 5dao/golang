package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(url.ParseQuery(r.URL.RequestURI()))

	fmt.Fprintln(w, "Tmh3je9zbnHAvPfwwHhQsFSJmKkeRTtKqmV", r.URL.RequestURI())
}

func main() {
	http.HandleFunc("/", Indexhandler)

	// 设置静态目录
	fsh := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))

	http.ListenAndServe("0.0.0.0:6060", nil)
}
