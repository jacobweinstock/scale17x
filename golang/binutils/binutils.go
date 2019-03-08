package binutils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jacobweinstock/scale17x/golang/extmodules"
)

const (
	// PythonBinaryName - final python binary name
	PythonBinaryName = "scale17x-py"
)

func location() string {
	dir, derr := filepath.Abs(filepath.Dir(os.Args[0]))
	if derr != nil {
		log.Fatal(derr)
	}

	return fmt.Sprintf("%s/%s", dir, PythonBinaryName)
}

// WriteToDisk - write binary from virtual filesystem to local filesystem
func WriteToDisk() {
	log.Println("Writing binary to disk")
	b, err := extmodules.ReadFile(fmt.Sprintf("extmodules/%s", PythonBinaryName))
	if err != nil {
		log.Fatal(err)
	}
	loc := location()
	wErr := ioutil.WriteFile(loc, b, 0770)
	if wErr != nil {
		log.Fatal(wErr)
	}
}

// DeleteFromDisk - delete the binary from the local filesystem
func DeleteFromDisk() {
	log.Println("Cleaning up binary")
	loc := location()
	err := os.Remove(loc)
	if err != nil {
		log.Fatal(err)
	}
}

// RunCMD - call the python binary
func RunCMD(arg string) (outStr, errStr []byte) {
	dir, derr := filepath.Abs(filepath.Dir(os.Args[0]))
	if derr != nil {
		log.Fatal(derr)
	}
	cmds := dir + "/" + PythonBinaryName + " " + arg
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
