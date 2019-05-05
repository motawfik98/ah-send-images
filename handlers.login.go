package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func showLoginPage(c *gin.Context) {
	mainDepartments := findAllMainDepartments()
	secondaryDepartments := findAllSecondaryDepartments()
  c.JSON(http.StatusOK, gin.H{
    "title":                "تسجيل دخول",
    "mainDepartments":      mainDepartments,
    "secondaryDepartments": secondaryDepartments,
  })
}

func performLogin(c *gin.Context) {
	var loginData, user User
	c.ShouldBindJSON(&loginData)
	db.Where("username = ? AND main_department_id = ? AND secondary_department_id = ?",
		loginData.Username, loginData.MainDepartmentID, loginData.SecondaryDepartmentID).First(&user)
	if user.ID == 0 {
    c.JSON(http.StatusTemporaryRedirect, gin.H{
      "flashStatus": "failure",
      "flashMessage": "بيانات الدخول ليست صحيحه",
      "url": "login",
    })
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
      c.JSON(http.StatusTemporaryRedirect, gin.H{
        "flashStatus": "failure",
        "flashMessage": "بيانات الدخول ليست صحيحه",
        "url": "login",
      })
		} else {
      c.JSON(http.StatusTemporaryRedirect, gin.H{
        "url": "/",
      })
    }
	}
}

func logout(c *gin.Context) {
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusTemporaryRedirect, gin.H{
    "url": "/login",
  })
}
