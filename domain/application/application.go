package application

import (
	"burger-api/domain/middleware"
	"burger-api/domain/repository"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/sethvargo/go-limiter/memorystore"
)

type Application struct {
	Router     *mux.Router
	Repository repository.BurgerRepository
}

func NewApplication(connStr string) *Application {
	repo, err := repository.NewBurgerRepository(connStr)
	if err != nil {
		log.Fatal(err)
	}
	app := Application{
		Router:     mux.NewRouter(),
		Repository: repo,
	}
	return &app
}

func (app *Application) Start(port string) {
	store, err := memorystore.New(&memorystore.Config{
		// Number of tokens allowed per interval.
		Tokens: 15,

		// Interval until tokens reset.
		Interval: time.Minute,
	})
	if err != nil {
		log.Fatal(err)
	}
	ratelimit, err := httplimit.NewMiddleware(store, httplimit.IPKeyFunc())
	if err != nil {
		log.Fatal(err)
	}
	app.InitRoutes()
	app.Router.Use(ratelimit.Handle)
	app.Router.Use(middleware.ContentTypeMiddleware)
	app.Router.Use(middleware.CorsMiddleware)
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}
