/**
  create by yy on 2019-10-22
*/

package go_mongodb_yy

import (
	"context"
	"errors"
	"github.com/guaidashu/go_mongodb_yy/libs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"time"
)

var MDBPoolSize = 2

type MDBPool struct {
	pool   chan *mongo.Database
	client *mongo.Client
}

type ClientOpts struct {
	Opt *options.ClientOptions
	Uri string
}

type MDBInfo struct {
	Database string
	Host     string
	Port     string
}

func (m *MDBPool) Collection(collection string) *MongoCollection {
	return &MongoCollection{
		mdbPool:        m,
		CollectionName: collection,
	}
}

func NewClient(ct ...ClientOpts) *MDBPool {
	mdb := &MDBPool{}

	mdb.initMongoDB(ct...)

	return mdb
}

func (m *MDBPool) initMongoDB(ct ...ClientOpts) {
	m.pool = make(chan *mongo.Database, MDBPoolSize)
	if len(ct) < 1 {
		libs.DebugPrint(libs.NewReportError(errors.New("clientOpts' size is nil")).Error())
		return
	}

	m.client = m.getConnect(ct[0].Opt.ApplyURI(ct[0].Uri))

	cs, err := connstring.Parse(ct[0].Uri)
	if err != nil {
		libs.DebugPrint(libs.NewReportError(err).Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = m.client.Connect(ctx)
	if err != nil {
		libs.DebugPrint(libs.NewReportError(err).Error())
		return
	}

	for i := 0; i < MDBPoolSize; i++ {
		m.pool <- m.client.Database(cs.Database)
	}
}

func (m *MDBPool) getConnect(opts ...*options.ClientOptions) *mongo.Client {
	client, err := mongo.NewClient(opts...)

	if err != nil {
		libs.DebugPrint(libs.GetErrorString(libs.NewReportError(err)))
	}

	return client
}

func (m *MDBPool) Close() {

}

//func (m *MDBPool) getApplyUrl() (applyUrl string, err error) {
//	//"mongodb://localhost:27017"
//	if config.Config.Mongodb.Host == "" || config.Config.Mongodb.Port == "" {
//		return "mongodb://localhost:27017", libs.NewReportError(errors.New("mongodb error: nil host or nil port"))
//	}
//	if config.Config.Mongodb.Username == "" {
//		applyUrl = fmt.Sprintf("mongodb://%v:%v", config.Config.Mongodb.Host, config.Config.Mongodb.Port)
//	} else {
//		applyUrl = fmt.Sprintf("mongodb://%v:%v@%v:%v", config.Config.Mongodb.Username, config.Config.Mongodb.Password, config.Config.Mongodb.Host, config.Config.Mongodb.Port)
//	}
//	return
//}
