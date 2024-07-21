package aequa

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

// Aequa is the overall type exported by this package. All members are available to any application that
// exports it.
type Aequa struct {
	Name     string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   config
}

// creates a new instance of Aequa, creates necessary files anf folders requires to run this application, i.e
// a .env file and folders such as middlewares, logs e.t.c.
func (a *Aequa) New(rootPath string) error {
	pathConfig := initPaths{
		rootName: rootPath,
		folders:  folderData,
	}

	err := a.Init(pathConfig)
	if err != nil {
		return err
	}

	err = a.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	infoLog, errorLog := a.newLogger()
	a.InfoLog = infoLog
	a.ErrorLog = errorLog
	a.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	a.Version = version
	a.RootPath = rootPath

	a.config = config{
		port:   os.Getenv("PORT"),
		render: os.Getenv("RENDERER"),
	}

	return nil
}

func (a *Aequa) Init(path initPaths) error {
	root := path.rootName
	for _, folders := range path.folders {
		err := a.CreateDirIfNotExist(root + "/" + folders)
		if err != nil {
			return err
		}
	}
	return nil
}

// BuildServer builds the server that will be used in Aequa's module.
func (a *Aequa) BuildServer() error {
	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:", port),
		ErrorLog:     a.ErrorLog,
		Handler:      a.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 100 * time.Second,
	}

	a.InfoLog.Printf("Listening on port %s", port)
	err := srv.ListenAndServe()
	if err != nil {
		a.ErrorLog.Fatal(err)
	}

	return nil
}

// check if a .env file exists
func (a *Aequa) checkDotEnv(path string) error {
	err := a.CreateFileIfNotExist(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

// creates a new logger for errors and general info i.e time and environment
func (a *Aequa) newLogger() (*log.Logger, *log.Logger) {
	var InfoLog *log.Logger
	var ErrorLog *log.Logger

	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return InfoLog, ErrorLog
}
