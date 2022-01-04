package security

import (
	"echo_rest/helpers"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var Access = GetAccess()

func GetAccess() map[string]interface{} {
	data, err := os.ReadFile("security/access.json")
	check(err)

	m := map[string]interface{}{}
	json.Unmarshal(data, &m)
	return m
}

// Update akses permission jika file access berubah
func UpdateAccess() {
	Access = GetAccess()
}

// check access permission
func CheckAccess(fungsi string, c echo.Context) bool {
	data := Access[fungsi]

	//convert type interface to string and then split string
	data_str := fmt.Sprintf("%v", data)
	permitted_roles := strings.Split(data_str, ", ")

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*helpers.JwtCustomClaims)
	user_roles := claims.Roles
	timeUTC := time.Unix(claims.ExpiresAt, 0)
	fmt.Println(timeUTC)

	if contain(permitted_roles, user_roles) {
		return contain(permitted_roles, user_roles)
	}

	return false
}

func contain(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

// Mendapatkan nama fungsi untuk permissivon
func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	caller := fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	caller_split := strings.Split(caller, ".")
	return caller_split[len(caller_split)-1]
}
