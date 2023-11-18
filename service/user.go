package service

import (
	"strconv"

	"github.com/loveyandex/go-mongo-celestia-scan/db"
	"github.com/loveyandex/go-mongo-celestia-scan/model"
	"github.com/loveyandex/go-mongo-celestia-scan/mongoCtx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Col *mongo.Collection
}

// Delete implements IServiceFLow
func (ser *UserService) Delete(id string) error {
	err := mongoCtx.Delete(id, ser.Col)
	return err
}

// Get implements IServiceFLow
func (ser *UserService) Get(id string) (*model.User, error) {

	var m model.User
	err := mongoCtx.GetByID(id, ser.Col, &m)
	return &m, err

}

// Get implements IServiceFLow
func (ser *UserService) GetUserByPhone(id string) (*model.User, error) {

	var m model.User
	err := mongoCtx.GetByQuery(primitive.M{"phone": id}, ser.Col, &m)
	return &m, err

}

// Get implements IServiceFLow
func (ser *UserService) GetUserByInstaId(pk int64) (*model.User, error) {

	var m model.User
	err := mongoCtx.GetByQuery(primitive.M{"insta_user.pk": pk}, ser.Col, &m)
	return &m, err

}

// Get implements IServiceFLow
func (ser *UserService) GetUserByInstaUsername(un string) (*model.User, error) {
	var m model.User
	err := mongoCtx.GetByQuery(primitive.M{"insta_user.username": un}, ser.Col, &m)
	return &m, err

}

// GetAll implements IServiceFLow
func (ser *UserService) GetAll(pgnumber string, limit string) ([]model.User, error) {

	pg, err := strconv.ParseInt(pgnumber, 10, 32)
	if err != nil {
		return nil, err
	}
	lim, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return nil, err
	}
	var m []model.User
	err = mongoCtx.GetAll(pg, lim, ser.Col, &m)
	return m, err
}

// GetAll implements IServiceFLow
func (ser *UserService) CountAll() (int64, error) {
	m, err := mongoCtx.CountAll(ser.Col)
	return m, err
}

// Post implements IServiceFLow
func (ser *UserService) Post(obj interface{}) error {
	err := mongoCtx.Create(obj, ser.Col)
	return err

}
func (ser *UserService) PostRetId(obj interface{}) (interface{}, error) {
	r, err := mongoCtx.CreateRetId(obj, ser.Col)
	return r, err

}

// Put implements IServiceFLow
func (ser *UserService) Put(obj interface{}) error {
	panic("unimplemented")
}

// Put implements IServiceFLow
func (ser *UserService) PushNewWalletById(id string, obj interface{}) error {

	hexcid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = mongoCtx.Update(bson.M{"_id": hexcid}, bson.M{"$push": bson.M{"wallets": obj}}, ser.Col)
	return err

}

func NewUserService() *UserService {
	return &UserService{
		Col: db.Collection("user")}

}
