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

func InsertForm(c echo.Context) error {
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

			errCheckIsDraft := 0
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Id_form")
			mandatory_input = append(mandatory_input,"Entry_user")
			mandatory_input = append(mandatory_input,"Entry_name")
			mandatory_input = append(mandatory_input,"Is_draft")

			if input["Is_draft"] == "1" {

				// set mandatory input
				mandatory_input = append(mandatory_input,"Judul_form")
				mandatory_input = append(mandatory_input,"Deskripsi_form")
				mandatory_input = append(mandatory_input,"Link_form")
				mandatory_input = append(mandatory_input,"Start_date")
				mandatory_input = append(mandatory_input,"End_date")
				mandatory_input = append(mandatory_input,"Target")
				mandatory_input = append(mandatory_input,"Approval_form")
				mandatory_input = append(mandatory_input,"List_pertanyaan")
				mandatory_input = append(mandatory_input,"Approval_posisi")
				mandatory_input = append(mandatory_input,"Approval_list")
				// set mandatory input

			} else if input["Is_draft"] == "2" {
				mandatory_input = append(mandatory_input,"Judul_form")
			} else {
				errCheckIsDraft = errCheckIsDraft + 1
			}

			if errCheckIsDraft == 0 {

				cek_mandatory := 0
				for _, val := range mandatory_input {
					if input[val] == "" || input[val] == nil {
						cek_mandatory++
					}
				}
				
				if cek_mandatory == 0 {

					err := models.InsertForm(input)

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
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "InsertForm", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func GetFormAll(c echo.Context) error {
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
				
				Datas, count, errorCode, err, errs := models.GetFormAll(input)
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
func UpdateForm(c echo.Context) error {
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

			errCheckIsDraft := 0
			
			var mandatory_input []string
			mandatory_input = append(mandatory_input,"Id_form_existing")
			mandatory_input = append(mandatory_input,"Id_form")
			mandatory_input = append(mandatory_input,"Entry_user")
			mandatory_input = append(mandatory_input,"Entry_name")
			mandatory_input = append(mandatory_input,"Is_draft")

			if input["Is_draft"] == "1" {

				// set mandatory input
				mandatory_input = append(mandatory_input,"Judul_form")
				mandatory_input = append(mandatory_input,"Deskripsi_form")
				mandatory_input = append(mandatory_input,"Link_form")
				mandatory_input = append(mandatory_input,"Start_date")
				mandatory_input = append(mandatory_input,"End_date")
				mandatory_input = append(mandatory_input,"Target")
				mandatory_input = append(mandatory_input,"Approval_form")
				mandatory_input = append(mandatory_input,"List_pertanyaan")
				mandatory_input = append(mandatory_input,"Approval_posisi")
				mandatory_input = append(mandatory_input,"Approval_list")
				// set mandatory input

			} else if input["Is_draft"] == "2" {
				mandatory_input = append(mandatory_input,"Judul_form")
			} else {
				errCheckIsDraft = errCheckIsDraft + 1
			}

			if errCheckIsDraft == 0 {

				cek_mandatory := 0
				for _, val := range mandatory_input {
					if input[val] == "" || input[val] == nil {
						cek_mandatory++
					}
				}
				
				if cek_mandatory == 0 {

					err := models.UpdateForm(input)

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
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "UpdateForm", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func ApprovalForm(c echo.Context) error {
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

				err := models.ApprovalForm(input)

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
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "ApprovalForm", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}
func ActiveForm(c echo.Context) error {
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
			mandatory_input = append(mandatory_input,"Key")

			cek_mandatory := 0
			for _, val := range mandatory_input {
				if input[val] == "" || input[val] == nil {
					cek_mandatory++
				}
			}
			
			if cek_mandatory == 0 {

				err := models.ActiveForm(input)

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
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "ActiveForm", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}

func InsertJawabanForm(c echo.Context) error {
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

				err := models.InsertJawabanForm(input)

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
	AuditTrail(c.RealIP(), c.Request().Header.Get("User"), "InsertJawabanForm", string(req), string(response), "")
	// create log auditrail

	return c.JSONPretty(http.StatusOK, dataResponse, "  ")
}