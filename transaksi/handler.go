package transaksi

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Handler for transaksi API
type Handler struct {
	Handler Object
}

// ProvideTransaksiAPI to provide API object
func ProvideTransaksiAPI(db *gorm.DB) Handler {
	model := ProvideObject(db)
	return Handler{Handler: model}
}

// Add transaksi
func (h Handler) Add() gin.HandlerFunc {
	var transaksi RequestMapper
	var tagihan float64
	var pelanggan string

	return func(c *gin.Context) {
		err := c.BindJSON(&transaksi)
		if err != nil {
			log.Panicln("transaksi handler.go: ", err)
			c.Status(http.StatusBadRequest)
			return
		}

		tagihan = h.getTagihan(transaksi.Detail)
		pelanggan = h.getNamaPelanggan(transaksi.PelangganID)
		transaksi.Tagihan = tagihan

		// save new transaksi
		newTransaksi := h.Handler.Save(toTableTransaksi(transaksi))

		// save each item of transaksi (detail)
		detail := make([]ItemResponseMapper, len(transaksi.Detail))
		for i, v := range transaksi.Detail {
			v.TransaksiID = newTransaksi.ID
			newDetail := h.Handler.SaveDetail(toTableItem(v))
			newItem := h.getLayananDetail(v.LayananID)
			detail[i] = ItemResponseMapper{
				LayananID: newDetail.LayananID,
				Qty:       newDetail.Qty,
				Layanan:   newItem,
			}
		}

		ret := &TransaksiResponseMapper{
			ID:        newTransaksi.ID,
			Pelanggan: pelanggan,
			Unit:      newTransaksi.Unit,
			Tagihan:   tagihan,
			Detail:    detail,
		}

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "success", "data": ret})
	}
}

// Read a single transaksi
func (h Handler) Read() gin.HandlerFunc {
	var pelanggan string

	return func(c *gin.Context) {
		idTransaksi, _ := strconv.ParseUint(c.Params.ByName("idTransaksi"), 36, 64)

		dataTransaksi := h.Handler.Read(uint(idTransaksi))
		if dataTransaksi == (Transaksi{}) {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "failed", "message": "transaksi not found"})
			return
		}
		dataDetail := h.Handler.ReadDetailByTransaksi(uint(idTransaksi))

		detail := make([]ItemResponseMapper, len(dataDetail))
		for i, v := range dataDetail {
			layanan := h.getLayananDetail(v.LayananID)
			detail[i] = ItemResponseMapper{
				LayananID: v.LayananID,
				Qty:       v.Qty,
				Layanan:   layanan,
			}
		}

		pelanggan = h.getNamaPelanggan(dataTransaksi.PelangganID)

		ret := &TransaksiResponseMapper{
			ID:        dataTransaksi.ID,
			Pelanggan: pelanggan,
			Unit:      dataTransaksi.Unit,
			Tagihan:   dataTransaksi.Tagihan,
			Detail:    detail,
		}

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "success", "data": ret})
	}
}
