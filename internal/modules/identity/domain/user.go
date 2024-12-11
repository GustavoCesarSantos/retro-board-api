package domain

import "time"

type User struct {
    ID int64
    Name string
    Email string
    Version int
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
