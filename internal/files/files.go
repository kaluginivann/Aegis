package files

import (
	"os"

	"github.com/kaluginivann/Aegis/internal/configs"
)

func CheckExistsFile(conf *configs.Config) (os.FileInfo, error) {
	FileInfo, err := os.Stat(conf.FilePath)
	if err != nil {
		conf.Logger.Error("File does not exists", "error", err)
		return nil, err
	}
	return FileInfo, nil
}
