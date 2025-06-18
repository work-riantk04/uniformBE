package main

import (
	control "api/uniform/controllers"
	"fmt"
	"strings"

	// mdl "api/uniform/middlewares"
	model "api/uniform/models"
	st "api/uniform/struct"
	sha256 "crypto/sha256"
	hex "encoding/hex"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	jwt "github.com/dgrijalva/jwt-go"
)

func isAuthorized(endpoint echo.HandlerFunc) echo.HandlerFunc {
	var dataResponse st.DataResponse

	return func(c echo.Context) error {
		if c.Request().Header.Get("signature") != "" && c.Request().Header.Get("appname") != "" {

			code, desc, signature := model.CekLogin(c.RealIP(), c.Request().Header.Get("appname"))

			if code == "000" {
				h := sha256.New()
				h.Write([]byte(signature))
				signingKey := hex.EncodeToString(h.Sum(nil))

				if signingKey == c.Request().Header.Get("signature") {
					endpoint(c)
				} else {
					dataResponse = st.ResponseNil{
						ErrorCode:    "ES-200",
						ResponseCode: "002",
						ResponseDesc: "Signature Invalid",
					}

					return c.JSONPretty(http.StatusOK, dataResponse, "  ")
				}
			} else {
				dataResponse = st.ResponseNil{
					ErrorCode:    code,
					ResponseCode: "001",
					ResponseDesc: desc,
				}

				return c.JSONPretty(http.StatusOK, dataResponse, "  ")
			}
		} else {
			dataResponse = st.ResponseNil{
				ErrorCode:    "ES-201",
				ResponseCode: "002",
				ResponseDesc: "Not Authorized",
			}

			return c.JSONPretty(http.StatusOK, dataResponse, "  ")
		}
		dataResponse = st.ResponseNil{
			ErrorCode:    "ES-201",
			ResponseCode: "002",
			ResponseDesc: "Not Authorized",
		}
		return echo.NewHTTPError(http.StatusInternalServerError, dataResponse, "  ")
	}
}

func checkSession(endpoint echo.HandlerFunc) echo.HandlerFunc {
	var dataResponse st.DataResponse
	return func(c echo.Context) error {
		
		if c.Request().Header.Get("Authorization") != "" && c.Request().Header.Get("User") != "" {
			h := sha256.New()
			h.Write([]byte(c.Request().Header.Get("User")))
			signingKey := hex.EncodeToString(h.Sum(nil))

			stringAuth := c.Request().Header.Get("Authorization")
			splitAuth := strings.Split(stringAuth, " ")
			if stringAuth != "" {
				if splitAuth[0] == "Bearer" && len(splitAuth) > 1 && splitAuth[1] != "null" {
					token, _ := jwt.ParseWithClaims(splitAuth[1], &st.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
						return []byte(signingKey), nil
					})

					// exp := err.Error()[:16]
					// if exp == "token is expired" {
					// 	models.Logout(c.Request().Header.Get("User"))
					// 	endpoint(c)
					// } else {
					if token.Valid {
						endpoint(c)
					} else {
						dataResponse = st.Response{
							Response: st.ResponseNil{
								ErrorCode:    "ES-200",
								ResponseCode: "002",
								ResponseDesc: "Token Invalid",
							},
						}
						return c.JSONPretty(http.StatusOK, dataResponse, "  ")
					}
					// }
				} else {
					dataResponse = st.Response{
						Response: st.ResponseNil{
							ErrorCode:    "ES-200",
							ResponseCode: "002",
							ResponseDesc: "Token Invalid",
						},
					}

					return c.JSONPretty(http.StatusOK, dataResponse, "  ")
				}
			} else {
				dataResponse = st.Response{
					Response: st.ResponseNil{
						ErrorCode:    "ES-200",
						ResponseCode: "002",
						ResponseDesc: "Token Invalid",
					},
				}

				return c.JSONPretty(http.StatusOK, dataResponse, "  ")
			}
		} else {
			dataResponse = st.Response{
				Response: st.ResponseNil{
					ErrorCode:    "ES-200",
					ResponseCode: "002",
					ResponseDesc: "Token Invalid",
				},
			}

			return c.JSONPretty(http.StatusOK, dataResponse, "  ")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
}

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		var dataResponse st.DataResponse

		dataResponse = st.Response{
			Response: st.ResponseNil{
				ErrorCode:    "ER-900",
				ResponseCode: "099",
				ResponseDesc: "Method Not Found",
			},
		}

		return c.JSONPretty(http.StatusNotFound, dataResponse, "  ")
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		// AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// e.POST("/LoginBristars", control.LoginBristars)
	e.POST("/Login", control.Login)
	e.POST("/Logout", control.Logout, checkSession)
	e.POST("/GetMenu", control.GetMenu, checkSession)
	e.POST("/GetStatus", control.GetStatus, checkSession)
	e.POST("/GetUserAll", control.GetUserAll, checkSession)
	e.POST("/GetUserById", control.GetUserAll, checkSession)
	e.POST("/GetDivisiAll", control.GetDivisiAll, checkSession)
	e.POST("/GetParamAll", control.GetParamAll, checkSession)
	
	//NOTIF
	// e.POST("/GetNotifAll", control.GetNotifAll, checkSession)
	// e.POST("/GetNotifById", control.GetNotifById, checkSession)
	// e.POST("/UpdateNotif", control.UpdateNotif, checkSession)

	// Dashboard
	e.POST("/GetSummary", control.GetSummary, checkSession)

	// Form
	e.POST("/InsertForm", control.InsertForm, checkSession)
	e.POST("/GetFormAll", control.GetFormAll, checkSession)
	e.POST("/UpdateForm", control.UpdateForm, checkSession)
	e.POST("/ApprovalForm", control.ApprovalForm, checkSession)
	e.POST("/ActiveForm", control.ActiveForm, checkSession)

	e.POST("/InsertJawabanForm", control.InsertJawabanForm, checkSession)

	// Response
	e.POST("/GetResponseAll", control.GetResponseAll, checkSession)
	e.POST("/GetSummaryJawabanList", control.GetSummaryJawabanList, checkSession)
	e.POST("/GetJawabanByFormId", control.GetJawabanByFormId, checkSession)
	e.POST("/GetJawabanAll", control.GetJawabanAll, checkSession)
	e.POST("/ApprovalJawaban", control.ApprovalJawaban, checkSession)

	e.POST("/UpdateJawabanForm", control.UpdateJawabanForm, checkSession)

	e.POST("/RequestDownloadAttachment", control.RequestDownloadAttachment, checkSession)
	e.POST("/GetRequestAttachmentAll", control.GetRequestAttachmentAll, checkSession)
	
	e.POST("/JobGetRequestAttachmentDetail", control.JobGetRequestAttachmentDetail)
	e.POST("/JobUpdateRequestAttachment", control.JobUpdateRequestAttachment)
	

	fmt.Println("starting API uniform")
	e.Start(":3300")
}
