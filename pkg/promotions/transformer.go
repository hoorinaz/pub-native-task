package promotions

import (
	"strconv"
	"time"
)

const (
	defaultTimeLayout  = "2006-01-02 15:04:05 -0700 MST"
	responseTimeLayout = "2006-01-02 15:04:05"
)

func recordToPromotion(r []string, c int) (*Promotion, error) {

	t, err := time.Parse(defaultTimeLayout, r[2])
	if err != nil {
		return nil, err
	}
	p, err := strconv.ParseFloat(r[1], 64)
	if err != nil {
		return nil, err
	}

	pModel := &Promotion{
		RowID:          int64(c),
		ID:             r[0],
		Price:          p,
		ExpirationDate: t,
	}

	return pModel, nil
}

func promotionResponse(p *Promotion) *PromotionResponse {
	return &PromotionResponse{
		ID:             p.ID,
		Price:          p.Price,
		ExpirationDate: p.ExpirationDate.Format(responseTimeLayout),
	}

}
