package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"


type userE string

func (e userE) Error() string {
	return e.Massage()
}

func (e userE) Massage() string {
	return string(e)
}

func HandleFileList (writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userE(fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return  nil
}
