package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type SigninToken struct {
    Token string `json:"token"`
    ExpiresIn time.Time `json:"expires_in"`
}

type User struct {
    ID int64
    Name string
    Email string
    Version int
    SigninToken SigninToken
    CreatedAt time.Time
    UpdatedAt *time.Time
}

func NewUser(
    name string,
    email string,
) *User {
    return &User{
        Name: name,
        Email: email,
        CreatedAt: time.Now(),
    }
}

func (s SigninToken) Value() (driver.Value, error) {
    return json.Marshal(s)
}

func (s *SigninToken) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(b, &s)
}
