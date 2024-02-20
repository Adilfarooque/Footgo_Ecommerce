package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

// func LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
// 	var details domain.Admin
// 	if err := db.DB.Raw("SELECT * FROM users WHERE email = ? AND isadmin = ture", adminDetails).Scan(&details).Error; err != nil {
// 		return domain.Admin{}, err
// 	}
// 	return details, nil
// }


func LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
    var details domain.Admin
    if err := db.DB.Raw("SELECT * FROM users WHERE email = ? AND isadmin = true", adminDetails.Email).Scan(&details).Error; err != nil {
        return domain.Admin{}, err
    }
    return details, nil
}


