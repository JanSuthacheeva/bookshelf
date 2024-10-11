package models

import (
  "database/sql"
  "time"
)

type User struct {
  ID              int
  Name            string
  Email           string
  HashedPassword  string
  Created         time.Time
}

type UserModel struct {
  DB *sql.DB
}

func (m *UserModel) Insert(name, email, hashedPassword string, created time.Time) (int, error) {
  return 0, nil
}

func (m *UserModel) Get(id int) (User, error) {
  return User{}, nil
}
