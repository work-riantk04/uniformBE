package models

import (
	str "strconv"
    db 	"api/uniform/config"
    st 	"api/uniform/struct"
	_ 	"assets/mysql"

	"fmt"
	"strings"
	
	// "time"
	// sha256 "crypto/sha256"
	// hex "encoding/hex"
)

func InsertForm(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	if param["Is_draft"].(string) == "1" { //submit
		var countExist int
		db.QueryRow("SELECT count(id_form) FROM form WHERE link_form=?", param["Link_form"].(string)).Scan(&countExist)
		if countExist > 0 {
			return fmt.Errorf("Gagal menyimpan form! Link sudah digunakan!")
		}
	}

	qry_insert_form := "INSERT INTO form (id_form, judul_form, deskripsi_form, link_form, start_date, end_date, target, approval_form, entry_user, entry_name, entry_time, approval_posisi, approval_list, status)"
	if param["Is_draft"].(string) == "1" { //submit
		qry_insert_form = qry_insert_form + "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ?, ?, 0)"
	} else {
		qry_insert_form = qry_insert_form + "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ?, ?, 1)"
	}

	qry_insert_form_pertanyaan := "INSERT INTO form_pertanyaan (id_form_pertanyaan, id_form, pertanyaan, pertanyaan_image, jenis_pertanyaan, list_pilihan_jawaban, mandatory) VALUES %s"

	valueStrings	:= []string{}
	valueArgs 		:= []interface{}{}

	DataPertanyaanAll 	:= param["List_pertanyaan"].([]interface{})
	limit		 		:= len(DataPertanyaanAll)

	for i := 0; i < int(limit); i++ {
		DataDetail 		:= DataPertanyaanAll[i].(map[string]interface{})
		valueStrings 	= append(valueStrings, "(?, ?, ?, ?, ?, ?, ?)")
		valueArgs 		= append(valueArgs,DataDetail["id"].(string))
		valueArgs 		= append(valueArgs,DataDetail["id_form"].(string))
		valueArgs 		= append(valueArgs,DataDetail["pertanyaan"].(string))
		valueArgs 		= append(valueArgs,DataDetail["pertanyaan_image"].(string))
		valueArgs 		= append(valueArgs,DataDetail["jenis_pertanyaan"].(string))
		valueArgs 		= append(valueArgs,DataDetail["list_pilihan_jawaban"].(string))
		valueArgs 		= append(valueArgs,DataDetail["mandatory"].(string))
	}
	qry_insert_form_pertanyaan = fmt.Sprintf(qry_insert_form_pertanyaan, strings.Join(valueStrings, ","))

	tx, err := db.Begin()
	if err != nil { 
		return err 
	}

	insert_form, err_qry_insert_form 						:= tx.Prepare(qry_insert_form)
	insert_form_pertanyaan, err_qry_insert_form_pertanyaan 	:= tx.Prepare(qry_insert_form_pertanyaan)

	if err_qry_insert_form == nil && err_qry_insert_form_pertanyaan == nil {
		defer insert_form.Close()
		defer insert_form_pertanyaan.Close()

		_, err_insert_form 				:= insert_form.Exec(param["Id_form"].(string), param["Judul_form"].(string), param["Deskripsi_form"].(string), param["Link_form"].(string), param["Start_date"].(string), param["End_date"].(string), param["Target"].(string), param["Approval_form"].(string), param["Entry_user"].(string), param["Entry_name"].(string), param["Approval_posisi"].(string), param["Approval_list"].(string))
		_, err_insert_form_pertanyaan 	:= insert_form_pertanyaan.Exec(valueArgs...)

		if err_insert_form == nil && err_insert_form_pertanyaan == nil {
			tx.Commit()
			return nil
		} else {
			if err_insert_form != nil {
				err = err_insert_form
			}
	
			if err_insert_form_pertanyaan != nil {
				err = err_insert_form_pertanyaan
			}
	
			tx.Rollback()
			return err
		}
	} else {
		if err_qry_insert_form != nil {
			err = err_qry_insert_form
		}

		if err_qry_insert_form_pertanyaan != nil {
			err = err_qry_insert_form_pertanyaan
		}

		tx.Rollback()
		return err
	}
}
func GetFormAll(param map[string]interface{}) (st.Forms, int, string, error, error) {
	var Datas = st.Forms{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	selected_column := "a.id_form, a.judul_form, a.deskripsi_form, a.link_form, a.start_date, a.end_date, a.target, a.approval_form, a.entry_user, a.entry_name, a.entry_time, a.approval_posisi, a.approval_list, a.status, b.status_desc, a.is_active"
	
	sql_query 		:= "SELECT "+selected_column+" FROM form a INNER JOIN mst_status b ON a.status=b.id_status"
	sql_detail 		:= "SELECT id_form_pertanyaan, id_form, pertanyaan, pertanyaan_image, jenis_pertanyaan, list_pilihan_jawaban, mandatory FROM form_pertanyaan WHERE id_form=?"
	sql_count 		:= "SELECT count(1) FROM form a INNER JOIN mst_status b ON a.status=b.id_status"

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

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}
func UpdateForm(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(id_form) FROM form WHERE (status=1 OR status=2) AND id_form=?", param["Id_form_existing"].(string)).Scan(&countExist)

	if countExist > 0 {
		if param["Is_draft"].(string) == "1" { //submit
			db.QueryRow("SELECT count(id_form) FROM form WHERE link_form=? AND id_form!=?", param["Link_form"].(string), param["Id_form_existing"].(string)).Scan(&countExist)
			if countExist > 0 {
				return fmt.Errorf("Gagal menyimpan form! Link sudah digunakan!")
			}
		}
	
		qry_insert_form := "INSERT INTO form (id_form, judul_form, deskripsi_form, link_form, start_date, end_date, target, approval_form, entry_user, entry_name, entry_time, approval_posisi, approval_list, status)"
		if param["Is_draft"].(string) == "1" { //submit
			qry_insert_form = qry_insert_form + "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ?, ?, 0)"
		} else {
			qry_insert_form = qry_insert_form + "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ?, ?, 1)"
		}
	
		qry_insert_form_pertanyaan 	:= "INSERT INTO form_pertanyaan (id_form_pertanyaan, id_form, pertanyaan, pertanyaan_image, jenis_pertanyaan, list_pilihan_jawaban, mandatory) VALUES %s"
	
		valueStrings	:= []string{}
		valueArgs 		:= []interface{}{}
	
		DataPertanyaanAll 	:= param["List_pertanyaan"].([]interface{})
		limit		 		:= len(DataPertanyaanAll)
	
		for i := 0; i < int(limit); i++ {
			DataDetail 		:= DataPertanyaanAll[i].(map[string]interface{})
			valueStrings 	= append(valueStrings, "(?, ?, ?, ?, ?, ?, ?)")
			valueArgs 		= append(valueArgs,DataDetail["id"].(string))
			valueArgs 		= append(valueArgs,DataDetail["id_form"].(string))
			valueArgs 		= append(valueArgs,DataDetail["pertanyaan"].(string))
			valueArgs 		= append(valueArgs,DataDetail["pertanyaan_image"].(string))
			valueArgs 		= append(valueArgs,DataDetail["jenis_pertanyaan"].(string))
			valueArgs 		= append(valueArgs,DataDetail["list_pilihan_jawaban"].(string))
			valueArgs 		= append(valueArgs,DataDetail["mandatory"].(string))
		}
		qry_insert_form_pertanyaan = fmt.Sprintf(qry_insert_form_pertanyaan, strings.Join(valueStrings, ","))
	
		tx, err := db.Begin()
		if err != nil { 
			return err 
		}
	
		insert_form, err_qry_insert_form 						:= tx.Prepare(qry_insert_form)
		insert_form_pertanyaan, err_qry_insert_form_pertanyaan 	:= tx.Prepare(qry_insert_form_pertanyaan)

		if err_qry_insert_form == nil && err_qry_insert_form_pertanyaan == nil {
			defer insert_form.Close()
			defer insert_form_pertanyaan.Close()

			_, err_insert_form 				:= insert_form.Exec(param["Id_form"].(string), param["Judul_form"].(string), param["Deskripsi_form"].(string), param["Link_form"].(string), param["Start_date"].(string), param["End_date"].(string), param["Target"].(string), param["Approval_form"].(string), param["Entry_user"].(string), param["Entry_name"].(string), param["Approval_posisi"].(string), param["Approval_list"].(string))
			_, err_insert_form_pertanyaan	:= insert_form_pertanyaan.Exec(valueArgs...)

			if err_insert_form == nil && err_insert_form_pertanyaan == nil {

				qry_del_form_exist 				:= "DELETE FROM form WHERE id_form=?"
				qry_del_form_pertanyaan_exist 	:= "DELETE FROM form_pertanyaan WHERE id_form=?"

				del_form_exist, err_qry_del_form_exist 							:= tx.Prepare(qry_del_form_exist)
				del_form_pertanyaan_exist, err_qry_del_form_pertanyaan_exist 	:= tx.Prepare(qry_del_form_pertanyaan_exist)

				if err_qry_del_form_exist == nil && err_qry_del_form_pertanyaan_exist == nil {
					defer del_form_exist.Close()
					defer del_form_pertanyaan_exist.Close()

					_, err_del_form_exist	 			:= del_form_exist.Exec(param["Id_form_existing"].(string))
					_, err_del_form_pertanyaan_exist 	:= del_form_pertanyaan_exist.Exec(param["Id_form_existing"].(string))

					if err_del_form_exist == nil && err_del_form_pertanyaan_exist == nil {
						tx.Commit()
						return nil
					} else {
						if err_del_form_exist != nil {
							err = err_del_form_exist
						}

						if err_del_form_pertanyaan_exist != nil {
							err = err_del_form_pertanyaan_exist
						}

						tx.Rollback()
						return err
					}
				} else {
					if err_qry_del_form_exist != nil {
						err = err_qry_del_form_exist
					}

					if err_qry_del_form_pertanyaan_exist != nil {
						err = err_qry_del_form_pertanyaan_exist
					}

					tx.Rollback()
					return err
				}
			} else {
				if err_insert_form != nil {
					err = err_insert_form
				}
	
				if err_insert_form_pertanyaan != nil {
					err = err_insert_form_pertanyaan
				}
	
				tx.Rollback()
				return err
			}
		} else {
			if err_qry_insert_form != nil {
				err = err_qry_insert_form
			}

			if err_qry_insert_form_pertanyaan != nil {
				err = err_qry_insert_form_pertanyaan
			}

			tx.Rollback()
			return err
		}
	} else {
		return fmt.Errorf("Form tidak ditemukan")
	}
}
func ApprovalForm(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(id_form) FROM form WHERE status=0 AND approval_posisi=? AND id_form=?", param["Approval_now"].(string), param["Id_form"].(string)).Scan(&countExist)

	if countExist > 0 {

		tx, err := db.Begin()
		if err != nil { 
			return err 
		}

		qry_update_form 					:= "UPDATE form SET status=?, approval_posisi=?, approval_list=? WHERE status=0 AND approval_posisi=? AND id_form=?"
		update_form, err_qry_update_form	:= tx.Prepare(qry_update_form)

		if err_qry_update_form == nil {

			defer update_form.Close()
			_, err_update_form := update_form.Exec(param["Status"].(string), param["Approval_next"].(string), param["Approval_list"].(string), param["Approval_now"].(string), param["Id_form"].(string))

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
func ActiveForm(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(id_form) FROM form WHERE id_form=?", param["Key"].(string)).Scan(&countExist)

	if countExist > 0 {
		tx, err := db.Begin()
		if err != nil { 
			return err 
		}

		var isActive int
		db.QueryRow("SELECT is_active FROM form WHERE id_form=?", param["Key"].(string)).Scan(&isActive)

		var qry_update_form string
		if isActive == 0 {
			qry_update_form = "UPDATE form SET is_active=1 WHERE id_form=?"
		} else {
			qry_update_form = "UPDATE form SET is_active=0 WHERE id_form=?"
		}
		
		update_form, err_qry_update_form := tx.Prepare(qry_update_form)

		if err_qry_update_form == nil {

			defer update_form.Close()
			_, err_update_form := update_form.Exec(param["Key"].(string))

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

func InsertJawabanForm(param map[string]interface{}) error {
	db, err := db.Default()
	if err != nil {
		return err
	}
	defer db.Close()

	var countExist int
	db.QueryRow("SELECT count(*) FROM form_jawaban WHERE id_form=? AND entry_user=?", param["Id_form"].(string), param["Entry_user"].(string)).Scan(&countExist)
	if countExist > 0 {
		return fmt.Errorf("Gagal menyimpan form! Anda sudah mengisi form ini sebelumnya!")
	}

	qry_insert_jawaban 			:= "INSERT INTO form_jawaban (id_form, id_jawaban, entry_user, entry_name, entry_time, approval_posisi, approval_list) VALUES (?,?,?,?,NOW(),?,?)"
	qry_insert_jawaban_detail	:= "INSERT INTO form_jawaban_detail (id_form, id_form_pertanyaan, id_jawaban, id_jawaban_detail, jawaban) VALUES %s"

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

	insert_jawaban, err_qry_insert_jawaban 					:= tx.Prepare(qry_insert_jawaban)
	insert_jawaban_detail, err_qry_insert_jawaban_detail 	:= tx.Prepare(qry_insert_jawaban_detail)

	if err_qry_insert_jawaban == nil && err_qry_insert_jawaban_detail == nil {
		defer insert_jawaban.Close()
		defer insert_jawaban_detail.Close()

		_, err_insert_jawaban 			:= insert_jawaban.Exec(param["Id_form"].(string), param["Id_jawaban"].(string), param["Entry_user"].(string), param["Entry_name"].(string), param["Approval_posisi"].(string), param["Approval_list"].(string))
		_, err_insert_jawaban_detail 	:= insert_jawaban_detail.Exec(valueArgs...)

		if err_insert_jawaban == nil && err_insert_jawaban_detail == nil {
			tx.Commit()
			return nil
		} else {
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
		if err_qry_insert_jawaban != nil {
			err = err_qry_insert_jawaban
		}

		if err_qry_insert_jawaban_detail != nil {
			err = err_qry_insert_jawaban_detail
		}
		
		tx.Rollback()
		return err
	}
}