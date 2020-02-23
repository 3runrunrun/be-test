package transaksi

import (
	"encoding/json"
	"log"

	"github.com/3runrunrun/be-test/layanan"
	"github.com/3runrunrun/be-test/user"
)

func (h Handler) initLayananGateway() layanan.ModulGateway {
	mg := layanan.ProvideLayananGateway(h.Handler.DB)
	return mg
}

func (h Handler) initUserGateway() user.ModulGateway {
	mg := user.ProvideUserGateway(h.Handler.DB)
	return mg
}

func (h Handler) getTagihan(detail []ItemRequestMapper) float64 {
	var hargaLayanan, tagihan float64

	layananGateway := h.initLayananGateway()

	for _, v := range detail {
		hargaLayanan = layananGateway.GetHarga(v.LayananID)
		tagihan += (hargaLayanan * v.Qty)
	}

	return tagihan
}

func (h Handler) getNamaPelanggan(idPelanggan uint) string {
	var namaUser string

	userGateway := h.initUserGateway()

	namaUser = userGateway.ReadNamaUser(idPelanggan)
	return namaUser
}

func (h Handler) getLayananDetail(idLayanan uint) LayananResponseMapper {
	var layananResp LayananResponseMapper
	var layanan layanan.LayananMapper

	layananGateway := h.initLayananGateway()

	layanan = layananGateway.Read(idLayanan)

	// log.Println(layanan)

	j, err := json.MarshalIndent(layanan, "", "\t")
	if err != nil {
		log.Panicln("transaksi controller.go (marshal): ", err)
	}

	// log.Printf("%s\n\n", j)

	err = json.Unmarshal(j, &layananResp)
	if err != nil {
		log.Panicln("transaksi controller.go (unmarshal): ", err)
	}

	// log.Println(layananResp)
	return layananResp
}
