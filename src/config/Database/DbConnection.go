package Database

import (
	"flightApi/src/config"
	"fmt"
	"gopkg.in/mgo.v2"
)

func DbConn() *mgo.Session {

	conn, err := mgo.Dial(fmt.Sprintf(config.StringConnDb))
	if err != nil {
		panic(err)
	}
	return conn

}
