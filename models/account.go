package models

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/yuriygr/go-mlh/formatter"
	"github.com/yuriygr/go-posledstvie/utils"
)

// Account - Структура пользователя
type Account struct {
	ID       uint32 `json:"account_id"  db:"a.account_id"`
	Email    string `json:"email"       db:"a.email"`
	Password string `json:"-"           db:"a.password"`
	Status   bool   `json:"-"           db:"a.status"`
	Created  string `json:"-"           db:"a.created"`

	Subsite *User `json:"-"          db:"s"`
}

// Bind - Bind HTTP request data and validate it
func (c *Account) Bind(r *http.Request) error {

	if r.FormValue("name") == "" {
		return errors.New("name must be filled")
	}

	if r.FormValue("email") == "" {
		return errors.New("email must be filled")
	}

	if r.FormValue("password") == "" {
		return errors.New("password must be filled")
	}

	name := formatter.EscapeString(r.FormValue("name"))
	email := formatter.EscapeString(r.FormValue("email"))
	password, err := utils.HashPassword(r.FormValue("password"))
	if err != nil {
		return err
	}

	c.Email = email
	c.Password = password
	c.Status = true

	fmt.Println(name)

	//c.Subsite.Name = name
	//c.Subsite.Slug = name
	return nil
}
