package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func ReadFileForDay(day int, fileName string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(currentFile))

	dayFolder := "day" + fmt.Sprintf("%02d", day)
	filePath := filepath.Join(projectRoot, baseDir, dayFolder, fileName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
