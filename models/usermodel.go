package models

import (
	"database/sql"
	"fmt"

	"github.com/fahrigunadi/go-crud/config"
	"github.com/fahrigunadi/go-crud/entities"
)

type UserModel struct {
	conn *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConnection()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		conn: conn,
	}
}

func (u *UserModel) FindAll() ([]entities.User, error) {
	rows, err := u.conn.Query("select * from users")
	if err != nil {
		fmt.Println(err)
		return []entities.User{}, err
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.Phone)

		users = append(users, user)
	}

	return users, nil
}

func (u *UserModel) Create(user entities.User) bool {
	result, err := u.conn.Exec("insert into users (name, username, phone) values (?, ?, ?)", user.Name, user.Username, user.Phone)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (u *UserModel) Find(id int64, user *entities.User) error {
	return u.conn.QueryRow("select * from users where id = ?", id).Scan(&user.Id, &user.Name, &user.Username, &user.Phone)
}

func (u *UserModel) Update(user entities.User) error {
	_, err := u.conn.Exec("update users set name = ?, username = ?, phone = ? where id = ?", user.Name, user.Username, user.Phone, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserModel) Delete(id int64) {
	u.conn.Exec("delete from users where id = ?", id)
}
