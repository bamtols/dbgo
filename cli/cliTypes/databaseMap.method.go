package cliTypes

import (
	"github.com/bamtols/dbgo/func/errFn"
)

func (x *DatabaseMap) GetDatabase(nm DatabaseNm) (*Database, error) {
	db, isOk := (*x)[nm]
	if !isOk {
		return nil, errFn.BindF(ErrNotFound, "Database : databaseNm=%s", nm.String())
	}
	return &db, nil
}

func (x *DatabaseMap) DefaultDBKey() (DatabaseNm, error) {
	if len(*x) != 1 {
		return "", errFn.BindF(ErrNotFound, "DefaultDatabase")
	}

	for key := range *x {
		return key, nil
	}

	return "", errFn.BindF(ErrNotFound, "DefaultDatabase")
}

func (x *DatabaseMap) DefaultDB() (*Database, error) {
	key, err := x.DefaultDBKey()
	if err != nil {
		return nil, err
	}
	db := (*x)[key]
	return &db, nil
}
