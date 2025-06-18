package models

import (
	"fmt"
	// sha256 "crypto/sha256"
	// hex "encoding/hex"
	str "strconv"
    db "api/uniform/config"
    st "api/uniform/struct"
	_ "assets/mysql"
)

func GetMenu() (st.Menus, error) {
	var Datas = st.Menus{}
	db, err := db.Default()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sql_query := "select menu_id,menu_parent,menu_icon,menu_title,IFNULL(menu_link, '') as menu_link,level_user from menu where menu_active='1' order by menu_id ASC"
	
	//query data all
	rows, err := db.Query(sql_query)

	for rows.Next() {
		var each st.Menu
		var err = rows.Scan(&each.Menu_id, &each.Menu_parent, &each.Menu_icon, &each.Menu_title, &each.Menu_link, &each.Level_user)
		
        if err != nil {
			return nil, err
        }

        Datas = append(Datas, each)
    }
	return Datas, err
}

func GetUserAll(param map[string]interface{}) (st.Users, int, string, error, error) {
	var Datas = st.Users{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	// selected_column := "PERNR, SNAME, IFNULL(JGPG, '') JGPG, IFNULL(ESELON, '') ESELON, WERKS , WERKS_TX , BTRTL , BTRTL_TX , KOSTL , KOSTL_TX , ORGEH , ORGEH_TX , STELL , STELL_TX , PLANS , PLANS_TX , HILFM, HTEXT, BRANCH , MAINBR , IS_PEMIMPIN , ADMIN_LEVEL , ORGEH_PGS , ORGEH_PGS_TX , PLANS_PGS , PLANS_PGS_TX , BRANCH_PGS , HILFM_PGS , HTEXT_PGS , TIPE_UKER , REKENING , NPWP, REGION , RGDESC , BRDESC , MBDESC"

	selected_column := "PERNR, SNAME, IFNULL(JGPG, '') JGPG, IFNULL(ESELON, '') ESELON, WERKS, WERKS_TX, BTRTL, BTRTL_TX, KOSTL, KOSTL_TX, ORGEH, ORGEH_TX, IFNULL(STELL, '') STELL, IFNULL(STELL_TX, '') STELL_TX, IFNULL(PLANS, '') PLANS, IFNULL(PLANS_TX, '') PLANS_TX, IFNULL(HILFM, '') HILFM, HTEXT, BRANCH, MAINBR, IS_PEMIMPIN, IFNULL(ADMIN_LEVEL, '') ADMIN_LEVEL , IFNULL(ORGEH_PGS, '') ORGEH_PGS, IFNULL(ORGEH_PGS_TX, '') ORGEH_PGS_TX, IFNULL(PLANS_PGS, '') PLANS_PGS, IFNULL(PLANS_PGS_TX, '') PLANS_PGS_TX, IFNULL(BRANCH_PGS, '') BRANCH_PGS, IFNULL(HILFM_PGS, '') HILFM_PGS, IFNULL(HTEXT_PGS, '') HTEXT_PGS, TIPE_UKER, REKENING, IFNULL(NPWP, '') NPWP, REGION, RGDESC, BRDESC, MBDESC"

	sql_count := "SELECT count(1) FROM user"
	sql_query := "SELECT "+selected_column+" FROM user"

	// sql_count = sql_count + param["Where"].(string)
	// sql_query = sql_query + param["Where"].(string)

	if param["Where"].(string) != "" {
		sql_query = sql_query + " where "+ param["Where"].(string)
		sql_count = sql_count + " where "+ param["Where"].(string)
	}
	
	if param["Order"].(string) != "" {
		sql_query = sql_query +" order by "+ param["Order"].(string)
	} else {
		sql_query = sql_query +" order by SNAME ASC"
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
		var each st.User
		var err = rows.Scan(&each.PERNR, &each.SNAME, &each.JGPG, &each.ESELON, &each.WERKS , &each.WERKS_TX , &each.BTRTL , &each.BTRTL_TX , &each.KOSTL , &each.KOSTL_TX , &each.ORGEH , &each.ORGEH_TX , &each.STELL , &each.STELL_TX , &each.PLANS , &each.PLANS_TX , &each.HILFM, &each.HTEXT, &each.BRANCH , &each.MAINBR , &each.IS_PEMIMPIN , &each.ADMIN_LEVEL , &each.ORGEH_PGS , &each.ORGEH_PGS_TX , &each.PLANS_PGS , &each.PLANS_PGS_TX , &each.BRANCH_PGS , &each.HILFM_PGS , &each.HTEXT_PGS , &each.TIPE_UKER , &each.REKENING , &each.NPWP, &each.REGION , &each.RGDESC , &each.BRDESC , &each.MBDESC)
		
		if err != nil {
			return Datas, 0, "", err, err
		}

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}

func GetUserById(param map[string]interface{}) (st.User, string, error, error) {
	var each = st.User{}
	db, err := db.Default()
	if err != nil {
		return each, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	selected_column := "PERNR, SNAME, IFNULL(JGPG, '') JGPG, IFNULL(ESELON, '') ESELON, WERKS, WERKS_TX, BTRTL, BTRTL_TX, KOSTL, KOSTL_TX, ORGEH, ORGEH_TX, IFNULL(STELL, '') STELL, IFNULL(STELL_TX, '') STELL_TX, IFNULL(PLANS, '') PLANS, IFNULL(PLANS_TX, '') PLANS_TX, IFNULL(HILFM, '') HILFM, HTEXT, BRANCH, MAINBR, IS_PEMIMPIN, IFNULL(ADMIN_LEVEL, '') ADMIN_LEVEL , IFNULL(ORGEH_PGS, '') ORGEH_PGS, IFNULL(ORGEH_PGS_TX, '') ORGEH_PGS_TX, IFNULL(PLANS_PGS, '') PLANS_PGS, IFNULL(PLANS_PGS_TX, '') PLANS_PGS_TX, IFNULL(BRANCH_PGS, '') BRANCH_PGS, IFNULL(HILFM_PGS, '') HILFM_PGS, IFNULL(HTEXT_PGS, '') HTEXT_PGS, TIPE_UKER, REKENING, IFNULL(NPWP, '') NPWP, REGION, RGDESC, BRDESC, MBDESC"

	sql_query := "SELECT "+selected_column+" FROM user Where PERNR=?"

	//query count data
	err = db.QueryRow(sql_query, param["USER"].(string)).Scan(&each.PERNR, &each.SNAME, &each.JGPG, &each.ESELON, &each.WERKS , &each.WERKS_TX , &each.BTRTL , &each.BTRTL_TX , &each.KOSTL , &each.KOSTL_TX , &each.ORGEH , &each.ORGEH_TX , &each.STELL , &each.STELL_TX , &each.PLANS , &each.PLANS_TX , &each.HILFM, &each.HTEXT, &each.BRANCH , &each.MAINBR , &each.IS_PEMIMPIN , &each.ADMIN_LEVEL , &each.ORGEH_PGS , &each.ORGEH_PGS_TX , &each.PLANS_PGS , &each.PLANS_PGS_TX , &each.BRANCH_PGS , &each.HILFM_PGS , &each.HTEXT_PGS , &each.TIPE_UKER , &each.REKENING , &each.NPWP, &each.REGION , &each.RGDESC , &each.BRDESC , &each.MBDESC)
	
	if err != nil {
		return each, "ER-099", fmt.Errorf("Username atau password salah"), err
	}
	
	return each, "", err, err
}

func GetDivisiAll(param map[string]interface{}) (st.DataDivisis, int, string, error, error) {
	var Datas = st.DataDivisis{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	sql_count := "SELECT count(distinct KOSTL, KOSTL_TX) FROM `user` where WERKS='KP00' and HILFM != '098' and KOSTL not in ('PS98000','PS98200')"
	sql_query := "select distinct KOSTL, KOSTL_TX from `user` where WERKS='KP00' and HILFM != '098' and KOSTL not in ('PS98000','PS98200')"

	if param["Where"].(string) != "" {
		sql_query = sql_query + " and "+ param["Where"].(string)
		sql_count = sql_count + " and "+ param["Where"].(string)
	}
	
	if param["Order"].(string) != "" {
		sql_query = sql_query +" order by "+ param["Order"].(string)
	} else {
		sql_query = sql_query +" order by HILFM ASC"
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
		var each st.DataDivisi
		var err = rows.Scan(&each.KOSTL, &each.KOSTL_TX)
		
		if err != nil {
			return Datas, 0, "", err, err
		}

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}

func GetStatus(param map[string]interface{}) (st.DataStatuss, int, string, error, error) {
	var Datas = st.DataStatuss{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	sql_count := "SELECT count(1) FROM mst_status"
	sql_query := "select id_status, nama_status from mst_status"

	//query count data
	db.QueryRow(sql_count).Scan(&count)
	rows, err := db.Query(sql_query)

	if err != nil {
		return Datas, 0, "ER-099", fmt.Errorf("General Error"), err
	}
	
	defer rows.Close()
	for rows.Next() {
		var each st.DataStatus
		var err = rows.Scan(&each.IdStatus, &each.StatusDesc)
		
		if err != nil {
			return Datas, 0, "ED-004", err, err
		}

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}

func GetParamAll(param map[string]interface{}) (st.DataParams, int, string, error, error) {
	var Datas = st.DataParams{}
	db, err := db.Default()
	if err != nil {
		return Datas, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var count int

	selected_column := "vendor, userid, userpass, userurl"

	sql_count := "SELECT count(1) FROM mst_param"
	sql_query := "SELECT "+selected_column+" FROM mst_param"

	if param["Where"].(string) != "" {
		sql_query = sql_query + " where "+ param["Where"].(string)
		sql_count = sql_count + " where "+ param["Where"].(string)
	}
	
	if param["Order"].(string) != "" {
		sql_query = sql_query +" order by "+ param["Order"].(string)
	} else {
		sql_query = sql_query +" order by vendor ASC"
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
		var each st.DataParam
		var err = rows.Scan(&each.Vendor, &each.Userid, &each.Userpass, &each.Userurl)
		
		if err != nil {
			return Datas, 0, "", err, err
		}

		Datas = append(Datas, each)
	}
	
	return Datas, count, "", err, err
}

func GetSummary(param map[string]interface{}) (st.DataSummary, int, string, error, error) {
	var Data = st.DataSummary{}
	db, err := db.Default()
	if err != nil {
		return Data, 0, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	var countFormTotal int
	db.QueryRow("SELECT count(1) FROM form WHERE status='3' AND entry_user=?",param["User"].(string)).Scan(&countFormTotal)

	rows, err := db.Query("SELECT id_form, target FROM form WHERE entry_user=? AND status=3",param["User"].(string))
	if err != nil {
		return Data, 0, "ER-099", fmt.Errorf("General Error"), err
	}
	defer rows.Close()

	var sumPercentage float64

	for rows.Next() {
		var idForm	string
		var target	int
		var err = rows.Scan(&idForm, &target)
		if err != nil {
			return Data, 0, "ER-099", fmt.Errorf("General Error"), err
		}

		var totalFeedback int
		db.QueryRow("SELECT count(1) FROM form_jawaban WHERE id_form=?",idForm).Scan(&totalFeedback)

		sumPercentage = sumPercentage + ((float64(totalFeedback) / float64(target)) * 100)
	}

	var countApprovalForm int
	db.QueryRow("SELECT count(1) FROM form WHERE status = 0 AND approval_posisi=?",param["User"].(string)).Scan(&countApprovalForm)

	var countApprovalResponse int
	db.QueryRow("SELECT count(1) FROM form_jawaban WHERE status = 0 AND approval_posisi=?",param["User"].(string)).Scan(&countApprovalResponse)

	finalPercentage := "0.00"
	if sumPercentage > 0 {
		finalPercentage = fmt.Sprintf("%.2f", (sumPercentage / float64(countFormTotal)))
	}

	Data.Form_total 						= countFormTotal
	Data.Form_feedback_average_percentage 	= finalPercentage
	Data.ApprovalForm_total 				= countApprovalForm
	Data.ApprovalResponse_total 			= countApprovalResponse

	return Data, 0, "", err, err
}