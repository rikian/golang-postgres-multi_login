package request

type Tb_users struct {
	User_id       string `gorm:"column:user_id;primayKey"`
	User_name     string
	User_email    string
	User_password string
	User_status   string
	Create_date   string
	User_session  string
}