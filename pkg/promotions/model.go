package promotions

import "time"

type (
	Promotion struct {
		RowID          int64     `db:"row_id"`
		ID             string    `db:"id" json:"id"`
		Price          float64   `db:"price" json:"price"`
		ExpirationDate time.Time `db:"expiration_date" json:"expiration_date"`
	}
	PromotionResponse struct {
		ID             string  `json:"id"`
		Price          float64 `json:"price"`
		ExpirationDate string  `json:"expiration_date"`
	}
)
