package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/akhilesh-ge/jwt-email/middleware"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type Detail struct {
	Email string `json:"email" binding:"required"`
}

type Login struct {
	Otp int `json:"otp" binding:"required"`
}

var otp int

func generateOTP() int {
	rand.Seed(time.Now().UnixNano())
	min_num := 1111
	max_num := 9999
	num := rand.Intn(max_num-min_num+1) + min_num
	return num
}

func SignIn(ctx *gin.Context) {
	var mailId Detail
	ctx.ShouldBindJSON(&mailId)
	to_mail := mailId.Email
	fmt.Printf("The mail id is %v\n", to_mail)
	fmt.Printf("The type is %T\n", to_mail)

	// get user password
	user := os.Getenv("USER")
	password := os.Getenv("PASS")

	// message template
	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", to_mail)
	m.SetHeader("Subject", "OTP Verification")
	otp = generateOTP()
	body := fmt.Sprintf("OTP for Signin: %v", otp)
	m.SetBody("text/plain", body)

	// Send message
	d := gomail.NewDialer("smtp.gmail.com", 587, user, password)

	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OTP sent Successfully",
	})
}

func VerifyOTP(ctx *gin.Context) {
	var verify Login
	ctx.ShouldBindJSON(&verify)
	if otp == verify.Otp {
		token := middleware.GenerateToken()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully SignedIN",
			"token":   token,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OTP is incorrect",
		})
	}
}
