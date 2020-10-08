package main

import (
	"fmt"
	"net/http"

	"github.com/vidu171/clean-architecture-go/infrastructure/db"

	"github.com/vidu171/clean-architecture-go/usecases"

	"github.com/vidu171/clean-architecture-go/interface/controllers"

	"github.com/vidu171/clean-architecture-go/infrastructure/logger"
	"github.com/vidu171/clean-architecture-go/infrastructure/router"
	"github.com/vidu171/clean-architecture-go/interface/repository"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	Logger                   = logger.Logger{}
	dbHandler  db.DBHandler
)

func getBookController() controllers.BookController {
	bookRepo := repository.NewBookRepo(dbHandler)
	bookInteractor := usecases.NewBookInteractor(bookRepo, Logger)
	bookController := controllers.NewBookController(bookInteractor)
	return *bookController
}

func getAuthorController() controllers.AuthorController {
	authorRepo := repository.NewAuthorRepo(dbHandler)
	authorInteractor := usecases.NewAuthorInteractor(authorRepo, Logger)
	authorController := controllers.NewAuthorController(authorInteractor)
	return *authorController
}

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "App is up and running..")
	})
	var err error
	dbHandler, err = db.NewDBHandler("mongodb://localhost:27017", "bookstore")
	if err != nil {
		Logger.Log("Unable to connect to the DataBase")
		return
	}
	bookController := getBookController()
	authorController := getAuthorController()
	httpRouter.POST("/book/add", bookController.Add)
	httpRouter.GET("/book", bookController.FindAll)
	httpRouter.POST("/author/add", authorController.Add)
	httpRouter.SERVE(":8000")
}
