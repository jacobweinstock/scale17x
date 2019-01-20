package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jacobweinstock/scale17x/golang/pythonbinary"
	log "github.com/sirupsen/logrus"
)

type pythonResponse struct {
	MSG  string `json:"msg"`
	DATE string `json:"date"`
}

type name struct {
	Name string `json:"name"`
}

func main() {
	pythonbinary.WriteToDisk()
	defer pythonbinary.DeleteFromDisk()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Printf("Caught: %v\n", sig)
		done <- true
	}()

	http.HandleFunc("/hello", hello)

	fmt.Println("Starting up...")
	go http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on Port 8080")
	<-done
	fmt.Println("Cleaning up and exiting")
}

func logPythonResponse(out []byte) {
	pResponse := pythonResponse{}
	err := json.Unmarshal(out, &pResponse)
	if err != nil {
		log.Debug("jErr: %v", err)
	}
	fmt.Printf("msg: %s\ndate: %s\n", pResponse.MSG, pResponse.DATE)
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
	result, err := pythonbinary.RunCMD(rr.Name)
	if string(err) != "" {
		fmt.Println(string(err))
	}
	go logPythonResponse(result)
	w.Write([]byte(result))
}
