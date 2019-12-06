package repository

import (
	"database/sql"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)
// fungsi get outlet
func GetOutlet(page *int, size *int, id *int) []model.Outlet {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var outlet []model.Outlet
	var rows *sql.Rows
	var err error
	var total int


	if id != nil {
		if page != nil && size != nil {
			fmt.Println("2")
			rows, err = db.Where("merchant_id = ?", id).Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		}
	} else {
		if page != nil && size != nil{
			rows, err = db.Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		} else {
			rows, err = db.Find(&outlet).Rows()
			if err != nil{
				panic(err)
			}
		}
	}

	result := make([]model.Outlet, 0)

	for rows.Next() {
		o := new(model.Outlet)
		err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.Deleted_by,
			&o.OutletName,
			&o.OutletAddress,
			&o.OutletPhone,
			&o.OutletCity,
			&o.OutletProvince,
			&o.OutletPostalCode,
			&o.OutletLongitude,
			&o.OutletLatitude,
			&o.OutletDay,
			&o.OutletHour,
			&o.MerchantId,
			)

		result = append(result,*o)
	}

	db.Close()
	return result
}
