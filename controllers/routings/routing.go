package routings

import (
	"encoding/json"
	"fmt"
	"golang/config"
	"golang/controllers/admin"
	"golang/entity/request"
	"golang/utils"
	"golang/views"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type BodyLogin struct {
	Email string
	Password string
}

type BodyRegister struct {
	User_name     string
	User_email    string
	User_password string
	User_phone	  string
	User_status   string
}

func SlashRouting(res http.ResponseWriter, req *http.Request) {
	_id, _idError := req.Cookie("id")
	_tkn, _tknError := req.Cookie("_tkn")
	url := req.URL.String()
	method := req.Method
	if _tknError == nil && _idError == nil && _id.Value != "" && _tkn.Value != "" {
		isUserOrAdmin, isUserOrAdminError := admin.GetUserAfterLogin(config.DB, _id.Value, _tkn.Value)
		if !isUserOrAdminError {
			// user login as admin
			if isUserOrAdmin.User_status == "admin" {
				if method == "GET" {
					views.AdminViews(res, isUserOrAdmin, url)
					return
				}

				if method == "POST" {
					if url == "/logout" {
						logout := admin.Logout(config.DB, isUserOrAdmin.User_id, isUserOrAdmin.User_email)
						if !logout {
							res.Header().Set("Content-Type", "application/json")
							res.WriteHeader(500)
							fmt.Fprintln(res, `{"message": "failed"}`)
							return
						}
						http.SetCookie(res, &http.Cookie {
							Name: "id",
							Value: "",
							Path: "/",
							HttpOnly: true,
						})
						http.SetCookie(res, &http.Cookie {
							Name: "_tkn",
							Value: "",
							Path: "/",
							HttpOnly: true,
						})
						res.Header().Set("Content-Type", "application/json")
						res.WriteHeader(200)
						fmt.Fprintln(res, `{"message": "success"}`)
						return
					}
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(401)
				fmt.Fprintln(res, `{"message": "method not allowed"}`)
				return
			}
			// user login as user
			if isUserOrAdmin.User_status == "user" {
				if method == "GET" {
					views.UserViews(res, isUserOrAdmin, url)
					return
				}
				if method == "POST" {
					if url == "/logout" {
						logout := admin.Logout(config.DB, isUserOrAdmin.User_id, isUserOrAdmin.User_email)
						if !logout {
							res.Header().Set("Content-Type", "application/json")
							res.WriteHeader(500)
							fmt.Fprintln(res, `{"message": "failed"}`)
							return
						}
						http.SetCookie(res, &http.Cookie {
							Name: "id",
							Value: "",
							Path: "/",
							HttpOnly: true,
						})
						http.SetCookie(res, &http.Cookie {
							Name: "_tkn",
							Value: "",
							Path: "/",
							HttpOnly: true,
						})
						res.Header().Set("Content-Type", "application/json")
						res.WriteHeader(200)
						fmt.Fprintln(res, `{"message": "success"}`)
						return
					}
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(401)
				fmt.Fprintln(res, `{"message": "method not allowed"}`)
				return
			}
		}
	}
	
	// user no login
	if method == "GET" {
		views.LoginViews(res, url)
		return
	}

	headers := req.Header
	if method == "POST" {
		switch (url) {
			case "/login":
				ct := headers.Get("Content-Type")
				cl, cl_err := strconv.Atoi(headers.Get("Content-Length"))
				if ct != "application/json" || cl_err != nil || cl > 128 {
					log.Println("invalid headers")
					break
				}
				var bodyLogin BodyLogin
				bodyLoginError := json.NewDecoder(req.Body).Decode(&bodyLogin)
				if bodyLoginError != nil || bodyLogin.Email == "" || bodyLogin.Password == "" {
					log.Println("invalid body")
					break
				}
				isEmail := utils.IsValidEmail(bodyLogin.Email)
				if !isEmail {
					log.Println("invalid email")
					break
				}
				dataUserLogin := request.Tb_users {
					User_email: bodyLogin.Email,
					User_password: utils.SHA256(bodyLogin.Password),
				}
				login, loginError := admin.Login(config.DB, &dataUserLogin)
				if loginError {
					break
				}
				// generate token login using jwt
				gtl, gtlError := utils.EncryptToken(120)
				if gtlError {
					break
				}
				// save part jwt to session
				saveTokenLogin := admin.SaveTokenLogin(config.DB, gtl, bodyLogin.Email, utils.SHA256(bodyLogin.Password))
				if !saveTokenLogin {
					break
				}
				http.SetCookie(res, &http.Cookie {
					Name: "id",
					Value: login.User_id,
					Path: "/",
					HttpOnly: true,
				})
				http.SetCookie(res, &http.Cookie {
					Name: "_tkn",
					Value: strings.Split(gtl, ".")[2],
					Path: "/",
					HttpOnly: true,
				})
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintln(res, `{"message": "success"}`)
				return
			case "/register":
				ct := headers.Get("Content-Type")
				cl, cl_err := strconv.Atoi(headers.Get("Content-Length"))
				if ct != "application/json" || cl_err != nil || cl > 256 {
					log.Println("invalid headers")
					break
				}
				var bodyRegister BodyRegister
				bodyLoginError := json.NewDecoder(req.Body).Decode(&bodyRegister)
				if bodyLoginError != nil || bodyRegister.User_name == "" || bodyRegister.User_email == "" || bodyRegister.User_phone == "" ||  bodyRegister.User_password == "" || bodyRegister.User_status != "admin" {
					if bodyRegister.User_status != "user" {
						log.Println("invalid body")
						break
					}
				}
				isEmail := utils.IsValidEmail(bodyRegister.User_email)
				if !isEmail {
					log.Println("invalid email")
					break
				}
				dataRegisterUser := request.Tb_users {
					User_id: uuid.NewString(),
					User_name: bodyRegister.User_name,
					User_email: bodyRegister.User_email,
					User_password: utils.SHA256(bodyRegister.User_password),
					User_status: bodyRegister.User_status,
					Create_date: time.Now().Format("2000-01-02 12:34:56"),
				}
				register := admin.Register(config.DB, &dataRegisterUser)
				if !register {
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(400)
					fmt.Fprintln(res, `{"message": "email already exist"}`)
					return	
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintln(res, `{"message": "registrasi success/nPlease login to continue..."}`)
				return
			default:
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(404)
				fmt.Fprintln(res, `{"message": "not found"}`)
				return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(400)
		fmt.Fprintln(res, `{"message": "bad request"}`)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(401)
	fmt.Fprintln(res, `{"message": "method not allowed"}`)
}