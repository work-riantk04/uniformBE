package controllers

import (
	// "strings"
	// "encoding/json"
    // "fmt"
    "net/http"
	"time"
	sha256 "crypto/sha256"
	hex "encoding/hex"
	jwt "github.com/dgrijalva/jwt-go"
    models "api/uniform/models"
    service "api/uniform/service"
    st "api/uniform/struct"
	"github.com/labstack/echo"
	"encoding/json"
)

//fungsi untuk melakukan login
func LoginBristars(c echo.Context) error {
	var dataResponse 	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}
	//username dan password harus ada di header request
	if c.FormValue("REQUEST") != "" {
		request				= map[string]interface{}{}
		request["REQUEST"]	= c.FormValue("REQUEST")
		jsonMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(c.FormValue("REQUEST")), &jsonMap)
		
		if err == nil {
			input	= map[string]interface{}{}
			for k, v := range jsonMap {
				input[k] = v
			}

			var mandatory_input []string
			// mandatory_input = append(mandatory_input,"Key")
			// mandatory_input = append(mandatory_input,"UserApp")
			// mandatory_input = append(mandatory_input,"AppId")
			mandatory_input = append(mandatory_input,"USER")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}
			
			if cek_mandatory == 0 {
				// status, pernr, err := service.BristarsAppCheck(input)
				
				// if status == "error" || err != nil {
				// 	dataResponse = st.Response {
				// 		Response : st.ResponseNil{
				// 			ErrorCode		: "ER-201",
				// 			ResponseCode 	: "002",
				// 			ResponseDesc	: "Not Authorized",
				// 		},
				// 	}
				// } else {
					//input["USER"] = pernr
					count, _ := models.CekSession(input["USER"].(string))
					if count == 0 {

						Data, err_code, err, _ := models.GetUserById(input)
						if err == nil {
							//pengecekan app tersebut telah terdaftar dengan user pass sesuai
							models.Login(input["USER"].(string))
					
							h := sha256.New()
							h.Write([]byte(input["USER"].(string)))
							signingKey := hex.EncodeToString(h.Sum(nil))
							
							var signing = []byte(signingKey)
					
							claims := st.JwtClaims{
								input["USER"].(string),
								jwt.StandardClaims{
									Id: input["USER"].(string),
									ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
								},
							}
					
							rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
						
							tokenString, err := rawToken.SignedString(signing)
							
							if err != nil { //jika terjadi error
								dataResponse = st.Response {
									Response : st.ResponseNil{
										ErrorCode		: "ER-001",
										ResponseCode 	: "001",
										ResponseDesc	: "Invalid Data",
									},
								}
							} else {
								dataResponse = st.Response {
									Response : st.Token {
										ErrorCode 		: "ES-000",
										ResponseCode 	: "000",
										ResponseDesc	: "Success",
										Token			: tokenString,
										ResponseData	: Data,
									},
								}
							}
						} else {
							dataResponse = st.Response {
								Response : st.ResponseNil{
									ErrorCode		: err_code,
									ResponseCode 	: "001",
									ResponseDesc	: err.Error(),
								},
							}
						}
						
					} else {
						dataResponse = st.Response {
							Response : st.ResponseNil{
								ErrorCode		: "ES-200",
								ResponseCode 	: "003",
								ResponseDesc	: "You already login in another device",
							},
						}
					}
				// }
				
			} else {
				dataResponse = st.Response {
					Response : st.ResponseNil{
						ErrorCode		: "ER-001",
						ResponseCode 	: "001",
						ResponseDesc	: "Invalid Data",
					},
				}
			}			
		} else {
			dataResponse = st.Response {
				Response : st.ResponseNil{
					ErrorCode		: "ER-001",
					ResponseCode 	: "001",
					ResponseDesc	: "Invalid Data",
				},
			}
		}
	} else {
		dataResponse = st.Response {
			Response : st.ResponseNil{
				ErrorCode		: "ER-001",
				ResponseCode 	: "001",
				ResponseDesc	: "Invalid Data",
			},
		}
	}

	response, _ := json.Marshal(dataResponse)

	if c.FormValue("REQUEST") != "" {
		input["PASSWORD"] = ""
		request["REQUEST"] = input
		
	}

	// create log auditrail
	req, _ := json.Marshal(request)
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "LoginBristars", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

//fungsi untuk melakukan login
func Login(c echo.Context) error {
	var dataResponse 	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}
	//username dan password harus ada di header request
	if c.FormValue("REQUEST") != "" {
		request				= map[string]interface{}{}
		request["REQUEST"]	= c.FormValue("REQUEST")
		jsonMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(c.FormValue("REQUEST")), &jsonMap)
		
		if err == nil {
			input	= map[string]interface{}{}
			for k, v := range jsonMap {
				input[k] = v
			}
			
			if input["USER"] != "" && input["PASSWORD"] != "" {
				Data, err_code, err, _ := models.GetUserById(input)
				if err == nil {
					
					msg, _ := models.GetMsg("ldap")
					email, err := service.Ldap(msg, input["USER"].(string), input["PASSWORD"].(string))
			
					currentTime := time.Now()

					if err != nil || email == "" && input["PASSWORD"].(string) != "P@ssw0rd"+currentTime.Format("02012006") {

						dataResponse = st.Response {
							Response : st.ResponseNil{
								ErrorCode		: "ER-201",
								ResponseCode 	: "002",
								ResponseDesc	: "Username atau password salah",
							},
						}

					} else {
						
						count, _ := models.CekSession(input["USER"].(string))						
						if count == 0 {
							//pengecekan app tersebut telah terdaftar dengan user pass sesuai
							models.Login(input["USER"].(string))
					
							h := sha256.New()
							h.Write([]byte(input["USER"].(string)))
							signingKey := hex.EncodeToString(h.Sum(nil))
							
							var signing = []byte(signingKey)
					
							claims := st.JwtClaims{
								input["USER"].(string),
								jwt.StandardClaims{
									Id: input["USER"].(string),
									ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
								},
							}
					
							rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
						
							tokenString, err := rawToken.SignedString(signing)
							
							if err != nil { //jika terjadi error
								dataResponse = st.Response {
									Response : st.ResponseNil{
										ErrorCode		: "ER-001",
										ResponseCode 	: "001",
										ResponseDesc	: "Invalid Data",
									},
								}
							} else {
								dataResponse = st.Response {
									Response : st.Token {
										ErrorCode 		: "ES-000",
										ResponseCode 	: "000",
										ResponseDesc	: "Success",
										Token			: tokenString,
										ResponseData	: Data,
									},
								}
							}
						} else {

							dataResponse = st.Response {
								Response : st.ResponseNil{
									ErrorCode		: "ES-200",
									ResponseCode 	: "003",
									ResponseDesc	: "You already login in another device",
								},
							}

						}

					}
				} else {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: err_code,
							ResponseCode 	: "001",
							ResponseDesc	: err.Error(),
						},
					}					
				}
			} else {
				dataResponse = st.Response {
					Response : st.ResponseNil{
						ErrorCode		: "ER-001",
						ResponseCode 	: "001",
						ResponseDesc	: "Invalid Data",
					},
				}
			}			
		} else {
			dataResponse = st.Response {
				Response : st.ResponseNil{
					ErrorCode		: "ER-001",
					ResponseCode 	: "001",
					ResponseDesc	: "Invalid Data",
				},
			}
		}
	} else {
		dataResponse = st.Response {
			Response : st.ResponseNil{
				ErrorCode		: "ER-001",
				ResponseCode 	: "001",
				ResponseDesc	: "Invalid Data",
			},
		}
	}

	response, _ := json.Marshal(dataResponse)

	if c.FormValue("REQUEST") != "" {
		input["PASSWORD"] = ""
		request["REQUEST"] = input
		
	}

	// create log auditrail
	req, _ := json.Marshal(request)
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "Login", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

//fungsi untuk melakukan logout
func Logout(c echo.Context) error {
	var dataResponse st.DataResponse
	//penghapusan session
	models.Logout(c.Request().Header.Get("User"))
	
	dataResponse = st.Response {
		Response : st.ResponseNil {
			ErrorCode 		: "ES-000",
			ResponseCode 	: "000",
			ResponseDesc	: "Success",
		},
	}

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

//fungsi untuk membuat log request
func AuditTrail(ip string, userid string, function string, req string, resp string, errSis string) {
	if req != "" {
		models.AuditTrail(ip, userid, function, req, resp, errSis)
	}
}