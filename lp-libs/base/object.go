package base

import (
	"encoding/json"
)

type OPHMediaUpdate struct {
	MediaId    int64   `json:"mid"`
	Summary    string  `json:"summary"`
	Header     byte    `json:"header"`
	Priority   byte    `json:"prio"`
	PlayMode   byte    `json:"mode"`
	PlayTime   []int32 `json:"time"`
	RepeatType byte    `json:"repeat"`
	RepeatDays int32   `json:"days"`
	CreateTime int64   `json:"created"`
	StartTime  int64   `json:"start"`
	ExpireTime int64   `json:"expired"`
}

type OPHMediaData struct {
	MediaId       int64   `json:"mid"`
	Header        byte    `json:"header"`
	Priority      byte    `json:"prio"`
	PlayMode      byte    `json:"mode"`
	PlayTime      []int32 `json:"time"`
	RepeatType    byte    `json:"repeat"`
	RepeatDays    int32   `json:"days"`
	AudioSize     int64   `json:"size"`
	AudioCodec    byte    `json:"codec"`
	AudioFormat   byte    `json:"format"`
	AudioChecksum uint32  `json:"cs"`
	AudioData     string  `json:"audio"`
}

type OPHMediaData2 struct {
	MediaId       int64   `json:"mid"`
	Header        int     `json:"header"`
	Priority      byte    `json:"prio"`
	Summary       string  `json:"summary"`
	CreateTime    int64   `json:"created"`
	StartTime     int64   `json:"start"`
	ExpireTime    int64   `json:"expired"`
	PlayMode      byte    `json:"mode"`
	PlayTime      []int32 `json:"ts"`
	RepeatType    byte    `json:"repeat"`
	RepeatDays    int32   `json:"days"`
	AudioSize     int64   `json:"size"`
	AudioCodec    byte    `json:"codec"`
	AudioFormat   byte    `json:"format"`
	AudioChecksum uint32  `json:"cs"`
	Url           string  `json:"url"`
	Auth          string  `json:"auth"`
}

type OPHMediaEdit struct {
	MediaId  int64   `json:"mid"`
	PlayMode byte    `json:"mode"`
	PlayTime []int32 `json:"time"`
}

type OPHMediaEdit2 struct {
	MediaId  int64   `json:"mid"`
	PlayMode byte    `json:"mode"`
	PlayTime []int32 `json:"ts"`
}

type OPHMediaDelete struct {
	MediaId int64 `json:"mid"`
}

type OPHVolume struct {
	VolumeLevel int64 `json:"volume_level"`
}

type OPHFMCfg struct {
	Volume int `json:"volume"`
	Auto   int `json:"auto"`
}

type OPHPlayHeader struct {
	HeaderIndex    int `json:"hid"`
	HeaderPriority int `json:"prio"`
}

type OPHSensor struct {
	SensorId   int64   `json:"sid"`
	Enable     byte    `json:"enable"`
	Name       string  `json:"name"`
	Thresholds []int32 `json:"threshold"`
}

type OPHAlarm struct {
	EventId int64 `json:"id"`
	//EventType    byte   `json:"type"`
	Name         string `json:"name"`
	SensorId     int64  `json:"sid"`
	AlarmState   byte   `json:"alarm"`
	Mode         byte   `json:"mode"`
	ActiveTime   int16  `json:"active"`
	InactiveTime int16  `json:"inactive"`
	AutoDays     byte   `json:"days"`
	PlayFile     int64  `json:"mid"`
}

type OPUGeneric struct {
	GroupList       []int64     `json:"group"`
	MediaIdLastest  int64       `json:"mid"`
	Volume          byte        `json:"volume"`
	LocalIp         string      `json:"ip"`
	VCode           string      `json:"vcode"`
	PhoneNumber     string      `json:"phone"`
	CameraList      []OPUCamera `json:"camera"`
	SensorList      []OPUSensor `json:"sensor"`
	AlarmList       []OPUAlarm  `json:"alarm"`
	FirmwareVersion string      `json:"fvers"`
	FMVolume        int32       `json:"fmvolume"`
	FMAuto          int32       `json:"fmauto"`
	TxType          int32       `json:"txtype"`
	ConnStatus      byte
	ConnTime        int64
	WanIP           string
}

type OPUCamera struct {
	CameraId       int64  `json:"id"`
	CameraName     string `json:"name"`
	CameraLocalIp  string `json:"ip"`
	CameraTypeId   int32  `json:"type"`
	CameraHttpPort int32  `json:"http"`
	CameraRtspPort int32  `json:"rtsp"`
	CameraUsername string `json:"user"`
	CameraPassword string `json:"pass"`
}

type OPUPhoneNumber struct {
	PhoneNumber string `json:"phone_number"`
}

type OPUSensor struct {
	SensorId    int64   `json:"id"`
	Enable      byte    `json:"enable"`
	Name        string  `json:"name"`
	Type        int32   `json:"type"`
	State       byte    `json:"state"`
	Value       int64   `json:"value"`
	Thresholds  []int32 `json:"threshold"`
	CreatedTime int64   `json:"ctime"`
}

type OPUAlarm struct {
	EventId      int64  `json:"id"`
	EventType    int32  `json:"type"`
	Name         string `json:"name"`
	State        byte   `json:"state"`
	SensorId     int64  `json:"sid"`
	SensorState  byte   `json:"alarm"`
	Mode         byte   `json:"mode"`
	ActiveTime   int16  `json:"active"`
	InactiveTime int16  `json:"inactive"`
	AutoDays     byte   `json:"days"`
	PlayFile     int64  `json:"mid"`
	OccurTime    int64  `json:"occur"`
}

type OPULog struct {
	LogId       int64 `json:"id"`
	CreatedTime int64 `json:"time"`
	LogType     byte  `json:"type"`
	MediaId     int64 `json:"mid"`
	EventId     int64 `json:"eid"`
	SensorId    int64 `json:"sid"`
	State       byte  `json:"state"`
	Value       int64 `json:"value"`
}

type OPUStatus struct {
	TxType     byte   `json:"conn"`
	Temp       int16  `json:"temp"`
	SpeakerErr byte   `json:"spkerr"`
	SpeakerSta uint16 `json:"spksta"`
	MCsq       byte   `json:"csqm"`
	WiCsq      byte   `json:"csqw"`
	FMStatus   byte   `json:"fmsta"`
}

type MqttMessagePublish struct {
	Id     int64           `json:"id"`
	OpCode byte            `json:"opcode"`
	Data   json.RawMessage `json:"data"`
}

type OPUConnStatus struct {
	McuId  int64  `json:"mcu_id"`
	Status byte   `json:"status"`
	McuIp  string `json:"mcu_ip"`
}

type MqttClientConnection struct {
	ClientId     string `json:"clientid"`
	UserName     string `json:"username"`
	IpAddress    string `json:"ipaddress"`
	CleanSession bool   `json:"clean_sess"`
	Protocal     int    `json:"protocal"`
	ConnAck      int32  `json:"connack"`
	Ts           int64  `json:"ts"`
	Reason       string `json:"reason"`
}

type OPHFM struct {
	FMId             int64   `json:"mid"`
	FMHeader         int32   `json:"header"`
	FMFrequence      int32   `json:"frequency"`
	FMPlayMode       int32   `json:"mode"`
	FMAutoSwitchTime int32   `json:"swtime"`
	FMPriority       int32   `json:"prio"`
	FMPlayTime       []int32 `json:"time"`
	FMPlayRepeatType int32   `json:"repeat"`
	FMPlayRepeatDays int32   `json:"days"`
	FMPlayStart      int64   `json:"start"`
	FMPlayExpire     int64   `json:"expired"`
	FMPlayDuration   int32   `json:"duration"`
}

type OPHFM2 struct {
	FMId             int64   `json:"mid"`
	FMHeader         int32   `json:"header"`
	FMFrequence      int32   `json:"frequency"`
	FMPlayMode       int32   `json:"mode"`
	FMAutoSwitchTime int32   `json:"swtime"`
	FMPriority       int32   `json:"prio"`
	FMPlayTime       []int32 `json:"ts"`
	FMPlayRepeatType int32   `json:"repeat"`
	FMPlayRepeatDays int32   `json:"days"`
	FMPlayStart      int64   `json:"start"`
	FMPlayExpire     int64   `json:"expired"`
	FMPlayDuration   int32   `json:"duration"`
}
