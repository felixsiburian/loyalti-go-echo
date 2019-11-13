package repository

import (
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)

func GetProgram() []model.Program {
	db := database.ConnectionDB()

	var program []model.Program
	db.Find(&program)
	db.Close()
	return program
}
