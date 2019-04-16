package main

import (
	"learngo/errhandling/filelistingsever/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper (handler appHandler) func(http.ResponseWriter,*http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()


		err := handler(writer, request)

		if  err != nil {
			log.Printf("Error handling request: %s",err.Error())

			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Massage(), http.StatusBadGateway)
				return
			}


			code := http.StatusOK
			switch  {
			case os.IsExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				log.Println("===4===")
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Massage() string
}



func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}