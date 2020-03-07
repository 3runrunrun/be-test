package user

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/3runrunrun/be-test/authenticator"
	"github.com/3runrunrun/be-test/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Handler API object
type Handler struct {
	Handler Object
}

type response struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ProvideUserAPI to provide the API object
func ProvideUserAPI(db *gorm.DB) Handler {
	model := ProvideObject(db)
	handler := Handler{Handler: model}
	return handler
}

// Register new user
func (h Handler) Register() gin.HandlerFunc {
	var user UserMapper
	return func(c *gin.Context) {
		err := c.BindJSON(&user)
		if err != nil {
			log.Panicln("user handler.go: ", err)
			c.Status(http.StatusBadRequest)
			return
		}

		h.Handler.Save(toTable(user))
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "success", "message": "berhasil terdaftar"})
	}
}

// Login user
func (h Handler) Login() gin.HandlerFunc {

	// var as request body binder
	var user UserMapper

	return func(c *gin.Context) {

		// bind request body to &user
		err := c.BindJSON(&user)
		if err != nil {
			log.Panicln("user handler.go: ", err)
			c.Status(http.StatusBadRequest)
			return
		}

		// authenticate username and password
		account := h.Handler.ReadByUserAndPassword(user.Username, user.Password)
		if account == (User{}) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "status": "failed", "message": "user not found!"})
			return
		}

		// create JWT token
		username := account.Username    // user
		iss := helpers.GetKey("APPKEY") // issuer of JWT
		token := authenticator.SetToken(username, iss)

		// sign JWT token
		secretkeystring := helpers.GetKey("SIGNATUREKEY")
		tokenString, err := token.SignedString([]byte(secretkeystring))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "token_generation_failed"})
			return
		}

		// set return struc
		ret := toLoginResponse(account)
		ret.Token = tokenString

		// set client cookie
		cookieexptime := time.Now().Add(time.Hour * time.Duration(48))
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "token",
			Value:   ret.Token,
			Expires: cookieexptime,
		})
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "id_pelanggan",
			Value:   strconv.FormatUint(uint64(ret.ID), 36),
			Expires: cookieexptime,
		})

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "status": "success", "data": ret})
	}
}

// Show all users
func (h Handler) Show() gin.HandlerFunc {

	return func(c *gin.Context) {
		users := h.Handler.Show()
		if len(users) <= 0 {
			log.Println("user handler.go: users is empty")
			c.Status(http.StatusOK)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": toMultipleResponse(users)})
	}
}
