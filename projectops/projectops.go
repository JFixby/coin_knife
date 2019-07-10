package projectops

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ClearProject(target string) {
	pin.D("clear", target)
	files, err := ioutil.ReadDir(target)
	lang.CheckErr(err)

	for _, f := range files {
		fileName := f.Name()
		filePath := filepath.Join(target, fileName)
		if fileName == ".git" {
			pin.D("  skip", filePath)
			continue
		}
		if fileName == "vendor" {
			pin.D("  skip", filePath)
			continue
		}
		pin.D("delete", filePath)
		err := os.RemoveAll(filePath)
		lang.CheckErr(err)
	}
	pin.D("")

}

func ListInputProjectFiles(target string) []string {
	if fileops.IsFile(target) {
		lang.ReportErr("This is not a folder: %v", target)
	}

	files, err := ioutil.ReadDir(target)
	lang.CheckErr(err)
	result := []string{}
	for _, f := range files {
		fileName := f.Name()
		filePath := filepath.Join(target, fileName)

		if fileName == ".git" {
			continue
		}
		if fileName == "vendor" {
			continue
		}

		if fileops.IsFolder(filePath) {
			children := ListInputProjectFiles(filePath)
			result = append(result, children...)
			continue
		}

		if fileops.IsFile(filePath) {
			result = append(result, filePath)
			continue
		}
	}
	result = append(result, target)
	lang.CheckErr(err)
	return result
}
