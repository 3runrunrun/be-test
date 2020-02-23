package main

import (
	"log"

	"github.com/3runrunrun/be-test/database"
	"github.com/3runrunrun/be-test/layanan"
	"github.com/3runrunrun/be-test/middleware"
	"github.com/3runrunrun/be-test/transaksi"
	"github.com/3runrunrun/be-test/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	db = initDB()

	userHandler := initUserAPI(db)
	layananHandler := initLayananAPI(db)
	transaksiHandler := initTransaksiAPI(db)

	server := gin.Default()

	rootdir := server.Group("/")
	{
		rootdir.POST("register", userHandler.Register())
		rootdir.POST("/login", userHandler.Login())
		rootdir.GET("show", middleware.Authorization(), userHandler.Show())
		// rootdir.POST("login", userHandler.Login())
	}

	api := server.Group("/api/v1")
	{
		api.Use(middleware.Authorization())
		api.GET("/layanan/show", layananHandler.Show())
		api.POST("/layanan/add", layananHandler.Add())

		api.POST("/transaksi/add", transaksiHandler.Add())
	}

	err := server.Run(":3680")
	if err != nil {
		log.Panic("error")
	}
}

func initDB() *gorm.DB {
	sql := database.ConnectSQL()

	if !sql.HasTable(&user.User{}) && !sql.HasTable(&layanan.Layanan{}) && !sql.HasTable(&transaksi.Transaksi{}) && !sql.HasTable(&transaksi.Item{}) {
		sql.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user.User{})
		sql.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&layanan.Layanan{})
		sql.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&transaksi.Transaksi{})
		sql.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&transaksi.Item{})

		// set foreign key
		sql.Model(&transaksi.Transaksi{}).AddForeignKey("id_pelanggan", "users(id)", "CASCADE", "CASCADE")
		sql.Model(&transaksi.Item{}).AddForeignKey("id_transaksi", "transaksis(id)", "CASCADE", "CASCADE")
		sql.Model(&transaksi.Item{}).AddForeignKey("id_layanan", "layanans(id)", "CASCADE", "CASCADE")
	}

	return sql
}

func initUserAPI(db *gorm.DB) user.Handler {
	handler := user.ProvideUserAPI(db)
	return handler
}

func initLayananAPI(db *gorm.DB) layanan.Handler {
	handler := layanan.ProvideLayananAPI(db)
	return handler
}

func initTransaksiAPI(db *gorm.DB) transaksi.Handler {
	handler := transaksi.ProvideTransaksiAPI(db)
	return handler
}
