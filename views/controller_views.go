package views

import (
	"fmt"
	"golang/entity/response"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func LoginViews(res http.ResponseWriter, url string) {
	switch (url) {
		case "/":
			var htmlParse, htmlError = template.ParseFiles("views/login/login.html")
			if htmlError != nil {
				break
			}
			res.Header().Set("Content-Type", "text/html")
			res.WriteHeader(200)
			htmlParse.Execute(res, nil)
			return
		case "/login.css":
			var cssParse, cssError = template.ParseFiles("views/login/login.css")
			if cssError != nil {
				break
			}
			res.Header().Set("Content-Type", "text/css")
			res.WriteHeader(200)
			cssParse.Execute(res, nil)
			return
		case "/login.js":
			var jsParse, jsError = template.ParseFiles("views/login/login.js")
			if jsError != nil {
				break
			}
			res.Header().Set("Content-Type", "text/javascript")
			res.WriteHeader(200)
			jsParse.Execute(res, nil)
			return
		default:
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(404)
			fmt.Fprintln(res, `{"message": "not found"}`)
			return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(500)
	fmt.Fprintln(res, `{"message": "internal server error"}`)
}

func AdminViews(res http.ResponseWriter, admin response.Tb_users, url string) {
	switch url {
		case "/":
			var htmlParse, htmlError = template.ParseFiles("views/admin/admin.html")
			if htmlError != nil {
				break
			}
			res.Header().Set("Content-Type", "text/html")
			res.WriteHeader(200)
			htmlParse.Execute(res, &admin)
			return
		case "/admin.js":
			var jsParse, jsError = template.ParseFiles("views/admin/admin.js")
			if jsError != nil {
				break
			}
			res.Header().Set("Content-Type", "text/javascript")
			res.WriteHeader(200)
			jsParse.Execute(res, &admin)
			return
		case "/images/profile.jpg":
			var imgParse, imgError = os.Open("views/admin/images/profile.jpg")
			if imgError != nil {
				log.Println(imgError)
				break
			}
			defer imgParse.Close()
			res.Header().Set("Content-Type", "images/jpg")
			res.WriteHeader(200)
			io.Copy(res, imgParse)
			return
	}

	var htmlParse, htmlError = template.ParseFiles("views/admin/admin.html")
	if htmlError != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(401)
		fmt.Fprintln(res, `{"message": "not found"}`)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(200)
	htmlParse.Execute(res, &admin)
}

func UserViews(res http.ResponseWriter, user response.Tb_users, url string) {
	switch url {
		case "/":
			var htmlParse, htmlError = template.ParseFiles("views/user/user.html")
			if htmlError != nil {
				break
			}
			res.Header().Set("Content-Type", "text/html")
			res.WriteHeader(200)
			htmlParse.Execute(res, &user)
			return
		case "/user.js":
			var jsParse, jsError = template.ParseFiles("views/user/user.js")
			if jsError != nil {
				break
			}
			res.Header().Set("Content-Type", "text/javascript")
			res.WriteHeader(200)
			jsParse.Execute(res, &user)
			return
		case "/images/user.jpg":
			var imgParse, imgError = os.Open("views/user/images/user.jpg")
			if imgError != nil {
				log.Println(imgError)
				break
			}
			defer imgParse.Close()
			res.Header().Set("Content-Type", "images/jpg")
			res.WriteHeader(200)
			io.Copy(res, imgParse)
			return
	}

	var htmlParse, htmlError = template.ParseFiles("views/user/user.html")
	if htmlError != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(401)
		fmt.Fprintln(res, `{"message": "not found"}`)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(200)
	htmlParse.Execute(res, &user)
}