package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func showLoginPage(c* gin.Context) {
  mainDepartments := findAllMainDepartments()
  secondaryDepartments := findAllSecondaryDepartments()
  c.HTML (
    http.StatusOK,
    "login.html",
    gin.H {
      "title": "تسجيل دخول",
      "mainDepartments": mainDepartments,
      "secondaryDepartments": secondaryDepartments,
    },
  )
}
