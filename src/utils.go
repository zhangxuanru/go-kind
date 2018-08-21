package main

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
	"net/http"
	"html/template"
	"strings"
	"data"
	"errors"
)

type Configuration struct {
	Address  string
	ReadTimeout  int64
	WriteTimeout int64
	Static string
}

var config Configuration
var logger *log.Logger

func init()  {
     loadConfig()
}

func loadConfig()  {
	file, e := os.Open("config.json")
	if e != nil{
		 log.Fatalln("Cannot open config file",e)
	}
	config = Configuration{}
	err := json.NewDecoder(file).Decode(&config)
    if err != nil{
    	log.Fatalln("Cannot get configuration from file",err)
	}
}


func generateHTML(w http.ResponseWriter,data interface{},files ...string)  {
    var fileList []string
    for _,file := range files{
    	 fileList = append(fileList,fmt.Sprintf("templates/%s.html",file))
	}
	must := template.Must(template.ParseFiles(fileList...))
	must.ExecuteTemplate(w,"layout",data)
}

func error_message(w http.ResponseWriter,r *http.Request,msg string)  {
     url := []string{"/err/?msg=",msg}
	 http.Redirect(w,r,strings.Join(url,""),302)
}

func session(w http.ResponseWriter,r *http.Request)(sess data.Session,err error)  {
	cookie, err := r.Cookie("_cookie")
	if err == nil{
		sess = data.Session{Uuid: cookie.Value}
		if ok,_ := sess.Check(); !ok{
             err = errors.New("Invalid session")
		}
	}
	return
}

func parseTemplateFiles(filenames ...string)(t *template.Template)  {
    var files []string
	t = template.New("layout")
	for _,file := range filenames{
		files = append(files,fmt.Sprintf("templates/%s.html",file))
	}
	t = template.Must(t.ParseFiles(files...))
    return
}




func p(a ...interface{})  {
	fmt.Println(a)
}

func warning(args ...interface{})  {
	 logger.SetPrefix("WARNING")
	 log.Println(args...)
}

func danger(args ...interface{})  {
	logger.SetPrefix("ERROR")
	logger.Println(args...)
}


func version() string {
	return "1.0"
}


