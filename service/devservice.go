package service

import (
	"github.com/loveyandex/go-mongo-celestia-scan/db"
	"github.com/loveyandex/go-mongo-celestia-scan/model"
)

type TiaTxnSrv struct {
	CmnSrv *CmnSrv[model.TiaTxn]
}

func NewTiaTxnSrv() *TiaTxnSrv {
	return &TiaTxnSrv{CmnSrv: &CmnSrv[model.TiaTxn]{Xcol: db.Collection("tia-txn")}}
}

type TiaTxnEventSrv struct {
	CmnSrv *CmnSrv[model.EventTxn]
}

func NewTiaTxnEventSrv() *TiaTxnEventSrv {
	return &TiaTxnEventSrv{CmnSrv: &CmnSrv[model.EventTxn]{Xcol: db.Collection("tia-txn-event")}}
}
