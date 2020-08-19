package main //This is part of program to run when we start our application,init we go first(not tokicjstart application)

//way to pull in packages
import (
	"fmt"
	"net/http"
)

//where our requests will go and write
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Welcome to my good site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch,please send an email to <a href = \"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> You are at an invalid page </h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc) //used buitlin serve mux
	http.ListenAndServe(":3000", nil) //used buitin serve mux
}
