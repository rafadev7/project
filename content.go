package main

import (
	"encoding/json"
	"github.com/dchest/uniuri"
	"log"
	"net/http"
	"time"

	"github.com/resourcerest/api"
)

type Content struct {
	ContentID    string    `db:"contentid" json:"contentid"`       // *PK max: 20
	UrlID        string    `db:"urlid" json:"urlid"`               // *FK max: 5
	CategoryID   string    `db:"categoryid" json:"categoryid"`     // *FK max: 20
	Title        string    `db:"title" json:"title"`               // *NN  max: 255 (250)
	Slug         string    `db:"slug" json:"slug"`                 // *NN  max: 255 (250)
	Description  string    `db:"description" json:"description"`   // *NN  max: 255
	Host         string    `db:"host" json:"host"`                 // *NN  max: 20
	UserID       string    `db:"userid" json:"userid"`             // *FK max: 20
	ImageID      string    `db:"imageid" json:"imageid"`           // *FK max: 20
	ImageMaxSize string    `db:"imagemaxsize" json:"imagemaxsize"` // *NN ENUM('small', 'medium', 'large')
	LikeCount    int       `db:"likecount" json:"likecount"`       // default: 0
	CommentCount int       `db:"commentcount" json:"commentcount"` // default: 0
	Ranking      int       `db:"ranking" json:"ranking"`           // default: 0
	Creation     time.Time `db:"creation" json:"creation"`         // *NN
	LastUpdate   time.Time `db:"lastupdate" json:"lastupdate"`     // *NN
	Deleted      bool      `db:"deleted" json:"-"`                 // default: 0
	ILike        bool      `db:"-" json:"ilike"`                   //  *NOT IN THE VIEW
}

type Contents []Content

func (c *Content) Init(id *api.ID) {
	c.Title = "Testing"
	c.ContentID = id.String()
}

func (c *Content) GET() *Content {
	return c
}

type ContentPOST struct {
	URL string
	Content
}

func (_ *Contents) POST(req *http.Request, db *DB, cat *Category, errs []error) (*ContentPOST, []error) {

	if len(errs) > 0 {
		//log.Fatalln("Error selecting the category:", err.Error())
		return nil, errs
	}

	log.Println("CATEGORY:", cat.CategorySlug)

	decoder := json.NewDecoder(req.Body)
	var data ContentPOST
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	urlID := uniuri.NewLen(5)

	res, err := db.Exec("INSERT INTO url (urlid, fullurl, userid, creation, viewcount) VALUES (?,?,?,?,?)",
		urlID, data.URL, "anonymous", time.Now(), 0)
	if err != nil {
		log.Fatalln("Error inserting url:", err.Error())
	}

	rw, err := res.RowsAffected()
	if err != nil {
		log.Fatalln("Error Getting Rows Affected", err.Error())
	}

	if rw < 1 {
		log.Fatalln("URL not inserted!")
	}

	/*
		c := &Content{
		ContentID: uniuri.NewLen(5),
		UrlID: urlID,
		CategoryID: cat.CategoryID,
		Title: ,
		Slug: ,
		Description: ,
		Host: ,
		UserID: ,
		ImageID: ,
		ImageMaxSize: ,
		LikeCount: ,
		CommentCount: ,
		Ranking: ,
		Creation: ,
		LastUpdate: ,
		Deleted: ,
		ILike: ,
		}

		_, err = db.Exec("INSERT INTO content (contentid, urlid, categoryid, title, slug, description, host, userid, imageid, imagemaxsize, likecount, commentcount, ranking, creation, lastupdate) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);",
			, data.Title, data.Slug+uniuri.NewLen(5), data.Description, "default", "anonymous", "default", urlID)
		if err != nil {
			log.Fatalln("Error inserting content:", err.Error())
		}
	*/

	return &data, nil
}
