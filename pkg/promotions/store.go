package promotions

import (
	"github.com/gocraft/dbr/v2"
	"github.com/hoorinaz/pub-native-task/shared/postgres"
	"log"
)

const (
	maxRowSize = 1000
)

type (
	Store interface {
		AddBulkPromotion([]*Promotion) error
		GetPromotion(*Promotion) error
		TruncatePromotionTable() error
	}
	store struct {
		SQL *dbr.Session
	}
)

var (
	promotionTableName        = "pubnative.promotions"
	promotionTableColumnNames = []string{
		"row_id",
		"id",
		"price",
		"expiration_date",
	}
)

func (s *store) AddBulkPromotion(promotions []*Promotion) error {
	insertStmt := s.SQL.InsertInto(promotionTableName).
		Columns(promotionTableColumnNames...)

	for _, v := range promotions {
		insertStmt.Record(v)
	}

	_, err := insertStmt.Exec()
	if err != nil {
		log.Println("error inserting bulk promotions", err.Error())
		return err
	}

	return nil
}

func (s *store) GetPromotion(p *Promotion) error {
	return s.SQL.Select(promotionTableColumnNames...).From(promotionTableName).
		Where("row_id = ?", p.RowID).LoadOne(p)
}

func (s *store) TruncatePromotionTable() error {
	_, err := s.SQL.DeleteFrom(promotionTableName).Exec()
	if err != nil {
		log.Println("error truncating promotion table")
		return err
	}

	return nil
}

func NewPromotionStore() Store {
	return &store{SQL: postgres.NewSession()}

}
