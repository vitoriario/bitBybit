package main

import (
	db "github.com/masisiliani/bitBybit/db"
	"github.com/masisiliani/bitBybit/api/router"
	"github.com/masisiliani/bitBybit/pkg/user"
	"github.com/masisiliani/bitBybit/pkg/post"
	"os"
	"net/http"
	"fmt"
)



func main(){
	if err := db.InitDatabase(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	database, err := db.Connect_SQL()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	r := router.Router{
		UserController: user.UserController{
			Repository: &db.MySQLRepository{
				DB: database,
			},
		},
		PostController: post.PostController{
			Repository: &db.MySQLRepository{
				DB: database,
			},
		},
	}
	http.HandleFunc("/login", r.Login)
	http.HandleFunc("/register", r.Register)
	http.HandleFunc("/changePassword", r.ChangePassword)

	fmt.Println("Starting server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
        os.Exit(1)
    }
}