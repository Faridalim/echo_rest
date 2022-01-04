package lain2

import (
	"bytes"
	"echo_rest/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/labstack/echo/v4"
)

func postExternal(c echo.Context) error {
	url := "https://portal.migascepu.id/sertifikasi/api/testJson"

	var jsonString = []byte(
		`{
			"token"	: "12345",
			"nama_peserta" : "Rahardian",
			"tanggal_lahir" : "22-05-2021",
			"sttk" : "Pemboran"
		}`,
	)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authentication", "Bearer 12345")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("resp status : ", resp.Status)
	fmt.Println("resp headers : ", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resp body : ", string(body))
	return c.HTML(200, string(body))
}

func kirimEmail(c echo.Context) error {
	to := []string{"owner.xmedia@gmail.com"}
	cc := []string{"asiswanto254@gmail.com"}
	subject := "Test mail"
	message := "Anda mendapatkan surat tagihan"

	err := sendMail(to, cc, subject, message)

	if err != nil {
		panic("pengiriman gagal")
	}

	return c.String(200, "kirim email sukses")

}

func sendMail(to []string, cc []string, subject string, message string) error {
	conf := config.GetConfig()

	body := "From: " + conf.SMTP_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc:" + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", conf.SMTP_USERNAME, conf.SMTP_PASSWORD, conf.SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%s", conf.SMTP_HOST, conf.SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, conf.SMTP_USERNAME, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
