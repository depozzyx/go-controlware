package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"controlware/shared"

	"github.com/gorilla/mux"
)

var (
	commands  []shared.Command = make([]shared.Command, 0)
	lastClear int64            = time.Now().UnixMilli()
)

func Run(host string) {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/commands/get", get)
	r.HandleFunc("/commands/add", add).Queries("pwd", "dpx")
	r.HandleFunc("/commands/clear", clear).Queries("pwd", "dpx")
	r.HandleFunc("/commands/output", output).Methods("POST")
	r.HandleFunc("/build", build).Queries("file", "{file}").Methods("GET")

	fmt.Printf("Server running on %s\n", host)
	http.ListenAndServe(host, r)
}

func build(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("file")
	if f, err := os.Open("build/" + path); err == nil {
		if bytes, err := ioutil.ReadAll(f); err == nil {
			w.Write(bytes)
			return
		}
	}

	files, _ := ioutil.ReadDir("build/")
	str := ""
	for _, file := range files {
		str += file.Name() + ", "
	}
	w.Write([]byte(fmt.Sprintf("file not found, available %v", str)))
}
func output(w http.ResponseWriter, r *http.Request) {
	if bytes, err := ioutil.ReadAll(r.Body); err == nil {
		f, err := os.Create("outputs/" + fmt.Sprintf("%v-%v.log", time.Now().Unix(), r.Host))
		if err != nil {
			log.Fatal(err)
		}
		f.Write(bytes)
	}
	w.Write([]byte("fuck you"))
}

func get(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(commands)
	if err != nil {
		w.Write([]byte("fuck"))
	} else {
		w.Write(bytes)
	}
}
func clear(w http.ResponseWriter, r *http.Request) {
	commands = []shared.Command{}
	lastClear = time.Now().UnixMilli()
	w.Write([]byte("cock"))
}
func add(w http.ResponseWriter, r *http.Request) {
	if bytes, err := ioutil.ReadAll(r.Body); err == nil {
		data := make([]string, 0)
		if err := json.Unmarshal(bytes, &data); err == nil {
			for _, cmd := range data {
				commands = append(commands, shared.Command{
					Id:  int(lastClear) + len(commands),
					Cmd: cmd,
				})
			}
			w.Write([]byte("cock"))
			return
		}
	}
	w.Write([]byte("format is json list of strings"))
}
