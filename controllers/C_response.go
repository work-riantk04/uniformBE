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

func GetResponseAll(c echo.Context) error {
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
				
				Datas, count, errorCode, err, errs := models.GetResponseAll(input)
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
	// response, _ := json.Marshal(dataResponse)

	// create log auditrail
	// req, _ := json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetFormAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func GetSummaryJawabanList(c echo.Context) error {
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
			mandatory_input = append(mandatory_input,"IdForm")
			mandatory_input = append(mandatory_input,"IdFormPertanyaan")
			mandatory_input = append(mandatory_input,"ArrList")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				
				Datas, errorCode, err, errs := models.GetSummaryJawabanList(input)
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
	// response, _ := json.Marshal(dataResponse)

	// create log auditrail
	// req, _ := json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetSummaryJawabanList", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func GetJawabanByFormId(c echo.Context) error {
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
				
				Datas, count, errorCode, err, errs := models.GetJawabanByFormId(input)
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
	// response, _ := json.Marshal(dataResponse)

	// create log auditrail
	// req, _ := json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetFormAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func GetJawabanAll(c echo.Context) error {
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
				
				Datas, count, errorCode, err, errs := models.GetJawabanAll(input)
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
	// response, _ := json.Marshal(dataResponse)

	// create log auditrail
	// req, _ := json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetFormAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func ApprovalJawaban(c echo.Context) error {
	var dataResponse 	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

	if c.FormValue("REQUEST") != "" {
		request				= map[string]interface{}{}
		request["REQUEST"]	= c.FormValue("REQUEST")
		jsonMap 			:= make(map[string]interface{})
		err 				:= json.Unmarshal([]byte(c.FormValue("REQUEST")), &jsonMap)
		
		if err == nil {
			input = map[string]interface{}{}
			for k, v := range jsonMap {
				input[k] = v
			}
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Id_jawaban")
			mandatory_input = append(mandatory_input,"Approval_now")
			mandatory_input = append(mandatory_input,"Approval_next")
			mandatory_input = append(mandatory_input,"Approval_list")
			mandatory_input = append(mandatory_input,"Status")

			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" || input[val] == nil {
					cek_mandatory++
				}
			}
			
			if cek_mandatory == 0 {

				err := models.ApprovalJawaban(input)

				if err == nil {
					dataResponse = st.Response {
						Response : st.ResponseNil {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-004",
							ResponseCode 	: "004",
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

	// create log auditrail
	req, _ := json.Marshal(request)
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "ApprovalJawaban", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func UpdateJawabanForm(c echo.Context) error {
	var dataResponse 	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

	if c.FormValue("REQUEST") != "" {
		request				= map[string]interface{}{}
		request["REQUEST"]	= c.FormValue("REQUEST")
		jsonMap 			:= make(map[string]interface{})
		err 				:= json.Unmarshal([]byte(c.FormValue("REQUEST")), &jsonMap)
		
		if err == nil {
			input = map[string]interface{}{}
			for k, v := range jsonMap {
				input[k] = v
			}
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Id_form")
			mandatory_input = append(mandatory_input,"Id_jawaban")
			mandatory_input = append(mandatory_input,"Entry_user")
			mandatory_input = append(mandatory_input,"Entry_name")

			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" || input[val] == nil {
					cek_mandatory++
				}
			}
			
			if cek_mandatory == 0 {

				err := models.UpdateJawabanForm(input)

				if err == nil {
					dataResponse = st.Response {
						Response : st.ResponseNil {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-004",
							ResponseCode 	: "004",
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

	// create log auditrail
	req, _ := json.Marshal(request)
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "UpdateJawabanForm", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func RequestDownloadAttachment(c echo.Context) error {
	var dataResponse 	st.DataResponse
	var request			map[string]interface{}
	var input 			map[string]interface{}

	if c.FormValue("REQUEST") != "" {
		request				= map[string]interface{}{}
		request["REQUEST"]	= c.FormValue("REQUEST")
		jsonMap 			:= make(map[string]interface{})
		err 				:= json.Unmarshal([]byte(c.FormValue("REQUEST")), &jsonMap)
		
		if err == nil {
			input = map[string]interface{}{}
			for k, v := range jsonMap {
				input[k] = v
			}
			input["User"] = c.Request().Header.Get("User")
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Request_user")
			mandatory_input = append(mandatory_input,"Request_name")
			mandatory_input = append(mandatory_input,"Id_form")
			mandatory_input = append(mandatory_input,"Judul_form")

			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" || input[val] == nil {
					cek_mandatory++
				}
			}
			
			if cek_mandatory == 0 {

				err := models.RequestDownloadAttachment(input)

				if err == nil {
					dataResponse = st.Response {
						Response : st.ResponseNil {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
						},
					}
				} else {
					dataResponse = st.Response {
						Response : st.ResponseNil{
							ErrorCode		: "ER-004",
							ResponseCode 	: "004",
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

	// create log auditrail
	req, _ := json.Marshal(request)
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "RequestDownloadAttachment", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func GetRequestAttachmentAll(c echo.Context) error {
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
				
				Datas, count, errorCode, err, errs := models.GetRequestAttachmentAll(input)
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
				ResponseDesc	: "Invalid Data",
			},
		}
	}
	// response, _ := json.Marshal(dataResponse)

	// create log auditrail
	// req, _ := json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetFormAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func JobGetRequestAttachmentDetail(c echo.Context) error {
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
			
			// set mandatory input
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Source")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 && input["Source"] == "Job" {		
				
				Datas, errorCode, err, errs := models.JobGetRequestAttachmentDetail(input)
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
	// response, _ := json.Marshal(dataResponse)

	// create log auditrail
	// req, _ := json.Marshal(request)
	// AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "GetFormAll", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func JobUpdateRequestAttachment(c echo.Context) error {
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
			
			// set mandatory input
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Id")
            mandatory_input = append(mandatory_input,"Url_attachment")
            mandatory_input = append(mandatory_input,"Status")
			// set mandatory input
			
			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" {
					cek_mandatory++
				}
			}

			if cek_mandatory == 0 {		
				
				err := models.JobUpdateRequestAttachment(input)
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
						Response : st.ResponseNil {
							ErrorCode 		: "ES-000",
							ResponseCode 	: "000",
							ResponseDesc	: "Success",
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

	// create log auditrail
	req, _ := json.Marshal(request)
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "JobUpdateRequestAttachment", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}