package data

import (
	"database/sql"
	"os"
	"log"
	"encoding/json"
	_ "github.com/lib/pq"
	"fmt"
	"crypto/rand"
	"crypto/sha1"
)
type DriverConfig struct {
	DriverName  string
	Host        string
	Port        string
	User        string
	Password    string
    Dbname      string
    Sslmode     string
}

var config DriverConfig

var Db *sql.DB

func init()  {
	var err error
	loadDriverConfig()
	params := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",config.Host,config.Port,config.User,config.Password,config.Dbname,config.Sslmode)
	Db, err = sql.Open("postgres", params)
	if err != nil{
		log.Fatal("db open error:",err)
	}
    return
}


func loadDriverConfig()  {
	file, e := os.Open("mysql.config.json")
	if e != nil{
		log.Fatalln("sql driver error:",e)
	}
	config = DriverConfig{}
	decode := json.NewDecoder(file).Decode(&config)
	if decode != nil{
		log.Fatalln("driver json error:",decode)
	}
}

func createUUID() (uuid string)  {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
