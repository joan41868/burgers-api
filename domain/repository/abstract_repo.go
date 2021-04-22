package repository

import "burger-api/domain/model"

type BurgerRepository interface {
	GetByID(id int) (*model.Burger, error)
	GetByName(name string) ([]*model.Burger, error)
	GetPaginated(pageNum, perPage uint) ([]*model.Burger, error)
	GetPaginatedByName(name string, pageNum, perPage uint) ([]*model.Burger, error)

	CreateOne(*model.Burger) (*model.Burger, error)
	Count() int
}

func GetById(id int, r BurgerRepository) (*model.Burger, error) {
	return r.GetByID(id)
}

func GetByName(name string, r BurgerRepository) ([]*model.Burger, error) {
	return r.GetByName(name)
}

func CreateOne(mdl *model.Burger, r BurgerRepository) (*model.Burger, error) {
	return r.CreateOne(mdl)
}

func GetPaginated(pageNum, perPage uint, r BurgerRepository) ([]*model.Burger, error) {
	return r.GetPaginated(pageNum, perPage)
}

func Count(r BurgerRepository) int {
	return r.Count()
}

func GetPaginatedByName(name string, pageNum, perPage uint, r BurgerRepository) ([]*model.Burger, error) {
	return r.GetPaginatedByName(name, pageNum, perPage)
}
