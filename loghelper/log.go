package loghelper

import (
	"io"
	"log"
	"os"
)

var (
	// Info : Discard
	Info *log.Logger
	// Warning : Stdout
	Warning *log.Logger
	// Error : Stderr
	Error *log.Logger
)

var errlog *os.File
var infolog *os.File

func set(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init() {
	errlog = getErrLogFile()
	infolog = getLogFile()
	set(infolog, os.Stdout, errlog)
	//Info.Println("Special Information")
	//Warning.Println("There is something you need to know about")
	Info.Println("Start up Info log")
	// Warning.Println("There is something you need to know about")
	Error.Println("Start up Error log")
}

// Free : close log file
func Free() {
	errlog.Close()
	infolog.Close()
}

func getErrLogFile() *os.File {
	//logPath := filepath.Join(os.Getenv("GOPATH"), "/src/Agenda/data/error.log")
	logPath := "data/error.log"
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	return file
	// defer file.Close()
}

func getLogFile() *os.File {
	logPath := "data/info.log"
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	return file
}
