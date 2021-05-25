package services

import (
	"tttm_cms_api/lp-libs/base"
	"tttm_cms_api/lp-libs/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/silkeh/senml"
	"strconv"
	"time"
)
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

