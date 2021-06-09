package settings

import (
	"encoding/json"
	"io/ioutil"
)

type GlogConfigs struct {
	LogDir  string
	MaxSize uint64
	V       int
}

type Settings struct {
	RestfulApiPort       int
	RestfulApiHost       string
	GlogConfig           *GlogConfigs
	CmsInfo              *CmsInfo
	MqttInfo             *MqttInfo
	RedisInfo            *RedisInfo
	UserNameAuth         string
	PasswordAuth         string
	KeyAuth				 string
	TopicDeviceMainflux  string
	ThingAuthPush *ThingAuthPush
	FileServerPath       string
	FileServerPublicPath string
	TTTTMInfo *TTTMInfo

}

type CmsInfo struct {
	ServerAddress string
	ServerPort    int
	BasePath      string
	UserName      string
	Password      string
}
type ThingAuthPush struct {
	ID string
	Key string
}
type MqttInfo struct {
	ServerAddress   string
	ServerAddressHTTP   string
	ServerPort      int
	ServerPortAuth      int
	NodeName        string
	UserName        string
	Password        string
	HttpApiPort     int
	HttpApiUserName string
	HttpApiPassword string
}
func GetThingAuthPush() *ThingAuthPush{
	return settings.ThingAuthPush
}


func GetTopicDeviceMainflux() string  {
	return settings.TopicDeviceMainflux

}
type RedisInfo struct {
	Host1       string
	Host2       string
	Host3       string
}
type TTTMInfo struct{
	TTTTMMainFluxAddress string
	Port int
}
var settings Settings = Settings{}

func init() {
	//C:\Users\Admin\go\src\api_cv\mcu_api\setting.json
	//C:\Users\Admin\go\src\api_cv\lp-cms-client\setting.json
	//	content, err := ioutil.ReadFile("C:\\Users\\Admin\\go\\src\\api_cv\\mcu_api\\setting.json")
	//	content, err := ioutil.ReadFile("C:\\Users\\Admin\\go\\src\\api_cv\\lp-cms-client\\setting.json")
	//	abc:= filepath.Join(NamePackage)

	//absPath, _ := filepath.Abs(filepath.Join("mcu_api","setting.json"))
	//fmt.Println(abc)
	content, err := ioutil.ReadFile("setting.json")
	if err != nil {
		panic(err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		panic(jsonErr)
	}
}

func GetRestfulApiPort() int {
	return settings.RestfulApiPort
}
func GetRestfulApiHost() string {
	return settings.RestfulApiHost
}

func GetCmsInfo() *CmsInfo {
	return settings.CmsInfo
}

func GetMqttInfo() *MqttInfo {
	return settings.MqttInfo
}

func GetRedisInfo() *RedisInfo {
	return settings.RedisInfo
}

func GetTTTMInfo() *TTTMInfo  {
	return settings.TTTTMInfo
}

func GetGlogConfig() *GlogConfigs {
	return settings.GlogConfig
}

func GetUserNameAuth() string {
	return settings.UserNameAuth
}
func GetKeyAuth() string{
	return settings.KeyAuth
}
func GetPasswordAuth() string {
	return settings.PasswordAuth
}

func GetFileServerPath() string {
	return settings.FileServerPath
}


func GetFileServerPublicPath() string {
	return settings.FileServerPublicPath
}
