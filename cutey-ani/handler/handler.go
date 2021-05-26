package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muyisz/cutey-ani/data"
)

const (
	// CookieExpireDuration is cookie's valid duration.
	CookieExpireDuration = 7200
	// CookieAccessScope is cookie's scope.
	CookieAccessScope = "127.0.0.1"
	// FileStorageDirectory is where these files storage.
	FileStorageDirectory = "views/img/"
	// DownloadUrlBase decide the base url of file's url.
	DownloadUrlBase = "http://127.0.0.1:8090/download"
)

func CheckCookie(db *data.MySQL) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("phone")
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{})
			c.Abort()
		}
		c.Next()
	}
}

func PostChatRoom(db *data.MySQL) gin.HandlerFunc {
	return func(c *gin.Context) {
		phone, _ := c.Cookie("phone")
		u, _ := db.GetUser(phone)
		msg := c.PostForm("msg")
		err := db.PostRoomMsg(u, msg)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"pass": false})
		}
		c.JSON(http.StatusOK, gin.H{"pass": true})
	}
}

func GetChatRoom(db *data.MySQL) gin.HandlerFunc {
	fmt.Println(1)
	return func(c *gin.Context) {
		msg, err := db.GetRoomMsg()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"pass": false})
		}
		c.JSON(http.StatusOK, gin.H{"pass": true, "msg": msg})
	}
}

func PostLoginData(db *data.MySQL) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("phone")
		password := c.PostForm("password")
		user := data.User{
			Email:    email,
			Password: password,
		}
		yes, _ := db.CheckUsers(&user)
		if yes {
			c.SetCookie("phone", email, CookieExpireDuration, "/", CookieAccessScope, false, true)
			c.JSON(http.StatusOK, gin.H{"pass": true, "phone": email})
		} else {
			c.JSON(http.StatusOK, gin.H{"pass": false})
		}
	}
}

func PostPhoto(db *data.MySQL) gin.HandlerFunc {
	return func(c *gin.Context) {
		photo, err := c.FormFile("photo")
		supp, _ := c.GetPostForm("supp")
		if err != nil || photo == nil {
			c.JSON(http.StatusOK, gin.H{"pass": false})
		}
		ext := c.PostForm("ext")
		num, _ := db.GetPhotoNum()
		url := FileStorageDirectory + strconv.Itoa(num+1) + ext
		dst := path.Join("./views/img/", strconv.Itoa(num+1)+ext)
		c.SaveUploadedFile(photo, dst)
		db.SetPhoto(num+1, url, supp)
		c.JSON(http.StatusOK, gin.H{"pass": true})
	}
}

func PostRegisterData(db *data.MySQL) gin.HandlerFunc {
	return func(c *gin.Context) {
		phone := c.PostForm("phone")
		passWord := c.PostForm("password")
		nickname := c.PostForm("nickname")
		address := c.PostForm("address")
		u := data.User{
			Email:    phone,
			Password: passWord,
			NickName: nickname,
			Address:  address,
		}
		err := db.JoinUsers(&u)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"pass": false, "phone": phone})
		}
		c.JSON(http.StatusOK, gin.H{"pass": true, "phone": phone})
	}
}

func GetPhoto(db *data.MySQL) gin.HandlerFunc {
	return func(c *gin.Context) {
		num, _ := db.GetPhotoNum()
		var url [11]string
		var supp [11]string
		var n [11]int
		rand.Seed(time.Now().Unix())
		for i := 0; i < 11; i++ {
			for {
				k := rand.Intn(num)
				flag := 1
				for j := 0; j < i; j++ {
					if k == n[j] {
						flag = 0
					}
				}
				if flag == 1 && k != 0 {
					fmt.Println(k)
					n[i] = k
					url[i], supp[i], _ = db.GetUrl(k)
					break
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"pass": true, "url": url, "supp": supp})
	}
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func GetRoom(c *gin.Context) {
	c.HTML(http.StatusOK, "message.html", gin.H{})
}
func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func GetInfo(db *data.MySQL) gin.HandlerFunc {
	return func(c *gin.Context) {
		phone, _ := c.Cookie("phone")
		u, _ := db.GetUser(phone)
		c.JSON(http.StatusOK, gin.H{"pass": true, "phone": u.Email, "address": u.Address, "nickname": u.NickName})
	}
}

func GetDash(c *gin.Context) {
	c.HTML(http.StatusOK, "dash.html", gin.H{"pass": true})
}
