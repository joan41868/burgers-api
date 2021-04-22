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

func handleHttpErr(err error, writer http.ResponseWriter){
	log.Println(err)
	writer.Write([]byte(err.Error()))
	return
}

func getRequestQuery(request *http.Request) (string,  int, int){
	name := request.URL.Query().Get("name")
	pageAsStr := request.URL.Query().Get("pageNum")
	pageNum, err := strconv.Atoi(pageAsStr)
	if err != nil{
		return name,  1, 15
	}

	perPageStr := request.URL.Query().Get("perPage")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil{
		return name,  pageNum, 15
	}
	log.Println("Params", name, pageNum, perPage)

	return name,  pageNum, perPage
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
			return
		}else {
			var saved *model.Burger
			saved, err = app.repository.CreateOne(&burger)
			if err != nil {
				handleHttpErr(err, writer)
				return
			}
			json.NewEncoder(writer).Encode(saved)
		}
	}).Methods("POST")

	//app.router.HandleFunc("/burgers", func(writer http.ResponseWriter, request *http.Request) {
	//	name := request.URL.Query().Get("name")
	//	found, err := app.repository.GetByName(name)
	//	if err != nil{
	//		handleHttpErr(err, writer)
	//		return
	//	}else{
	//		json.NewEncoder(writer).Encode(found)
	//	}
	//}).Methods("GET")



	app.router.HandleFunc("/burgers/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idAsStr := mux.Vars(request)["id"]
		id, err := strconv.Atoi(idAsStr)
		if err != nil{
			handleHttpErr(err, writer)
			return
		}else{
			found, err := app.repository.GetByID(id)
			if err != nil{
				handleHttpErr(err, writer)
				return
			}else{
				json.NewEncoder(writer).Encode(found)
			}
		}

	}).Methods("GET")


	app.router.HandleFunc("/burgers", func(writer http.ResponseWriter, request *http.Request) {
		burgerName,_, _ := getRequestQuery(request)

		countBurgers := app.repository.Count()
		lst := pagination.GetPaginatedListFromRequest(request, countBurgers)
		burgers, err  := app.repository.GetPaginatedByName(burgerName, uint(lst.Page), uint(lst.PerPage))
		if err != nil{
			log.Fatal(err)
		}
		json.NewEncoder(writer).Encode(burgers)
		//if perPage == 0 {
		//	perPage = 15
		//}
		//if pageNum == 0{
		//	pageNum = 1
		//}
		//if burgerName == ""{
		//	burgers, err := app.repository.GetPaginated(uint(pageNum), uint(perPage))
		//	if err != nil{
		//		handleHttpErr(err, writer)
		//	}
		//	err = json.NewEncoder(writer).Encode(burgers)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//}else{
		//	burgers, err := app.repository.GetByName(burgerName)
		//	if err != nil{
		//		handleHttpErr(err, writer)
		//	}
		//	burgers = repository.PaginateSliceOfBurgers(burgers, uint(pageNum), uint(perPage))
		//	err = json.NewEncoder(writer).Encode(burgers)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//}

	}).Methods("GET")
}
