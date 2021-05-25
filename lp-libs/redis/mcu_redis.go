package redis

import (
	"encoding/json"
	"strconv"
	"tttm_cms_api/lp-libs/base"
)

var mcuInfoFields = []string{
	Mcu_Vcode,
	Mcu_ConnectionStatus,
	Mcu_ConnectionTime,
	Mcu_PhoneNumber,
	Mcu_LocalIp,
	Mcu_WanIp,
	Mcu_MediaIdLastest,
	Mcu_Volume,
	Mcu_GroupList,
	Mcu_SensorList,
	Mcu_CameraList,
	Mcu_AlarmList,
	Mcu_Version,
	Mcu_FM_Volume,
	Mcu_FM_Auto,
}

func ExistsMcuId(mcuId int64) bool {
	cmd := iClient.Exists(strconv.FormatInt(mcuId, 10) + "TTTM")

	if cmd.Err() != nil || cmd.Val() == false {
		return false
	} else {
		return true
	}
}

func GetMcuInfo(mcuId int64) *base.OPUGeneric {
	result, err := iClient.HMGet(strconv.FormatInt(mcuId, 10) + "TTTM", mcuInfoFields...).Result()
	if err != nil {
		return nil
	} else {
		if result == nil {
			return nil
		}

		vcode := ""
		if result[0] != nil {
			vcode = result[0].(string)
		}

		connStatus := base.MCU_DISCONNECTED
		if result[1] != nil {
			i, _ := strconv.Atoi(result[1].(string))
			connStatus = byte(i)
		}
		connTime := int64(0)
		if result[2] != nil {
			i, _ := strconv.ParseInt(result[2].(string), 10, 64)
			connTime = i
		}
		phone := ""
		if result[3] != nil {
			phone = result[3].(string)
		}
		localIp := ""
		if result[4] != nil {
			localIp = result[4].(string)
		}
		wanIp := ""
		if result[5] != nil {
			wanIp = result[5].(string)
		}
		mId := int64(0)
		if result[6] != nil {
			i, _ := strconv.ParseInt(result[6].(string), 10, 64)
			mId = i
		}
		vol := byte(0)
		if result[7] != nil {
			i, _ := strconv.Atoi(result[7].(string))
			vol = byte(i)
		}
		groups := []int64{}
		if result[8] != nil {
			json.Unmarshal([]byte(result[8].(string)), &groups)
		}
		sensors := []base.OPUSensor{}
		if result[9] != nil {
			json.Unmarshal([]byte(result[9].(string)), &sensors)
		}
		cameras := []base.OPUCamera{}
		if result[10] != nil {
			json.Unmarshal([]byte(result[10].(string)), &cameras)
		}
		alarms := []base.OPUAlarm{}
		if result[11] != nil {
			json.Unmarshal([]byte(result[11].(string)), &alarms)
		}
		version := ""
		if result[12] != nil {
			version = result[12].(string)
		}
		fmvol := int32(0)
		if result[13] != nil {
			i, _ := strconv.Atoi(result[13].(string))
			fmvol = int32(i)
		}
		fmauto := int32(0)
		if result[14] != nil {
			i, _ := strconv.Atoi(result[14].(string))
			fmauto = int32(i)
		}
		mcu := &base.OPUGeneric{
			VCode:           vcode,
			ConnStatus:      connStatus,
			ConnTime:        connTime,
			PhoneNumber:     phone,
			LocalIp:         localIp,
			WanIP:           wanIp,
			MediaIdLastest:  mId,
			Volume:          vol,
			GroupList:       groups,
			SensorList:      sensors,
			CameraList:      cameras,
			AlarmList:       alarms,
			FirmwareVersion: version,
			FMVolume:        fmvol,
			FMAuto:          fmauto,
		}
		return mcu
	}
}

func UpdateMcuInfo(mcuId int64, generic *base.OPUGeneric) error {
	groupListStrJ := `[]`
	cameraListStrJ := `[]`
	sensorListStrJ := `[]`
	alarmListStrJ := `[]`
	if len(generic.GroupList) > 0 {
		groupsJson, err := json.Marshal(&generic.GroupList)
		if err != nil {
			return err
		}
		groupListStrJ = string(groupsJson)
	}

	if len(generic.SensorList) > 0 {
		sensorsJson, err := json.Marshal(&generic.SensorList)
		if err != nil {
			return err
		}
		sensorListStrJ = string(sensorsJson)
	}

	if len(generic.CameraList) > 0 {
		camerasJson, err := json.Marshal(&generic.CameraList)
		if err != nil {
			return err
		}
		cameraListStrJ = string(camerasJson)
	}
	if len(generic.AlarmList) > 0 {
		alarmsJson, err := json.Marshal(&generic.AlarmList)
		if err != nil {
			return err
		}
		alarmListStrJ = string(alarmsJson)
	}
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM",
		Mcu_Vcode, generic.VCode,
		Mcu_PhoneNumber, generic.PhoneNumber,
		Mcu_LocalIp, generic.LocalIp,
		Mcu_MediaIdLastest, strconv.FormatInt(generic.MediaIdLastest, 10),
		Mcu_Volume, strconv.Itoa(int(generic.Volume)),
		Mcu_GroupList, groupListStrJ,
		Mcu_SensorList, sensorListStrJ,
		Mcu_CameraList, cameraListStrJ,
		Mcu_AlarmList, alarmListStrJ,
		Mcu_Version, generic.FirmwareVersion,
		Mcu_FM_Volume, strconv.Itoa(int(generic.FMVolume)),
		Mcu_FM_Auto, strconv.Itoa(int(generic.FMAuto)),
		Mcu_TxType, strconv.Itoa(int(generic.TxType))).Err()
}

func UpdateMcuCameraList(mcuId int64, cameras []base.OPUCamera) error {
	cameraListStrJ := `[]`
	if len(cameras) > 0 {
		camerasJson, err := json.Marshal(&cameras)
		if err != nil {
			return err
		}
		cameraListStrJ = string(camerasJson)
	}

	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM", Mcu_CameraList, cameraListStrJ).Err()
}

func UpdateMcuPhone(mcuId int64, phone string) error {
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM", Mcu_PhoneNumber, phone).Err()
}

func UpdateMcuConnectionStatus(mcuId int64, status byte, connTime int64, ip string) error {
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM",
		Mcu_ConnectionStatus, strconv.Itoa(int(status)), Mcu_WanIp, ip, Mcu_ConnectionTime, strconv.FormatInt(connTime, 10)).Err()
}

func UpdateMcuVolume(mcuId int64, volume int32) error {
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM", Mcu_Volume, strconv.Itoa(int(volume))).Err()
}

func UpdateMcuGroupList(mcuId int64, groups []int64) error {
	groupListStrJ := `[]`
	if len(groups) > 0 {
		groupsJson, err := json.Marshal(&groups)
		if err != nil {
			return err
		}
		groupListStrJ = string(groupsJson)
	}
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM", Mcu_GroupList, groupListStrJ).Err()
}

func UpdateMcuCfgAlarmList(mcuId int64, alarms []base.OPHAlarm) error {
	alarmListStrJ := `[]`
	if len(alarms) > 0 {
		alarmsJson, err := json.Marshal(&alarms)
		if err != nil {
			return err
		}
		alarmListStrJ = string(alarmsJson)
	}
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM", Mcu_AlarmList, alarmListStrJ).Err()
}

func UpdateMcuCfgSensorList(mcuId int64, sensors []base.OPHSensor) error {
	cfgSensorListStrJ := `[]`
	if len(sensors) > 0 {
		cfgSensorJson, err := json.Marshal(&sensors)
		if err != nil {
			return err
		}
		cfgSensorListStrJ = string(cfgSensorJson)
	}
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM", Mcu_AlarmList, cfgSensorListStrJ).Err()
}

func UpdateMcuStatus(mcuId int64, status *base.OPUStatus) error {
	return iClient.HMSet(strconv.FormatInt(mcuId, 10) + "TTTM",
		Mcu_TxType, strconv.Itoa(int(status.TxType)),
		Mcu_SpeakerErr, strconv.Itoa(int(status.SpeakerErr)),
		Mcu_SpeakerCurrent, strconv.Itoa(int(status.SpeakerSta/10)),
		Mcu_MobileCsq, strconv.Itoa(int(status.MCsq)),
		Mcu_WifiCsq, strconv.Itoa(int(status.WiCsq)),
		Mcu_Temp, strconv.Itoa(int(status.Temp/10))).Err()
}

func UpdateMcuClientId(mcuId string) error {
	err := SetAclMcuUser(mcuId)
	err = iClient.HMSet(mcuId +"TTTM", Mcu_Vcode, "").Err()
	return err
}

func DeleteMcuById(id int64) error {
	return iClient.Del(strconv.FormatInt(id, 10)).Err()
}

func GetAllMcuIds() ([]string, error) {
	return iClient.Keys("*").Result()
}
