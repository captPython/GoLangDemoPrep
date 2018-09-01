package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	/*	fmt.Fprintf(w, "<h1> Hello world web </h1>")
		fmt.Fprintf(w, "<p> Welcome to go lang tutorial </p>")
		fmt.Fprintf(w, "<b> Go is good for dapps </b>")  */

	fmt.Fprintf(w, `<h1> Hello world web 2 </h1>
	<p> Welcome to go lang tutorial </p>
	<b> Go is good for dapps </b>`)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
