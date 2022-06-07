package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func outputFile(w http.ResponseWriter, r *http.Request) {
	fileHandler(w, r, "outputs/")
}

func buildFile(w http.ResponseWriter, r *http.Request) {
	fileHandler(w, r, "build/")
}

func fileHandler(w http.ResponseWriter, r *http.Request, folder string) {
	path := r.URL.Query().Get("file")
	if f, err := os.Open(folder + path); err == nil {
		if bytes, err := ioutil.ReadAll(f); err == nil {
			w.Write(bytes)
			return
		}
	}

	files, _ := ioutil.ReadDir(folder)
	str := ""
	for _, file := range files {
		str += fmt.Sprintf(
			"<li><a href='%s?file=%s'>%s</a></li>",
			r.URL.RawPath,
			file.Name(),
			file.Name(),
		)
	}
	w.Header().Add("content-type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("file not found, available files <ul>%v</ul>", str)))
}
