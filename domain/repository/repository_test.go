package repository

import (
	"burger-api/domain/model"
	"burger-api/helper/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByName(t *testing.T) {
	burgerName := "TheBurger"
	burger := &model.Burger{
		ID:   1,
		Name: burgerName,
	}
	// create repo with 1 burger inside - the burger from above
	repo := &mocks.BurgerRepository{Burgers: []*model.Burger{burger}}
	assert.Equal(t, len(repo.Burgers), 1, "Wtf, Repo should have been initialized")
	found, err := BurgerRepository.GetByName(repo, burgerName)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, len(found), 1, "Wtf, they should be the same")
}

func TestGetById(t *testing.T) {
	burgerName := "TheBurger"
	burger := &model.Burger{
		ID:   1,
		Name: burgerName,
	}
	repo := &mocks.BurgerRepository{Burgers: []*model.Burger{burger}}
	assert.Equal(t, len(repo.Burgers), 1, "Wtf, Repo should have been initialized")
	found, err := BurgerRepository.GetByID(repo, 1)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, found.ID, 1, "Wtf, it must be present")
}

func TestCreateOne(t *testing.T) {
	burgerName := "TheBurger"
	burger := &model.Burger{
		ID:   1,
		Name: burgerName,
	}
	repo := &mocks.BurgerRepository{Burgers: []*model.Burger{burger}}
	assert.Equal(t, len(repo.Burgers), 1, "Wtf, Repo should have been initialized")
	burger.ID = 2
	_, err := BurgerRepository.CreateOne(repo, burger)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 2, len(repo.Burgers))
}

func TestGetPaginated(t *testing.T) {
	burgerName := "TheBurger"
	burger := &model.Burger{
		ID:   1,
		Name: burgerName,
	}
	repo := &mocks.BurgerRepository{Burgers: []*model.Burger{burger}}
	assert.Equal(t, len(repo.Burgers), 1, "Wtf, Repo should have been initialized")
	burger.ID = 2
	_, err := BurgerRepository.CreateOne(repo, burger)
	if err != nil {
		panic(err)
	}

	paginated, err := BurgerRepository.GetPaginated(repo, 1, 1)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(paginated))

	paginated2, err := BurgerRepository.GetPaginated(repo, 1, 2) //nolint:typechecking
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 2, len(paginated2))
}
