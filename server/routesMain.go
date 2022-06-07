package server

import (
	"controlware/shared"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
	lastClear = shared.LastTimestamp()
	w.Write([]byte("cock"))
}

func add(w http.ResponseWriter, r *http.Request) {
	if bytes, err := ioutil.ReadAll(r.Body); err == nil {
		data := make([]string, 0)
		if err := json.Unmarshal(bytes, &data); err == nil {
			for _, cmd := range data {
				commands = append(commands, shared.Command{
					Cmd: cmd,
					Id:  shared.LastTimestampId(lastClear, len(commands)),
				})
			}
			w.Write([]byte("cock"))
			return
		}
	}
	w.Write([]byte("format is json list of strings"))
}
