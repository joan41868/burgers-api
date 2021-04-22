package application

import (
	"burger-api/domain/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func handleHttpErr(err error, writer http.ResponseWriter){
	log.Println(err)
	writer.Write([]byte(err.Error()))
	return
}

func (app *Application) InitRoutes() {

	app.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello"))
	})

	app.router.HandleFunc("/burger", func(writer http.ResponseWriter, request *http.Request) {
		var burger model.Burger
		err := json.NewDecoder(request.Body).Decode(&burger)
		if err != nil{
			handleHttpErr(err, writer)
		}else {
			var saved *model.Burger
			saved, err = app.repository.CreateOne(&burger)
			if err != nil {
				handleHttpErr(err, writer)
			}
			json.NewEncoder(writer).Encode(saved)
		}
	}).Methods("POST")

	app.router.HandleFunc("/burger/name/{name}", func(writer http.ResponseWriter, request *http.Request) {
		name := mux.Vars(request)["name"]
		found, err := app.repository.GetByName(name)
		if err != nil{
			handleHttpErr(err, writer)
		}else{
			json.NewEncoder(writer).Encode(found)
		}
	}).Methods("GET")

	app.router.HandleFunc("/burger/id/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idAsStr := mux.Vars(request)["id"]
		id, err := strconv.Atoi(idAsStr)
		if err != nil{
			handleHttpErr(err, writer)
		}else{
			found, err := app.repository.GetByID(id)
			if err != nil{
				handleHttpErr(err, writer)
			}else{
				json.NewEncoder(writer).Encode(found)
			}
		}

	}).Methods("GET")


	app.router.HandleFunc("/burgers/pageNum/{pageNum}/perPage/{perPage}", func(writer http.ResponseWriter, request *http.Request) {
		pageAsStr := mux.Vars(request)["pageNum"]
		pageNum, err := strconv.Atoi(pageAsStr)
		if err != nil{
			handleHttpErr(err, writer)
		}

		perPageStr := mux.Vars(request)["perPage"]
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil{
			handleHttpErr(err, writer)
		}

		burgers, err := app.repository.GetPaginated(uint(pageNum), uint(perPage))
		if err != nil{
			handleHttpErr(err, writer)
		}
		json.NewEncoder(writer).Encode(burgers)
	}).Methods("GET")
}
