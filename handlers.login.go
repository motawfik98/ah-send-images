package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
  "github.com/gorilla/sessions"
  "golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("1298498081"))

func showLoginPage(c* gin.Context) {
  mainDepartments := findAllMainDepartments()
  secondaryDepartments := findAllSecondaryDepartments()
  session, err := store.Get(c.Request, "login")
  if err != nil {
    http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
    return
  }
  flashStatus := session.Flashes("status")
  flashMessage := session.Flashes("message")
  session.Options.MaxAge = -1
  session.Save(c.Request, c.Writer)
  c.HTML (
    http.StatusOK,
    "login.html",
    gin.H {
      "title": "تسجيل دخول",
      "mainDepartments": mainDepartments,
      "secondaryDepartments": secondaryDepartments,
      "status": flashStatus,
      "message": flashMessage,
    },
  )
}

func performLogin(c *gin.Context) {
  var loginData, user User
  c.ShouldBindWith(&loginData, binding.Form)
  session, err := store.Get(c.Request, "login")
  db.Where("username = ? AND main_department_id = ? AND secondary_department_id = ?",
          loginData.Username, loginData.MainDepartmentID, loginData.SecondaryDepartmentID).First(&user)
  if user.ID == 0 {
    if err != nil {
      http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
      return
    }
    addFlashMessage(session, "failure", "بيانات الدخول ليست صحيحه", c)
    http.Redirect(c.Writer, c.Request, "/login", 301)
  } else {
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
      addFlashMessage(session, "failure", "بيانات الدخول ليست صحيحه", c)
      http.Redirect(c.Writer, c.Request, "/login", 301)
    }
  }
}
