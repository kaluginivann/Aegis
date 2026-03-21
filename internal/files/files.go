package files

import (
	"os"

	"github.com/kaluginivann/Aegis/internal/configs"
)

func CheckExistsFile(conf *configs.Config) error {
	if _, err := os.Stat(conf.FilePath); err != nil {
		conf.Logger.Error("File does not exists", "error", err)
		return err
	}
	return nil
}
