package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tttm_cms_api/lp-libs/base"
	"tttm_cms_api/lp-libs/models"
	"tttm_cms_api/lp-libs/redis"
	"tttm_cms_api/lp-libs/settings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/glog"
)

var mqttClient MQTT.Client
var MqttClientPushLog MQTT.Client
var mqttInfo *settings.MqttInfo = settings.GetMqttInfo()
var nodes []models.NodeDetail

var topicLogMainflux string = settings.GetTopicDeviceMainflux()
var onMessage MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	glog.Infof("Topic: %s, Message: %s\n", msg.Topic(), msg.Payload())
	obj := &base.MqttMessagePublish{}
	err := json.Unmarshal(msg.Payload(), obj)

	var MCUID int64 = obj.Id
	if err != nil {
		glog.Error("onMessage/json.Unmarshal err: ", err)
		glog.Error("Raw message: ", msg.Payload())
		return
	}
	switch obj.OpCode {
	case base.OPU_GENERIC:
		generic := &base.OPUGeneric{}
		err := json.Unmarshal(obj.Data, generic)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		//TOTO: save to redis
		err = redis.UpdateMcuInfo(MCUID, generic)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/redis.UpdateMcuInfo err: %v", MCUID, err)
			return
		}
		//TODO: req CMS API
		success, message, err := CmsUpdateMcuConfig(MCUID, generic)
		if err != nil {
			glog.Errorf("onMesses/OPU_GENERIC/%d/CsmUpdateMcuConfig err: %v", MCUID, err)
			return
		}
		if success {
			glog.V(8).Infof("MCU %d PUT OPU_GENERIC to CMS ---> success", MCUID)
		} else {
			glog.V(8).Infof("MCU %d PUT OPU_GENERIC to CMS ---> failure (%s)", MCUID, message)
		}
		//TODO: convert to SenML, insert InfluxDB;
		fmt.Println("Ver2")
		_, err = ConvertJsonToSenMLVer2(MCUID, generic, base.OPU_GENERIC, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", MCUID, err)
			return
		}

		success, message, err = collectLogMCU(MCUID, generic)
		if err != nil {
			glog.Errorf("api/v1/tttm/collect-log err: %v", MCUID, err)
			return
		}
		if success {
			glog.V(8).Infof("api/v1/tttm/collect-log ---> success", MCUID)
		} else {
			glog.V(8).Infof("api/v1/tttm/collect-log ---> failure (%s)", MCUID, message)
		}
		break
	case base.OPU_CAMERA:
		cameras := []base.OPUCamera{}
		err := json.Unmarshal(obj.Data, &cameras)
		if err != nil {
			glog.Errorf("onMessage/OPU_CAMERA/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		//TOTO: save to redis
		err = redis.UpdateMcuCameraList(MCUID, cameras)
		if err != nil {
			glog.Errorf("onMessage/OPU_CAMERA/%d/redis.UpdateMcuCameraList err: %v", MCUID, err)
			return
		}
		//TODO: req CMS API
		success, message, err := CmsUpdateCameraInfo(MCUID, cameras)
		if err != nil {
			glog.Errorf("onMessage/OPU_CAMERA/%d/CmsUpdateCameraInfo err: %v", MCUID, err)
			return
		} else {
			if success {
				glog.V(8).Infof("MCU %d PUT OPU_CAMERA to CMS ---> success", MCUID)
			} else {
				glog.V(8).Infof("MCU %d PUT OPU_CAMERA to CMS ---> failure (%s)", MCUID, message)
			}
		}
		//TODO: convert to SenML, insert InfluxDB
		_, err = ConvertJsonToSenMLVer2(MCUID, cameras, base.OPU_CAMERA, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_PHONE:
		phone := ""
		err := json.Unmarshal(obj.Data, &phone)
		if err != nil {
			glog.Errorf("onMessage/OPU_PHONE/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		//TOTO: save to redis
		err = redis.UpdateMcuPhone(MCUID, phone)
		if err != nil {
			glog.Errorf("onMessage/OPU_PHONE/%d/redis.UpdateMcuPhone err: %v", MCUID, err)
			return
		}
		//TODO: req CMS API
		success, message, err := CmsUpdateMcuTelInfo(MCUID, phone)
		if err != nil {
			glog.Errorf("onMessage/OPU_PHONE/%d/CmsUpdateMcuTelInfo err: %v", MCUID, err)
			return
		} else {
			if success {
				glog.V(8).Infof("MCU %d PUT OPU_PHONE to CMS ---> success", MCUID)
			} else {
				glog.V(8).Infof("MCU %d PUT OPU_PHONE to CMS ---> failure (%s)", MCUID, message)
			}
		}
		//TODO: convert to SenML, insert InfluxDB
		_, err = ConvertJsonToSenMLVer2(MCUID, phone, base.OPU_PHONE, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_SENSOR:
		sensors := []base.OPUSensor{}
		err := json.Unmarshal(obj.Data, &sensors)
		if err != nil {
			glog.Errorf("onMessage/OPU_SENSOR/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		//TODO: req CMS API
		success, message, err := CmsSensorUpdateStatus(MCUID, false, sensors)
		if err != nil {
			glog.Errorf("onMessage/OPU_SENSOR/%d/CmsSensorUpdateStatus err: %v", MCUID, err)
			return
		} else {
			if success {
				glog.V(8).Infof("MCU %d PUT OPU_SENSOR to CMS ---> success", MCUID)
			} else {
				glog.V(8).Infof("MCU %d PUT OPU_SENSOR to CMS ---> failure (%s)", MCUID, message)
			}
		}
		//TODO: convert to SenML, insert InfluxDB
		_, err = ConvertJsonToSenMLVer2(MCUID, sensors, base.OPU_SENSOR, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_ALARM:
		alarms := []base.OPUAlarm{}
		err := json.Unmarshal(obj.Data, &alarms)
		if err != nil {
			glog.Errorf("onMessage/OPU_ALARM/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		//TODO: req CMS API
		success, message, err := CmsPushNotification(MCUID, false, alarms)
		if err != nil {
			glog.Errorf("onMessage/OPU_ALARM/%d/CmsPushNotification err: %v", MCUID, err)
			return
		} else {
			if success {
				glog.V(8).Infof("MCU %d PUT OPU_ALARM to CMS ---> success", MCUID)
			} else {
				glog.V(8).Infof("MCU %d PUT OPU_ALARM to CMS ---> failure (%s)", MCUID, message)
			}
		}
		//TODO: convert to SenML, insert InfluxDB
		_, err = ConvertJsonToSenMLVer2(MCUID, alarms, base.OPU_ALARM, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_STATUS:
		status := base.OPUStatus{}
		err := json.Unmarshal(obj.Data, &status)
		if err != nil {
			glog.Errorf("onMessage/OPU_STATUS/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		//TOTO: save to redis
		err = redis.UpdateMcuStatus(MCUID, &status)
		if err != nil {
			glog.Errorf("onMessage/OPU_STATUS/%d/redis.UpdateMcuStatus err: %v", MCUID, err)
			return
		}
		//TODO: req CMS API
		success, message, err := CmsPostOperatingState(MCUID, status)
		if err != nil {
			glog.Errorf("onMessage/OPU_STATUS%d/%d/CmsPostOperatingState err: %v", MCUID, err)
			return
		} else {
			if success {
				glog.V(8).Infof("MCU %d POST OPU_STATUS to CMS ---> success", MCUID)
			} else {
				glog.V(8).Infof("MCU %d POST OPU_STATUS to CMS ---> failure (%s)", MCUID, message)
			}
		}
		//TODO: convert to SenML, insert InfluxDB
		_, err = ConvertJsonToSenMLVer2(MCUID, status, base.OPU_STATUS, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_MEDIA:
		medias := [][]int64{}
		err := json.Unmarshal(obj.Data, &medias)
		if err != nil {
			glog.Errorf("onMessage/OPU_MEDIA/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		items := models.ConvertOPUMediaToStruct(medias, MCUID)
		success, message, err := CmsRecordStat(items)
		if err != nil {
			glog.Errorf("onMessage/OPU_MEDIA%d/%d/CmsRecordStat err: %v", MCUID, err)
			return
		} else {
			if success {
				glog.V(8).Infof("MCU %d POST OPU_MEDIA to CMS ---> success", MCUID)
			} else {
				glog.V(8).Infof("MCU %d POST OPU_MEDIA to CMS ---> failure (%s)", MCUID, message)
			}
		}
		//TODO: convert to SenML, insert InfluxDB
		_, err = ConvertJsonToSenMLVer2(MCUID, items, base.OPU_MEDIA, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_LOG:

		logs := []base.OPULog{}
		err := json.Unmarshal(obj.Data, &logs)
		if err != nil {
			glog.Errorf("onMessage/OPU_LOG/%d/json.Unmarshal err: %v", MCUID, err)
			return
		}
		//TODO: save to redis
		var logType string
		for _, log := range logs {
			switch log.LogType {
			case base.LOG_RECORD_CREATE:
				logType = "LOG_RECORD_CREATE"
			case base.LOG_RECORD_EDIT:
				logType = "LOG_RECORD_EDIT"
			case base.LOG_RECORD_PLAY:
				logType = "LOG_RECORD_PLAY"
			case base.LOG_RECORD_DELETE:
				logType = "LOG_RECORD_DELETE"
			case base.LOG_NEW_EVENT:
				logType = "LOG_NEW_EVENT"
			case base.LOG_SENSOR_STATUS_CHANGE:
				logType = "LOG_SENSOR_STATUS_CHANGE"
			case base.LOG_RECORD_DOWNLOAD_ERR:
				logType = "LOG_RECORD_DOWNLOAD_ERR"
			case base.LOG_MGW_REBOOT:
				logType = "LOG_MGW_REBOOT"
			case base.LOG_MGW_CONNECTED:
				logType = "LOG_MGW_CONNECTED"
			case base.LOG_MGW_DISCONNECTED:
				logType = "LOG_MGW_DISCONNECTED"
			case base.LOG_FILE_ERROR:
				logType = "LOG_FILE_ERROR"
			case base.LOG_STOP_RECORD:
				logType = "LOG_STOP_RECORD"
			case base.LOG_ADD_FM_TO_PLAY_LIST:
				logType = "LOG_ADD_FM_TO_PLAY_LIST"
			case base.LOG_AUTO_PLAY_FM:
				logType = "LOG_AUTO_PLAY_FM"
			case base.LOG_STOP_RECORD_EXPIRE:
				logType = "LOG_STOP_RECORD_EXPIRE"
			case base.LOG_MQTT_PACKAGE:
				logType = "LOG_MQTT_PACKAGE"
			case base.LOG_TCP_STACK_ERROR:
				logType = "LOG_TCP_STACK_ERROR"
			case base.LOG_LIVE_STREAM_START:
				logType = "LOG_LIVE_STREAM_START"
			case base.LOG_LIVE_STREAM_END:
				logType = "LOG_LIVE_STREAM_END"
			default:
				logType = "LOG_UNKNOWN"
			}

			success, message, err := CmsCreateLogHandle(MCUID, log)
			if err != nil {
				glog.Error("onMessage/OPU_LOG/%d/CmsCreateLogHandle err: %v", MCUID, err)
				return
			}
			if success {
				glog.V(8).Infof("MCU %d OPU_LOG/%s logID %d ---> success", MCUID, logType, log.LogId)
			} else {
				glog.V(8).Infof("MCU %d OPU_LOG/%s logID %d ---> failure (%s)", MCUID, logType, log.LogId, message)
			}
		}
		//TODO: convert to SenML, insert InfluxDB
		_, err = ConvertJsonToSenMLVer2(MCUID, logs, base.OPU_LOG, topicLogMainflux)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/convertJsonToSenML: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_TIME_GET:
		err := publishTimeMessage(MCUID)
		if err != nil {
			glog.Error("onMessage/OPU_TIME_GET/%d/publishTimeMessage err: %v", MCUID, err)
		} else {
			glog.V(8).Infof("publish OPH_TIME to %d ---> success", MCUID)
		}
		break
	}
}

var onMessageClientConnection MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	connection := &base.MqttClientConnection{}
	err := json.Unmarshal(msg.Payload(), connection)
	if err != nil {
		glog.Error("onMessageClientConnection/json.Unmarshal err: ", err)
		return
	}

	if inorgeLPClient(connection.ClientId, connection.UserName) {
		return
	}
	mcuId, err := strconv.ParseInt(connection.ClientId, 10, 64)
	if err != nil {
		glog.V(16).Infof("onMessageClientConnection/%s/Client is not m-gateway", msg.Topic())
		return
	}
	if strings.Contains(msg.Topic(), "disconnected") {
		//disconnected
		//TOTO: save to redis
		err = redis.UpdateMcuConnectionStatus(mcuId, base.MCU_DISCONNECTED, connection.Ts, connection.IpAddress)
		if err != nil {
			glog.Errorf("onMessageClientConnection/%d/redis.UpdateConnectionStatus err: %v", mcuId, err)
			return
		}
		//TODO: call api update connection status
		success, message, err := CmsUpdateMcuConnectStatus(mcuId, base.MCU_DISCONNECTED, connection.IpAddress, connection.Ts, 0)
		if err != nil {
			glog.Errorf("onMessageClientConnection/disconnected/%d/CmsUpdateMcuConnectStatus err: %v", mcuId, err)
			return
		} else {
			if success {
				glog.Infof("PUT %d disconnected status to CMS ---> success", mcuId)
			} else {
				glog.Infof("PUT %d disconnected status to CMS ---> failure (%s)", mcuId, message)
			}
		}

	} else {
		//TOTO: save to redis
		err = redis.UpdateMcuConnectionStatus(mcuId, base.MCU_CONNECTED, connection.Ts, connection.IpAddress)
		if err != nil {
			glog.Errorf("onMessageClientConnection/%d/redis.UpdateConnectionStatus err: %v", mcuId, err)
			return
		}
		//TODO: call api update connection status
		success, message, err := CmsUpdateMcuConnectStatus(mcuId, base.MCU_CONNECTED, connection.IpAddress, connection.Ts, 0)
		if err != nil {
			glog.Errorf("onMessageClientConnection/connected/%d/CmsUpdateMcuConnectStatus err: %v", mcuId, err)
			return
		} else {
			if success {
				glog.Infof("PUT %d connected status to CMS ---> success", mcuId)
			} else {
				glog.Infof("PUT %d connected status to CMS ---> failure (%s)", mcuId, message)
			}
		}
	}
}

func SetSuperUser() error {
	username := base.LP_CMS_CLIENT
	password := base.LP_CMS_CLIENT
	if len(mqttInfo.UserName) > 0 {
		username = mqttInfo.UserName
		password = mqttInfo.Password
	}
	return redis.SetAclSuperUser(username, password)
}

func ConnectMqtt() error {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("-------------RECOVER err: ", err)
		}
	}()
	ip, err := resolveHostIp()
	if err != nil {
		return err
	}
	server := fmt.Sprintf("tcp://%s:%d", mqttInfo.ServerAddress, mqttInfo.ServerPort)
	opts := MQTT.NewClientOptions().AddBroker(server)
	opts.SetClientID(fmt.Sprintf("%s@%s", base.LP_CMS_CLIENT, ip))
	opts.SetCleanSession(false)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(onConnectionLost)
	opts.SetOnConnectHandler(onConnected)
	mqttClient = MQTT.NewClient(opts)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
func ConnectMQTTOpts() error {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("-------------RECOVER err: ", err)
		}
	}()
	ip, err := resolveHostIp()
	if err != nil {
		return err
	}
	server := fmt.Sprintf("tcp://%s:%d", mqttInfo.ServerAddress, mqttInfo.ServerPortAuth)
	opts := MQTT.NewClientOptions().AddBroker(server)
	opts.SetClientID(fmt.Sprintf("%s@%s", "ClientPushSenMLCMS", ip))
	//opts.SetDefaultPublishHandler(f)
	opts.SetUsername(settings.GetThingAuthPush().ID)
	opts.SetPassword(settings.GetThingAuthPush().Key)
	opts.SetCleanSession(false)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetAutoReconnect(true)
	MqttClientPushLog = MQTT.NewClient(opts)
	if token := MqttClientPushLog.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
func DisconnectMqtt() {
	mqttClient.Unsubscribe("h/s/#")
	//if len(nodes) > 0 {
	//	for _, node := range nodes {
	//		connectionTopic := fmt.Sprintf("$SYS/brokers/%s/clients/#", node.Name)
	//		mqttClient.Unsubscribe(connectionTopic)
	//	}
	//}
	mqttClient.Disconnect(250)
}
func publishMessageSenML(topic string, qos byte, retain bool, msgSenML string) error {
	token := MqttClientPushLog.Publish(topic, qos, retain, msgSenML)
	return token.Error()
}

var onConnectionLost MQTT.ConnectionLostHandler = func(client MQTT.Client, reason error) {
	glog.Infof("onConnectionLost/server(%s:%d)", mqttInfo.ServerAddress, mqttInfo.ServerPort)
}

var onConnected MQTT.OnConnectHandler = func(client MQTT.Client) {
	CheckMqttStatusClients()
	//Thiết bị sẽ tự phải pub msg vào topic h2/s/# để cms có thấy lấy dữ liệu như mCuId, VCode,..v để lưu vào Redis cho MCU dùng
	if token := client.Subscribe("h2/s/#", 1, onMessage); token.Wait() && token.Error() != nil {
		glog.Error("onConnected/client.Subscribe(h2/s/#) err: ", token.Error())
		return
	}
	//glog.Infof("onConnected/client.Subscribe((h/s/#) --->success")
	glog.Infof("onConnected/client.Subscribe(h2/s/#) --->success")
	//fmt.Println(<-c)
}

var mcuTopicFormat string = "h/d/%d"

func publishTimeMessage(mcuId int64) error {
	//TODO: publish to broker
	timeJson, _ := json.Marshal(time.Now().Unix())
	msg := &base.MqttMessagePublish{
		Id:     mcuId,
		OpCode: base.OPH_TIME,
		Data:   json.RawMessage(timeJson),
	}
	topic := fmt.Sprintf(mcuTopicFormat, mcuId)
	return publishMessage(topic, 0, false, msg)
}

func publishMessage(topic string, qos byte, retain bool, msg *base.MqttMessagePublish) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	} else {
		token := mqttClient.Publish(topic, qos, retain, payload)
		return token.Error()
	}
}

func getBrokerNodes(path string) ([]models.NodeDetail, error) {

	client := &http.Client{}
	url := fmt.Sprintf("http://%s@%s:%d/%s", settings.GetKeyAuth(), mqttInfo.ServerAddressHTTP, mqttInfo.HttpApiPort, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []models.NodeDetail{}, err
	}
	//Không cần Auth vì ở trên đường linh đã có Key Token giúp có thể Auth
	//req.SetBasicAuth(mqttInfo.HttpApiUserName, mqttInfo.HttpApiPassword)

	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}

	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {

				data, _ := ioutil.ReadAll(res.Body)
				nodes := models.NodeVenmqttd{}
				err := json.Unmarshal(data, &nodes)
				if err != nil {
					return nil, err
				}
				return nodes.Nodes, nil
			}
			return []models.NodeDetail{}, err
		}
	} else {
		return []models.NodeDetail{}, err
	}
	return []models.NodeDetail{}, nil
}

func getBrokerNodesV2(path string) ([]models.NodeEmqttd, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%d/%s", mqttInfo.ServerAddress, mqttInfo.HttpApiPort, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []models.NodeEmqttd{}, err
	}
	req.SetBasicAuth(mqttInfo.HttpApiUserName, mqttInfo.HttpApiPassword)

	res, err := client.Do(req)
	defer res.Body.Close()

	if err == nil {
		if res != nil {
			if res.StatusCode == http.StatusOK {
				data, _ := ioutil.ReadAll(res.Body)
				result := models.V2Nodes{}
				err := json.Unmarshal(data, &result)
				if err != nil {
					return nil, err
				}
				return result.Result, nil
			}
			return []models.NodeEmqttd{}, err
		}
	} else {
		return []models.NodeEmqttd{}, err
	}
	return []models.NodeEmqttd{}, nil
}

func CheckMqttStatusClients() {
	ids, err := redis.GetAllMcuIds()
	if err != nil {
		glog.Error("checkMqttStatusClients/redis.GetAllMcuIds err: ", err)
		return
	}

	for _, id := range ids {

		//fmt.Println(id)
		mcuId, err := strconv.ParseInt(id, 10, 64)
		//fmt.Println("MCUID: ",mcuId)
		if err != nil {
			continue
		}
		if id == "88171961786838220"{
			fmt.Println("QUANG: ",mcuId)
		}
		path := fmt.Sprintf("api/v1/session/show/?--client_id=%d", id)
		conns, err := getMqttStatusClients(path)
		if err != nil {
			glog.Error("checkMqttStatusClients/getMqttStatusClients err: ", err)
			return
		}

		isOnl := false
		if len(conns) != 0 {
			for _, conn := range conns {
				if id == "88171961786836606" && conn.ClientId == "88171961786836606" && conn.IsOnline == false {
					fmt.Println("===sonnh===id:"+id+" conn.ClientId"+conn.ClientId+" conn.IsOnline:"+strconv.FormatBool(conn.IsOnline))
					glog.Error("===sonnh===id:"+id+" conn.ClientId"+conn.ClientId+" conn.IsOnline:"+strconv.FormatBool(conn.IsOnline))
				}

				if id == conn.ClientId && conn.IsOnline == true &&  len(conn.PeerHost) > 0 {
					isOnl = true
					//TODO: call api update connection status
					//layout := "2006-01-02 15:04:05"
					//t, err := time.Parse(layout, conn.ConnectedAt)
					clientts, err := getMQTTClientts("api/v1/session/show?--session_started_at&--queue_started_at&--client_id=" + conn.ClientId)
					if err != nil {
						glog.Error("Fail get status client -- CheckMqttStatusClients/ getMQTTClientts")
						return
					}
					if len(clientts) > 0 {
						var t int64 = clientts[0].Session_started_at / 1000
						//if err != nil {
						//	continue
						//}

						success, message, err := CmsUpdateMcuConnectStatus(mcuId, base.MCU_CONNECTED, conn.PeerHost, t, 0)
						if err != nil {
							glog.Errorf("checkMqttStatusClients/connected/%d/CmsUpdateMcuConnectStatus err: %v", mcuId, err)
							continue
						} else {
							if success {
								fmt.Println("CONNECT: ",strconv.FormatInt(mcuId,10))
								ConvertJsonToSenMLVer2(mcuId,nil,base.STATE_DEVICE_CONNECTED,settings.GetTopicDeviceMainflux())
								glog.Infof("PUT %d connected status to CMS ---> success", mcuId)
							} else {
								ConvertJsonToSenMLVer2(mcuId,nil,base.STATE_DEVICE_CONNECTED,settings.GetTopicDeviceMainflux())
								glog.Infof("PUT %d connected status to CMS ---> failure (%s)", mcuId, message)
							}
						}
						//TOTO: save to redis
						err = redis.UpdateMcuConnectionStatus(mcuId, base.MCU_CONNECTED, t, conn.PeerHost)
						if err != nil {
							glog.Errorf("checkMqttStatusClients/%d/redis.UpdateConnectionStatus err: %v", mcuId, err)
							continue
						}
					}

				}
			}
		}

		if !isOnl {

			//TODO: call api update connection status
			client := redis.GetMcuInfo(mcuId)
			if client != nil {

				success, message, err := CmsUpdateMcuConnectStatus(mcuId, base.MCU_DISCONNECTED, "", client.ConnTime, 0)
				if err != nil {
					glog.Errorf("onMessageClientConnection/disconnected/%d/CmsUpdateMcuConnectStatus err: %v", mcuId, err)
					return
				} else {
					if success {

						ConvertJsonToSenMLVer2(mcuId,nil,base.STATE_DEVICE_DISCONNECTED,settings.GetTopicDeviceMainflux())
						glog.Infof("PUT %d disconnected status to CMS ---> success", mcuId)
					} else {
						ConvertJsonToSenMLVer2(mcuId,nil,base.STATE_DEVICE_DISCONNECTED,settings.GetTopicDeviceMainflux())
						glog.Infof("PUT %d disconnected status to CMS ---> failure (%s)", mcuId, message)
					}
				}
				//TOTO: save to redis
				err = redis.UpdateMcuConnectionStatus(mcuId, base.MCU_DISCONNECTED, client.ConnTime, "")
				if err != nil {
					glog.Errorf("onMessageClientConnection/%d/redis.UpdateConnectionStatus err: %v", mcuId, err)
					break
				}
			}
		}
	}
}

func getMqttStatusClients(path string) ([]models.MqttClientStatus, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s@%s:%d/%s", settings.GetKeyAuth(), mqttInfo.ServerAddressHTTP, mqttInfo.HttpApiPort, path)

	//url := fmt.Sprintf("http://%s:%d/%s", mqttInfo.ServerAddress, mqttInfo.HttpApiPort, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []models.MqttClientStatus{}, err
	}
	//req.SetBasicAuth(mqttInfo.HttpApiUserName, mqttInfo.HttpApiPassword)
	//fmt.Println(mqttInfo.HttpApiUserName, mqttInfo.HttpApiPassword)
	res, err := client.Do(req)
	//If err returned -> auto close - dont need to close
	if res != nil {
		defer res.Body.Close()
	}
	if err == nil {
		if res != nil {
			//fmt.Println("IN")
			if res.StatusCode == http.StatusOK {
				data, _ := ioutil.ReadAll(res.Body)
				result := models.MQTTClientList{}
				err := json.Unmarshal(data, &result)
				if err != nil {
					return nil, err
				}
				return result.MQTTClient, nil
			}
			return []models.MqttClientStatus{}, err
		}
	} else {
		//fmt.Println("IN2")
		return []models.MqttClientStatus{}, err
	}
	return []models.MqttClientStatus{}, nil
}

func getMQTTClientts(path string) ([]models.MQTTClientTs, error) {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s@%s:%d/%s", settings.GetKeyAuth(), mqttInfo.ServerAddressHTTP, mqttInfo.HttpApiPort, path)

	//url := fmt.Sprintf("http://%s:%d/%s", mqttInfo.ServerAddress, mqttInfo.HttpApiPort, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []models.MQTTClientTs{}, err
	}
	//req.SetBasicAuth(mqttInfo.HttpApiUserName, mqttInfo.HttpApiPassword)
	//fmt.Println(mqttInfo.HttpApiUserName, mqttInfo.HttpApiPassword)
	res, err := client.Do(req)
	//If err returned -> auto close - dont need to close
	if res != nil {
		defer res.Body.Close()
	}
	if err == nil {
		if res != nil {
			//fmt.Println("IN")
			if res.StatusCode == http.StatusOK {

				data, _ := ioutil.ReadAll(res.Body)
				result := models.MQTTClientTsList{}
				err := json.Unmarshal(data, &result)
				if err != nil {
					return []models.MQTTClientTs{}, err
				}
				return result.MQTTClientts, nil
			}
			return []models.MQTTClientTs{}, err
		}
	} else {
		//fmt.Println("IN2")
		return []models.MQTTClientTs{}, err
	}
	return []models.MQTTClientTs{}, nil
}

func resolveHostIp() (string, error) {
	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip := networkIp.IP.String()
			return ip, nil
		}
	}
	return "", err
}

func inorgeLPClient(clientId, username string) bool {
	return strings.Contains(clientId, base.LP_MCU_API) || strings.Contains(clientId, base.LP_CMS_CLIENT) || username == base.LP_MCU_API || username == base.LP_CMS_CLIENT
}
func updateMcuConnStatus() {
	connections, err := getMqttStatusClients("api/v1/session/show")
	//err := json.Unmarshal(msg.Payload(), connection)
	if err != nil {
		glog.Error("onMessageClientConnection/json.Unmarshal err: ", err)
		return
	}

	for _, conn := range connections {
		if inorgeLPClient(conn.ClientId, conn.User) {
			continue
		}

		clientts, err := getMQTTClientts("api/v1/session/show?--session_started_at&--queue_started_at&--client_id=" + conn.ClientId)
		//Vernemq's ClientId: mqttclientid+timestamp connect mqtt broker (from 2001 to 2286 has 13 digits)
		mcuId, err := strconv.ParseInt(conn.ClientId, 10, 64)
		if err != nil {
			glog.V(16).Infof("onMessageClientConnection/%s/Client is not m-gateway", conn.ClientId)
			return
		}
		var ts int64 = clientts[0].Session_started_at
		if conn.IsOnline == false {
			//disconnected
			//TOTO: save to redis
			err = redis.UpdateMcuConnectionStatus(mcuId, base.MCU_DISCONNECTED, ts, conn.PeerHost)
			if err != nil {
				glog.Errorf("onMessageClientConnection/%d/redis.UpdateConnectionStatus err: %v", mcuId, err)
				return
			}
			//TODO: call api update connection status
			success, message, err := CmsUpdateMcuConnectStatus(mcuId, base.MCU_DISCONNECTED, conn.PeerHost, ts, 0)
			if err != nil {
				glog.Errorf("onMessageClientConnection/disconnected/%d/CmsUpdateMcuConnectStatus err: %v", mcuId, err)
				return
			} else {
				if success {
					glog.Infof("PUT %d disconnected status to CMS ---> success", mcuId)
				} else {
					glog.Infof("PUT %d disconnected status to CMS ---> failure (%s)", mcuId, message)
				}
			}

		} else {
			//TOTO: save to redis
			err = redis.UpdateMcuConnectionStatus(mcuId, base.MCU_CONNECTED, ts, conn.PeerHost)
			if err != nil {
				glog.Errorf("onMessageClientConnection/%d/redis.UpdateConnectionStatus err: %v", mcuId, err)
				return
			}
			//TODO: call api update connection status
			success, message, err := CmsUpdateMcuConnectStatus(mcuId, base.MCU_CONNECTED, conn.PeerHost, ts, 0)
			if err != nil {
				glog.Errorf("onMessageClientConnection/connected/%d/CmsUpdateMcuConnectStatus err: %v", mcuId, err)
				return
			} else {
				if success {
					glog.Infof("PUT %d connected status to CMS ---> success", mcuId)
				} else {
					glog.Infof("PUT %d connected status to CMS ---> failure (%s)", mcuId, message)
				}
			}
		}
	}
}

func collectLogMCU(mcuId int64, mcuGeneric *base.OPUGeneric) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsUpdateMcuConfig err: ", err)
		}
	}()
	var camerasListString string = ""
	camGeneric := mcuGeneric.CameraList
	for _, cam := range camGeneric {
		camerasListString = camerasListString + "," + string(cam.CameraId)
	}
	camerasListStringLen := len(camerasListString)
	if camerasListStringLen > 0 {
		camerasListString = camerasListString[1:camerasListStringLen]
	}

	var sensorsListString string = ""
	sensorGeneric := mcuGeneric.SensorList
	for _, ssr := range sensorGeneric {
		sensorsListString = sensorsListString + "," + string(ssr.SensorId)
	}
	sensorsListStringLen := len(sensorsListString)
	if sensorsListStringLen > 0 {
		sensorsListString = sensorsListString[1:sensorsListStringLen]
	}

	var eventsListString string = ""
	eventGeneric := mcuGeneric.AlarmList
	for _, evt := range eventGeneric {
		eventsListString = eventsListString + "," + string(evt.EventId)
	}
	eventsListStringLen := len(eventsListString)
	if eventsListStringLen > 0 {
		eventsListString = eventsListString[1:eventsListStringLen]
	}

	var groupsListString string = ""
	var groupList []int64 = mcuGeneric.GroupList
	for _, id := range groupList {
		groupsListString = groupsListString + "," + strconv.Itoa(int(id))
	}
	groupsListStringLen := len(groupsListString)
	if groupsListStringLen > 0 {
		groupsListString = groupsListString[1:groupsListStringLen]
	}
	mcuConfig := &models.McuCfgLog{
		McuId:              strconv.Itoa(int(mcuId)),
		McuTelNumber:       mcuGeneric.PhoneNumber,
		McuVolume:          strconv.Itoa(int(int32(mcuGeneric.Volume))),
		McuIp:              mcuGeneric.LocalIp,
		McuFirmwareVersion: mcuGeneric.FirmwareVersion,
		McuFMVolume:        strconv.Itoa(int(mcuGeneric.FMVolume)),
		McuFMAuto:          strconv.Itoa(int(mcuGeneric.FMAuto)),
		McuTxType:          strconv.Itoa(int(mcuGeneric.TxType)),
		CmsGroupIdList:     groupsListString,
		McuCfgCameraList:   camerasListString,
		McuCfgSensorList:   sensorsListString,
		McuCfgEventList:    eventsListString,
	}

	client := &http.Client{}
	buf, err := json.Marshal(mcuConfig)
	if err != nil {
		return false, "", err
	}
	var tttmInfo *settings.TTTMInfo = settings.GetTTTMInfo()
	var serverTTTM string = fmt.Sprintf("http://%s:%d", tttmInfo.TTTTMMainFluxAddress, tttmInfo.Port)
	req, err := http.NewRequest(http.MethodPost, serverTTTM+"/api/v1/tttm/collect-log", bytes.NewBuffer(buf))
	if err != nil {
		return false, "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	if err == nil {
		if res != nil {
			data, _ := ioutil.ReadAll(res.Body)
			responseMSG := new(models.ResponseMSGTTTM)
			json.Unmarshal(data, &responseMSG)
			s, _ := (json.Marshal(responseMSG))
			fmt.Println(string(s))
			if responseMSG.Status == 1 {
				return true, "", nil
			} else {
				msgFail := string(responseMSG.Message)
				switch responseMSG.Code {
				case "SET_GROUP_ERROR_0001":
					return false, msgFail, nil
				case "SYSTEM_ERROR":
					return false, msgFail, nil
				default:
					return false, msgFail, nil
				}
			}

		}
	} else {
		return false, "", err
	}
	return false, "", nil
}
