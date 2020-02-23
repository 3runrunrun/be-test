package layanan

import (
	"github.com/jinzhu/gorm"
)

// ModulGateway is a getway object to be used inter-modul
type ModulGateway struct {
	Gateway Object
}

// ProvideLayananGateway provide layanan gateway to other modul
func ProvideLayananGateway(db *gorm.DB) ModulGateway {
	model := ProvideObject(db)
	return ModulGateway{Gateway: model}
}

// GetHarga a single layanan
func (mg *ModulGateway) GetHarga(id uint) float64 {
	layanan := mg.Gateway.Read(id)
	return layanan.Harga
}

func (mg *ModulGateway) Read(id uint) LayananMapper {
	layanan := mg.Gateway.Read(id)
	ret := toResponse(layanan)
	return ret
}
