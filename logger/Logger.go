package logger

import (
	"os"
	"log"
	"io"
	"litilegame/config"
)
var Debug bool
var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)


func InitLog() {

	tracefile, err := os.OpenFile(config.Log.RootPath+"/traces.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open traces log file:", err)
	}

	infofile, err := os.OpenFile(config.Log.RootPath+"/infos.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open infos log file:", err)
	}

	warningfile, err := os.OpenFile(config.Log.RootPath+"/warnings.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open warings log file:", err)
	}

	errfile, err := os.OpenFile(config.Log.RootPath+"/errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	if Debug {
		Trace = log.New(io.MultiWriter(tracefile,os.Stdout),
			"TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Info = log.New(io.MultiWriter(infofile,os.Stdout),
			"INFO: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Warning = log.New(io.MultiWriter(warningfile,os.Stdout),
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Error = log.New(io.MultiWriter(errfile, os.Stderr),
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}else {
		Trace = log.New(io.MultiWriter(os.Stdout),
			"TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Info = log.New(io.MultiWriter(os.Stdout),
			"INFO: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Warning = log.New(io.MultiWriter(os.Stdout),
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Error = log.New(io.MultiWriter(os.Stderr),
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}

}

