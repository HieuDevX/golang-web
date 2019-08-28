package main

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type Product struct {
	ID      int `sql: ",pk"`
	Name    string
	Score   float64
	Country []string `pg:",array"`
	Price   float64
}

func main() {
	fmt.Println("hello")

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Database: "postgres",
		Addr:     "localhost:5432",
	})

	fmt.Println(db)

	db.Exec("CREATE SCHEMA IF NOT EXISTS test1")
	db.Exec("CREATE SCHEMA IF NOT EXISTS test2")

	type Post1 struct {
		TableName []byte `sql:"test1.post1"`
		ID        int
		name      string
	}

	type Post2 struct {
		TableName []byte `sql:"test2.post2"`
		ID        int
		name      string
	}

	var post1 Post1
	var post2 Post2

	err1 := db.CreateTable(&post1, &orm.CreateTableOptions{
		IfNotExists: true,
	})

	err2 := db.CreateTable(&post2, &orm.CreateTableOptions{
		IfNotExists: true,
	})

	// var product Product

	// err := db.CreateTable(&product, &orm.CreateTableOptions{
	// 	Temp:          false,
	// 	FKConstraints: true,
	// 	IfNotExists:   true,
	// })

	var postType1 = Post1{
		ID:   2,
		name: "hello",
	}

	var postType2 = Post2{
		ID:   3,
		name: "hello",
	}

	db.Insert(&postType1, &postType2)

	if err1 != nil || err2 != nil {
		panic(err1)
	}

	var post1s []Post1
	errQuery := db.Model(&post1s).Column("id", "name").Select()

	if errQuery != nil {
		panic(errQuery)
	}
	fmt.Println(post1s)

	defer db.Close()
}
