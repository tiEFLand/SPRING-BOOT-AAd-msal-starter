
package utils

import (
	"gopkg.in/mgo.v2"
	"time"
)

var logger = GetLogger()
var GlobalMgoSession *mgo.Session

func init() {
	url, err := GetConfig().Get("mongo.url")
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	globalMgoSession, err := mgo.DialWithTimeout(url, 5*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(600)
}

func NewMongoSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}