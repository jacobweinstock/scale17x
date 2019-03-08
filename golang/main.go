package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jacobweinstock/scale17x/golang/binutils"
)

type pythonResponse struct {
	MSG  string `json:"msg"`
	DATE string `json:"date"`
}

type name struct {
	Name string `json:"name"`
}

func main() {
	// write python binary to disk
	binutils.WriteToDisk()
	defer binutils.DeleteFromDisk()

	var runtime string
	if len(os.Args) < 2 {
		runtime = ""
	} else {
		runtime = os.Args[1]
	}

	switch {
	case strings.Contains(runtime, "web"):
		// handle clean up and serve http
		done := make(chan bool, 1)
		go sigs(done)
		serveHTTP(done)
	default:
		cli(runtime)
	}
}

func cli(arg string) {
	result, err := binutils.RunCMD(arg)
	if string(err) != "" {
		log.Println(string(err))
	}
	logPythonResponse(result)
}

func serveHTTP(done chan bool) {
	http.HandleFunc("/hello", hello)

	log.Println("Starting up...")
	go http.ListenAndServe(":8080", nil)
	log.Println("Listening on Port 8080")
	<-done
	log.Println("Cleaning up and exiting")
}

func sigs(done chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	log.Printf("\nCaught: %v\n", sig)
	done <- true
}

func logPythonResponse(out []byte) {
	pResponse := pythonResponse{}
	err := json.Unmarshal(out, &pResponse)
	if err != nil {
		log.Fatalf("jErr: %v", err)
	}
	log.Printf("msg: %s date: %s\n", pResponse.MSG, pResponse.DATE)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json;charset=utf-8")
	decoder := json.NewDecoder(r.Body)
	var rr name
	dErr := decoder.Decode(&rr)
	if dErr != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result, err := binutils.RunCMD(rr.Name)
	if string(err) != "" {
		log.Println(string(err))
	}
	go logPythonResponse(result)
	w.Write([]byte(result))
}
