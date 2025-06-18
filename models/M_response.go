package models

import (
	"fmt"
	"strings"
	// "time"
	// sha256 "crypto/sha256"
	// hex "encoding/hex"
	str "strconv"
    db "api/uniform/config"
    st "api/uniform/struct"
	_ "assets/mysql"
)

func GetResponseAll(param map[string]interface{}) (st.Forms, int, string, error, error) {
	var Datas = st.Forms{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	selected_column := "a.id_form, a.judul_form, a.deskripsi_form, a.link_form, a.start_date, a.end_date, a.target, a.approval_form, a.entry_user, a.entry_name, a.entry_time, a.approval_posisi, a.approval_list, a.status, b.status_desc, a.is_active"
	
	sql_query 			:= "SELECT "+selected_column+" FROM form a INNER JOIN mst_status b ON a.status=b.id_status"
	sql_detail 			:= "SELECT id_form_pertanyaan, id_form, pertanyaan, pertanyaan_image, jenis_pertanyaan, list_pilihan_jawaban, mandatory FROM form_pertanyaan WHERE id_form=?"
	sql_count 			:= "SELECT count(1) FROM form a INNER JOIN mst_status b ON a.status=b.id_status"
	sql_count_response 	:= "SELECT count(1) FROM form_jawaban WHERE id_form=?"

	if param["Where"].(string) != "" {
		sql_query = sql_query + " where "+ param["Where"].(string)
		sql_count = sql_count + " where "+ param["Where"].(string)
	}
	
	if param["Order"].(string) != "" {
		sql_query = sql_query +" order by "+ param["Order"].(string)
	} else {
		sql_query = sql_query +" order by a.entry_time DESC"
	}

	if (param["Limit"].(string) != "" && param["Page"].(string) != "") && (param["Limit"].(string) !="0" && param["Page"].(string) !="0") {
		int_page, _ := str.Atoi(param["Page"].(string))
		int_limit, _ := str.Atoi(param["Limit"].(string))
		start_page := int_limit*(int_page-1)
		sql_query = sql_query+" limit "+str.Itoa(start_page)+","+param["Limit"].(string)
	} else if param["Limit"].(string) != "" && param["Limit"].(string) !="0" {
		sql_query = sql_query+" limit 0,"+param["Limit"].(string)
	}
	//query count data
	db.QueryRow(sql_count).Scan(&count)
	rows, err := db.Query(sql_query)

	if err != nil {
		return Datas, 0, "ER-099", fmt.Errorf("General Error"), err
	}
	
	defer rows.Close()
	for rows.Next() {
		var each st.Form
		var err = rows.Scan(&each.Id_form, &each.Judul_form, &each.Deskripsi_form, &each.Link_form, &each.Start_date, &each.End_date, &each.Target, &each.Approval_form, &each.Entry_user, &each.Entry_name, &each.Entry_time, &each.Approval_posisi, &each.Approval_list, &each.Status, &each.Status_desc, &each.Is_active)
		
		if err != nil {
			return Datas, 0, "", err, err
		}

		Form_pertanyaans := each.Form_pertanyaan

		rowsDetail, err := db.Query(sql_detail, each.Id_form)

		if err != nil {
			return Datas, 0, "", err, err
		}

		for rowsDetail.Next() {
			var each1 st.Form_pertanyaan
			var err = rowsDetail.Scan(&each1.Id_form_pertanyaan, &each1.Id_form, &each1.Pertanyaan, &each1.Pertanyaan_image, &each1.Jenis_pertanyaan, &each1.List_pilihan_jawaban, &each1.Mandatory)
			if err != nil {
				return Datas, 0, "", err, err
			}

			Form_pertanyaans = append(Form_pertanyaans, each1)
		}
		each.Form_pertanyaan = Form_pertanyaans

		var countResponse string
		db.QueryRow(sql_count_response, each.Id_form).Scan(&countResponse)
		each.Total_response = countResponse

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}
func GetSummaryJawabanList(param map[string]interface{}) (st.JawabanSummarys, string, error, error) {
	var Datas = st.JawabanSummarys{}
	db, err := db.Default()
	if err != nil {
		return Datas, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var countJawabanTotal int
	db.QueryRow("SELECT COUNT(*) FROM form_jawaban_detail WHERE id_form=? AND id_form_pertanyaan=?",param["IdForm"].(string),param["IdFormPertanyaan"].(string)).Scan(&countJawabanTotal)


	ArrList := param["ArrList"].([]interface{})
	limit	:= len(ArrList)

	for i := 0; i < int(limit); i++ {
		var each st.JawabanSummary

		each.Jawaban 		= ArrList[i].(string)
		each.JawabanTotal 	= countJawabanTotal

		var countJawabanMatch int
		qryCountJawabanMatch := "SELECT COUNT(*) FROM form_jawaban_detail WHERE id_form="+param["IdForm"].(string)+" AND id_form_pertanyaan="+param["IdFormPertanyaan"].(string)+" AND jawaban LIKE '%"+ArrList[i].(string)+"%'"
		db.QueryRow(qryCountJawabanMatch).Scan(&countJawabanMatch)

		each.JawabanMatch 			= countJawabanMatch
		each.JawabanMatchPercentage	= fmt.Sprintf("%.2f", ((float64(countJawabanMatch) / float64(countJawabanTotal)) * 100))

		Datas = append(Datas, each)
	}

	return Datas, "", err, err
}
func GetJawabanByFormId(param map[string]interface{}) (st.JawabanByIds, int, string, error, error) {
	var Datas = st.JawabanByIds{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	selected_column := "b.entry_user, b.entry_name, b.entry_time, a.id_form, a.id_form_pertanyaan, a.id_jawaban, a.id_jawaban_detail, a.jawaban"
	
	sql_count 	:= "SELECT count(1) FROM form_jawaban_detail a INNER JOIN form_jawaban b ON a.id_jawaban = b.id_jawaban"
	sql_query 	:= "SELECT "+selected_column+" FROM form_jawaban_detail a INNER JOIN form_jawaban b ON a.id_jawaban = b.id_jawaban"

	if param["Where"].(string) != "" {
		sql_count = sql_count + " where "+ param["Where"].(string)
		sql_query = sql_query + " where "+ param["Where"].(string)
	}
	
	if param["Order"].(string) != "" {
		sql_query = sql_query +" order by "+ param["Order"].(string)
	} else {
		sql_query = sql_query +" order by b.entry_time DESC"
	}

	if (param["Limit"].(string) != "" && param["Page"].(string) != "") && (param["Limit"].(string) !="0" && param["Page"].(string) !="0") {
		int_page, _ := str.Atoi(param["Page"].(string))
		int_limit, _ := str.Atoi(param["Limit"].(string))
		start_page := int_limit*(int_page-1)
		sql_query = sql_query+" limit "+str.Itoa(start_page)+","+param["Limit"].(string)
	} else if param["Limit"].(string) != "" && param["Limit"].(string) !="0" {
		sql_query = sql_query+" limit 0,"+param["Limit"].(string)
	}
	//query count data
	db.QueryRow(sql_count).Scan(&count)
	rows, err := db.Query(sql_query)

	if err != nil {
		return Datas, 0, "ER-099", fmt.Errorf("General Error"), err
	}
	
	defer rows.Close()
	for rows.Next() {
		var each st.JawabanById
		var err = rows.Scan(&each.Entry_user, &each.Entry_name, &each.Entry_time, &each.Id_form, &each.Id_form_pertanyaan, &each.Id_jawaban, &each.Id_jawaban_detail, &each.Jawaban)
		
		if err != nil {
			return Datas, 0, "", err, err
		}

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}
func GetJawabanAll(param map[string]interface{}) (st.Jawabans, int, string, error, error) {
	var Datas = st.Jawabans{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	selected_column := "a.id_form, c.judul_form, c.deskripsi_form, c.link_form, a.id_jawaban, a.entry_user, a.entry_name, a.entry_time, a.approval_posisi, a.approval_list, a.status, b.status_desc"
	
	sql_query 			:= "SELECT "+selected_column+" FROM form_jawaban a INNER JOIN mst_status b ON a.status=b.id_status INNER JOIN form c ON a.id_form = c.id_form"
	sql_detail 			:= "SELECT a.id_form, a.id_form_pertanyaan, b.pertanyaan, b.pertanyaan_image, b.jenis_pertanyaan, b.list_pilihan_jawaban, b.mandatory, a.id_jawaban, a.id_jawaban_detail, a.jawaban FROM form_jawaban_detail a INNER JOIN form_pertanyaan b ON a.id_form_pertanyaan=b.id_form_pertanyaan WHERE a.id_jawaban=?"
	sql_count 			:= "SELECT count(1) FROM form_jawaban a INNER JOIN mst_status b ON a.status=b.id_status INNER JOIN form c ON a.Id_form = c.Id_form"

	if param["Where"].(string) != "" {
		sql_query = sql_query + " where "+ param["Where"].(string)
		sql_count = sql_count + " where "+ param["Where"].(string)
	}
	
	if param["Order"].(string) != "" {
		sql_query = sql_query +" order by "+ param["Order"].(string)
	} else {
		sql_query = sql_query +" order by a.entry_time DESC"
	}

	if (param["Limit"].(string) != "" && param["Page"].(string) != "") && (param["Limit"].(string) !="0" && param["Page"].(string) !="0") {
		int_page, _ := str.Atoi(param["Page"].(string))
		int_limit, _ := str.Atoi(param["Limit"].(string))
		start_page := int_limit*(int_page-1)
		sql_query = sql_query+" limit "+str.Itoa(start_page)+","+param["Limit"].(string)
	} else if param["Limit"].(string) != "" && param["Limit"].(string) !="0" {
		sql_query = sql_query+" limit 0,"+param["Limit"].(string)
	}
	//query count data
	db.QueryRow(sql_count).Scan(&count)
	rows, err := db.Query(sql_query)

	if err != nil {
		return Datas, 0, "ER-099", fmt.Errorf("General Error"), err
	}
	
	defer rows.Close()
	for rows.Next() {
		var each st.Jawaban
		var err = rows.Scan(&each.Id_form, &each.Judul_form, &each.Deskripsi_form, &each.Link_form, &each.Id_jawaban, &each.Entry_user, &each.Entry_name, &each.Entry_time, &each.Approval_posisi, &each.Approval_list, &each.Status, &each.Status_desc)
		
		if err != nil {
			return Datas, 0, "", err, err
		}

		Jawaban_details := each.Jawaban_detail

		jawabanDetail, err := db.Query(sql_detail, each.Id_jawaban)

		if err != nil {
			return Datas, 0, "", err, err
		}

		for jawabanDetail.Next() {
			var each1 st.Jawaban_detail
			var err = jawabanDetail.Scan(&each1.Id_form,&each1.Id_form_pertanyaan,&each1.Pertanyaan,&each1.Pertanyaan_image,&each1.Jenis_pertanyaan,&each1.List_pilihan_jawaban,&each1.Mandatory,&each1.Id_jawaban,&each1.Id_jawaban_detail,&each1.Jawaban)
			
			if err != nil {
				return Datas, 0, "", err, err
			}

			Jawaban_details = append(Jawaban_details, each1)
		}
		each.Jawaban_detail = Jawaban_details

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}
func ApprovalJawaban(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(id_jawaban) FROM form_jawaban WHERE status=0 AND approval_posisi=? AND id_jawaban=?", param["Approval_now"].(string), param["Id_jawaban"].(string)).Scan(&countExist)

	if countExist > 0 {

		tx, err := db.Begin()
		if err != nil { 
			return err 
		}

		qry_update_form 					:= "UPDATE form_jawaban SET status=?, approval_posisi=?, approval_list=? WHERE status=0 AND approval_posisi=? AND id_jawaban=?"
		update_form, err_qry_update_form	:= tx.Prepare(qry_update_form)

		if err_qry_update_form == nil {

			defer update_form.Close()
			_, err_update_form := update_form.Exec(param["Status"].(string), param["Approval_next"].(string), param["Approval_list"].(string), param["Approval_now"].(string), param["Id_jawaban"].(string))

			if err_update_form == nil {

				tx.Commit()
				return nil

			} else {

				tx.Rollback()
				return err_update_form	

			}
		} else {

			tx.Rollback()
			return err_qry_update_form	

		}
	} else {

		return fmt.Errorf("Form tidak ditemukan")
		
	}
}

func UpdateJawabanForm(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(*) FROM form_jawaban WHERE id_form=? AND entry_user=? AND status=2", param["Id_form"].(string), param["Entry_user"].(string)).Scan(&countExist)
	if countExist > 0 {
		qry_delete_old_jawaban 			:= "DELETE FROM form_jawaban WHERE id_form=?"
		qry_delete_old_jawaban_detail	:= "DELETE FROM form_jawaban_detail WHERE id_form=?"
		qry_insert_jawaban 				:= "INSERT INTO form_jawaban (id_form, id_jawaban, entry_user, entry_name, entry_time, approval_posisi, approval_list) VALUES (?,?,?,?,NOW(),?,?)"
		qry_insert_jawaban_detail		:= "INSERT INTO form_jawaban_detail (id_form, id_form_pertanyaan, id_jawaban, id_jawaban_detail, jawaban) VALUES %s"

		valueStrings	:= []string{}
		valueArgs 		:= []interface{}{}

		ArrJawaban 	:= param["Arr_jawaban"].([]interface{})
		limit		:= len(ArrJawaban)

		for i := 0; i < int(limit); i++ {
			DataDetail 		:= ArrJawaban[i].(map[string]interface{})
			valueStrings 	= append(valueStrings, "(?, ?, ?, ?, ?)")
			valueArgs 		= append(valueArgs,DataDetail["id_form"].(string))
			valueArgs 		= append(valueArgs,DataDetail["id_form_pertanyaan"].(string))
			valueArgs 		= append(valueArgs,DataDetail["id_jawaban"].(string))
			valueArgs 		= append(valueArgs,DataDetail["id_jawaban_detail"].(string))
			valueArgs 		= append(valueArgs,DataDetail["Jawaban"].(string))
		}
		qry_insert_jawaban_detail = fmt.Sprintf(qry_insert_jawaban_detail, strings.Join(valueStrings, ","))

		tx, err := db.Begin()
		if err != nil { 		
			return err 
		}

		delete_old_jawaban, err_qry_delete_old_jawaban					:= tx.Prepare(qry_delete_old_jawaban)
		delete_old_jawaban_detail, err_qry_delete_old_jawaban_detail	:= tx.Prepare(qry_delete_old_jawaban_detail)
		insert_jawaban, err_qry_insert_jawaban 							:= tx.Prepare(qry_insert_jawaban)
		insert_jawaban_detail, err_qry_insert_jawaban_detail 			:= tx.Prepare(qry_insert_jawaban_detail)

		if err_qry_delete_old_jawaban == nil && err_qry_delete_old_jawaban_detail == nil && err_qry_insert_jawaban == nil && err_qry_insert_jawaban_detail == nil {
			defer delete_old_jawaban.Close()
			defer delete_old_jawaban_detail.Close()
			defer insert_jawaban.Close()
			defer insert_jawaban_detail.Close()

			_, err_delete_old_jawaban 			:= delete_old_jawaban.Exec(param["Id_form"].(string))
			_, err_delete_old_jawaban_detail 	:= delete_old_jawaban_detail.Exec(param["Id_form"].(string))
			_, err_insert_jawaban 				:= insert_jawaban.Exec(param["Id_form"].(string), param["Id_jawaban"].(string), param["Entry_user"].(string), param["Entry_name"].(string), param["Approval_posisi"].(string), param["Approval_list"].(string))
			_, err_insert_jawaban_detail 		:= insert_jawaban_detail.Exec(valueArgs...)

			if err_delete_old_jawaban == nil && err_delete_old_jawaban_detail == nil && err_insert_jawaban == nil && err_insert_jawaban_detail == nil {
				tx.Commit()
				return nil
			} else {
				if err_delete_old_jawaban != nil {
					err = err_delete_old_jawaban
				}

				if err_delete_old_jawaban_detail != nil {
					err = err_delete_old_jawaban_detail
				}

				if err_insert_jawaban != nil {
					err = err_insert_jawaban
				}
		
				if err_insert_jawaban_detail != nil {
					err = err_insert_jawaban_detail
				}
		
				tx.Rollback()
				return err
			}
		} else {
			if err_qry_delete_old_jawaban != nil {
				err = err_qry_delete_old_jawaban
			}

			if err_qry_delete_old_jawaban_detail != nil {
				err = err_qry_delete_old_jawaban_detail
			}

			if err_qry_insert_jawaban != nil {
				err = err_qry_insert_jawaban
			}

			if err_qry_insert_jawaban_detail != nil {
				err = err_qry_insert_jawaban_detail
			}
			
			tx.Rollback()
			return err
		}
		
	} else {
		return fmt.Errorf("Gagal menyimpan form! Form tidak ditemukan!")
	}
}

func RequestDownloadAttachment(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(*) FROM form_request_attachment WHERE id_form=? AND request_user=?", param["Id_form"].(string), param["Request_user"].(string)).Scan(&countExist)
	if countExist == 0 {
		qry_insert_request := "INSERT INTO form_request_attachment (request_user, request_name, request_time, id_form, judul_form, status) VALUES (?,?,NOW(),?,?,0)"

		tx, err := db.Begin()
		if err != nil { 		
			return err 
		}

		insert_request, err_qry_insert_request := tx.Prepare(qry_insert_request)		

		if err_qry_insert_request == nil {
			defer insert_request.Close()
			_, err_insert_request := insert_request.Exec(param["Request_user"].(string), param["Request_name"].(string), param["Id_form"].(string), param["Judul_form"].(string))

			if err_insert_request == nil {
				tx.Commit()
				return nil
			} else {
				if err_insert_request != nil {
					err = err_insert_request
				}
		
				tx.Rollback()
				return err
			}
		} else {
			if err_qry_insert_request != nil {
				err = err_qry_insert_request
			}
			
			tx.Rollback()
			return err
		}
		
	} else {
		return fmt.Errorf("Gagal request attachment form! Maksimal 1X Request!")
	}
}
func GetRequestAttachmentAll(param map[string]interface{}) (st.DataRequestAttachments, int, string, error, error) {
	var Datas = st.DataRequestAttachments{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	selected_column := "id, request_user, request_name, request_time, id_form, judul_form, IFNULL(url_attachment, '') url_attachment, status"
	
	sql_query 			:= "SELECT "+selected_column+" FROM form_request_attachment"
	sql_count 			:= "SELECT count(1) FROM form_request_attachment"

	if param["Where"].(string) != "" {
		sql_query = sql_query + " where "+ param["Where"].(string)
		sql_count = sql_count + " where "+ param["Where"].(string)
	}
	
	if param["Order"].(string) != "" {
		sql_query = sql_query +" order by "+ param["Order"].(string)
	} else {
		sql_query = sql_query +" order by id DESC"
	}

	if (param["Limit"].(string) != "" && param["Page"].(string) != "") && (param["Limit"].(string) !="0" && param["Page"].(string) !="0") {
		int_page, _ := str.Atoi(param["Page"].(string))
		int_limit, _ := str.Atoi(param["Limit"].(string))
		start_page := int_limit*(int_page-1)
		sql_query = sql_query+" limit "+str.Itoa(start_page)+","+param["Limit"].(string)
	} else if param["Limit"].(string) != "" && param["Limit"].(string) !="0" {
		sql_query = sql_query+" limit 0,"+param["Limit"].(string)
	}
	//query count data
	db.QueryRow(sql_count).Scan(&count)
	rows, err := db.Query(sql_query)

	if err != nil {
		return Datas, 0, "ER-099", fmt.Errorf("General Error"), err
	}
	
	defer rows.Close()
	for rows.Next() {
		var each st.DataRequestAttachment
		var err = rows.Scan(&each.Id, &each.Request_user, &each.Request_name, &each.Request_time, &each.Id_form, &each.Judul_form, &each.Url_attachment, &each.Status)
		
		if err != nil {
			return Datas, 0, "", err, err
		}

		if each.Status == "0" {
			each.StatusDesc = "Proses"
		} else {
			each.StatusDesc = "Done"
		}

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}

func JobGetRequestAttachmentDetail(param map[string]interface{}) (st.DataJobRequestAttachments, string, error, error) {
	var Datas = st.DataJobRequestAttachments{}
	db, err := db.Default()
	if err != nil {
		return Datas, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()
	
	sql_query 			:= "SELECT id, request_user, request_name, request_time, id_form, judul_form, IFNULL(url_attachment, '') url_attachment FROM form_request_attachment WHERE status = 0 ORDER BY request_time ASC"
	sql_query_detail	:= "SELECT c.entry_name, a.jawaban FROM form_jawaban_detail a INNER JOIN form_pertanyaan b ON a.id_form_pertanyaan = b.id_form_pertanyaan INNER JOIN form_jawaban c ON a.id_jawaban = c.id_jawaban WHERE b.jenis_pertanyaan = 5 AND a.id_form = ?"
	rows, err := db.Query(sql_query)
	if err != nil {
		return Datas, "ER-099", fmt.Errorf("General Error"), err
	}

	defer rows.Close()
	for rows.Next() {
		var each st.DataJobRequestAttachment
		var err = rows.Scan(&each.Id, &each.Request_user, &each.Request_name, &each.Request_time, &each.Id_form, &each.Judul_form, &each.Url_attachment)
		
		if err != nil {
			return Datas, "", err, err
		}

		Request_details := each.Request_detail

		rowsDetail, err := db.Query(sql_query_detail, each.Id_form)
		if err != nil {
			return Datas, "", err, err
		}
		for rowsDetail.Next() {
			var each1 st.DataJobRequestAttachmentDetail
			var err = rowsDetail.Scan(&each1.Entry_name, &each1.Jawaban)
			if err != nil {
				return Datas, "", err, err
			}

			Request_details = append(Request_details, each1)
		}
		each.Request_detail = Request_details

		Datas = append(Datas, each)
	}
	
	return Datas, "", err, err
}
func JobUpdateRequestAttachment(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(id) FROM form_request_attachment WHERE status=0 AND id=?", param["Id"].(string)).Scan(&countExist)

	if countExist > 0 {

		tx, err := db.Begin()
		if err != nil { 
			return err 
		}

		qry_update_request 						:= "UPDATE form_request_attachment SET url_attachment=?, status=? WHERE status=0 AND id=?"
		update_request, err_qry_update_request	:= tx.Prepare(qry_update_request)

		if err_qry_update_request == nil {

			defer update_request.Close()
			_, err_update_request := update_request.Exec(param["Url_attachment"].(string), param["Status"].(float64), param["Id"].(string))
			if err_update_request == nil {
				tx.Commit()
				return nil
			} else {
				tx.Rollback()
				return err_update_request	
			}
		} else {
			tx.Rollback()
			return err_qry_update_request	
		}
	} else {
		return fmt.Errorf("Request tidak ditemukan")
	}
}