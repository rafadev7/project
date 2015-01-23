package main

import (
	"time"
	//"database/sql"
	"log"

	"github.com/dchest/uniuri"
	"github.com/extemporalgenome/slug"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	dbname = "project"
	dbuser = "app"
	dbpass = "SecretPassword!"
)

type DB struct {
	*sqlx.DB
}

// The only one DB instance
var db *DB = &DB{}

func init() {
	log.Println("Connecting to database...")
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	//var err error
	dbx, err := sqlx.Open("mysql", dbuser+":"+dbpass+"@/"+dbname+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalf("DB init Error:", err.Error())
	}
	//defer db.Close() // I DUNNO IF IT WORKS HERE, LETS TEST
	log.Println("Database connected!")
	db.DB = dbx

	log.Println("Start routine to create the default values of our datas ...")

	checkAndCreateDefaultPic(db)

	checkAndCreateDefaultImage(db)

	checkAndCreateAnonymousUser(db)

	checkAndCreateCategories(db)

	log.Println("All default values has been created.")
}

func checkAndCreateCategories(db *DB) {
	for i, categoryName := range categoryList {
		categorySlug := slug.Slug(categoryName)
		count := 0
		err := db.Get(&count, "select count(*) from category where categoryslug=?", categorySlug)
		if err != nil {
			log.Fatalf("Error searching for the category with categorySlug %s. Err: %s\n", categorySlug, err.Error())
		}

		if count == 0 {

			categoryId := uniuri.NewLen(20)

			if i == 0 { // "Sem Categoria" is my default category
				categoryId = "default"
			}

			_, err := db.Exec("INSERT INTO category (categoryid, categoryname, categoryslug, likecount) VALUES (?,?,?,?)",
				categoryId, categoryName, categorySlug, 0)
			if err != nil {
				log.Fatalf("Error when creating the category %s . Err: %s\n", categoryName, err)
			} else {
				log.Printf("Category %s created!\n", categoryName)
			}
		}
	}
}

func checkAndCreateDefaultPic(db *DB) {
	// Adding default value to pic
	// pic/default.png !
	count := 0
	err := db.Get(&count, "select count(*) from pic where picid=?", "default")
	if err != nil {
		log.Fatalf("Error searching for the default pic. Err: %s\n", err.Error())
	}
	if err == nil {
		if count == 0 {
			_, err := db.Exec("INSERT INTO pic (picid, creation, deleted) VALUES (?,?,?)",
				"default", time.Now(), false)
			if err != nil {
				log.Printf("Error creating default pic. Err: %s\n", err)
			} else {
				log.Println("Default pic created!")
			}
		}
	} else {
		log.Printf("Error searching for default pic. %s\n", err)
	}
}

func checkAndCreateDefaultImage(db *DB) {
	// Adding default value to image
	// img/default-small.png
	// img/default-medium.png
	// img/default-large.png
	count := 0
	err := db.Get(&count, "select count(*) from image where imageid=?", "default")
	if err != nil {
		log.Fatalf("Error searching for the default images. Err: %s\n", err.Error())
	}
	if err == nil {
		if count == 0 {
			_, err := db.Exec("INSERT INTO image (imageid, creation, deleted) VALUES (?,?,?)",
				"default", time.Now(), false)
			if err != nil {
				log.Printf("Error creating default image. Err: %s\n", err)
			} else {
				log.Printf("Default image created!\n")
			}
		}
	} else {
		log.Printf("Error searching for default image. %s\n", err)
	}
}

func checkAndCreateAnonymousUser(db *DB) {
	count := 0
	err := db.Get(&count, "select count(*) from user where userid=?", "anonymous")
	if err != nil {
		log.Fatalf("Error searching for the default images. Err: %s\n", err.Error())
	}
	if err == nil {
		if count == 0 {

			_, err := db.Exec("INSERT INTO user (userid, username, picid, fullname, likecount, creation, lastupdate, deleted, admin) VALUES (?,?,?,?,?,?,?,?,?)",
				"anonymous", "Anonimo", "default", "Usuario Anonimo", 0, time.Now(), time.Now(), false, false)
			if err != nil {
				log.Printf("Error creating the anonymous user. Err: %s\n", err)
			} else {
				log.Println("User anonymous created!")
			}
		}
	} else {
		log.Printf("Error searching for user anonymous. %s\n", err)
	}
}

var categoryList = []string{
	"Sem categoria",
	"Animais",
	"Arte e Cultura",
	"Beleza e Estilo",
	"Carros e Motos",
	"Casa e Decoração",
	"Ciência e Tecnologia",
	"Comidas e Bebidas",
	"Crianças",
	"Curiosidades",
	"Downloads",
	"Educação",
	"Entretenimento",
	"Esporte",
	"Eventos",
	"Família",
	"Filmes",
	"Fotos",
	"Futebol",
	"Humor",
	"Internacional",
	"Internet",
	"Jogos",
	"Livro",
	"Meio ambiente",
	"Mulher",
	"Música",
	"Negócios",
	"Notícias",
	"Pessoas e Blogs",
	"Política",
	"Saúde",
	"Turismo e Viagem",
	"Vídeos",
}
