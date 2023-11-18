package service

import (
	"strconv"
	"strings"

	"github.com/loveyandex/go-mongo-celestia-scan/mongoCtx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CmnSrv[T any] struct {
	Xcol *mongo.Collection
}

func (cmn *CmnSrv[T]) Search(s string, pgnumber string, limit string) ([]T, error) {
	qs := strings.Split(s, "&")
	filter := bson.D{}
	for _, v := range qs {
		s2 := strings.Split(v, "=")
		if len(s2) < 2 {
			continue
		}
		if s2[0] == "pg" || s2[0] == "lim" || s2[0] == "page" || s2[0] == "limit" {
			continue
		}
		if strings.Contains(s2[0], "_id") {
			hexcid, _ := primitive.ObjectIDFromHex(s2[1])
			filter = append(filter, bson.E{Key: s2[0], Value: hexcid})
			continue
		}
		filter = append(filter, bson.E{Key: s2[0], Value: s2[1]})
	}

	pg, err := strconv.ParseInt(pgnumber, 10, 32)
	if err != nil {
		return nil, err
	}
	lim, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return nil, err
	}
	var m []T
	err = mongoCtx.GetAllByQuery(filter, pg, lim, cmn.Xcol, &m)

	if err != nil {
		return nil, err
	}
	return m, nil

}

func (cmn *CmnSrv[T]) Create(what *T) (interface{}, error) {
	i, err := mongoCtx.CreateRetId(what, cmn.Xcol)
	return i, err
}

func (cmn *CmnSrv[T]) Createold(what interface{}) (interface{}, error) {
	i, err := mongoCtx.CreateRetId(what, cmn.Xcol)
	return i, err
}

// Delete implements IServiceFLow
func (cmn *CmnSrv[T]) Delete(id string) error {
	err := mongoCtx.Delete(id, cmn.Xcol)
	return err
}

// Get implements IServiceFLow
func (ser *CmnSrv[T]) Get(id string) (T, error) {
	var m T
	err := mongoCtx.GetByID(id, ser.Xcol, &m)
	return m, err

}

// GetAll implements IServiceFLow
func (cmn *CmnSrv[T]) GetAll(pgnumber string, limit string) ([]T, error) {

	pg, err := strconv.ParseInt(pgnumber, 10, 32)
	if err != nil {
		return nil, err
	}
	lim, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return nil, err
	}
	var m []T
	err = mongoCtx.GetAll(pg, lim, cmn.Xcol, &m)
	return m, err
}

func (ser *CmnSrv[T]) CountAll() (int64, error) {
	m, err := mongoCtx.CountAll(ser.Xcol)
	return m, err
}
