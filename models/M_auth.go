package models

import (
	db "api/uniform/config"

	"fmt"
	"time"

	// _ "assets/mysql"
	// sha256 "crypto/sha256"
	// hex "encoding/hex"
    // st "api/uniform/struct"
)

//CekLogin is
func CekLogin(ip string, appname string) (CODE string, DESC string, signature string) {
	db, err := db.Default()
	if err != nil {
		CODE = "ER-004"
		DESC = "Connection Database Failed"
		signature	= ""
		return CODE, DESC, signature
	}
	defer db.Close()

	var clientId	string
	var clientKey	string
	
	_ = db.QueryRow("select client_id, client_key from user_api where user_ip=? and appname = ?", ip, appname).Scan(&clientId, &clientKey)
	
	if clientId != "" {
		CODE 		= "ES-000"
		DESC 		= "Login success"
		signature	= clientId+clientKey

		return CODE, DESC, signature
	} else {
		CODE 		= "ES-201"
		DESC 		= "Not Authorized"
		signature	= ""

		return CODE, DESC, signature
	}
}

//CekSession is
func CekSession(pn string) (int, error) {
	var Count	int

	db, err := db.Default()
		
	err = db.QueryRow("select count(*) count from user_session where userid=? and DATE_ADD(date_insert, INTERVAL 2 HOUR) > NOW()", pn).Scan(&Count)
	
	return Count, err
}

func Login(User string) {
	db, err := db.Default()
	if err != nil {
		return
	}
	defer db.Close()

	tx, err := db.Begin()

	Delete, err := tx.Prepare("delete from user_session where userid=?")
	insert, err := tx.Prepare("insert into user_session (userid, date_insert) values (?, NOW())")
	
	if err != nil {
		return
	} else {
		defer Delete.Close()
		defer insert.Close()

		_, errDel := Delete.Exec(User)
		_, errIns := insert.Exec(User)

		if errDel != nil || errIns != nil {
			tx.Rollback()
			return
		} else {
			tx.Commit()
			return
		}
		return
	}	
}

func Logout(User string) {
	db, err := db.Default()
	if err != nil {
		return
	}
	defer db.Close()

	Delete, err := db.Prepare("delete from user_session where userid=?")

	if err != nil {
		return
	} else {
		defer Delete.Close()
		Delete.Exec(User)
		return
	}	
}

func GetUrl(vendor string,  jenis string, sender string) (string, string, string, string, error, error) {
	var userid		string
	var userpass	string
	var userurl		string

	db, err := db.Default()
	if err != nil {
		return userid, userpass, userurl, "ED-004", fmt.Errorf("Connection Database Failed"), err
	}
	defer db.Close()

	if vendor != "" {
		err = db.QueryRow("SELECT userid, userpass, userurl FROM mst_param WHERE vendor=?", vendor).Scan(&userid, &userpass, &userurl)
		if err != nil {
			return userid, userpass, userurl, "ER-099", fmt.Errorf("User tidak ditemukan"), err
		}
	} else {
		return userid, userpass, userurl, "ER-001", fmt.Errorf("Invalid Data"), err
	}
	
	return userid, userpass, userurl, "", err, err
}

func GetMsg(vendor string) (string, error) {
	var msg_vendor		string

	db, err := db.Default()
	if err != nil {
		return msg_vendor, err
	}
	defer db.Close()

	db.QueryRow("select msg_service from mst_msg where vendor=?", vendor).Scan(&msg_vendor)
	
	return msg_vendor, err
}

func AuditTrail(Ip string, userid string, Function string, Body string, Resp string, errSis string) {
	db, err := db.Default()
	if err != nil {
		return
	}
	defer db.Close()
	
	currentTime 	:= time.Now()
	yearMonthNow 	:= currentTime.Format("200601")
	db.QueryRow("CREATE TABLE IF NOT EXISTS log_request_"+ yearMonthNow +" LIKE log_request")

	insert, err := db.Prepare("insert into log_request_"+ yearMonthNow +" (ipaddress, userid, `function`, body, response, errordesc, time_request) values (?, ?, ?, ?, ?, ?, NOW())")

	if err != nil {
		return
	} else {
		defer insert.Close()
		insert.Exec(Ip, userid, Function, Body, Resp, errSis)
		return
	}
}