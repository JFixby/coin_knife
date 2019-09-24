package coinknife

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"os"
	"path/filepath"
	"strings"
)

func TransferFiles(set *Settings) {
	detector := set.IsFileProcessable
	fileNameProc := set.FileNameProcessor
	DoNotProcessFiles := set.DoNotProcessAnyFiles
	fileContentProc := set.FileContentProcessor
	from := set.PathToInputRepo
	to := set.PathToOutputRepo

	inputFiles := ListInputProjectFiles(set.PathToInputRepo, set)
	for _, f := range inputFiles {
		postfix := strings.TrimPrefix(f, from)
		postfix = fileNameProc(postfix)
		//pin.D("postfix", postfix)
		newpath := filepath.Join(to, postfix)
		//pin.D("newpath", newpath)
		if fileops.IsFolder(f) {
			err := os.MkdirAll(newpath, 0700)
			lang.CheckErr(err)
			pin.D("make", newpath)
			continue
		}
		if fileops.IsFile(f) {
			ProcessFile(f, newpath, detector, DoNotProcessFiles, fileContentProc)
			continue
		}
	}

}

func ProcessFile(from string, to string, detector FileToProcessDetector, DoNotProcessFiles bool, fileContentProc StringProcessor) {
	if fileops.IsFolder(from) {
		lang.ReportErr("This is not a file: %v", from)
	}
	if detector(from) {
		data := fileops.ReadFileToString(from)
		if !DoNotProcessFiles {
			data = fileContentProc(data)
		}
		fileops.WriteStringToFile(to, data)
	} else {
		fileops.Copy(from, to)
	}
}

func Replace(data, from, to string) string {
	return strings.Replace(data, from, to, -1)
}
