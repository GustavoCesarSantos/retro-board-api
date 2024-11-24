package domain

import "time"

type User struct {
    ID int64
    Name string
    Email string
    Version int
    CreatedAt time.Time
}

func NewUser(
    id int64,
    name string,
    email string,
) *User {
    return &User{
        ID: id,
        Name: name,
        Email: email,
        Version: 1,
        CreatedAt: time.Now(),
    }
}
