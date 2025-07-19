package repository_main

import (
	model_main "go-template/src/model/main"

	"gorm.io/gorm"
)

type ITestTableRepository interface {
	GetListTestTable() ([]model_main.TestTable, error)
	GetDetailTestTable(id int) (*model_main.TestTable, error)
}

type TestTableRepository struct {
	DB *gorm.DB
}

func (r *TestTableRepository) GetListTestTable() ([]model_main.TestTable, error) {
	var listTestTable []model_main.TestTable

	err := r.DB.Find(&listTestTable).Error
	if err != nil {
		return nil, err
	}

	return listTestTable, nil
}

func (r *TestTableRepository) GetDetailTestTable(id int) (*model_main.TestTable, error) {
	var detailData model_main.TestTable

	err := r.DB.Preload("Detail").First(&detailData, "id", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &detailData, nil
}

func NewMainRepository(db *gorm.DB) ITestTableRepository {
	return &TestTableRepository{
		DB: db,
	}
}
