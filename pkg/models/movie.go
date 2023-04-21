package models

import (
	"fmt"
	"movielist/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Title  string `gorm:"" json:"title"`
	Rating string `json:"rating"`
}

func init() {
	config.Connect()
	fmt.Println("database connected successfully")
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func (b *Movie) CreateMovie() *Movie {
	db.Create(&b)
	return b
}

func GetMovie() []Movie {
	var movies []Movie
	db.Find(&movies)
	return movies
}

func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db := db.Where("Id = ?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(Id int64) Movie {
	var movie Movie
	db.Where("Id=?", Id).Delete(movie)
	return movie
}
