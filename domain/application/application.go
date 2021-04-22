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
	router     *mux.Router
	repository repository.BurgerRepository
}

func NewApplication(connStr string) *Application {
	repo, err := repository.NewBurgerRepository(connStr)
	if err != nil {
		log.Fatal(err)
	}
	app := Application{
		router:     mux.NewRouter(),
		repository: repo,
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
	// routes
	app.InitRoutes()
	app.router.PathPrefix("/").Handler(http.FileServer(http.Dir("domain/static/"))) // serve files under the static dir - images & docs
	// middlewares
	app.router.Use(ratelimit.Handle)
	app.router.Use(middleware.ContentTypeMiddleware) // attaches a Content-type : application/json header
	app.router.Use(middleware.CorsMiddleware)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, app.router))
}
