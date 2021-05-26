package data

import "time"

type User struct {
	Email    string
	Password string
	NickName string
	Address  string
}

type Msg struct {
	Time    time.Time `bson:"time"`
	Content string    `bson:"msg"`
	From    string    `bson:"from"`
}
