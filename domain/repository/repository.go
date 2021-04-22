package repository

import (
	"burger-api/domain/model"
	"burger-api/domain/pagination"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// abstract

// concrete

type BurgerRepositoryImpl struct {
	Db *gorm.DB
}

// NewBurgerRepository creates a new instance of BurgerRepositoryImpl
func NewBurgerRepository(connStr string) (*BurgerRepositoryImpl, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	repo := BurgerRepositoryImpl{
		Db: db,
	}

	db.AutoMigrate(&model.Burger{})
	return &repo, nil
}

func (repo *BurgerRepositoryImpl) CreateOne(mdl *model.Burger) (*model.Burger, error) {
	tx := repo.Db.Create(mdl)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return mdl, nil
}

func (repo *BurgerRepositoryImpl) GetByName(name string) ([]*model.Burger, error) {
	var brg []*model.Burger
	tx := repo.Db.Where("name like ?", "%"+name+"%").Find(&brg)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return brg, nil
}

func (repo *BurgerRepositoryImpl) GetByID(id int) (*model.Burger, error) {
	var brg model.Burger
	tx := repo.Db.First(&brg, "id = ?", id)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return &brg, nil
}

func (repo *BurgerRepositoryImpl) GetPaginated(pageNum, perPage uint) ([]*model.Burger, error) {

	pgn := pagination.Paginator{
		Limit:   perPage,
		Offset:  (pageNum - 1) * perPage, // Offset 1 will start from the second row, this is why i did this hack
		PageNum: pageNum - 1,
	}
	var burgers []*model.Burger
	tx := repo.Db.Limit(int(pgn.Limit)).Offset(int(pgn.Offset)).Find(&burgers)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return burgers, nil
}

func (repo *BurgerRepositoryImpl) Count() int {
	var count int64
	repo.Db.Model(&model.Burger{}).Count(&count)
	return int(count)
}

func (repo *BurgerRepositoryImpl) GetPaginatedByName(name string, pageNum, perPage uint) ([]*model.Burger, error) {
	pgn := pagination.Paginator{
		Limit:   perPage,
		Offset:  (pageNum - 1) * perPage, // Offset 1 will start from the second row, this is why i did this hack
		PageNum: pageNum - 1,
	}
	var burgers []*model.Burger
	tx := repo.Db.Where("name like ?", "%"+name+"%").Limit(int(pgn.Limit)).Offset(int(pgn.Offset)).Find(&burgers)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return burgers, nil
}
