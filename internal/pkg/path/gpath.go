package path

import (
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// HandleHomedirOrPwd replace ~ with homeDir or . with wdDir
func HandleHomedirOrPwd(filePath string) string {
	const homeDir = "~"
	const pwdDir = "."
	if strings.HasPrefix(filePath, homeDir) {
		home, err := homedir.Dir()
		if err != nil {
			return filePath
		}
		result := strings.Replace(filePath, homeDir, home, 1)
		return result
	}

	if strings.HasPrefix(filePath, pwdDir) {
		pwd, err := os.Getwd()
		if err != nil {
			return filePath
		}
		result := strings.Replace(filePath, pwdDir, pwd, 1)
		return result
	}

	return filePath
}
