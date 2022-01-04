package middlewares

// func AccessControl(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// fmt.Println("from middleware one")
// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(*helpers.JwtCustomClaims)
// 		name := claims.Name
// 		// fmt.Println(name)

// 		return next(c)
// 	}
// }
