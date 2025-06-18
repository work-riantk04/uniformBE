package service

import (
	// "encoding/json"
	"fmt"
	"net/http"
    // "net/url"
	"net/http/httputil"
	"io/ioutil"
	"encoding/xml"
	"bytes"
	"crypto/tls"
	"strings"
	// str "strconv"
    models "api/uniform/models"
    // "crypto/md5"
	// "encoding/hex"
)

type MyRespEnvelope struct {
    XMLName xml.Name
    Body    Body
}

type Body struct {
    XMLName     xml.Name
    GetResponse Getda `xml:"ns1:validate_aduserResponse"`
}

type Getda struct {
    XMLName     xml.Name
    GetComplete completeResponse `xml:"WSSendSMSResult"`
}

type completeResponse struct {
	Status             string   `xml:"Status"`
	Reason             string   `xml:"Reason"`
	TimeRequest        string   `xml:"TimeRequest"`
	TimeReply          string   `xml:"TimeReply"`
}

/*
func GetSessionBristars(pernr string, key string) (string, error){
	type Response 	map[string]interface{}
	var DataUser 	string
	msgkey 			:= "password"
	msg 			:= map[string]interface{}{}
	msg["pernr"]	= pernr
	msg[msgkey]		= key
	request, _		:= json.Marshal(msg)

	userid, userpass, userurl, _, err, _ := models.GetUrl("digest", "", "")
	
	uri 	:= "/bristars_api/api/digest/pekerja/login"
	user	:= userid
	keyapp	:= userpass
	realm	:= "API BRISTARS"
	nonce	:= ""
	nc		:= "" 
	cnonce	:= "" 
	qop		:= ""

	hash_key := md5.New()
    hash_key.Write([]byte(user+":"+realm+":"+keyapp))
	resp_key := hex.EncodeToString(hash_key.Sum(nil))

	hash_method := md5.New()
    hash_method.Write([]byte("POST:"+uri))
	resp_method := hex.EncodeToString(hash_method.Sum(nil))
	hasher := md5.New()
    hasher.Write([]byte(resp_key+":"+nonce+":"+nc+":"+cnonce+":"+qop+":"+resp_method))
	response := hex.EncodeToString(hasher.Sum(nil))
	
	auth := "Digest username='"+ user +"', realm='API BRISTARS', nonce='', nc='', cnonce='', qop='', uri='"+uri+"', algorithm='MD5', response='"+response+"'"

	req, err := http.NewRequest("POST", userurl, bytes.NewBuffer(request))
	//req, err := http.NewRequest("POST", "http://10.35.65.148:6001/invoke/bri.core.services.v1:inquiryGL", bytes.NewBuffer(request))
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header["Authorization"] = []string{auth}
	req.Header["Cache-Control"] = []string{"no-cache"}
	if err != nil {
		return DataUser, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return DataUser, err
	}
	
	responseBody := make(Response)
	// err = json.NewDecoder(resp.Body).Decode(&responseBody)
	// if err != nil {
	// 	return data, err
	// }
	// responseData	:= responseBody["RESPONSE"].(map[string]interface{})

	// if responseData["RESPONSE_CODE"].(string) == "00" {
	// 	data.ResponseCode	= responseData["RESPONSE_CODE"].(string)
	// 	data.ResponseDesc	= responseData["RESPONSE_MESSAGE"].(string)
	// 	data.AcctNo			= responseData["ACCT_NO"].(string)
	// 	data.Name 			= responseData["DESCRIPTION"].(string)
	// 	data.Currency		= responseData["CURRENCY"].(string)
	// 	data.Balance		= ""
	// } else {
	// 	data.ResponseCode	= responseData["RESPONSE_CODE"].(string)
	// 	data.ResponseDesc	= responseData["RESPONSE_MESSAGE"].(string)
	// 	data.AcctNo			= param["rekening"].(string)
	// 	data.Name 			= ""
	// 	data.Currency		= ""
	// 	data.Balance		= ""
	// }

	return DataUser, nil
}

func BristarsAppCheck(param map[string]interface{}) (string, string, error){
	type Response 	map[string]interface{}
	var status 		string
	var pernr		string

	_, _, apiurl, _, _, _ 	:= models.GetUrl("appcheck", "", "")

    data := url.Values{}
    data.Set("key", param["Key"].(string))
    data.Set("user", param["UserApp"].(string))
    data.Set("app_id", param["AppId"].(string))

	u, _ := url.ParseRequestURI(apiurl)
    urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	
	if err == nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")	
		req.Header.Add("Content-Length", str.Itoa(len(data.Encode())))

		client := &http.Client{}
		resp, err := client.Do(req)
		
		if err == nil {
			if resp != nil {
				defer resp.Body.Close()
				responseBody := make(Response)
				err = json.NewDecoder(resp.Body).Decode(&responseBody)
				
				if err == nil {
					if responseBody["status"].(string) == "Success" {
						dataProfile		:= responseBody["profile"].(map[string]interface{})

						pernr			= dataProfile["pernr"].(string)
						// dataResp.SNAME			= dataProfile["nama"].(string)
						// dataResp.JGPG			= dataProfile["jgpg"].(string)
						// dataResp.ESELON			= dataProfile["eselon"].(string)
						// dataResp.WERKS 			= ""
						// dataResp.WERKS_TX 		= ""
						// dataResp.BTRTL 			= ""
						// dataResp.BTRTL_TX 		= ""
						// dataResp.KOSTL 			= dataProfile["cost_center"].(string)
						// dataResp.KOSTL_TX 		= dataProfile["desc_cost_center"].(string)
						// dataResp.ORGEH 			= dataProfile["organisasi_unit"].(string)
						// dataResp.ORGEH_TX 		= ""
						// dataResp.STELL 			= dataProfile["stell"].(string)
						// dataResp.STELL_TX 		= dataProfile["stell_tx"].(string)
						// dataResp.PLANS 			= dataProfile["plans"].(string)
						// dataResp.PLANS_TX 		= dataProfile["plans_tx"].(string)
						// dataResp.HILFM			= dataProfile["hilfm"].(string)
						// dataResp.HTEXT			= dataProfile["htext"].(string)
						// dataResp.BRANCH 		= dataProfile["branch_code"].(string)
						// dataResp.MAINBR 		= ""
						// dataResp.IS_PEMIMPIN 	= ""
						// dataResp.ADMIN_LEVEL 	= ""
						// dataResp.ORGEH_PGS 		= dataProfile["organisasi_unit_pgs"].(string)
						// dataResp.ORGEH_PGS_TX 	= ""
						// dataResp.PLANS_PGS 		= dataProfile["plans_pgs"].(string)
						// dataResp.PLANS_PGS_TX 	= dataProfile["plans_pgs_tx"].(string)
						// dataResp.BRANCH_PGS 		= dataProfile["branch_code_pgs"].(string)
						// dataResp.HILFM_PGS 		= dataProfile["hilfm_pgs"].(string)
						// dataResp.HTEXT_PGS 		= dataProfile["htext_pgs"].(string)
						// dataResp.TIPE_UKER 		= dataProfile["tipe_uker"].(string)
						// dataResp.REKENING 		= dataProfile["pernr"].(string)
						// dataResp.NPWP			= dataProfile["pernr"].(string)
						// dataResp.REGION 		= ""
						// dataResp.RGDESC 		= ""
						// dataResp.BRDESC 		= ""
						// dataResp.MBDESC			= ""

						status 	= responseBody["status"].(string)
					} else {
						status 	= responseBody["status"].(string)
						err		= fmt.Errorf(responseBody["message"].(string))
					}
				} else {
					status  = "error"
					err		= fmt.Errorf("error session bristars")
				}
			} else {
				status 	= "error"
				err		= fmt.Errorf("error session bristars")
			}
			
		} else {
			status 	= "error"
			err		= fmt.Errorf("error session bristars")
		}
	} else {
		err		= fmt.Errorf("error session bristars")
		status 	= "error"
	}
	
	return status, pernr, err
}
*/

func Ldap(msgService string, pernr string, key string) (string, error){
	// Responses := Response{}

	_, _, userurl, _, _, _ := models.GetUrl("ldap", "", "")

	msg := fmt.Sprintf(msgService, pernr, key)
	
	param := []byte(msg)
	
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	
	req, err := http.NewRequest("POST", userurl, bytes.NewBuffer(param))
	if err != nil {
		return "", err
	}
	req.Header["Content-Type"] = []string{"text/xml"}
	
	_, err = httputil.DumpRequestOut(req, true)
	if err != nil {
		return "", err
	}

	client := &http.Client{Transport: tr}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	
	if err != nil {
			return "", err
	}	

	defer response.Body.Close()
	Soap 		:= []byte(bodyBytes)
	email_start := strings.Index(string(Soap), "<return")
	email_end 	:= strings.Index(string(Soap), "</return>")

	return string(Soap)[email_start+30:email_end], nil
	
}