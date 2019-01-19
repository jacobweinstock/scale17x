package pythonbinary

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jacobweinstock/scale17x/golang/extmodules"
	log "github.com/sirupsen/logrus"
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
func WriteToDisk() error {
	log.Debug("Writing binary to disk")
	b, err := extmodules.ReadFile(fmt.Sprintf("extmodules/%s", PythonBinaryName))
	if err != nil {
		log.Fatal(err)
		return err
	}
	loc := location()
	wErr := ioutil.WriteFile(loc, b, 0770)
	if wErr != nil {
		log.Fatal(wErr)
		return wErr
	}
	return nil
}

// DeleteFromDisk - delete the binary from the local filesystem
func DeleteFromDisk() error {
	log.Debug("Cleaning up binary")
	loc := location()
	err := os.Remove(loc)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
