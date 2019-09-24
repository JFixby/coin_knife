package coin_knife

import (
	"github.com/jfixby/pin"
	"github.com/picfight/coin_knife/fileproc"
	"github.com/picfight/coin_knife/injector"
	"github.com/picfight/coin_knife/projectops"
)

type Settings struct {
	PathToInputRepo  string
	PathToOutputRepo string

	DoNotProcessAnyFiles bool
	FileContentProcessor fileproc.StringProcessor
	FileNameProcessor    fileproc.StringProcessor
	IsFileProcessable    fileproc.FileToProcessDetector

	IgnoredFiles  map[string]bool
	InjectorsPath string
}

func Build(set *Settings) {
	pin.D(" Input", set.PathToInputRepo)
	pin.D("Output", set.PathToOutputRepo)
	pin.D("")

	projectops.ClearProject(set.PathToOutputRepo, set.IgnoredFiles)

	fileproc.TransferFiles(
		set.IsFileProcessable,
		set.FileNameProcessor,
		set.DoNotProcessAnyFiles,
		set.FileContentProcessor,
		set.PathToInputRepo,
		set.PathToOutputRepo,
	)

	injector.PerformInjections(set.PathToOutputRepo, set.InjectorsPath)

	//FixSecp256k1Checksum(set.PathToOutputRepo)

	projectops.AppendGitIgnore(set.PathToOutputRepo)

	projectops.GoFmt(set.PathToOutputRepo)

	pin.D("Done!")
}
