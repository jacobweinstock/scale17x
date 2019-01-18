package pythonbinary

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/scale17x/golang/extmodules"
	log "github.com/sirupsen/logrus"
)

// WriteToDisk - write cliproxy from virtual filesystem to client filesystem
func WriteToDisk() error {
	log.Debug("Writing cliproxy binary to disk")
	b, err := extmodules.ReadFile("extmodules/run")
	if err != nil {
		log.Fatal(err)
		return err
	}

	dir, derr := filepath.Abs(filepath.Dir(os.Args[0]))
	if derr != nil {
		log.Fatal(derr)
	}

	location := dir + "/run"

	wErr := ioutil.WriteFile(location, b, 0770)
	if wErr != nil {
		log.Fatal(wErr)
		return wErr
	}
	return nil
}

// DeleteFromDisk - delete the cliproxy from the client filesystem
func DeleteFromDisk() error {
	log.Debug("Cleaning up cliproxy binary")

	dir, derr := filepath.Abs(filepath.Dir(os.Args[0]))
	if derr != nil {
		log.Fatal(derr)
	}

	location := dir + "/run"

	err := os.Remove(location)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
