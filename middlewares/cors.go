package middlewares

import "github.com/labstack/echo"

func CORSMiddlewareHandler() echo.MiddlewareFunc {
	return func(c *echo.Context) {
		c.Header().Set("Content-Type", "application/json")
		c.Header().Set("Access-Control-Allow-Origin", "*")
		c.Header().Set("Access-Control-Max-Age", "86400")
		c.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
	}
}