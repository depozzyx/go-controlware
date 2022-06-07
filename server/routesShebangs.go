package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func output(w http.ResponseWriter, r *http.Request) {
	if bytes, err := ioutil.ReadAll(r.Body); err == nil {
		t := time.Now().Format("15:04.05 2006-01-02")
		f, err := os.Create("outputs/" + fmt.Sprintf("%v %s.log", t, r.RemoteAddr))
		if err != nil {
			log.Fatal(err)
		}
		f.Write(bytes)
	}
	w.Write([]byte("fuck you"))
}
