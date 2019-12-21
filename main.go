package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"./builder"
)

type Process struct {
	Executer int   `json:"exec"`
	Process  []int `json:"process"`
}
var p Process

func main() {
	loadProcess()

	b := builder.Builder{}
	b.SetExecuter(p.Executer)	

	b.AddActions(p.Process)
	b.SetActions()

	b.Exec.BuildPart()
}

func loadProcess() {
	file, err := ioutil.ReadFile("resources/process.json")
	if err != nil {
		log.Fatalf("Error reading from file : %v", err)
	}

	err = json.Unmarshal(file, &p)
	if err != nil {
		log.Fatalf("Error on convert JSON  to object: %v", err)
	}
}