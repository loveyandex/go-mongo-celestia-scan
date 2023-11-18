package celenium

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/loveyandex/go-mongo-celestia-scan/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestXxx(t *testing.T) {

	tts := service.NewTiaTxnSrv()
	ttes := service.NewTiaTxnEventSrv()

	for i := 32100; i < 100000000; i += 100 {
		fmt.Printf("i: %v\n", i)
		c := NewTiaApi()
		tt, err := c.Txs(i)

		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		for _, tt2 := range tt {

			// fmt.Printf("tt2: %v\n", tt2)
			_, err2 := tts.CmnSrv.Create(&tt2)
			if err2 != nil {

				fmt.Printf("err2: %v\n", err2)

			}
			et, err3 := c.EventTxs(tt2.Hash)
			if err3 != nil {

				fmt.Printf("err3: %v\n", err3)
			}
			for _, et2 := range et {

				_, err4 := ttes.CmnSrv.Create(&et2)

				if err4 != nil {
					fmt.Printf("err4: %v\n", err4)
				}

			}

		}

	}

}

func TestTsfrEvents(t *testing.T) {

	var TXAmounts []TXAmount

	ttes := service.NewTiaTxnEventSrv()

	et, err := ttes.CmnSrv.GetAll("0", "1000000000")

	if err != nil {

	}
	for _, et2 := range et {
		if et2.Type == "transfer" {
			te := et2.Data.(primitive.D)
			for _, v := range te {
				if v.Key == "amount" {
					xxxxxxx := v.Value.(string)

					fmt.Printf("te[0].Value.(string): %v\n", xxxxxxx)
					s := strings.Split(xxxxxxx, "utia")
					i, err2 := strconv.ParseInt(s[0], 10, 64)
					if err2 != nil {
						fmt.Printf("err2: %v\n", err2)

					}
					fmt.Printf("i: %v\n", i)
					TXAmounts = append(TXAmounts, TXAmount{
						TxId:   et2.TxID,
						Amount: i,
					})
				}

			}

			fmt.Printf("te: %v\n", te)
		}

	}
	fmt.Printf("len(TXAmounts): %v\n", len(TXAmounts))

	sort.Slice(TXAmounts, func(i, j int) bool {
		return TXAmounts[i].Amount > TXAmounts[j].Amount
	  })

	  for _, v := range TXAmounts {
		fmt.Println(v)
	  }
}

type TXAmount struct {
	TxId   int
	Amount int64
}
