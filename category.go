package main

import (
	"github.com/resourcerest/api"
)

type Category struct {
	CategoryID   string `db:"categoryid" json:"categoryid"`     // *PK max: 20
	CategoryName string `db:"categoryname" json:"categoryname"` //  max: 20
	CategorySlug string `db:"categoryslug" json:"categoryslug"` // *UQ max: 20 Index
	LikeCount    int    `db:"likecount" json:"likecount"`       // default: 0

	Contents *Contents
}

type Categories []Category

func (c *Category) Init(id *api.ID, db *DB) (*Category, error) {
	err := db.Get(c, "SELECT * FROM category WHERE categoryslug=?", id.String())
	if err != nil {
		return nil, err
	}

	return c, nil
}
