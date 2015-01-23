package main

import (
	"time"
)

type User struct {
	UserID     string    `db:"userid" json:"userid"`       // *PK max: 20
	UserName   string    `db:"username" json:"username"`   // *UQ max: 20
	PicID      string    `db:"picid" json:"picid"`         // *PK max: 20
	FullName   string    `db:"fullname" json:"fullname"`   // *UQ max: 20
	LikeCount  int       `db:"likecount" json:"likecount"` // default: 0
	Creation   time.Time `db:"creation" json:"-"`          // *NN
	LastUpdate time.Time `db:"lastupdate" json:"-"`        // *NN
	Deleted    bool      `db:"deleted" json:"-"`           // default: 0
	Admin      bool      `db:"admin" json:"-"`             // default: 0
}
