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
	Ctx            context.Context
}

func (m *MongoCollection) InsertOne(document interface{},
	opts ...*options.InsertOneOptions) (result *mongo.InsertOneResult, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.InsertOne(m.Ctx, document, opts...)
	})

	return
}

func (m *MongoCollection) UpdateOne(filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.UpdateOne(m.Ctx, filter, update, opts...)
	})

	return
}

func (m *MongoCollection) UpdateMany(filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (result *mongo.UpdateResult, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.UpdateMany(m.Ctx, filter, update, opts...)
	})

	return
}

func (m *MongoCollection) InsertMany(documents []interface{},
	opts ...*options.InsertManyOptions) (result *mongo.InsertManyResult, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.InsertMany(m.Ctx, documents, opts...)
	})

	return
}

func (m *MongoCollection) DeleteOne(filter interface{},
	opts ...*options.DeleteOptions) (result *mongo.DeleteResult, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.DeleteOne(m.Ctx, filter, opts...)
	})

	return
}

func (m *MongoCollection) DeleteMany(filter interface{},
	opts ...*options.DeleteOptions) (result *mongo.DeleteResult, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.DeleteMany(m.Ctx, filter, opts...)
	})

	return
}

func (m *MongoCollection) ReplaceOne(filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (result *mongo.UpdateResult, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.ReplaceOne(m.Ctx, filter, replacement, opts...)
	})

	return
}

func (m *MongoCollection) Aggregate(ctx context.Context, pipeline interface{},
	opts ...*options.AggregateOptions) (result *mongo.Cursor, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.Aggregate(m.Ctx, pipeline, opts...)
	})

	return
}

func (m *MongoCollection) CountDocuments(ctx context.Context, filter interface{},
	opts ...*options.CountOptions) (result int64, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.CountDocuments(m.Ctx, filter, opts...)
	})

	return
}

func (m *MongoCollection) EstimatedDocumentCount(ctx context.Context,
	opts ...*options.EstimatedDocumentCountOptions) (result int64, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.EstimatedDocumentCount(m.Ctx, opts...)
	})

	return
}

func (m *MongoCollection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (result *mongo.Cursor, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.Find(m.Ctx, filter, opts...)
	})

	return
}

func (m *MongoCollection) FindOne(ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) (result *mongo.SingleResult) {

	m.Exec(func(collection *mongo.Collection) {
		result = collection.FindOne(m.Ctx, filter, opts...)
	})

	return
}

func (m *MongoCollection) FindOneAndDelete(ctx context.Context, filter interface{},
	opts ...*options.FindOneAndDeleteOptions) (result *mongo.SingleResult) {

	m.Exec(func(collection *mongo.Collection) {
		result = collection.FindOneAndDelete(m.Ctx, filter, opts...)
	})

	return
}

func (m *MongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) (result *mongo.SingleResult) {

	m.Exec(func(collection *mongo.Collection) {
		result = collection.FindOneAndUpdate(m.Ctx, filter, update, opts...)
	})

	return
}

func (m *MongoCollection) FindOneAndReplace(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.FindOneAndReplaceOptions) (result *mongo.SingleResult) {

	m.Exec(func(collection *mongo.Collection) {
		result = collection.FindOneAndReplace(m.Ctx, filter, replacement, opts...)
	})

	return
}

func (m *MongoCollection) Watch(ctx context.Context, pipeline interface{},
	opts ...*options.ChangeStreamOptions) (result *mongo.ChangeStream, err error) {

	m.Exec(func(collection *mongo.Collection) {
		result, err = collection.Watch(m.Ctx, pipeline, opts...)
	})

	return
}

func (m *MongoCollection) Database() (db *mongo.Database) {
	m.Exec(func(collection *mongo.Collection) {
		db = collection.Database()
	})
	return
}

func (m *MongoCollection) Exec(fun func(collection *mongo.Collection)) {
	db := <-m.mdbPool.pool

	// 进行回调
	fun(db.Collection(m.CollectionName))

	m.mdbPool.pool <- db
}
