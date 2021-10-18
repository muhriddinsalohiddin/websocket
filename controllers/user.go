package controllers

import (
	"log"
	"github.com/gin-gonic/gin"
	"app/utils"
	"app/models"
)


func Register (c *gin.Context) {

	db, err := utils.DB()

	if err != nil {
		log.Fatal(err)
	}

	var reg,user models.User
	
	c.BindJSON(&reg)

	db.Select("email").Where("email = ?",reg.Email).Find(&user)

	if user.Email == "" {

		pass := utils.RandString(5)
		db.Model(&reg).Create([]map[string]interface{}{
			{"email": reg.Email,"password": pass},
		})

		if err = utils.SendMail([]string{reg.Email},[]byte(pass)); err == nil{
			c.String(200,"Emailga parol yuborildi")
		}

	}else{
		c.String(400,"Siz ro'yxatdan o'tgansiz")
	}

}

func Login (c *gin.Context) {

	db, err := utils.DB()
	
	if err != nil {
		log.Fatal(err)
	}

	var user,login models.User

	c.BindJSON(&login)
	
	db.Where("email=? and password = ?",login.Email,login.Password).Find(&user)
	if user.Email == "" {

		c.String(400,"Wrong email or password")
	}else{

		c.String(200,utils.CreateToken(user))
	}

}

func Change(c *gin.Context) {

	var UpdateUser,checkuser models.ChangeUser
	var user models.User
	c.BindJSON(&UpdateUser)

	db, err := utils.DB()

	if err != nil {
		log.Fatal(err)
	}

	db.Model(&models.User{}).Select("email","password").Where("email = ? and password = ?",UpdateUser.Email,UpdateUser.Password).Find(&checkuser)

	if checkuser.Email == "" {
		c.String(400,"Wrong email or password")
		return
	}

	db.Model(&models.User{}).
	Where("email = ? and password = ?",UpdateUser.Email,UpdateUser.Password).
	Update("password",UpdateUser.NewPassword)
	
	db.Model(&models.User{}).Where("email = ?",UpdateUser.Email).Find(&user)

	c.String(200,utils.CreateToken(user))
}