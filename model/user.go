package model

import "encoding/json"

type Users []User

func (o *User) ToJsonRaw() json.RawMessage {
	b, _ := json.Marshal(o)
	return b
}

type User struct {
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}

func (o *User) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *User) IsValid() {
	// check the User is valid ....
}
