package models

type (
	userDB struct {
		id         string
		data       *User
		originData *User
	}
)

type (
	UserColumn string
)

const (
	UserColumnId        UserColumn = "Id"
	UserColumnPassword  UserColumn = "Password"
	UserColumnPoint     UserColumn = "Point"
	UserColumnUsername  UserColumn = "Username"
	UserColumnJoinAt    UserColumn = "JoinAt"
	UserColumnCreatedAt UserColumn = "CreatedAt"
	UserColumnUpdatedAt UserColumn = "UpdatedAt"
	UserColumnDeletedAt UserColumn = "DeletedAt"
)

func newUserDB(data *User) *userDB {
	return &userDB{
		data: data,
	}
}

func (x *userDB) Data() *User {
	return x.data
}
