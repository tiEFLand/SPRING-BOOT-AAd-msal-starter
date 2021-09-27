
package db

import (
	"crypto-user/utils"

	mgo "gopkg.in/mgo.v2"
)

type Page struct {
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
	Total    int         `json:"total"`
	Data     interface{} `json:"data"`
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	session := utils.NewMongoSession()
	coll := session.DB(db).C(collection)
	return session, coll
}

func Insert(db, collection string, doc interface{}) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Insert(doc)
}

func Update(db, collection string, selector, updateDoc interface{}) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Update(selector, updateDoc)
}

func UpdateAll(db, collection string, selector, updateDoc interface{}) (*mgo.ChangeInfo, error) {
	se, c := connect(db, collection)
	defer se.Close()
	return c.UpdateAll(selector, updateDoc)
}

func Upset(db, collection string, selector, upsetDoc interface{}) (*mgo.ChangeInfo, error) {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Upsert(selector, upsetDoc)
}

func FindOneById(db, collection string, id, result interface{}) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.FindId(id).One(result)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {