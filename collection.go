/**
  create by yy on 2019-10-22
*/

package go_mongodb_yy

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollection struct {
	mdbPool        *MDBPool
	CollectionName string
}

func (m *MongoCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (result *mongo.InsertOneResult, err error) {
	db := <-m.mdbPool.pool

	result, err = db.Collection(m.CollectionName).InsertOne(ctx, document, opts...)

	m.mdbPool.pool <- db
	return
}

func (m *MongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {
	db := <-m.mdbPool.pool

	result, err = db.Collection(m.CollectionName).UpdateOne(ctx, filter, update, opts...)

	m.mdbPool.pool <- db
	return
}
