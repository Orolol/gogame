package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/orolol/gogame/utils"
)

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://http://localhost:8080/todos

*/
func SignUp(c *gin.Context) {
	db, err := gorm.Open("mysql", ConnexionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var acc utils.Account
	fmt.Println("CREATE ACC")

	c.ShouldBind(&acc)

	fmt.Println(acc)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("PASS : ", acc.Password)
	acc.Password = string(hashedPassword)
	fmt.Println("FINAL PASS : ", acc.Password)

	acc.ELO = 1500

	stoken, err := utils.GenerateRandomString(22)
	if err != nil {
		fmt.Println("Error while generating token")
		return
	}
	token := utils.Token{Token: stoken, Status: "active"}

	if err := db.Create(&token).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error during account creation")
		return
	}

	acc.TokenID = token.ID

	if err := db.Create(&acc).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error during account creation")
		return
	} else {
		fmt.Println("CREATED ", acc.ID)
		fmt.Println("CREATED ", acc.Name)
		fmt.Println("CREATED ", acc.Login)
		fmt.Println("CREATED ", acc.Password)
		fmt.Println("CREATED ", acc.TokenID)

		c.String(http.StatusCreated, stoken)
	}

}

func EditAccount(c *gin.Context) {
	db, err := gorm.Open("mysql", ConnexionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var acc utils.Account
	var dbacc utils.Account

	c.ShouldBind(&acc)
	claims := jwt.ExtractClaims(c)
	if acc.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		acc.Password = string(hashedPassword)
	}
	if res := db.First(&dbacc, "Login = ?", claims["id"]); res.Error != nil {
		c.String(http.StatusInternalServerError, "Error during account edition")
		return
	} else {
		dbacc.Name = acc.Name
		dbacc.ProfilePic = acc.ProfilePic
		dbacc.Step = acc.Step
		if acc.Password != "" {
			dbacc.Password = acc.Password
		}
		db.Save(&dbacc)
		c.Status(http.StatusCreated)
	}

}

func GetProfileInfos(c *gin.Context) {
	var acc utils.Account
	var accApi utils.AccountApi
	claims := jwt.ExtractClaims(c)
	db, _ := gorm.Open("mysql", ConnexionString)
	db.Where("Login = ?", claims["id"]).First(&acc)

	accApi.ID = acc.ID
	accApi.Login = acc.Login
	accApi.Name = acc.Name
	accApi.ELO = acc.ELO
	accApi.ProfilePic = acc.ProfilePic
	accApi.Step = acc.Step
	accApi.SelectedCountry = acc.SelectedCountry
	c.JSON(http.StatusOK, accApi)
}
func GetEnemyInfos(c *gin.Context) {
	var acc utils.Account
	var accApi utils.AccountApi
	idquer := c.Param("id")
	c.ShouldBind(&acc)

	db, _ := gorm.Open("mysql", ConnexionString)
	db.Where("ID = ?", idquer).First(&acc)
	fmt.Println("acc", acc)
	accApi.Name = acc.Name
	accApi.ELO = acc.ELO
	accApi.ProfilePic = acc.ProfilePic
	fmt.Println("Name", accApi.Name)
	fmt.Println("ELO", accApi.ELO)
	fmt.Println("ProfilePic", accApi.ProfilePic)
	c.JSON(http.StatusOK, accApi)
}

func GetPP(c *gin.Context) {
	var list []utils.ProfilePic
	list = append(list,
		utils.ProfilePic{Availablity: "all", Name: "pp1"},
	)
	list = append(list,
		utils.ProfilePic{Availablity: "all", Name: "pp2"},
	)

	c.JSON(http.StatusOK, list)
}

func GetNews(c *gin.Context) {
	var list = utils.GetNews()
	c.JSON(http.StatusOK, list)
}
