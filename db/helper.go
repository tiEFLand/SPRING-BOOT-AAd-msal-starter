
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
	se, c := connect(db, collection)
	defer se.Close()
	return c.Find(query).Select(selector).All(result)
}

func FindAllLimit(db, collection string, query, selector, result interface{}, limit int) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Find(query).Select(selector).Limit(limit).All(result)
}

func FindAllLimitSort(db, collection string, query, selector, result interface{}, sort string, limit int) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Find(query).Sort(sort).Select(selector).Limit(limit).All(result)
}

func Aggregate(db, collection string, aggregate, result interface{}) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Pipe(aggregate).All(result)
}

func FindByPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

/**
page: 页码
limit: 每页的条数
sort: 字段排序
query: 查询条件
selector: 查找指定字段
*/
func FindByPageSort(db, collection string, page, limit int, sort string, query, selector, result interface{}) error {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Find(query).Select(selector).Sort(sort).Skip(page * limit).Limit(limit).All(result)
}

func FindCount(db, collection string, query interface{}) (int, error) {
	se, c := connect(db, collection)
	defer se.Close()
	return c.Find(query).Count()
}