package mocks

import (
	"burger-api/domain/model"
	"errors"
	"log"
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

func (repo *BurgerRepository) GetByName(name string) (*model.Burger, error) {
	if repo.ShouldError {
		return nil, errors.New("i should error")
	}

	for _, u := range repo.Burgers {
		if u.Name == name {
			return u, nil
		}
	}
	return nil, errors.New("burger not found")
}

func (repo *BurgerRepository) GetPaginated(pageNum, perPage uint) ([]*model.Burger, error){
	type paginator struct {
		Limit uint // PerPage
		Offset uint
		PageNum uint
	}
	pgn := paginator{
		Limit:   perPage,
		Offset:  (pageNum-1)*perPage, // Offset 1 will start from the second row, this is why i did this hack
		PageNum: pageNum-1,
	}
	burgersLen := uint(len(repo.Burgers))
	mp := make(map[uint] []*model.Burger)
	var cnt uint = 1 // actual page number
	for {
		burgersLen = burgersLen - perPage
		mp[cnt] = repo.Burgers[pgn.Offset:pgn.Limit]
		if burgersLen <= 0{
			break
		}
		cnt = cnt + 1
	}
	log.Println(mp)
	return mp[pageNum], nil
}
