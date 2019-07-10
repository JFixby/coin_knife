package fileproc

type StringProcessor func(input string) string

type FileToProcessDetector func(input string) bool
