package mocks

import (
	"burger-api/domain/model"
	"burger-api/domain/pagination"
	"errors"
	"log"
	"strings"
)

type BurgerRepository struct {
	Burgers     []*model.Burger
	ShouldError bool
}

func (repo *BurgerRepository) GetByID(id int) (*model.Burger, error) {
	if repo.ShouldError {
		return nil, errors.New("")
	}
	for _, u := range repo.Burgers {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("burger not found")
}

func (repo *BurgerRepository) CreateOne(mdl *model.Burger) (*model.Burger, error) {
	if repo.ShouldError {
		return nil, errors.New("saving failed")
	}
	repo.Burgers = append(repo.Burgers, mdl)
	return mdl, nil
}

func (repo *BurgerRepository) GetByName(name string) ([]*model.Burger, error) {
	if repo.ShouldError {
		return nil, errors.New("i should error")
	}
	var slc []*model.Burger
	for _, v := range repo.Burgers {
		if strings.Contains(v.Name, name) {
			slc = append(slc, v)
		}
	}
	return slc, nil
}

func (repo *BurgerRepository) GetPaginated(pageNum, perPage uint) ([]*model.Burger, error) {

	pgn := pagination.PaginatedList{
		PerPage: int(perPage),
		Page: int(pageNum),
	}
	burgersLen := uint(len(repo.Burgers))
	mp := make(map[uint][]*model.Burger)
	var cnt uint = 1 // actual page number
	for {
		burgersLen = burgersLen - perPage
		mp[cnt] = repo.Burgers[pgn.Offset():pgn.Limit()]
		if burgersLen <= 0 {
			break
		}
		cnt = cnt + 1
	}
	log.Println(mp)
	return mp[pageNum], nil
}

func (repo *BurgerRepository) Count() int {
	return len(repo.Burgers)
}

func (repo *BurgerRepository) GetPaginatedByName(name string, pageNum, perPage uint) ([]*model.Burger, error) {

	pgn := pagination.PaginatedList{
		PerPage: int(perPage),
		Page: int(pageNum),
	}
	burgersLen := uint(len(repo.Burgers))
	mp := make(map[uint][]*model.Burger)
	var cnt uint = 1 // actual page number
	for {
		burgersLen = burgersLen - perPage
		mp[cnt] = repo.Burgers[pgn.Offset():pgn.Limit()]
		if burgersLen <= 0 {
			break
		}
		cnt = cnt + 1
	}
	var slc []*model.Burger
	for _, v := range mp[pageNum] {
		if strings.Contains(v.Name, name) {
			slc = append(slc, v)
		}
	}
	return slc, nil
}
