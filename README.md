# Go Mongodb designed by yy.

## Usage

1. In Go file, you need to input these code.

```go
var MDB *go_mongodb_yy.MDBPool

func InitMongoDB() {
	MDB = getConnect()
}

func getConnect() *go_mongodb_yy.MDBPool {
	// 设置连接池 数量
	if config.Config.Mongodb.PoolSize != 0 {
		go_mongodb_yy.MDBPoolSize = config.Config.Mongodb.PoolSize
	}

	applyUrl, err := getApplyUrl()
	if err != nil {
		libs.DebugPrint(libs.GetErrorString(libs.NewReportError(err)))
		return nil
	}

	return go_mongodb_yy.NewClient(go_mongodb_yy.ClientOpts{
		Uri: applyUrl,
		Opt: options.Client(),
	})
}

func getApplyUrl() (applyUrl string, err error) {
	if config.Config.Mongodb.Host == "" || config.Config.Mongodb.Port == "" {
		return "mongodb://localhost:27017/admin", libs.NewReportError(errors.New("mongodb error: nil host or nil port"))
	}
	if config.Config.Mongodb.Username == "" {
		applyUrl = fmt.Sprintf("mongodb://%v:%v/%v",
			config.Config.Mongodb.Host,
			config.Config.Mongodb.Port,
			config.Config.Mongodb.Database)
	} else {
		applyUrl = fmt.Sprintf("mongodb://%v:%v@%v:%v/%v",
			config.Config.Mongodb.Username,
			config.Config.Mongodb.Password,
			config.Config.Mongodb.Host,
			config.Config.Mongodb.Port,
			config.Config.Mongodb.Database)
	}
	return
}

```

2. In Xml file, you need to input these code.

```xml
mongodb:
  database:
    "admin"
  port:
    27017
  host:
    "127.0.0.1"
  username:
    ""
  password:
    ""
  poolsize:
    1
```