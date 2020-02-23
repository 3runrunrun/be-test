package layanan

// LayananMapper table map
type LayananMapper struct {
	ID    uint    `json:"id,string,omitempty"`
	Nama  string  `json:"nama"`
	Unit  string  `json:"unit"`
	Harga float64 `json:"harga,string"`
}

func toTable(l LayananMapper) Layanan {
	return Layanan{Nama: l.Nama, Unit: l.Unit, Harga: l.Harga}
}

func toResponse(l Layanan) LayananMapper {
	return LayananMapper{ID: l.ID, Nama: l.Nama, Unit: l.Unit, Harga: l.Harga}
}

func toMultipleResponse(l []Layanan) []LayananMapper {
	ret := make([]LayananMapper, len(l))

	for k, v := range l {
		ret[k] = toResponse(v)
	}

	return ret
}
