package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/silkeh/senml"
	"strconv"
	"strings"
	"time"
	"tttm_cms_api/lp-libs/base"
	"tttm_cms_api/lp-libs/models"
)

const NULL = "NULL"
func arrayToString(a []int64, delim string) string {
	if len(a) ==0{
		return NULL
	}
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func arrayToString32(a []int32, delim string) string {
	if len(a) ==0{
		return NULL
	}
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func HandleNullObj(obj interface{},op_code byte)(interface{})  {
	switch op_code {
	case base.OPU_GENERIC:
		var generic *base.OPUGeneric = obj.(*base.OPUGeneric)
		if generic.LocalIp == ""{
			generic.LocalIp = NULL
		}
		if generic.VCode=="" {
			generic.VCode = NULL
		}
		if generic.PhoneNumber == ""{
			generic.PhoneNumber = NULL
		}
		if generic.FirmwareVersion == ""{
			generic.FirmwareVersion = NULL
		}
		return generic
	case base.OPU_CAMERA:
		var camera base.OPUCamera = obj.(base.OPUCamera)
		if camera.CameraName == ""{
			camera.CameraName = NULL
		}
		if camera.CameraLocalIp == ""{
			camera.CameraLocalIp = NULL
		}
		return camera
	case base.OPU_SENSOR:
		var sensor base.OPUSensor = obj.(base.OPUSensor)
		if sensor.Name == ""{
			sensor.Name = NULL
		}
		return sensor
	case base.OPU_ALARM:
		var alarm base.OPHAlarm = obj.(base.OPHAlarm)
		if alarm.Name == "" { 
			alarm.Name = NULL
		}
		return alarm
	default:
		return nil
	}
}
func ConvertJsonToSenMLVer2(mcu_id int64, obj interface{}, op_code byte,topicLogMainflux string)(string, error) {

	now := time.Now()
	switch op_code {
	case base.OPU_GENERIC:
		fmt.Println("IN OPU_GENERIC ver2")
		var generic *base.OPUGeneric = obj.(*base.OPUGeneric)
		generic = HandleNullObj(generic,base.OPU_GENERIC).(*base.OPUGeneric)
		volumn := float64(generic.Volume)
		if generic.LocalIp ==""{
			generic.LocalIp = "null"
		}
		list := []senml.Measurement{
			senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_Volume", volumn, "Volume", now, 0),
			senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_LocalIp", generic.LocalIp, "LocalIp", now, 0),
			senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_GroupList", arrayToString(generic.GroupList,","), "GroupList", now, 0),
			senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_PhoneNumber", generic.PhoneNumber, "PhoneNumber", now, 0),
			senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_VCode", generic.VCode, "VCode", now, 0),
			//senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC","base" , "BaseNameBaseName", now, 0),
			//senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_ConnStatus", string(generic.ConnStatus), "ConnStatus", now, 0),
			senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_MediaIdLastest", float64(generic.MediaIdLastest), "MediaIdLastest", now, 0),
			senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_FMVolume", float64(generic.FMVolume), "FMVolume", now, 0),
			senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_FMAuto", float64(generic.FMAuto), "FMAuto", now, 0),
			senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_TxType", float64(generic.TxType), "TxType", now, 0),
			//senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_ConnTime",strconv.FormatInt(generic.ConnTime, 10) , "ConnTime", now, 0),
			//senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_WanIP","generic.WanIP" , "WanIP", now, 0),
			senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_FirmwareVersion", generic.FirmwareVersion, "FirmwareVersion", now, 0),

		}
		var listOPUCamera []base.OPUCamera = generic.CameraList
		for _,camera :=range listOPUCamera{
			camera = HandleNullObj(camera,base.OPU_CAMERA).(base.OPUCamera)
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraName",
				camera.CameraName , "CameraName", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraLocalIp",
				camera.CameraLocalIp , "CameraLocalIp", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraTypeId",
				float64(camera.CameraTypeId) , "CameraTypeId", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraHttpPort",
				float64(camera.CameraHttpPort) , "CameraHttpPort", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraRtspPort",
				float64(camera.CameraRtspPort) , "CameraRtspPort", now, 0))
		}
		var listOPUSensor []base.OPUSensor = generic.SensorList
		for _,sensor :=range listOPUSensor{
			sensor = HandleNullObj(sensor,base.OPU_SENSOR).(base.OPUSensor)
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Enable",
				string(sensor.Enable) , "Enable", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Name",
				sensor.Name , "Name", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Type",
				float64(sensor.Type) , "Type", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Value",
				float64(sensor.Value) , "Value", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_State",
				string(sensor.State) , "State", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Thresholds",
				arrayToString32(sensor.Thresholds,",") , "Thresholds", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_CreatedTime",
				strconv.FormatInt(sensor.CreatedTime, 10), "CreatedTime", now, 0))
		}
		var listOPUAlarm []base.OPUAlarm = generic.AlarmList
		for _,alarm :=range listOPUAlarm{
			alarm = HandleNullObj(alarm,base.OPU_ALARM).(base.OPUAlarm)
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_Name",
				alarm.Name , "Name", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_State",
				string(alarm.State ), "State", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_SensorState",
				string(alarm.SensorState ), "SensorState", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_Mode",
				string(alarm.Mode ), "Mode", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_AutoDays",
				string(alarm.AutoDays ), "AutoDays", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_EventType",
				float64(alarm.EventType), "EventType", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_SensorId",
				float64(alarm.SensorId), "SensorId", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_ActiveTime",
				float64(alarm.ActiveTime), "ActiveTime", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_InactiveTime",
				float64(alarm.InactiveTime), "InactiveTime", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_PlayFile",
				float64(alarm.PlayFile), "PlayFile", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_OccurTime",
				float64(alarm.OccurTime), "OccurTime", now, 0))

		}
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		//stringData := strings.ReplaceAll(string(data),"\\","")

		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	case base.OPU_SENSOR:
		fmt.Println("ConvertJsonToSenML_InSensor")
		var sensors []base.OPUSensor = obj.([]base.OPUSensor)
		fmt.Println(len(sensors))
		list := []senml.Measurement{}
		for _,sensor :=range sensors{
			sensor = HandleNullObj(sensor,base.OPU_SENSOR).(base.OPUSensor)
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Enable",
				string(sensor.Enable) , "Enable", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Name",
				sensor.Name , "Name", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Type",
				float64(sensor.Type) , "Type", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Value",
				float64(sensor.Value) , "Value", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_State",
				string(sensor.State) , "State", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_Thresholds",
				arrayToString32(sensor.Thresholds,",") , "Thresholds", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR_SensorId:"+strconv.FormatInt(sensor.SensorId, 10)+"_CreatedTime",
				strconv.FormatInt(sensor.CreatedTime, 10), "CreatedTime", now, 0))
		}
		//rs,_:=json.Marshal(sensors)
		//list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR",string(rs), "SENSOR", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
		//lấy thông tin nhiệt độ bên trong hộp thiết bị
	case base.OPU_STATUS:
		fmt.Println("In OPUT_STATUS")
		var status base.OPUStatus = obj.(base.OPUStatus)
		list := []senml.Measurement{}
		list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS_Temp", float64(status.Temp), "Temp", now, 0))
		list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS_SpeakerSta", float64(status.SpeakerSta), "SpeakerSta", now, 0))
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS_TxType", string(status.TxType), "TxType", now, 0))
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS_MCsq", string(status.MCsq), "MCsq", now, 0))
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS_WiCsq", string(status.WiCsq), "WiCsq", now, 0))
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS_SpeakerErr", string(status.SpeakerErr), "SpeakerErr", now, 0))
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS_FMStatus", string(status.FMStatus), "FMStatus", now, 0))
		//rs,_:=json.Marshal(status)
		//list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS",string(rs), "STATUS", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	case base.OPU_CAMERA:
		fmt.Println("ConvertJsonToSenML_InCamera")
		var cameras []base.OPUCamera = obj.([]base.OPUCamera)
		fmt.Println(len(cameras))
		list := []senml.Measurement{}
		for _,camera :=range cameras{
			camera = HandleNullObj(camera,base.OPU_CAMERA).(base.OPUCamera)
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_CAMERA_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraName",
				camera.CameraName , "CameraName", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_CAMERA_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraLocalIp",
				camera.CameraLocalIp , "CameraLocalIp", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_CAMERA_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraTypeId",
				float64(camera.CameraTypeId) , "CameraTypeId", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_CAMERA_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraHttpPort",
				float64(camera.CameraHttpPort) , "CameraHttpPort", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_CAMERA_OPUCameraID:"+strconv.FormatInt(camera.CameraId, 10)+"_CameraRtspPort",
				float64(camera.CameraRtspPort) , "CameraRtspPort", now, 0))
		}
		
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//Lấy giá trị của Phone
	case base.OPU_PHONE:
		fmt.Println("ConvertJsonToSenML_InPhone")
		var phone string = obj.(string)
		//var cameras []base.OPUCamera = obj.([]base.OPUCamera)
		fmt.Println(len(phone))
		list := []senml.Measurement{}
		//rs,_:=json.Marshal(phone)
		if phone == ""{
			phone = NULL
		}
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_PHONE",phone, "PHONE", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
		//Lấy giá trị của sensor
		case base.OPU_ALARM:
			fmt.Println("ConvertJsonToSenML_InSensor")
			var alarms []base.OPUAlarm = obj.([]base.OPUAlarm)
			
			list := []senml.Measurement{}
			for _,alarm :=range alarms{
				alarm = HandleNullObj(alarm,base.OPU_ALARM).(base.OPUAlarm)
				list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_Name",
					alarm.Name , "Name", now, 0))
				list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_State",
					string(alarm.State ), "State", now, 0))
				list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_SensorState",
					string(alarm.SensorState ), "SensorState", now, 0))
				list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_Mode",
					string(alarm.Mode ), "Mode", now, 0))
				list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_AutoDays",
					string(alarm.AutoDays ), "AutoDays", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_EventType",
					float64(alarm.EventType), "EventType", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_SensorId",
					float64(alarm.SensorId), "SensorId", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_ActiveTime",
					float64(alarm.ActiveTime), "ActiveTime", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_InactiveTime",
					float64(alarm.InactiveTime), "InactiveTime", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_PlayFile",
					float64(alarm.PlayFile), "PlayFile", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM_EventId:"+strconv.FormatInt(alarm.EventId, 10)+"_OccurTime",
					float64(alarm.OccurTime), "OccurTime", now, 0))

			}
			data, err := senml.EncodeJSON(list)
			if err != nil {
				fmt.Print("Error encoding to JSON:", err)
				return "",err
			}
			fmt.Printf("%s\n", data)
			errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
			if errPublish != nil {
				return "", errPublish
			}
			return string(data), nil
	//Lấy giá trị của media
	case base.OPU_MEDIA:
		fmt.Println("ConvertJsonToSenML_InSensor")
		var items []models.RecordStat = obj.([]models.RecordStat)
		fmt.Println(len(items))
		list := []senml.Measurement{}
		for _,item := range  items{
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_McuId", float64(item.McuId) , "McuId", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecStatus", float64(item.RecStatus) , "RecStatus", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecPlayMode", float64(item.RecPlayMode) , "RecPlayMode", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecPriority", float64(item.RecPriority) , "RecPriority", now, 0))
			list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecPlayTime", arrayToString32(item.RecPlayTime,",") , "RecPlayTime", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecPlayRepeatType", float64(item.RecPlayRepeatType) , "RecPlayRepeatType", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecPlayRepeatDays", float64(item.RecPlayRepeatDays) , "RecPlayRepeatDays", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecCreateTime", float64(item.RecCreateTime) , "RecCreateTime", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecPlayStart", float64(item.RecPlayStart) , "RecPlayStart", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecPlayExpire", float64(item.RecPlayExpire) , "RecPlayExpire", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecAudioCodec", float64(item.RecAudioCodec) , "RecAudioCodec", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecAudioFormat", float64(item.RecAudioFormat) , "RecAudioFormat", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecSize", float64(item.RecSize) , "RecSize", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_RecChecksum", float64(item.RecChecksum) , "RecChecksum", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_FMFrequence", float64(item.FMFrequence) , "FMFrequence", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_FMPlayDuration", float64(item.FMPlayDuration) , "FMPlayDuration", now, 0))
			list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA_RecId:"+strconv.FormatInt(item.RecId, 10)+
				"_FMAutoSwitchTime", float64(item.FMAutoSwitchTime) , "FMAutoSwitchTime", now, 0))
		}
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
		//lấy thông tin nhiệt độ bên trong hộp thiết bị
		case base.OPU_LOG:
			fmt.Println("ConvertJsonToSenML_InSensor")
			var logs []base.OPULog = obj.([]base.OPULog)
			fmt.Println(len(logs))
			list := []senml.Measurement{}
			for _,log := range logs{
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG_LogID:"+strconv.FormatInt(log.LogId, 10)+
					"_CreatedTime",float64(log.CreatedTime), "CreatedTime", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG_LogID:"+strconv.FormatInt(log.LogId, 10)+
					"_MediaId",float64(log.MediaId), "MediaId", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG_LogID:"+strconv.FormatInt(log.LogId, 10)+
					"_EventId",float64(log.EventId), "EventId", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG_LogID:"+strconv.FormatInt(log.LogId, 10)+
					"_SensorId",float64(log.SensorId), "SensorId", now, 0))
				list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG_LogID:"+strconv.FormatInt(log.LogId, 10)+
					"_Value",float64(log.Value), "Value", now, 0))
				list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG_LogID:"+strconv.FormatInt(log.LogId, 10)+
					"_LogType",string(log.LogType), "LogType", now, 0))
				list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG_LogID:"+strconv.FormatInt(log.LogId, 10)+
					"_State",string(log.State), "State", now, 0))
			}
			data, err := senml.EncodeJSON(list)
			if err != nil {
				fmt.Print("Error encoding to JSON:", err)
				return "",err
			}
			fmt.Printf("%s\n", data)
			errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
			if errPublish != nil {
				return "", errPublish
			}
			return string(data), nil
	}
	return "", errors.New("Wrong Opcode to convertJsonToSenML")
}
func ConvertJsonToSenML(mcu_id int64, obj interface{}, op_code byte,topicLogMainflux string) (string, error) {
	now := time.Now()
	switch op_code {
	//Lấy giá trị âm thanh của loa, còn âm thanh FM chưa lấy
	case base.OPU_GENERIC:
		fmt.Println("IN OPU_GENERIC")
		var generic *base.OPUGeneric = obj.(*base.OPUGeneric)
		volumn := float64(generic.Volume)
		list := []senml.Measurement{
			senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_volume", volumn, senml.Decibel, now, 0),
		}
		rs,_:=json.Marshal(generic)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_GENERIC",string(rs), "GENERIC", now, 0))

		//fmt.Print(len(list))
		fmt.Println("jsonGeneric",string(rs))

		data, err := senml.EncodeJSON(list)

		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		//stringData := strings.ReplaceAll(string(data),"\\","")

		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//Lấy giá trị của cameras
	case base.OPU_CAMERA:
		fmt.Println("ConvertJsonToSenML_InCamera")
		var cameras []base.OPUCamera = obj.([]base.OPUCamera)
		fmt.Println(len(cameras))
		list := []senml.Measurement{}
		rs,_:=json.Marshal(cameras)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_CAMERA",string(rs), "CAMERA", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//Lấy giá trị của Phone
	case base.OPU_PHONE:
		fmt.Println("ConvertJsonToSenML_InPhone")
		var phone string = obj.(string)
		//var cameras []base.OPUCamera = obj.([]base.OPUCamera)
		fmt.Println(len(phone))
		list := []senml.Measurement{}
		rs,_:=json.Marshal(phone)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_PHONE",string(rs), "PHONE", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//Lấy giá trị của sensor
	case base.OPU_SENSOR:
		fmt.Println("ConvertJsonToSenML_InSensor")
		var sensors []base.OPUSensor = obj.([]base.OPUSensor)
		fmt.Println(len(sensors))
		list := []senml.Measurement{}
		for _, sensor := range sensors {
			createTime := time.Unix(0, sensor.CreatedTime)
			fmt.Print("CreateTime", createTime)
			/*McuID+typeSensor+sensorID*/
			/*SensorName+typeSensor+sensorID*/
			nameSensor := "TTTM_"+strconv.FormatInt(mcu_id, 10) + "_" + sensor.Name + "_" + strconv.FormatInt(sensor.SensorId, 10)
			list = append(list, senml.NewValue(nameSensor,
				float64(sensor.Value), "ValueSensor", createTime, 0))
		}
		rs,_:=json.Marshal(sensors)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_SENSOR",string(rs), "SENSOR", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//Lấy giá trị của alarm
	case base.OPU_ALARM:
		fmt.Println("ConvertJsonToSenML_InSensor")
		var alarms []base.OPUAlarm = obj.([]base.OPUAlarm)
		fmt.Println(len(alarms))
		list := []senml.Measurement{}
		rs,_:=json.Marshal(alarms)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_ALARM",string(rs), "ALARM", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//Lấy giá trị của media
	case base.OPU_MEDIA:
		fmt.Println("ConvertJsonToSenML_InSensor")
		var items []models.RecordStat = obj.([]models.RecordStat)
		fmt.Println(len(items))
		list := []senml.Measurement{}
		rs,_:=json.Marshal(items)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_MEDIA",string(rs), "MEDIA", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//lấy thông tin nhiệt độ bên trong hộp thiết bị
	case base.OPU_STATUS:
		fmt.Println("In OPUT_STATUS")
		var status base.OPUStatus = obj.(base.OPUStatus)
		list := []senml.Measurement{}
		list = append(list, senml.NewValue("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_Temp", float64(status.Temp), senml.Degree, now, 0))
		rs,_:=json.Marshal(status)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_STATUS",string(rs), "STATUS", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	//Lấy giá trị của media
	case base.OPU_LOG:
		fmt.Println("ConvertJsonToSenML_InSensor")
		var logs []base.OPULog = obj.([]base.OPULog)
		fmt.Println(len(logs))
		list := []senml.Measurement{}
		rs,_:=json.Marshal(logs)
		list = append(list, senml.NewString("TTTM_"+strconv.FormatInt(mcu_id, 10)+"_OPU_LOG",string(rs), "LOG", now, 0))
		data, err := senml.EncodeJSON(list)
		if err != nil {
			fmt.Print("Error encoding to JSON:", err)
			return "",err
		}
		fmt.Printf("%s\n", data)
		errPublish := publishMessageSenML(topicLogMainflux, 0, false, string(data))
		if errPublish != nil {
			return "", errPublish
		}
		return string(data), nil
	}
	return "", errors.New("Wrong Opcode to convertJsonToSenML")
}

