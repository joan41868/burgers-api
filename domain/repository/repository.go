package repository

import (
	"burger-api/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// abstract

type BurgerRepository interface {
	GetByID(id int) (*model.Burger, error)
	GetByName(name string) (*model.Burger, error)
	GetPaginated(pageNum, perPage uint) ([]*model.Burger, error)

	CreateOne(*model.Burger) (*model.Burger, error)
}

func GetById(id int, r BurgerRepository) (*model.Burger, error) {
	return r.GetByID(id)
}

func GetByName(name string, r BurgerRepository) (*model.Burger, error) {
	return r.GetByName(name)
}

func CreateOne(mdl *model.Burger, r BurgerRepository) (*model.Burger, error) {
	return r.CreateOne(mdl)
}

func GetPaginated(pageNum, perPage uint, r BurgerRepository) ([]*model.Burger, error){
	return r.GetPaginated(pageNum, perPage)
}

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

func (repo *BurgerRepositoryImpl) GetByName(name string) (*model.Burger, error) {
	var brg model.Burger
	tx := repo.Db.First(&brg, "name = ?", name)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return &brg, nil
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



func (repo *BurgerRepositoryImpl) GetPaginated(pageNum, perPage uint) ([]*model.Burger, error){
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
	var burgers [] *model.Burger
	tx := repo.Db.Limit(int(pgn.Limit)).Offset(int(pgn.Offset)).Find(&burgers)
	if tx.Error != nil{
		return nil, tx.Error
	}

	return burgers, nil
}
