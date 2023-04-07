package utils

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
)

func UseString(s *string) (value string) {
	if s == nil {
		value = ""
		return
	}
	value = *s
	return
}

func GetRefreshToken(c *gin.Context) (refreshToken string, err error) {
	refreshToken = c.Value("refresh_token").(string)
	if refreshToken == "" {
		return "", errors.New("error finding cookie")
	}
	return
}
func SetRefreshToken(value string, c *gin.Context) {
	c.SetCookie(
		"refresh_token",
		value,
		60*60*24*7,
		"/",
		"0.0.0.0",
		false,
		false,
	)
}

func SendEmail(subject, to, body string) (err error) {
	from := viper.GetString("mail_username")
	pass := viper.GetString("mail_password")

	e := email.NewEmail()
	e.From = "Robbo <" + from + ">"
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(body)

	auth := smtp.PlainAuth("", from, pass, viper.GetString("smtp_server_host"))
	err = e.Send(viper.GetString("smtp_server_address"), auth)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func Hash(s string) (hash string) {
	pwd := sha1.New()
	pwd.Write([]byte(s))
	pwd.Write([]byte(viper.GetString("auth_hash_salt")))
	hash = fmt.Sprintf("%x", pwd.Sum(nil))
	return
}
