package application

import (
	"burger-api/domain/model"
	"burger-api/domain/pagination"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func propagateErrorToClient(err error, writer http.ResponseWriter) {
	log.Println(err)
	writer.Write([]byte(err.Error()))
	return
}

// InitRoutes attaches handlers to app.router (*mux.Router)
func (app *Application) InitRoutes() {

	app.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello limes!"))
	})

	app.router.HandleFunc("/burger", func(writer http.ResponseWriter, request *http.Request) {
		var burger model.Burger
		err := json.NewDecoder(request.Body).Decode(&burger)
		if err != nil {
			propagateErrorToClient(err, writer)
			return
		} else {
			var saved *model.Burger
			saved, err = app.repository.CreateOne(&burger)
			if err != nil {
				propagateErrorToClient(err, writer)
				return
			}
			json.NewEncoder(writer).Encode(saved)
		}
	}).Methods("POST")

	app.router.HandleFunc("/burgers/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idAsStr := mux.Vars(request)["id"]
		id, err := strconv.Atoi(idAsStr)
		if err != nil {
			propagateErrorToClient(err, writer)
			return
		} else {
			found, err := app.repository.GetByID(id)
			if err != nil {
				propagateErrorToClient(err, writer)
				return
			} else {
				json.NewEncoder(writer).Encode(found)
			}
		}

	}).Methods("GET")

	// query stuff
	app.router.HandleFunc("/burgers", func(writer http.ResponseWriter, request *http.Request) {
		burgerName := request.URL.Query().Get("burger_name")
		countBurgers := app.repository.Count()
		lst := pagination.GetPaginatedListFromRequest(request, countBurgers)
		// handle missing burger name
		if burgerName == "" {
			burgers, err := app.repository.GetPaginated(uint(lst.Page), uint(lst.PerPage))
			if err != nil {
				propagateErrorToClient(err, writer)
				return
			}
			json.NewEncoder(writer).Encode(burgers)
			return
		}

		burgers, err := app.repository.GetPaginatedByName(burgerName, uint(lst.Page), uint(lst.PerPage))
		if err != nil {
			propagateErrorToClient(err, writer)
			return
		}
		json.NewEncoder(writer).Encode(burgers)
	}).Methods("GET")
}
