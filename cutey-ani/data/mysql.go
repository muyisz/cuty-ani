package data

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func CreatDB() *MySQL {
	return &MySQL{
		db: nil,
	}
}

const dsn = "root:256275@tcp(127.0.0.1:3306)/cuty_ani?charset=utf8mb4&parseTime=True"

func (m *MySQL) InitDatabase() error {
	con, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	m.db = con
	err = m.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQL) CloseDatabase() error {
	err := m.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQL) SetPhoto(id int, url string, supp string) error {
	sqlStr := "Insert into photo(url,id,supp) values(?,?,?)"
	_, err := m.db.Exec(sqlStr, url, id, supp)
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQL) JoinUsers(user *User) error {
	sqlStr := "Insert into user(email,password,nickname,address) values(?,?,?,?)"
	_, err := m.db.Exec(sqlStr, user.Email, user.Password, user.NickName, user.Address)
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQL) GetUrl(id int) (string, string, error) {
	sqlStr := "select url from photo where id=?"
	var u string
	err := m.db.QueryRow(sqlStr, id).Scan(&u)
	if err != nil {
		return "", "", err
	}
	sqlStr = "select supp from photo where id=?"
	var v string
	err = m.db.QueryRow(sqlStr, id).Scan(&v)
	if err != nil {
		return u, "", err
	}
	return u, v, nil
}

func (m *MySQL) CheckUsers(user *User) (bool, error) {
	sqlStr := "select password from user where email=?"
	var u string
	err := m.db.QueryRow(sqlStr, user.Email).Scan(&u)
	if err != nil {
		return false, err
	}
	if u == user.Password {
		return true, nil
	} else {
		return false, nil
	}
}

func (m *MySQL) GetPhotoNum() (int, error) {
	sqlStr := "select count(*) from photo"
	var num string
	err := m.db.QueryRow(sqlStr).Scan(&num)
	cnt, _ := strconv.Atoi(num)
	if err != nil {
		return cnt, err
	}
	return cnt, nil
}

func (m *MySQL) GetUser(phone string) (User, error) {
	var n string
	var a string
	sqlStr := "select nickname from user where email=?"
	var u User
	err := m.db.QueryRow(sqlStr, phone).Scan(&n)
	if err != nil {
		return u, err
	}
	sqlStr = "select address from user where email=?"
	err = m.db.QueryRow(sqlStr, phone).Scan(&a)
	u.Address = a
	u.NickName = n
	u.Email = phone
	if err != nil {
		return u, err
	}
	return u, nil
}

func (m *MySQL) PostRoomMsg(u User, msg string) error {
	time := time.Now()
	sqlStr := "Insert into chat_room values(?,?,?)"
	_, err := m.db.Exec(sqlStr, u.NickName, msg, time)
	fmt.Println("post")
	if err != nil {
		return err
	}
	return nil
}

func (m *MySQL) GetRoomMsg() ([]Msg, error) {
	sqlStr := "select fro,content,time from chat_room"
	rows, err := m.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var all = make([]Msg, 0)
	for rows.Next() {
		var v Msg
		rows.Scan(&v.From, &v.Content, &v.Time)
		all = append(all, v)
	}
	return all, nil
}
