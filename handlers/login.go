package handlers

import (
	"../models"
	"../repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type MyDB struct {
	GormDB *gorm.DB
}

func (db *MyDB) showLoginPage(c echo.Context) error {
	mainDepartments := repositories.FindAllMainDepartments(db.GormDB)
	secondaryDepartments := repositories.FindAllSecondaryDepartments(db.GormDB)
	return c.JSON(http.StatusOK, gin.H{
		"title":                "تسجيل دخول",
		"mainDepartments":      mainDepartments,
		"secondaryDepartments": secondaryDepartments,
	})
}

func (db *MyDB) performLogin(c echo.Context) error {
	var loginData, user models.User
	_ = c.Bind(&loginData)
	db.GormDB.Where("username = ? AND main_department_id = ? AND secondary_department_id = ?",
		loginData.Username, loginData.MainDepartmentID, loginData.SecondaryDepartmentID).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusTemporaryRedirect, gin.H{
			"flashStatus":  "failure",
			"flashMessage": "بيانات الدخول ليست صحيحه",
			"url":          "/login",
		})
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
			return c.JSON(http.StatusTemporaryRedirect, gin.H{
				"flashStatus":  "failure",
				"flashMessage": "بيانات الدخول ليست صحيحه",
				"url":          "/login",
			})
		} else {
			err := setCookie(&c, user.ID)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusTemporaryRedirect, gin.H{
				"url": "/",
			})
		}
	}
}

func setCookie(c *echo.Context, id uint) error {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	(*c).SetCookie(&http.Cookie{
		Name:  "Authorization",
		Value: t,
	})
	return nil
}

func logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:    "Authorization",
		Expires: time.Now(),
	})
	return c.JSON(http.StatusTemporaryRedirect, map[string]string{
		"url": "/login",
	})
}
