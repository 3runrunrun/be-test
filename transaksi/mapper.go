package transaksi

// TransaksiMapper for transaksi table
type TransaksiMapper struct {
	ID          uint    `json:"id,string,omitempty"`
	PelangganID uint    `json:"id_pelanggan,string"`
	Unit        string  `json:"unit"`
	Tagihan     float64 `json:"tagihan,string"`
}

// ItemMapper for item table
type ItemMapper struct {
	ID          uint    `json:"id,string,omitempty"`
	TransaksiID uint    `json:"id_transaksi,string"`
	Qty         float64 `json:"qty,string"`
	LayananID   uint    `json:"id_layanan,string"`
}

// RequestMapper for request catcher
type RequestMapper struct {
	PelangganID uint                `json:"id_pelanggan,string"`
	Unit        string              `json:"unit"`
	Tagihan     float64             `json:"tagihan,string,omitempty"`
	ItemID      uint                `json:"id_item,string,omitempty"`
	Detail      []ItemRequestMapper `json:"detail"`
}

// TransaksiDetail consist of items
type TransaksiDetail struct {
	Item []ItemRequestMapper
}

// ItemRequestMapper for item-requset catcher
type ItemRequestMapper struct {
	TransaksiID uint    `json:"id_transaksi,omitempty"`
	LayananID   uint    `json:"id_layanan,string"`
	Qty         float64 `json:"qty,string"`
}

// TransaksiResponseMapper for transaksi-response
type TransaksiResponseMapper struct {
	ID        uint                 `json:"id,string"`
	Pelanggan string               `json:"pelanggan"`
	Unit      string               `json:"unit"`
	Tagihan   float64              `json:"tagihan,string"`
	Detail    []ItemResponseMapper `json:"detail"`
}

// DetailResponseMapper for detail-layanan-response
type DetailResponseMapper struct {
	detail []ItemResponseMapper
}

// ItemResponseMapper for item-response
type ItemResponseMapper struct {
	LayananID uint                  `json:"id_layanan,string"`
	Qty       float64               `json:"qty,string"`
	Layanan   LayananResponseMapper `json:"layanan"`
}

// LayananResponseMapper for layanan-detail-response
type LayananResponseMapper struct {
	ID    uint    `json:"id,string"`
	Nama  string  `json:"nama"`
	Unit  string  `json:"unit"`
	Harga float64 `json:"harga,string"`
}

// FinalResponse for transaksi-response
type FinalResponse struct {
	ID        uint                 `json:"id,string"`
	Pelanggan string               `json:"pelanggan"`
	Unit      string               `json:"unit"`
	Tagihan   float64              `json:"tagihan,string"`
	Detail    []ItemResponseMapper `json:"detail"`
}

func toTableTransaksi(rm RequestMapper) Transaksi {
	return Transaksi{PelangganID: rm.PelangganID, Unit: rm.Unit, Tagihan: rm.Tagihan}
}

func toTableItem(im ItemRequestMapper) Item {
	return Item{TransaksiID: im.TransaksiID, LayananID: im.LayananID, Qty: im.Qty}
}
