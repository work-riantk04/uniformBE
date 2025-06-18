package controllers

import (
	// "fmt"
	// "strings"
	"encoding/json"
    "net/http"
    models "api/uniform/models"
    st "api/uniform/struct"
	"github.com/labstack/echo"
)

func GetMenu(c echo.Context) error {
	result, err := models.GetMenu()
	
	var dataResponse 	st.DataResponse

	if err != nil {
		dataResponse = st.Response {
			Response : st.ResponseNil{
				ErrorCode		: "ER-004",
				ResponseCode 	: "004",
				ResponseDesc	: err.Error(),
			},
		}
	}
	
	if result != nil {	
		var data st.Menu
		var datas st.Menus

		for _, n := range result {
			if n.Menu_id != "" {
				data.Menu_id 		= n.Menu_id
				data.Menu_parent 	= n.Menu_parent
				data.Menu_icon 		= n.Menu_icon
				data.Menu_title 	= n.Menu_title
				data.Menu_link 		= n.Menu_link
				data.Level_user 	= n.Level_user
			} else {
				data.Menu_id 		= ""
				data.Menu_parent 	= ""
				data.Menu_icon 		= ""
				data.Menu_title 	= ""
				data.Menu_link 		= ""
				data.Level_user 	= ""
			}

			datas = append(datas, data)
		}

		dataResponse = st.Response {
			Response : st.ResponseMenu {
				ErrorCode 		: "ES-000",
				ResponseCode 	: "000",
				ResponseDesc	: "Success",
				ResponseData	: datas,
			},
		}

		return c.JSONPretty(http.StatusOK, dataResponse, "  ")
	} else {
		dataResponse = st.Response {
			Response : st.ResponseNil{
				ErrorCode		: "ER-001",
				ResponseCode 	: "001",
				ResponseDesc	: "Data not found",
			},
		}
		
		return c.JSONPretty(http.StatusOK, dataResponse, "  ")

	}
}

func GetUserAll(c echo.Context) error {
	var dataResponse	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

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
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Limit")
			mandatory_input = append(mandatory_input,"User")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				
				Datas, count, errorCode, err, errs := models.GetUserAll(input)
				if err != nil {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-001",
							ResponseCode 	: errorCode,
							ResponseDesc	: err.Error()+ " - "+errs.Error(),
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.DataAll {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
							ResponseCount	: count,
							ResponseData	: Datas,
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

	// create log auditrail
	// response, _ 	:= json.Marshal(dataResponse)
	// req, _ 		:= json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetUserAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func GetUserById(c echo.Context) error {
	var dataResponse	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

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
			input["User"] = c.Request().Header.Get("User")

			var mandatory_input []string
			mandatory_input = append(mandatory_input,"PERNR")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				Datas, _, err, _ := models.GetUserById(input)
				if err != nil {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-004",
							ResponseCode 	: "004",
							ResponseDesc	: err.Error(),
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.Data {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
							ResponseData	: Datas,
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

	// create log auditrail
	// response, _ 	:= json.Marshal(dataResponse)
	// req, _ 		:= json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetUserById", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func GetDivisiAll(c echo.Context) error {
	var dataResponse	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

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
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Limit")
			mandatory_input = append(mandatory_input,"User")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				
				Datas, count, errorCode, err, errs := models.GetDivisiAll(input)
				if err != nil {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-001",
							ResponseCode 	: errorCode,
							ResponseDesc	: err.Error()+ " - "+errs.Error(),
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.DataAll {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
							ResponseCount	: count,
							ResponseData	: Datas,
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

	// create log auditrail
	// response, _ 	:= json.Marshal(dataResponse)
	// req, _ 		:= json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetDivisiAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func GetStatus(c echo.Context) error {
	var dataResponse	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

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
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"User")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				
				Datas, count, errorCode, err, errs := models.GetStatus(input)
				if err != nil {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-001",
							ResponseCode 	: errorCode,
							ResponseDesc	: err.Error()+ " - "+errs.Error(),
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.DataAll {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
							ResponseCount	: count,
							ResponseData	: Datas,
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

	// create log auditrail
	// response, _ 	:= json.Marshal(dataResponse)
	// req, _ 		:= json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetStatus", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func GetParamAll(c echo.Context) error {
	var dataResponse	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

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
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Limit")
			mandatory_input = append(mandatory_input,"User")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				
				Datas, count, errorCode, err, errs := models.GetParamAll(input)
				if err != nil {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-001",
							ResponseCode 	: errorCode,
							ResponseDesc	: err.Error()+ " - "+errs.Error(),
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.DataAll {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
							ResponseCount	: count,
							ResponseData	: Datas,
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

	// create log auditrail
	// response, _ 	:= json.Marshal(dataResponse)
	// req, _ 		:= json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetParamAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func GetSummary(c echo.Context) error {
	var dataResponse	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

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
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"User")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				
				Datas, count, errorCode, err, errs := models.GetSummary(input)
				if err != nil {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-001",
							ResponseCode 	: errorCode,
							ResponseDesc	: err.Error()+ " - "+errs.Error(),
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.DataAll {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
							ResponseCount	: count,
							ResponseData	: Datas,
						},
					}
				}
			} else {
				dataResponse = st.Response {
					Response : st.ResponseNil{
						ErrorCode		: "ER-001",
						ResponseCode 	: "001",
						ResponseDesc	: "Invalid Data1",
					},
				}
			}
		} else {
			dataResponse = st.Response {
				Response : st.ResponseNil{
					ErrorCode		: "ER-001",
					ResponseCode 	: "001",
					ResponseDesc	: "Invalid Data2",
				},
			}
		}
	} else {
		dataResponse = st.Response {
			Response : st.ResponseNil{
				ErrorCode		: "ER-001",
				ResponseCode 	: "001",
				ResponseDesc	: "Invalid Data3",
			},
		}
	}	

	// create log auditrail
	// response, _ 	:= json.Marshal(dataResponse)
	// req, _ 		:= json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetSummary", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}