package main

import (
	"fmt"
	"log"

	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

const (
	COLLECTION = "Timesheets"
)

type MongoServer struct {
	Server   string
	Database string
}

func (m *MongoServer) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

var server = MongoServer{Server: "localhost:27017", Database: "mongoteste"}

type TimeSheet struct {
	id    int
	Start time.Time
	End   time.Time
}

type TimeSheetDAO struct {
}

func (m *TimeSheetDAO) GetAll() ([]TimeSheet, error) {
	var timesheets []TimeSheet
	err := db.C(COLLECTION).Find(bson.M{}).All(&timesheets)
	return timesheets, err
}

func (m *TimeSheetDAO) Create(timesheet TimeSheet) error {
	err := db.C(COLLECTION).Insert(timesheet)
	return err
}

var dao = TimeSheetDAO{}

func main() {
	server.Connect()
	timesheet := TimeSheet{id: 0, Start: time.Now(), End: time.Now()}
	fmt.Printf("Done")
	dao.Create(timesheet)
}
