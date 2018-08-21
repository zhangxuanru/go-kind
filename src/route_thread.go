package main

import (
	"net/http"
	"data"
	"fmt"
)

func newThread(w http.ResponseWriter,r *http.Request)  {
	_, e := session(w, r)
	if e != nil{
        http.Redirect(w,r,"/login",302)
	}else{
	    generateHTML(w,nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(w http.ResponseWriter, r *http.Request)  {
	sess, e := session(w, r)
	if e != nil {
		http.Redirect(w,r,"/login",302)
	}else{
		err := r.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := r.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(w, r, "/", 302)
	}
}

func postThread(w http.ResponseWriter, r *http.Request)  {
	sess, e := session(w, r)
	if e != nil {
		http.Redirect(w,r,"/login",302)
	}else{
		err := r.ParseForm()
		if err != nil{
			danger(err, "Cannot parse form")
		}
		user, err2 := sess.User()
		if err2 != nil {
			danger(err, "Cannot get user from session")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(w, r, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(w, r, url, 302)
	}
}

func readThread(w http.ResponseWriter, r *http.Request)  {
	uuid := r.URL.Query().Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		error_message(w, r, "Cannot read thread")
	}else{
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}

}

