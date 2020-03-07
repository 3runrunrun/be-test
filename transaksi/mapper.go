package transaksi

// RequestMapper for request catcher
type RequestMapper struct {
	PelangganID uint                `json:"id_pelanggan,string"`
	Unit        string              `json:"unit"`
	Tagihan     float64             `json:"tagihan,string,omitempty"`
	ItemID      uint                `json:"id_item,string,omitempty"`
	Detail      []ItemRequestMapper `json:"detail"`
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

func toTableTransaksi(rm RequestMapper) Transaksi {
	return Transaksi{PelangganID: rm.PelangganID, Unit: rm.Unit, Tagihan: rm.Tagihan}
}

func toTableItem(im ItemRequestMapper) Item {
	return Item{TransaksiID: im.TransaksiID, LayananID: im.LayananID, Qty: im.Qty}
}

func toResponseMultipleItem(im ItemResponseMapper) []ItemResponseMapper {
	ret := make([]ItemResp, size)
}
