package models

import (
	"errors"
	"time"

	"github.com/4nmt/team-manage-api/db"
	"github.com/4nmt/team-manage-api/forms"
)

//User ...
type User struct {
	ID        int    `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

//UserModel ...
type UserModel struct{}

//Create ...
func (m UserModel) Create(form forms.UserForm) (user User, err error) {
	getDb := db.GetDB()

	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.users WHERE email=LOWER($1) LIMIT 1", form.Email)
	if err != nil {
		return user, err
	}

	if checkUser > 0 {
		return user, errors.New("User exists")
	}
	res, err := getDb.Exec("INSERT INTO public.users(email, name, updated_at, created_at) VALUES($1, $2, $3, $4) RETURNING id", form.Email, form.Name, time.Now().Unix(), time.Now().Unix())

	if res != nil && err == nil {
		err = getDb.SelectOne(&user, "SELECT id, email, name, updated_at, created_at FROM public.users WHERE email=LOWER($1) LIMIT 1", form.Email)
		if err == nil {
			return user, nil
		}
	}

	return user, errors.New("Not registered")
}

//Delete ...
func (m UserModel) Delete(userID int) error {
	_, err := m.One(userID)
	if err != nil {
		return err
	}

	if _, err = db.GetDB().Exec("DELETE FROM public.user_project WHERE user_id=$1", userID); err != nil {
		return err
	}
	_, err = db.GetDB().Exec("DELETE FROM public.users WHERE id=$1", userID)

	return err
}

//One ...
func (m UserModel) One(userID int) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.users WHERE id=$1", userID)
	return user, err
}

//All ...
func (m UserModel) All() (users []User, err error) {
	_, err = db.GetDB().Select(&users, "SELECT id, email, name FROM public.users")
	return users, err
}
