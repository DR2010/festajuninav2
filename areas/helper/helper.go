package helper

import (
	"encoding/json"
	"festajuninav2/areas/commonstruct"
	"fmt"
	"io/ioutil"

	"github.com/go-redis/redis"

	// "festajuninav2/areas/security"
)

var redisclient *redis.Client
var SYSID string
var databaseEV commonstruct.DatabaseX
var envirvar commonstruct.RestEnvVariables

// Resultado is a struct
type Resultado struct {
	ErrorCode        string // error code
	ErrorDescription string // description
	IsSuccessful     string // Y or N
	ReturnedValue    string
}


// Readfileintostruct is
func Readfileintostruct() commonstruct.RestEnvVariables {
	dat, err := ioutil.ReadFile("festajuninav2.ini")
	check(err)
	fmt.Print(string(dat))

	var restenv commonstruct.RestEnvVariables

	json.Unmarshal(dat, &restenv)

	return restenv
}

// RestEnvVariables = restaurante environment variables
//
type RestEnvVariables struct {
	APIMongoDBLocation    string // location of the database localhost, something.com, etc
	APIMongoDBDatabase    string // database name
	APIAPIServerPort      string // collection name
	APIAPIServerIPAddress string // apiserver name
	WEBDebug              string // debug
	RecordCurrencyTick    string // debug
	RunningFromServer     string // debug
	WEBServerPort         string // collection name
	ConfigFileFound       string // collection name
	ApplicationID         string // collection name
	UserID                string // collection name
	AppFestaJuninaEnabled string
	AppBelnorthEnabled    string
	AppBitcoinEnabled     string
}

// Claim is
type Claim struct {
	Type  string
	Value string
}

// Credentials is a struct
// ----------------------------------------------------
type Credentials struct {
	UserID        string // error code
	UserName      string // description
	KeyJWT        string
	JWT           string
	Expiry        string
	Roles         []string         // Y or N
	ClaimSet      []Claim // Y or N
	ApplicationID string           //
	IsAdmin       string           //
	CentroID      string           //
}


// GetSYSID is just returning the System ID directly from file
// It is happening to enable multiple usage of Redis Keys ("SYSID" + "APIURL" for instance)
func GetSYSID() string {

	if SYSID == "" {

		dat, err := ioutil.ReadFile("festajuninav2.ini")
		check(err)
		fmt.Print(string(dat))

		var restenv commonstruct.RestEnvVariables

		json.Unmarshal(dat, &restenv)

		SYSID = restenv.SYSID

		return restenv.SYSID
	}

	return SYSID

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetRedisPointer returns
func GetRedisPointer(bucket int) *redis.Client {

	bucket = 0

	if redisclient == nil {
		redisclient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",     // no password set
			DB:       bucket, // use default DB
		})
	}

	return redisclient
}

// Getvaluefromcache returns the value of a key from cache
func Getvaluefromcache(key string) string {

	// bucket is ZERO for now
	// I am allowing it to be setup now
	rp := GetRedisPointer(0)

	sysid := GetSYSID()

	valuetoreturn, _ := rp.Get(sysid + key).Result()
	fmt.Println("Getvaluefromcache key: " + key + "  valuetoreturn:" + valuetoreturn)

	return valuetoreturn
}

// GetDBParmFromCache returns the value of a key from cache
func GetDBParmFromCache(collection string) *commonstruct.DatabaseX {

	database := new(commonstruct.DatabaseX)

	database.Collection = Getvaluefromcache(collection)
	database.Database = Getvaluefromcache("API.MongoDB.Database")
	database.Location = Getvaluefromcache("API.MongoDB.Location")

	return database
}

// GetDBParmAllFromCache returns all data from cache
func GetDBParmAllFromCache() commonstruct.RestEnvVariables {

	if envirvar.APIAPIServerIPAddress == "" {
		envirvar.APIAPIServerIPAddress = Getvaluefromcache("Web.APIServer.IPAddress")
		envirvar.APIAPIServerPort = Getvaluefromcache("Web.APIServer.Port")
		envirvar.WEBServerPort = Getvaluefromcache("WEBServerPort")
		envirvar.WEBDebug = Getvaluefromcache("Web.Debug")
		envirvar.RecordCurrencyTick = Getvaluefromcache("RecordCurrencyTick")
		envirvar.RunningFromServer = Getvaluefromcache("RunningFromServer")
		envirvar.AppBelnorthEnabled = Getvaluefromcache("AppBelnorthEnabled")
		envirvar.AppBitcoinEnabled = Getvaluefromcache("AppBitcoinEnabled")
		envirvar.AppFestaJuninaEnabled = Getvaluefromcache("AppFestaJuninaEnabled")
		envirvar.MSAPImainPort = Getvaluefromcache("MSAPImainPort")
		envirvar.MSAPIdishesPort = Getvaluefromcache("MSAPIdishesPort")
		envirvar.MSAPIordersPort = Getvaluefromcache("MSAPIordersPort")
		envirvar.SecurityMicroservice = Getvaluefromcache("SecurityMicroservice")
		envirvar.SecurityMicroserviceURL = Getvaluefromcache("SecurityMicroserviceURL")
		envirvar.APIMongoDBLocation = Getvaluefromcache("APIMongoDBLocation")
		envirvar.APIMongoDBDatabase = Getvaluefromcache("APIMongoDBDatabase")
	}

	return envirvar

}
