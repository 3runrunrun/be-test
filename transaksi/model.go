package transaksi

import "github.com/jinzhu/gorm"

// Object for transaksi model
type Object struct {
	DB *gorm.DB
}

// Transaksi table structure
type Transaksi struct {
	gorm.Model
	PelangganID uint    `gorm:"column:id_pelanggan;type:int unsigned;NOT NULL" json:"id_pelanggan"`
	Unit        string  `gorm:"type:varchar(10);NOT NULL" json:"unit"`
	Tagihan     float64 `gorm:"type:float;NOT NULL" json:"tagihan"`
}

// Item table structure
type Item struct {
	gorm.Model
	TransaksiID uint    `gorm:"column:id_transaksi;type:int unsigned;NOT NULL" json:"id_transaksi"`
	Qty         float64 `gorm:"type:float;NOT NULL" json:"qty"`
	LayananID   uint    `gorm:"column:id_layanan;type:int unsigned;NOT NULL" json:"id_layanan"`
}

// ProvideObject for transaksi model
func ProvideObject(db *gorm.DB) Object {
	return Object{DB: db}
}

// Save new Transaction
func (o Object) Save(t Transaksi) Transaksi {
	o.DB.Create(&t)

	return t
}

// SaveDetail to save transaction detail
func (o Object) SaveDetail(i Item) Item {
	o.DB.Create(&i)

	return i
}
