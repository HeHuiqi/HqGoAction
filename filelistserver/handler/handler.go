package handler

import (
	"net/http"
	"strings"
	"os"
	"io/ioutil"
)

const prefix  = "/list/"


type userError string

func (e userError)Error() string  {
	return e.Message()
}
func (e userError) Message() string  {
	return string(e)
}
// HqFileList 文件处理器
func HqFileList(writer http.ResponseWriter, request *http.Request)  error {

	if strings.Index(request.URL.Path,prefix) != 0 {
		return  userError("path must start " + "with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file ,err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer,err.Error(),http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all,err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	writer.Write(all)

	return nil
}