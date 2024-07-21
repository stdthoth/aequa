package aequa

const version = "1.0.0"

type Aequa struct {
	Name    string
	Debug   bool
	Version string
}

func (a *Aequa) New(rootPath string) error {
	pathConfig := initPaths{
		rootName: rootPath,
		folders:  folderData,
	}

	err := a.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (a *Aequa) Init(path initPaths) error {
	root := path.rootName
	for _, folders := range path.folders {
		err := a.CreateDirIfNotExist(root + "/" + folders)
		if err != nil {
			return err
		}
	}
	return nil
}
