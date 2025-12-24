package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

func ParseInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Cannot cast string '%s' to int", s)
	}
	return value
}
