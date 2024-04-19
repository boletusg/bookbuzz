package model

import (
	"bookbuzz/app/server"
	"time"
)

type User struct {
	UserId        int       `db:"id_user"`
	LoginUser     string    `db:"login_user"`
	PasswordUser  string    `db:"password_user"`
	NameUser      string    `db:"name_user"`
	PhoneUser     string    `db:"phone_user"`
	Nickname      string    `db:"nickname"`
	StatusUser    string    `db:"status_user"`
	BiographyUser string    `db:"biography_user"`
	Links         string    `db:"links"`
	AvatarUser    []byte    `db:"avatar_user"`
	DateRegUser   time.Time `db:"datereg_user"`
}

type UserSummary struct {
	LoginUser string `db:"login_user"`
	NameUser  string `db:"name_user"`
}

func GetAllUsers() (users []UserSummary, err error) {
	query := `SELECT login_user, name_user FROM users`
	rows, err := server.Db.Queryx(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := UserSummary{}
		err = rows.StructScan(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func NewUser(name, login string) *User {
	return &User{NameUser: name, LoginUser: login}
}

func GetUserById(userId string) (u User, err error) {
	query := `SELECT * FROM users WHERE id_user = ?`
	err = server.Db.Get(&u, query, userId)
	return
}

func (u *User) Add() (err error) {
	query := `INSERT INTO users (name_user, login_user) VALUES (?, ?)`
	_, err = server.Db.Exec(query, u.NameUser, u.LoginUser)
	return
}
