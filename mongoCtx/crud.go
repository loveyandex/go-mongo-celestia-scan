package mongoCtx

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//create read update delete just for authenticated users no check identity here

func Create(what interface{}, coll *mongo.Collection) error {
	_, err2 := coll.InsertOne(context.TODO(), what)
	if err2 != nil {
		return err2
	}

	return nil

}
func CreateRetId(what interface{}, coll *mongo.Collection) (interface{}, error) {
	r, err2 := coll.InsertOne(context.TODO(), what)
	if err2 != nil {
		return "", err2
	}
	return r.InsertedID, nil
}
func GetByID(whatID string, coll *mongo.Collection, result interface{}) error {

	hexcid, _ := primitive.ObjectIDFromHex(whatID)

	err := coll.FindOne(context.Background(), bson.M{"_id": hexcid}).Decode(result)
	if err != nil {
		return err
	}
	return nil

}
func GetByQuery(bsonM bson.M, coll *mongo.Collection, result interface{}) error {
	err := coll.FindOne(context.Background(), bsonM).Decode(result)
	if err != nil {
		return err
	}
	return nil

}
func GetOneByManyQuery(filtet bson.D, coll *mongo.Collection, result interface{}) error {
	err := coll.FindOne(context.Background(), filtet).Decode(result)
	if err != nil {
		return err
	}
	return nil

}

func Update(filter interface{}, update interface{}, coll *mongo.Collection) (*mongo.UpdateResult, error) {
	ur, err := coll.UpdateOne(context.Background(), filter, update)

	return ur, err

}

func GetAllByQuery(query bson.D, page int64, limit int64, coll *mongo.Collection, results interface{}) error {

	opts := options.Find().SetSkip(page * limit).SetLimit(limit)

	cursor, err2 := coll.Find(context.TODO(), query, opts)

	if err2 != nil {
		return err2
	}

	f := cursor.All(context.Background(), results)
	if f != nil {
		return f
	}
	return nil

}
func GetAll(page int64, limit int64, coll *mongo.Collection, results interface{}) error {

	query := bson.D{}
	opts := options.Find().SetSkip(page * limit).SetLimit(limit)

	//sort by time added default 
	// Sort by `price` field descending
	opts.SetSort(bson.D{{"_id", -1}})

	cursor, err2 := coll.Find(context.TODO(), query, opts)
	if err2 != nil {
		return err2
	}

	f := cursor.All(context.Background(), results)
	if f != nil {
		return f
	}
	return nil

}
func CountAll(coll *mongo.Collection) (int64, error) {

	query := bson.D{}

	innn, err2 := coll.CountDocuments(context.Background(), query)

	return innn, err2

}

func Replace(what interface{}, coll *mongo.Collection) {

}

func Delete(whatID string, coll *mongo.Collection) error {
	return nil
}
