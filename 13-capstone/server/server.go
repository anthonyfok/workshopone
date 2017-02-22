package main

import (
	"log"
	"net/http"
	"os"
)

const (
	square     = "square"
	frequency  = "frequency"
	multiply   = "multiply"
	operations = "operations"
)

type dummyFrequencyGenerator struct{}

func (dfg *dummyFrequencyGenerator) FetchParagraph() string {
	return "one one two one two three one two three four"
}

func main() {
	host := os.Getenv("EDMONTONGO_WORKSHOPONE_HOST")
	if host == "" {
		host = ":8080"
	}

	tm := NewTaskMaster()
	tm.problems[square] = &SquareProblemGenerator{}
	tm.next[""] = square
	tm.problems[frequency] = &FrequencyProblemGenerator{&dummyFrequencyGenerator{}}
	tm.next[square] = frequency
	tm.problems[multiply] = &MultiplyProblemGenerator{}
	tm.next[frequency] = multiply
	tm.problems[operations] = &OperationsProblemGenerator{}
	tm.next[multiply] = operations

	http.Handle("/task/", http.StripPrefix("/task", tm.Mux()))
	http.Handle("/status/", http.StripPrefix("/status", tm.sk))

	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal("Could not launch HTTP server:", err.Error())
	}
}
