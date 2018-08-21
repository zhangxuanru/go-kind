package main

import (
	"net/http"
	"data"
)

func login(w http.ResponseWriter, r *http.Request)  {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w,nil)
}

func logout(w http.ResponseWriter, r *http.Request)  {
	cookie, e := r.Cookie("_cookie")
	if e != http.ErrNoCookie{
		warning(e, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
	    session.DeleteByUUID()
	}
	http.Redirect(w,r,"/",302)
}

func signup(w http.ResponseWriter, r *http.Request)  {
	generateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

func signupAccount(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil{
	 	danger(err, "Cannot parse form")
	}
	user := data.User{
		Name : r.PostFormValue("name"),
		Email: r.PostFormValue("email"),
		Password:r.PostFormValue("password"),
	}
	if e := user.Create(); e != nil{
		danger(err, "Cannot create user")
	}
	http.Redirect(w,r,"/login",302)
}

func authenticate(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil{
		danger(err, "Cannot find user")
	}
	user, e := data.UserByEmail(r.PostFormValue("email"))
	if e != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")){
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:"_cookie",
			Value:session.Uuid,
			HttpOnly:true,
		}
       http.SetCookie(w,&cookie)
       http.Redirect(w,r,"/",302)
	}else{
		http.Redirect(w,r,"/login",302)
	}
}


