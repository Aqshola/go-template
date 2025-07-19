package model_main

type TestTable struct {
	Id     int            `gorm:"column:id;primaryKey"`
	Name   string         `gorm:"column:name"`
	Detail *TestJoinTable `gorm:"foreignKey:Id;references:IdMaster"`
}
