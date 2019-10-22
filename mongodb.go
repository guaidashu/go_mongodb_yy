/**
  create by yy on 2019-10-22
*/

package go_mongodb_yy

import (
	"context"
	"github.com/guaidashu/go_mongodb_yy/libs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"time"
)

var MDBPoolSize = 2

type MDBPool struct {
	pool       chan *mongo.Database
	client     *mongo.Client
	clientOpts *ClientOpts
}

type ClientOpts struct {
	Opt      *options.ClientOptions
	Uri      string
	Database string
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
		Ctx:            context.Background(),
	}
}

func NewClient(ct *ClientOpts) *MDBPool {
	mdb := &MDBPool{
		clientOpts: ct,
	}

	mdb.initMongoDB()

	return mdb
}

func (m *MDBPool) initMongoDB() {
	m.pool = make(chan *mongo.Database, MDBPoolSize)

	m.client = m.getConnect(m.clientOpts.Opt.ApplyURI(m.clientOpts.Uri))

	if m.clientOpts.Database == "" {

		cs, err := connstring.Parse(m.clientOpts.Uri)
		if err != nil {
			libs.DebugPrint(libs.NewReportError(err).Error())
		}

		m.clientOpts.Database = cs.Database
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := m.client.Connect(ctx)
	if err != nil {
		libs.DebugPrint(libs.NewReportError(err).Error())
		return
	}

	for i := 0; i < MDBPoolSize; i++ {
		m.pool <- m.client.Database(m.clientOpts.Database)
	}
}

func (m *MDBPool) getConnect(opts ...*options.ClientOptions) *mongo.Client {
	client, err := mongo.NewClient(opts...)

	if err != nil {
		libs.DebugPrint(libs.GetErrorString(libs.NewReportError(err)))
	}

	return client
}

func (m *MDBPool) Close() error {
	return m.client.Disconnect(context.Background())
}

func (m *MDBPool) GetDatabase() string {
	return m.clientOpts.Database
}
