package user

import "github.com/jinzhu/gorm"

// ModulGateway is a getway object to be used inter-modul
type ModulGateway struct {
	Gateway Object
}

// ProvideUserGateway for other modul
func ProvideUserGateway(db *gorm.DB) ModulGateway {
	modul := ProvideObject(db)
	return ModulGateway{Gateway: modul}
}

// ReadNamaUser get a user name
func (mg *ModulGateway) ReadNamaUser(id uint) string {
	user := mg.Gateway.Read(id)
	return user.Nama
}
