package main

import (
	"github.com/hoorinaz/pub-native-task/pkg/promotions"
	"github.com/hoorinaz/pub-native-task/shared/server"
	"log"
	"net/http"
	"os"
	"time"
)

const FilePath = "FILE_PATH"

var (
	filePath string
)

func init() {
	filePath = os.Getenv(FilePath)
}

func main() {

	// execute CSV process to copy csv data to storage
	executeCSVProcess()
	ticker := time.NewTicker(30 * time.Minute)
	done := make(chan bool)

	// create a goroutine for CSV process (clean the table and import the csv data)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				executeCSVProcess()
			}
		}
	}()

	// create a goroutine for web API
	go func() {
		webHandlers := promotions.NewPromotionWeb()
		http.HandleFunc("/promotions/", webHandlers.GetPromotion)

		log.Println("starting http service on port 1321")
		panic(http.ListenAndServe(":1321", nil))
	}()

	sig := server.WaitSignal()

	ticker.Stop()
	done <- true

	//Blocking...
	log.Println("received signal " + sig.String() + ", exiting...")

}

func executeCSVProcess() {
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

	if err := file.Close(); err != nil {
		log.Panic(err)
	}
	log.Println("finished truncating and copying csv", time.Now())
}
