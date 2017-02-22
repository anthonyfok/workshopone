// Command server provides a basic RESTesque API for participants to practice
// writing simple Go programs against.
//
// Basic HTTP API sketch
//   GET  /task/                   - Provides documentation
//   GET  /task/?name=[name]       - Gets a new task token
//   POST /task/[token]            - Post form data in reply
//   GET  /task/square/            - Gets documentation
//   GET  /task/square/[token]     - Gets a number to square
//   POST /task/square/[token]     - Post your squared number
//   GET  /task/frequency/[token]  - Gets a paragraph to analyze
//   POST /task/frequency/[token]  - Post your top three words
//   GET  /task/multiply/[token]   - Gets integers to multiply
//   POST /task/multiply/[token]   - Post your integers multiplied
//   GET  /task/operations/[token] - Gets operations to perform
//   POST /task/operations/[token] - Post your results
//
//   GET  /status         - HTML showing status
//   GET  /status/updates - Websocket providing updates to status
//
package main

import (
	"fmt"
	"net/http"
)

func writeDocumentation(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(wr,
		`To complete this exercise, you must retrieve and answer several problems.
First, you have to get a token that will represent a taskset. This is done
with GET /task/new?name=YOURNAME.

Then you must confirm you have a token, simply with POST /task/TOKEN.
The next problems are at the following locations, and must be done in
the order provided, and completed only once per TOKEN:

  /task/square/
  /task/frequency/
  /task/multiply/
  /task/operations/

- Using GET on the above URLs will provide task specific instructions.
- A POST to the correct location with incorrect data will be responded
to with a 406 NotAcceptable.
- 404 NotFound is returned if a TOKEN is invalid, expired, or for a
different task.

To GET started, look up http.Get in godoc.
`)
}
