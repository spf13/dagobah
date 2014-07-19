// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodbSession *mgo.Session

func init() {
	CreateUniqueIndexes()
}

func DBSession() *mgo.Session {
	if mongodbSession == nil {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			uri = viper.GetString("mongodb_uri")

			if uri == "" {
				log.Fatalln("No connection uri for MongoDB provided")
			}
		}

		var err error
		mongodbSession, err = mgo.Dial(uri)
		if mongodbSession == nil || err != nil {
			log.Fatalf("Can't connect to mongo, go error %v\n", err)
		}

		mongodbSession.SetSafe(&mgo.Safe{})
	}
	return mongodbSession
}

func Items() *mgo.Collection {
	return DB().C("items")
}

func Channels() *mgo.Collection {
	return DB().C("channels")
}

func DB() *mgo.Database {
	return DBSession().DB(viper.GetString("dbname"))
}

func CreateUniqueIndexes() {
	idx := mgo.Index{
		Key:        []string{"key"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	if err := Items().EnsureIndex(idx); err != nil {
		fmt.Println(err)
	}

	if err := Channels().EnsureIndex(idx); err != nil {
		fmt.Println(err)
	}

	ftidx := mgo.Index{
		Key:        []string{"$text:fullcontent"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	if err := Items().EnsureIndex(ftidx); err != nil {
		fmt.Println(err)
	}
}

func AllChannels() []Chnl {
	var channels []Chnl
	results2 := Channels().Find(bson.M{}).Sort("-lastbuilddate")
	results2.All(&channels)
	return channels
}
