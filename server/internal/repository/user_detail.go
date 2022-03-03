package repository

import (
	"database/sql"
	"fmt"
)

type UserDetailEntity struct {
	Id        int64  `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserDetail interface {
	FindByEmail(email string) (*UserDetailEntity, error)
	FindAll() ([]UserDetailEntity, error)
}

type userDetail struct {
	db *sql.DB
}

func NewUserDetail(db *sql.DB) UserDetail {
	return &userDetail{
		db: db,
	}
}

func (repo userDetail) FindByEmail(email string) (*UserDetailEntity, error) {
	stmt, err := repo.db.Prepare("SELECT * FROM user_detail where email = ?")
	if err != nil {
		return nil, err
	}
	row, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		tmp := UserDetailEntity{}
		err = row.Scan(&tmp.Id, &tmp.Firstname, &tmp.Lastname, &tmp.Email)
		if err != nil {
			return nil, err
		}
		return &tmp, nil
	} else {
		return nil, nil
	}
}

func (repo userDetail) FindAll() ([]UserDetailEntity, error) {
	stmt, err := repo.db.Prepare("SELECT * FROM user_detail")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var r []UserDetailEntity
	for rows.Next() {
		var tmp UserDetailEntity
		err = rows.Scan(&tmp.Id, &tmp.Firstname, &tmp.Lastname, &tmp.Email)
		if err != nil {
			return nil, nil
		}
		r = append(r, tmp)
	}
	fmt.Printf("findAll: %v\n", len(r))
	return r, nil
}
