package filesys

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/deeper-x/ique/client"
	"github.com/deeper-x/ique/configuration"
	"github.com/deeper-x/ique/myutils"
	"github.com/fsnotify/fsnotify"
)

// OS is the main package interface
type OS interface {
	AddListener() error
	ReadFileContent(string) (string, error)
}

// FileManager is in charge for OS management
type FileManager struct {
	Pwd string
}

const name = configuration.QueueName

// ReadFileContent read file content and return it
func (fm FileManager) ReadFileContent(fileName string) (string, error) {
	filePath := fmt.Sprintf("%s/%s", fm.Pwd, fileName)
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// AddListener listen for file creation
func (fm FileManager) AddListener() error {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	go func() error {
		log.Printf("Monitoring %v ...\n", fm.Pwd)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return errors.New("Error: Watching event")
				}

				if event.Op.String() == "CREATE" {
					log.Printf("#TODO trigger push + deletion of %s\n", event.Name)
					// fileManager := filesys.FileManager{Pwd: configuration.MonitoredDir}
					dataContent, err := fm.ReadFileContent("test.txt")
					myutils.FailsOnError(err, "Failed to read file content...")

					pitch := client.Pitch{}
					err = client.Run(&pitch, name, dataContent)

					myutils.FailsOnError(err, "Failed running sender...")
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return err
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(fm.Pwd)
	if err != nil {
		log.Println(err)
		return err
	}

	<-done
	return nil
}

// RunListen is the listener runner
func RunListen(o OS) error {
	err := o.AddListener()

	if err != nil {
		return err
	}
	return nil
}
