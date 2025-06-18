package structure

type DataResponse interface{}

type Response struct {
	Response interface{}
}

type ErrorDesc struct {
	ErrorDesc string
}

type ResponseNil struct {
	ErrorCode    string
	ResponseCode string
	ResponseDesc string
}

type Token struct {
	ErrorCode    string
	ResponseCode string
	ResponseDesc string
	Token        string
	ResponseData interface{}
}

type Data struct {
	ErrorCode    string
	ResponseCode string
	ResponseDesc string
	ResponseData interface{}
}

type DataAll struct {
	ErrorCode     string
	ResponseCode  string
	ResponseDesc  string
	ResponseCount int
	ResponseData  interface{}
}

type DataStatus struct {
	IdStatus	int
	StatusDesc	string
}

type DataStatuss []DataStatus

type User struct {
	PERNR			string
	SNAME			string
	JGPG			string
	ESELON			string
	WERKS			string
	WERKS_TX		string
	BTRTL			string
	BTRTL_TX		string
	KOSTL			string
	KOSTL_TX		string
	ORGEH			string
	ORGEH_TX		string
	STELL			string
	STELL_TX		string
	PLANS			string
	PLANS_TX		string
	HILFM			string
	HTEXT			string
	BRANCH			string
	MAINBR			string
	IS_PEMIMPIN		string
	ADMIN_LEVEL		string
	ORGEH_PGS		string
	ORGEH_PGS_TX	string
	PLANS_PGS		string
	PLANS_PGS_TX	string
	BRANCH_PGS		string
	HILFM_PGS		string
	HTEXT_PGS		string
	TIPE_UKER		string
	REKENING		string
	NPWP			string
	REGION			string
	RGDESC			string
	BRDESC			string
	MBDESC			string
}

type Users []User

type Menu struct {
	Menu_id 		string
	Menu_parent 	string
	Menu_icon 		string
	Menu_title 		string
	Menu_link 		string
	Level_user 		string
}

type Menus []Menu

type ResponseMenu struct {
	ErrorCode     	string
	ResponseCode	string
	ResponseDesc	string
	ResponseData	Menus
}

type DataDivisi struct {
	KOSTL		string
	KOSTL_TX	string
}

type DataDivisis []DataDivisi


type Notif struct {
	IdNotif				string
	KeyNotif			string
	DescNotif			string
	UserPengirim		string
	UserNamaPengirim	string
	UserPenerima		string
	UserNamaPenerima	string
	EntryDate			string
	JenisNotif			string
	Status				string
	StatusApproval  	string
	DescStatus			string
	JenisDokumen		string
	Keterangan			string
	IdJenisSurat		string
	JenisSurat			string
}

type Notifs []Notif

type DataParam struct {
	Vendor		string
	Userid		string
	Userpass	string
	Userurl		string
}
type DataParams [] DataParam

type Form struct {
	Id_form			string
	Judul_form		string
	Deskripsi_form	string
	Link_form		string
	Start_date		string
	End_date		string
	Target			string
	Approval_form	string
	Entry_user		string
	Entry_name		string
	Entry_time		string
	Approval_posisi	string
	Approval_list	string
	Status			string
	Status_desc		string
	Is_active		string
	Total_response	string
	
	Form_pertanyaan Form_pertanyaans
	
}
type Forms [] Form

type Form_pertanyaan struct {
	Id_form_pertanyaan 		string
	Id_form 				string
	Pertanyaan 				string
	Pertanyaan_image 		string
	Jenis_pertanyaan 		string
	List_pilihan_jawaban	string
	Mandatory 				string
}
type Form_pertanyaans [] Form_pertanyaan

type JawabanSummary struct {
	Jawaban					string
	JawabanTotal			int
	JawabanMatch			int
	JawabanMatchPercentage	string
}
type JawabanSummarys [] JawabanSummary

type JawabanById struct {
	Entry_user			string
	Entry_name			string
	Entry_time			string
	Id_form				string
	Id_form_pertanyaan	string
	Id_jawaban			string
	Id_jawaban_detail	string
	Jawaban				string
}
type JawabanByIds [] JawabanById

type Jawaban struct {
	Id_form			string
	Judul_form 		string
	Deskripsi_form	string
	Link_form 		string
	Id_jawaban		string
	Entry_user		string
	Entry_name		string
	Entry_time		string
	Approval_posisi	string
	Approval_list	string
	Status			string
	Status_desc		string

	Jawaban_detail	Jawaban_details
}
type Jawabans [] Jawaban

type Jawaban_detail struct {
	Id_form					string
	Id_form_pertanyaan		string
	Pertanyaan 				string
	Pertanyaan_image 		string
	Jenis_pertanyaan 		string
	List_pilihan_jawaban	string
	Mandatory 				string
	Id_jawaban				string
	Id_jawaban_detail		string
	Jawaban					string
}
type Jawaban_details [] Jawaban_detail

type DataSummary struct {
	Form_total							int
	Form_feedback_average_percentage	string
	ApprovalForm_total					int
	ApprovalResponse_total				int
}
type DataSummarys [] DataSummary

type DataRequestAttachment struct {
	Id				string
	Request_user	string
	Request_name	string
	Request_time	string
	Id_form			string
	Judul_form		string
	Url_attachment	string
	Status			string
	StatusDesc		string
}
type DataRequestAttachments [] DataRequestAttachment

type DataJobRequestAttachment struct {
	Id				string
	Request_user	string
	Request_name	string
	Request_time	string
	Id_form			string
	Judul_form		string
	Url_attachment	string
	
	Request_detail	DataJobRequestAttachmentDetails
}
type DataJobRequestAttachments [] DataJobRequestAttachment

type DataJobRequestAttachmentDetail struct {
	Entry_name	string
	Jawaban		string
}
type DataJobRequestAttachmentDetails [] DataJobRequestAttachmentDetail
