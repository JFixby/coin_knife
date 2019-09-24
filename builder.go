package coinknife

import (
	"github.com/jfixby/pin"
)

type Settings struct {
	PathToInputRepo  string
	PathToOutputRepo string

	DoNotProcessAnyFiles   bool
	DoNotProcessSubfolders bool

	FileContentProcessor StringProcessor
	FileNameProcessor    StringProcessor
	IsFileProcessable    FileToProcessDetector

	IgnoredFiles  map[string]bool
	InjectorsPath string
}

func Build(set *Settings) {
	pin.D(" Input", set.PathToInputRepo)
	pin.D("Output", set.PathToOutputRepo)
	pin.D("")

	ClearProject(set.PathToOutputRepo, set.IgnoredFiles)

	TransferFiles(
		set,
	)

	PerformInjections(set.PathToOutputRepo, set.InjectorsPath)

	//FixSecp256k1Checksum(set.PathToOutputRepo)

	AppendGitIgnore(set.PathToOutputRepo)

	GoFmt(set.PathToOutputRepo)

	pin.D("Done!")
}
