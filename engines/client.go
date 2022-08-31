package engines

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

type (
	Client struct {
		root *sql.DB
	}

	INewClient struct {
		Username string
		Password string
		URL      string
		Params   *string
		Database string
		DBType   DBType
	}

	DBType string
)

const (
	DBTypeMySQL            DBType = "mysql"
	INewClientParamDefault        = "charset=utf8mb4&parseTime=True&loc=UTC"
)

func (x *DBType) String() string {
	return string(*x)
}

func (x *INewClient) host() string {
	host := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		x.Username,
		x.Password,
		x.URL,
		x.Database,
	)

	if x.Params != nil {
		host += fmt.Sprintf("?%s", *x.Params)
	}

	return host
}

func NewClient(i *INewClient) (*Client, error) {
	db, err := sql.Open(i.DBType.String(), i.host())
	if err != nil {
		return nil, err
	}

	res := &Client{
		root: db,
	}

	return res, nil
}
