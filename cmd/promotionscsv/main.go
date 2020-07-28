package main

import (
	"github.com/hoorinaz/pub-native-task/pkg/promotions"
	"log"
	"os"
)

const FilePath = "FILE_PATH"

var (
	filePath string
)

func init() {
	if filePath == "" {
		filePath = os.Getenv(FilePath)
	}
}
func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}

	csvProc := promotions.NewCSVProcessor()

	if err := csvProc.TruncatePromotionTable(); err != nil {
		panic(err)
	}

	if err := csvProc.CopyToStorage(file); err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Panic(err)
		}
	}()
}
