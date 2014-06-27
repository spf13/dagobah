// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"labix.org/v2/mgo"
)

var mongodbSession *mgo.Session

func init() {
	CreateUniqueIndexes()
}

func DBSession() *mgo.Session {
	connectString := viper.GetString("dbhost") + ":" + viper.GetString("dbport")
	var err error

	if mongodbSession == nil {
		mongodbSession, err = mgo.Dial(connectString)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		if !(viper.GetString("dbusername") == "" && viper.GetString("dbpassword") == "") {
			err = mongodbSession.DB(viper.GetString("dbname")).Login(viper.GetString("dbusername"), viper.GetString("dbpassword"))
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}
	}

	if mongodbSession == nil {
		fmt.Println("unable to connect to MongoDB at", connectString)
		os.Exit(-1)
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

	var err error

	err = Items().EnsureIndex(idx)
	if err != nil {
		fmt.Println(err)
	}

	err = Channels().EnsureIndex(idx)
	if err != nil {
		fmt.Println(err)
	}
}
