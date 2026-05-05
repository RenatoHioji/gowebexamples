package main

import "net/http"

func main() {

	/* Dynamic based on which folder you are executing the code. In this case, to access as http://localhost/static/home.html u need
	to run the go comand from inside the gowebexamples folder */
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
