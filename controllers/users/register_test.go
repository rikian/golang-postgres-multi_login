package users

import (
	"golang/config"
	"golang/entity/request"
	"log"
	"testing"

	"github.com/google/uuid"
)

func TestRegister(t *testing.T) {
	config.ConnectDB()
	user := request.Tb_users {
		User_id: uuid.NewString(),
		User_name: "Rikian Faisal",
		User_email: "rikianfaisal@gmail.com",
		User_password: "sha256(54n94t_r4h4514...)",
		User_status: "user",
		Create_date: "17-a",
	}
	
	register := Register(config.DB, &user)

	log.Println(register)
}