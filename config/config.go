package config

const ENV = "development"
// const ENV = "production"

func UrlDio() (string) {
	var UrlDio = "http://10.35.65.111/eoffice/index.php/restws_surat_server"

	if ENV == "development" {
		UrlDio = "http://10.35.65.111/eoffice/index.php/restws_surat_server"
	} else if ENV == "production" {
		UrlDio = ""
	}

	return UrlDio
}

func UrlBristars() (string) {
	var Url = "http://10.35.65.113/bristars/appcheck/iface_app?wsdl"

	if ENV == "development" {
		Url = "http://10.35.65.113/bristars/appcheck/iface_app?wsdl"
	} else if ENV == "production" {
		Url = ""
	}

	return Url
}

func UrlDigest() (string) {
	var Url = "http://10.35.65.88/bristars_api/api/digest/pekerja/login/readable"

	if ENV == "development" {
		Url = "http://10.35.65.88/bristars_api/api/digest/pekerja/login/readable"
	} else if ENV == "production" {
		Url = ""
	}

	return Url
}