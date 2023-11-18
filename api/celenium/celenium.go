package celenium

import (
	"encoding/json"
	"fmt"

	"github.com/loveyandex/go-mongo-celestia-scan/api"
	"github.com/loveyandex/go-mongo-celestia-scan/model" 
)

type TiaApi struct {
	api *api.Api
}

func NewTiaApi() *TiaApi {
	return &TiaApi{api: api.NewApi("https://api.celenium.io/v1", "")}
}

func (x *TiaApi) Txs(offset int)  ([]model.TiaTxn,error ) {

	var e []model.TiaTxn

 
	b, err := x.api.JsonGet(fmt.Sprintf("/tx?limit=100&offset=%d&sort=desc",offset))

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = json.Unmarshal(b, &e)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return e,err


}

func (x *TiaApi) EventTxs(tx string) ([]model.EventTxn, error) {

	var e []model.EventTxn
	b, err := x.api.JsonGet("/tx/" + tx + "/events")

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = json.Unmarshal(b, &e)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return e, nil

}
