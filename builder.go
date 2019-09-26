package coinknife

import (
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
)

type Settings struct {
	PathToInputRepo  string
	PathToOutputRepo string

	DoNotProcessAnyFiles   bool
	DoNotProcessSubfolders bool

	FileContentProcessor StringProcessor
	FileNameProcessor    StringProcessor
	IsFileProcessable    FileToProcessDetector

	IgnoredFiles    map[string]bool
	InjectorsPath   string
	AppendGitIgnore func(targetProject string)
	GoFmt           func(targetProject string)
}

func Build(set *Settings) {
	pin.D(" Input", set.PathToInputRepo)
	pin.D("Output", set.PathToOutputRepo)
	pin.D("")

	if fileops.FileExists(set.PathToOutputRepo) {
		ClearProject(set.PathToOutputRepo, set.IgnoredFiles)
	}
	TransferFiles(
		set,
	)

	PerformInjections(set.PathToOutputRepo, set.InjectorsPath)

	//FixSecp256k1Checksum(set.PathToOutputRepo)

	if set.AppendGitIgnore != nil {
		set.AppendGitIgnore(set.PathToOutputRepo)
	}

	if set.GoFmt != nil {
		set.GoFmt(set.PathToOutputRepo)
	}

	pin.D("Done!")
}
