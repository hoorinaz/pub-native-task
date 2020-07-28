package promotions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type (
	Web interface {
		GetPromotion(http.ResponseWriter, *http.Request)
	}

	web struct {
		store Store
	}
)

func (w *web) GetPromotion(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/promotions/")
	rowID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	promotion := &Promotion{RowID: rowID}

	if err := w.store.GetPromotion(promotion); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	pResponse := promotionResponse(promotion)
	pByte, err := json.Marshal(pResponse)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	_, err = fmt.Fprintln(rw, string(pByte))
	if err != nil {
		log.Println("error writing json data", err.Error())
		return
	}
}

func NewPromotionWeb() Web {
	return &web{store: NewPromotionStore()}
}
