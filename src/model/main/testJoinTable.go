package model_main

type TestJoinTable struct {
	Id       int    `gorm:"column:id;primaryKey"`
	IdMaster int    `gorm:"column:id_master"`
	IsDetail string `gorm:"column:is_detail"`
}
