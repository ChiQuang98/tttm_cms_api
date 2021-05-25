package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"tttm_cms_api/lp-libs/base"
	"tttm_cms_api/lp-libs/models"
	"tttm_cms_api/lp-libs/redis"
	"tttm_cms_api/lp-libs/settings"
	"github.com/golang/glog"
)

var cmsInfo *settings.CmsInfo = settings.GetCmsInfo()
var serverBasePath string = fmt.Sprintf("http://%s:%d/%s", cmsInfo.ServerAddress, cmsInfo.ServerPort, cmsInfo.BasePath)

func CmsUpdateMcuTelInfo(mcuId int64, phoneNumber string) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsUpdateMcuTelInfo err: ", err)
		}
	}()
	telInfo := &models.McuTelInfo{
		McuId:        mcuId,
		McuTelNumber: phoneNumber,
	}
	client := &http.Client{}
	buf, err := json.Marshal(telInfo)
	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPut, serverBasePath+"/mcu-update-tel-info", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			defer res.Body.Close()
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsUpdateMcuConfig(mcuId int64, mcuGeneric *base.OPUGeneric) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsUpdateMcuConfig err: ", err)
		}
	}()
	cameras := []models.McuCamera{}
	camGeneric := mcuGeneric.CameraList
	for _, cam := range camGeneric {
		camera := models.McuCamera{
			McuCameraId:       cam.CameraId,
			McuCameraName:     cam.CameraName,
			McuCameraTypeId:   cam.CameraTypeId,
			McuCameraLocalIp:  cam.CameraLocalIp,
			McuCameraHttpPort: cam.CameraHttpPort,
			McuCameraRtspPort: cam.CameraRtspPort,
			McuCameraUsername: cam.CameraUsername,
			McuCameraPassword: cam.CameraPassword,
		}
		cameras = append(cameras, camera)
	}

	sensors := []models.CfgSensor{}
	sensorGeneric := mcuGeneric.SensorList
	for _, ssr := range sensorGeneric {
		enable := true
		if ssr.Enable != 1 {
			enable = false
		}
		sensor := models.CfgSensor{
			McuSensorId:         ssr.SensorId,
			McuSensorType:       ssr.Type,
			McuSensorName:       ssr.Name,
			McuSensorEnable:     enable,
			McuSensorThresholds: ssr.Thresholds,
		}
		sensors = append(sensors, sensor)
	}

	events := []models.CfgEvent{}
	eventGeneric := mcuGeneric.AlarmList
	for _, evt := range eventGeneric {
		event := models.CfgEvent{
			McuEventId:               evt.EventId,
			McuEventType:             int32(evt.EventType),
			McuEventName:             evt.Name,
			McuEventSensorId:         evt.SensorId,
			McuEventSensorState:      int32(evt.SensorState),
			McuEventActiveMode:       int32(evt.Mode),
			McuEventActiveAutoTime:   int32(evt.ActiveTime),
			McuEventInactiveAutoTime: int32(evt.InactiveTime),
			McuEventActiveAutoDays:   int32(evt.AutoDays),
			McuEventPlayFile:         evt.PlayFile,
		}
		events = append(events, event)
	}

	mcuConfig := &models.McuCfg{
		McuId:              mcuId,
		McuTelNumber:       mcuGeneric.PhoneNumber,
		McuVolume:          int32(mcuGeneric.Volume),
		McuIp:              mcuGeneric.LocalIp,
		McuFirmwareVersion: mcuGeneric.FirmwareVersion,
		McuFMVolume:        mcuGeneric.FMVolume,
		McuFMAuto:          mcuGeneric.FMAuto,
		McuTxType:          mcuGeneric.TxType,
		CmsGroupIdList:     mcuGeneric.GroupList,
		McuCfgCameraList:   cameras,
		McuCfgSensorList:   sensors,
		McuCfgEventList:    events,
	}

	client := &http.Client{}
	buf, err := json.Marshal(mcuConfig)
	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPut, serverBasePath+"/mcu-update-config", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsUpdateCameraInfo(mcuId int64, opuCameras []base.OPUCamera) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsUpdateCameraInfo err: ", err)
		}
	}()
	cameras := []models.McuCamera{}
	for _, cam := range opuCameras {
		camera := models.McuCamera{
			McuCameraId:       cam.CameraId,
			McuCameraName:     cam.CameraName,
			McuCameraTypeId:   cam.CameraTypeId,
			McuCameraLocalIp:  cam.CameraLocalIp,
			McuCameraHttpPort: cam.CameraHttpPort,
			McuCameraRtspPort: cam.CameraRtspPort,
			McuCameraUsername: cam.CameraUsername,
			McuCameraPassword: cam.CameraPassword,
		}
		cameras = append(cameras, camera)
	}
	body := &models.McuCameraUpdate{
		McuId:         mcuId,
		McuCameraList: cameras,
	}
	client := &http.Client{}
	buf, err := json.Marshal(body)
	if err != nil {
		return false, "", err
	}

	req, err := http.NewRequest(http.MethodPut, serverBasePath+"/update-camera-info", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsSensorUpdateStatus(mcuId int64, isLog bool, opuSensors []base.OPUSensor) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsSensorUpdateStatus err: ", err)
		}
	}()
	statuss := []models.Sensor{}
	for _, sensor := range opuSensors {
		enable := true
		if sensor.Enable != 1 {
			enable = false
		}
		status := models.Sensor{
			McuSensorId:         sensor.SensorId,
			McuSensorType:       int32(sensor.Type),
			McuSensorName:       sensor.Name,
			McuSensorEnable:     enable,
			McuSensorState:      int32(sensor.State),
			McuSensorData:       sensor.Value,
			McuSensorUpdateTime: sensor.CreatedTime,
		}
		statuss = append(statuss, status)
	}
	body := &models.McuSensor{
		McuId:         mcuId,
		IsLog:         isLog,
		McuSensorList: statuss,
	}
	client := &http.Client{}
	buf, err := json.Marshal(body)

	if err != nil {
		return false, "", err
	}

	req, err := http.NewRequest(http.MethodPost, serverBasePath+"/update-sensor-status", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsCreateLogHandle(mcuId int64, log base.OPULog) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsCreateLogHandle err: ", err)
		}
	}()
	logType := log.LogType
	if logType == base.LOG_RECORD_CREATE || logType == base.LOG_RECORD_PLAY || logType == base.LOG_RECORD_EDIT || logType == base.LOG_RECORD_DELETE || logType == base.LOG_RECORD_DOWNLOAD_ERR || logType == base.LOG_FILE_ERROR || logType == base.LOG_STOP_RECORD || logType == base.LOG_ADD_FM_TO_PLAY_LIST || logType == base.LOG_AUTO_PLAY_FM || logType == base.LOG_STOP_RECORD_EXPIRE || logType == base.LOG_LIVE_STREAM_START || logType == base.LOG_LIVE_STREAM_END {
		record := models.RecordUpdate{
			RecordId:       log.MediaId,
			McuId:          mcuId,
			LogType:        int32(logType),
			LogTime:        log.CreatedTime,
			PlayReason:     0,
			PlayReasonTime: -1,
		}
		if logType == base.LOG_RECORD_PLAY {
			record.PlayReason = int32(log.State)
			if log.State == 1 { //schedule
				record.PlayReasonTime = int32(log.Value)
			}
		}
		success, message, err := CmsUpdateRecord(record)
		return success, message, err
	} else if logType == base.LOG_SENSOR_STATUS_CHANGE {
		sensors := []base.OPUSensor{
			base.OPUSensor{
				SensorId:    log.SensorId,
				State:       log.State,
				Value:       log.Value,
				CreatedTime: log.CreatedTime,
			},
		}
		//call cms api
		return CmsSensorUpdateStatus(mcuId, true, sensors)
	} else if logType == base.LOG_NEW_EVENT {
		events := []base.OPUAlarm{
			base.OPUAlarm{
				EventId:     log.EventId,
				SensorId:    log.SensorId,
				SensorState: log.State,
				OccurTime:   log.CreatedTime,
			},
		}
		return CmsPushNotification(mcuId, true, events)
	} else if logType == base.LOG_MGW_CONNECTED || logType == base.LOG_MGW_DISCONNECTED || logType == base.LOG_MGW_REBOOT {
		return CmsCreateMcuStat(mcuId, int32(logType), int32(log.Value), log.CreatedTime)
		//return CmsUpdateMcuConnectStatus(mcuId, logType, "", log.CreatedTime, int32(log.Value))
	} else if logType == base.LOG_MQTT_PACKAGE {
		return CmsCreateMcuMqttMessage(mcuId, int32(log.State), log.CreatedTime)
	} else if logType == base.LOG_TCP_STACK_ERROR {
		return CmsCreateMcuStat(mcuId, int32(logType), 0, log.CreatedTime)
	}
	return false, "", nil
}

func CmsPushNotification(mcuId int64, isLog bool, opuAlarms []base.OPUAlarm) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsPushNotification err: ", err)
		}
	}()
	notifications := []models.Event{}
	for _, al := range opuAlarms {
		alarm := models.Event{
			McuEventId:               al.EventId,
			McuEventType:             int32(al.EventType),
			McuEventName:             al.Name,
			McuEventSensorId:         al.SensorId,
			McuEventState:            int32(al.State),
			McuEventSensorState:      int32(al.SensorState),
			McuEventActiveMode:       int32(al.Mode),
			McuEventActiveAutoTime:   int32(al.ActiveTime),
			McuEventInactiveAutoTime: int32(al.InactiveTime),
			McuEventActiveAutoDays:   int32(al.AutoDays),
			McuEventPlayFile:         al.PlayFile,
			McuEventUpdateTime:       al.OccurTime,
		}
		notifications = append(notifications, alarm)
	}
	body := &models.McuNotification{
		McuId:        mcuId,
		IsLog:        isLog,
		McuEventList: notifications,
	}
	client := &http.Client{}
	buf, err := json.Marshal(body)

	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPost, serverBasePath+"/alarm-push-notification", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsUpdateMcuConnectStatus(mcuId int64, status byte, ip string, ts int64, value int32) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsUpdateMcuConnectStatus err: ", err)
		}
	}()
	body := &models.McuConnStatus{
		McuId:     mcuId,
		McuStatus: int32(status),
		McuIp:     ip,
		McuTs:     ts,
	}
	client := &http.Client{}
	buf, err := json.Marshal(body)
	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPut, serverBasePath+"/mcu-update-connect-status", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsCreateMcuStat(mcuId int64, statType int32, statValue int32, statTime int64) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsCreateMcuStat err: ", err)
		}
	}()
	body := &models.McuStat{
		McuId:        mcuId,
		McuStatType:  statType,
		McuStatValue: statValue,
		McuStatTime:  statTime,
	}
	client := &http.Client{}
	buf, err := json.Marshal(body)
	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPost, serverBasePath+"/mcu-stat", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsPostOperatingState(mcuId int64, opuStatus base.OPUStatus) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsPostOperatingState err: ", err)
		}
	}()
	body := &models.McuOperatingState{
		McuId:             mcuId,
		McuTxType:         int32(opuStatus.TxType),
		McuMobileCsq:      int32(opuStatus.MCsq),
		McuWifiCsq:        int32(opuStatus.WiCsq),
		McuSpeakerErr:     int32(opuStatus.SpeakerErr),
		McuSpeakerCurrent: float32(opuStatus.SpeakerSta / 10),
		McuTemp:           float32(opuStatus.Temp / 10),
		McuFMStatus:       int32(opuStatus.FMStatus),
	}
	client := &http.Client{}
	buf, err := json.Marshal(body)
	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPost, serverBasePath+"/post-operating-state", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func CmsCreateMcuMqttMessage(mcuId int64, opcode int32, time int64) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsCreateMcuMqttMessage err: ", err)
		}
	}()
	body := &models.McuMqttOpcode{
		McuId:         mcuId,
		McuMqttOpcode: opcode,
		McuMqttTime:   time,
	}
	client := &http.Client{}
	buf, err := json.Marshal(body)
	if err != nil {
		return false, "", err
	}
	req, err := http.NewRequest(http.MethodPost, serverBasePath+"/mcu-mqtt-message", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				return true, "", nil
			}
			data, _ := ioutil.ReadAll(res.Body)
			return false, string(data), nil
		}
	} else {
		return false, "", err
	}
	return false, "", nil
}

func Test() {
	client := &http.Client{}
	url := serverBasePath + "/testWS"
	fmt.Printf("Do request: %s", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf(" err: %v", err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf(" err: %v", err)
		return
	}
	defer res.Body.Close()
	if res != nil {
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Printf(" ----> result: %s", string(data))
	}
}

func Test1() {
	ids, _ := redis.GetAllMcuIds()
	fmt.Println(ids)
}
