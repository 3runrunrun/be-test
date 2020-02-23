package layanan

import "github.com/jinzhu/gorm"

// Object for layanan model
type Object struct {
	DB *gorm.DB
}

// Layanan table structure
type Layanan struct {
	gorm.Model
	Nama  string  `gorm:"type:varchar(50);NOT NULL" json:"nama"`
	Unit  string  `gorm:"type:varchar(10);NOT NULL" json:"unit"`
	Harga float64 `gorm:"type:float" json:"harga"`
}

// ProvideObject model
func ProvideObject(db *gorm.DB) Object {
	return Object{DB: db}
}

//Save new layanan
func (o Object) Save(l Layanan) Layanan {
	o.DB.Create(&l)

	return l
}

// Show all layanans
func (o Object) Show() []Layanan {
	var layanans []Layanan
	o.DB.Find(&layanans)

	return layanans
}

// Read single layanan
func (o Object) Read(id uint) Layanan {
	var layanan Layanan
	o.DB.First(&layanan, id)

	return layanan
}
