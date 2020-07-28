package promotions

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
)

type CSVProcessor struct {
	store Store
}

func (c *CSVProcessor) TruncatePromotionTable() error {
	return c.store.TruncatePromotionTable()
}

func (c *CSVProcessor) CopyToStorage(f *os.File) error {
	log.Println("start copying csv to storage")

	wg := new(sync.WaitGroup)

	reader := csv.NewReader(f)
	counter := 1
	var promotionList []*Promotion

	for {
		record, err := reader.Read()
		if err == io.EOF {
			wg.Add(1)
			go func() {
				err = c.store.AddBulkPromotion(promotionList)
				if err != nil {
					return
				}
				wg.Done()
			}()
			break
		}
		if err != nil {
			log.Println("error reading from file", err.Error())
			return err
		}

		pModel, err := recordToPromotion(record, counter)
		if err != nil {
			log.Println("error transforming record to promotion", err.Error())
			return err
		}

		promotionList = append(promotionList, pModel)

		if len(promotionList) == maxRowSize {
			wg.Add(1)
			go func(p []*Promotion) {
				err = c.store.AddBulkPromotion(p)
				if err != nil {
					return
				}
				wg.Done()
			}(promotionList)

			promotionList = nil

		}

		counter++
	}

	wg.Wait()
	log.Println("Finish copying csv to storage")

	return nil
}

func NewCSVProcessor() *CSVProcessor {
	return &CSVProcessor{
		store: NewPromotionStore(),
	}
}
