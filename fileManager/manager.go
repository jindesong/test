package fileManager

import (
	"fmt"
	"io/ioutil"
)

func GetFileByType(mType string, id string) string {
	var res string
	switch mType {
	case "android":
		path := "files/android"
		res = getFileById(path, id)

	case "ios":

		path := "files/ios"
		res = getFileById(path, id)
	case "web":
		path := "files/web"
		res = getFileById(path, id)

	case "data":
		path := "files/data"
		res = getFileById(path, id)
	case "server":
		path := "files/server"
		res = getFileById(path, id)

	}
	return res
}

func getFileById(filePath string, id string) string {

	realPath := filePath + "/" + id + ".md"
	f, err := ioutil.ReadFile(realPath)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)

}
