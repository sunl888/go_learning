package serve

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
