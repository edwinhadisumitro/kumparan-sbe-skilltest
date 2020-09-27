package helper

import (
	"flag"
	"log"
	"os"
)

// ProjectFolder : Absolute path of project
var ProjectFolder = flag.String("folder", "../../", "absolute path of project folder")

// WriteToLogFile : Function to write to log file
func WriteToLogFile(title string, description string) {
	f, err := os.OpenFile(*ProjectFolder+"logs/error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	//set output of logs to f
	log.SetOutput(f)

	log.Println(title + ", " + description)
}
