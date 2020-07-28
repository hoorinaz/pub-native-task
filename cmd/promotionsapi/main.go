package main

import (
	"github.com/hoorinaz/pub-native-task/pkg/promotions"
	"log"
	"net/http"
)

func main() {
	webHandlers := promotions.NewPromotionWeb()
	http.HandleFunc("/promotions/", webHandlers.GetPromotion)

	log.Println("starting http service on port :1321")
	panic(http.ListenAndServe(":1321", nil))

}
