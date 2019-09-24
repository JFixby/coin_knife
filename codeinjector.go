package coinknife

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
	"github.com/jfixby/pin/lang"
	"path/filepath"
	"strings"
)

var IgnoreUnderscore = func(file string) bool {
	_, f := filepath.Split(file)
	if strings.Index(f, "_") == 0 {
		pin.D("drop", file)
		return false
	}
	return true
}

var GoFile = func(file string) bool {
	return fileops.ExtensionIs(file, "go")
}

var FoldersOnlyAndIgnoreUnderscore = func(file string) bool {
	return IgnoreUnderscore(file) && fileops.FoldersOnly(file)
	//&& str.EndsWith(file, ".go")
}

func PerformInjections(outputPath string, injections string) {
	prefix := fileops.Abs(injections)
	filesList := fileops.ListFiles(prefix, FoldersOnlyAndIgnoreUnderscore, fileops.ALL_CHILDREN)
	pin.D("list", filesList)
	processFiles(outputPath, prefix, filesList)
}

func processFiles(outputPath string, prefix string, files []string) {
	for _, s := range files {
		tail := fileops.SplitPath(s, prefix)
		outputFile := filepath.Join(outputPath, tail)
		inputFile := filepath.Join(prefix, tail)
		processInjections(inputFile, outputFile)
	}
}
func processInjections(i string, o string) {
	list := fileops.ListFiles(i, IgnoreUnderscore, fileops.DIRECT_CHILDREN)
	for _, e := range list {
		if fileops.ExtensionIs(e, "reject") {
			rejectionFile := e
			injectionFile := filepath.Join(fileops.Parent(rejectionFile), fileops.NameWithoutExtention(e)+".inject")
			proc := filepath.Join(fileops.Parent(rejectionFile), fileops.NameWithoutExtention(e))
			pin.D("processing", proc)
			processInjection(o, rejectionFile, injectionFile)
		}
	}
}
func processInjection(out string, rej string, inj string) {
	lang.AssertValue(rej+" file exists", fileops.FileExists(rej), true)
	lang.AssertValue(inj+" file exists", fileops.FileExists(inj), true)
	lang.AssertValue(out+" file exists", fileops.FileExists(out), true)

	outData := fileops.ReadFileToString(out)

	rejData := fileops.ReadFileToString(rej)
	injData := fileops.ReadFileToString(inj)
	before := outData
	outData = strings.Replace(outData, rejData, injData, 1)
	after := outData

	if before == after {
		lang.ReportErr("Code injection failed for: %v\n"+
			"injector: %v",
			out,
			inj,
		)
	}

	fileops.WriteStringToFile(out, outData)
}

func ReplaceFromFile(data, in, out string) string {
	outData := fileops.ReadFileToString(out)
	inData := fileops.ReadFileToString(in)
	//i := IndexOf(data, inData, 0)
	//lang.AssertNot("indexof", i, -1)
	data = strings.Replace(data, inData, outData, 1)
	return data
}
