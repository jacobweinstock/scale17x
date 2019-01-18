package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/scale17x/golang/pythonbinary"
	log "github.com/sirupsen/logrus"
)

type pythonResponse struct {
	MSG  string `json:"msg"`
	DATE string `json:"date"`
}

func main() {
	fmt.Println("Hello Scale 17x!")
	_ = pythonbinary.WriteToDisk()
	out, err := GenericCMD()
	if string(err) != "" {
		fmt.Println(string(err))
	}

	pResponse := pythonResponse{}
	// marshall cli stdout to json
	jErr := json.Unmarshal(out, &pResponse)
	if err != nil {
		log.Debug("jErr: %v", jErr)
	}

	fmt.Printf("msg: %s\ndate: %s\n", pResponse.MSG, pResponse.DATE)
	_ = pythonbinary.DeleteFromDisk()
}

// GenericCMD - run the daemon
func GenericCMD() (outStr, errStr []byte) {
	dir, derr := filepath.Abs(filepath.Dir(os.Args[0]))
	if derr != nil {
		log.Fatal(derr)
	}
	cmds := dir + "/run"
	cmd := exec.Command("/bin/sh", "-c", cmds)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr = stdout.Bytes(), stderr.Bytes()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s", err)
	}

	return outStr, errStr
}
