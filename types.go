package aequa

var folderData = []string{"handlers", "middleware", "logs", "cache", "migrations", "views", "data", "tmp", "public"}

type initPaths struct {
	rootName string
	folders  []string
}
