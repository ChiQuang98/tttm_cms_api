package models

import (
	"time"
)

type McuRes struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`
}

type McuVCode struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Mã vcode
	VCode string `json:"vcode"`
}

type McuGroup struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Danh sách cms_group_id
	CmsGroupIdList []int64 `json:"cms_group_id_list"`
}

type McuVolume struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Giá trị âm lượng
	Volume int32 `json:"volume"`
}

type McuFMCfg struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Giá trị âm lượng FM
	Volume int32 `json:"volume"`

	// Tính năng kiểm tra có thu được tín hiệu FM, 0=disable auto detection, 1=enable auto detection
	Auto int32 `json:"auto"`
}

type McuTelInfo struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Số điện thoại hiện tại của MCU
	McuTelNumber string `json:"tel_number"`
}

type McuCamera struct {

	// ID camera
	McuCameraId int64 `json:"mcu_camera_id"`

	// Tên camera
	McuCameraName string `json:"mcu_camera_name"`

	// Địa chỉ IP local của camera IP
	McuCameraLocalIp string `json:"mcu_camera_local_ip"`

	// Loại camera
	McuCameraTypeId int32 `json:"mcu_camera_type_id"`

	// Cổng http của camera
	McuCameraHttpPort int32 `json:"mcu_camera_http_port"`

	// Cổng rtsp của camera
	McuCameraRtspPort int32 `json:"mcu_camera_rtsp_port"`

	// Thông tin đăng nhập camera
	McuCameraUsername string `json:"mcu_camera_username"`

	// Thông tin password đăng nhập camera
	McuCameraPassword string `json:"mcu_camera_password"`
}

type McuCameraUpdate struct {
	// ID của mcu
	McuId int64 `json:"mcu_id"`

	// Danh sách camera
	McuCameraList []McuCamera `json:"mcu_camera_list"`
}

type McuConnStatus struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Loại log (0 - mất kết nối, 1 - có kết nối, 8 - thiết bị khởi động, 9 - có kết nối với Server, 10- - mất kết nối tới server)
	McuStatus int32 `json:"mcu_status"`

	// Địa chỉ IP Internet hiện tại của MCU
	McuIp string `json:"mcu_ip"`

	// Thời điểm xảy ra log (epoch time)
	McuTs int64 `json:"mcu_ts"`

	// mcu_status = 8 (nguyên nhân khởi động lại [số nguyên 32-bit, ý nghĩa các bit như sau]; B25 - BOR [điện áp yếu], B27 - POR [bật nguồn lên], B28 - SFT [tự khởi động lại do lệnh], B29 - WDG [tự khởi động lại do bị treo CPU]), mcu_status = 9 (số lần thực hiện kết nối)
	McuStatValue int32 `json:"mcu_stat_value"`
}

type McuStat struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Loại log (8 - thiết bị khởi động, 9 - có kết nối với Server)
	McuStatType int32 `json:"mcu_stat_type"`

	// Địa chỉ IP Internet hiện tại của MCU
	McuStatValue int32 `json:"mcu_stat_value"`

	// Thời điểm thiết bị kết nối (epoch time)
	McuStatTime int64 `json:"mcu_stat_time"`
}

type McuMqttOpcode struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Opcode của gói tin nhận được
	McuMqttOpcode int32 `json:"mcu_mqtt_opcode"`

	// Thời điểm nhận gói tin
	McuMqttTime int64 `json:"mcu_mqtt_time"`
}

type Sensor struct {

	// ID sensor
	McuSensorId int64 `json:"mcu_sensor_id"`

	// Loại cảm biến (12 - Cửa, 263 - Chuyển động, 128 - nhiệt độ, 129 - độ ẩm, 160 - Khói)
	McuSensorType int32 `json:"mcu_sensor_type"`

	// Tên cảm biến
	McuSensorName string `json:"mcu_sensor_name"`

	// Kích hoạt hay không kích hoạt cảm biến.
	McuSensorEnable bool `json:"mcu_sensor_enable"`

	// Các ngưỡng thiết lập trạng thái cho cảm biến([20,30] 0->20 nhiệt độ thấp, 20->30 bình thường, >30 cao. Cảm biến DIN = [])
	McuSensorThresholds []int32 `json:"mcu_sensor_thresholds"`

	// Trạng thái hiện tại của cảm biến (DIN [0 - Đóng, 1 - Mở], AIN [0 - Thấp, 1 - Bình thường, 2 - Cao]
	McuSensorState int32 `json:"mcu_sensor_state"`

	// Dữ liệu cảm biến
	McuSensorData int64 `json:"mcu_sensor_data"`

	// Thời điểm sensor cập nhật dữ liệu (epoch time)
	McuSensorUpdateTime int64 `json:"mcu_sensor_update_time"`
}

type CfgSensor struct {
	// ID sensor
	McuSensorId int64 `json:"mcu_sensor_id"`

	// Loại cảm biến (12 - Cửa, 263 - Chuyển động, 128 - nhiệt độ, 129 - độ ẩm, 160 - Khói)
	McuSensorType int32 `json:"mcu_sensor_type"`

	// Tên cảm biến
	McuSensorName string `json:"mcu_sensor_name"`

	// Kích hoạt hay không kích hoạt cảm biến.
	McuSensorEnable bool `json:"mcu_sensor_enable"`

	// Các ngưỡng thiết lập trạng thái cho cảm biến([20,30] 0->20 nhiệt độ thấp, 20->30 bình thường, >30 cao. Cảm biến DIN = [])
	McuSensorThresholds []int32 `json:"mcu_sensor_thresholds"`
}

type McuSensor struct {

	// ID sensor
	McuId int64 `json:"mcu_id"`

	// Xác định xem có phải log của MCU gửi về hay không
	IsLog bool `json:"is_log"`

	// Danh sách trạng thái các cảm biến
	McuSensorList []Sensor `json:"sensor_status_list"`
}

type McuCfgSensor struct {
	// ID của MCU
	McuId int64 `json:"mcu_id"`

	//Danh sách thông tin cấu hình cảm biến
	McuCfgSensorList []CfgSensor `json:"mcu_sensor_list"`
}

type Event struct {

	// ID của sự kiện cảnh báo
	McuEventId int64 `json:"mcu_event_id"`

	// Loại cảnh báo
	McuEventType int32 `json:"mcu_event_type"`

	// Tên cảnh báo
	McuEventName string `json:"mcu_event_name"`

	// ID của cảm biến tạo ra cảnh báo
	McuEventSensorId int64 `json:"mcu_event_sensor_id"`

	//Trạng thái hiện tại của cảnh báo (0 – bình thường, 1 – đang cảnh báo)
	McuEventState int32 `json:"mcu_event_state"`

	// Trạng thái của cảm biến thiết lập cảnh báo
	McuEventSensorState int32 `json:"mcu_event_sensor_state"`

	// Chế độ kích hoạt kịch bản (0 – disable, 1 – auto, 2- always on)
	McuEventActiveMode int32 `json:"mcu_event_active_mode"`

	// Thời gian bắt đầu kích hoạt cảnh báo trong ngày
	McuEventActiveAutoTime int32 `json:"mcu_event_active_auto_time"`

	// Thời gian ngừng kích hoạt cảnh báo trong ngày
	McuEventInactiveAutoTime int32 `json:"mcu_event_inactive_auto_time"`

	// Ngày trong tuần kích hoạt cảnh báo
	McuEventActiveAutoDays int32 `json:"mcu_event_active_auto_days"`

	// Id của file header phát loa khi xảy ra cảnh báo (0 - không phát)
	McuEventPlayFile int64 `json:"mcu_event_play_file"`

	// Thời điểm xảy ra cảnh báo (epoch time)
	McuEventUpdateTime int64 `json:"mcu_event_update_time"`
}

type CfgEvent struct {

	// ID của sự kiện cảnh báo
	McuEventId int64 `json:"mcu_event_id"`

	// Loại cảnh báo
	McuEventType int32 `json:"mcu_event_type"`

	// Tên cảnh báo
	McuEventName string `json:"mcu_event_name"`

	// ID của cảm biến tạo ra cảnh báo
	McuEventSensorId int64 `json:"mcu_event_sensor_id"`

	// Trạng thái của cảm biến thiết lập cảnh báo
	McuEventSensorState int32 `json:"mcu_event_sensor_state"`

	// Chế độ kích hoạt kịch bản (0 – disable, 1 – auto, 2- always on)
	McuEventActiveMode int32 `json:"mcu_event_active_mode"`

	// Thời gian bắt đầu kích hoạt cảnh báo trong ngày
	McuEventActiveAutoTime int32 `json:"mcu_event_active_auto_time"`

	// Thời gian ngừng kích hoạt cảnh báo trong ngày
	McuEventInactiveAutoTime int32 `json:"mcu_event_inactive_auto_time"`

	// Ngày trong tuần kích hoạt cảnh báo
	McuEventActiveAutoDays int32 `json:"mcu_event_active_auto_days"`

	// Id của file header phát loa khi xảy ra cảnh báo (0 - không phát)
	McuEventPlayFile int64 `json:"mcu_event_play_file"`
}

type McuAlarm struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Danh sách thông tin cấu hình cảnh báo
	McuEventList []Event `json:"mcu_event_list"`
}

type McuCfgAlarm struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Danh sách thông tin cấu hình cảnh báo
	McuCfgEventList []CfgEvent `json:"mcu_event_list"`
}

type McuNotification struct {

	// ID của MCU có cảnh báo
	McuId int64 `json:"mcu_id"`

	// Xác định xem có phải log của MCU gửi về hay không
	IsLog bool `json:"is_log"`

	// Danh sách các cảnh báo đang xảy ra
	McuEventList []Event `json:"mcu_event_list"`
}

type Mcu struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Tel num của sim gắn vào MCU
	McuTelNumber string `json:"mcu_tel_number"`

	// Giá trị âm lượng
	McuVolume int32 `json:"mcu_volume"`

	// Địa chỉ IP Internet hiện tại của MCU
	McuIp string `json:"mcu_ip"`

	// Danh sách cms_group_id
	CmsGroupIdList []int64 `json:"cms_group_id_list"`

	// Danh sách camera
	McuCameraList []McuCamera `json:"mcu_camera_list"`

	// Danh sách cảm biến
	McuSensorList []Sensor `json:"mcu_sensor_list"`

	// Danh sách cấu hình cảnh báo của MCU
	McuEventList []Event `json:"mcu_event_list"`
}

type McuCfg struct {

	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// Tel num của sim gắn vào MCU
	McuTelNumber string `json:"mcu_tel_number"`

	// Giá trị âm lượng
	McuVolume int32 `json:"mcu_volume"`

	// Địa chỉ IP Internet hiện tại của MCU
	McuIp string `json:"mcu_ip"`

	// Firmware version của thiết bị
	McuFirmwareVersion string `json:"mcu_firmware_verion"`

	// Giá trị âm lượng FM
	McuFMVolume int32 `json:"mcu_fm_volume"`

	// Ato-Detection hiện tại của khối FM
	McuFMAuto int32 `json:"mcu_fm_auto"`

	// Đường truyền dẫn về Server (0 - Ethernet, 1 - 3G, 2 - WIFI)
	McuTxType int32 `json:"mcu_txtype"`

	// Danh sách cms_group_id
	CmsGroupIdList []int64 `json:"cms_group_id_list"`

	// Danh sách camera
	McuCfgCameraList []McuCamera `json:"mcu_camera_list"`

	// Danh sách cảm biến
	McuCfgSensorList []CfgSensor `json:"mcu_sensor_list"`

	// Danh sách cấu hình cảnh báo của MCU
	McuCfgEventList []CfgEvent `json:"mcu_event_list"`
}

type McuOperatingState struct {
	// ID của MCU
	McuId int64 `json:"mcu_id"`

	//Đường truyền dẫn về Server
	McuTxType int32 `json:"mcu_txtype"`

	//Xác định xem loa có hoạt động (0 - lỗi, 1 - ok)
	McuSpeakerErr int32 `json:"mcu_speaker_err"`

	//Dòng tiêu thụ hiện tại của loa
	McuSpeakerCurrent float32 `json:"mcu_speaker_current"`

	//Mức tín hiệu Mobile Network, từ 0 – 100%
	McuMobileCsq int32 `json:"mcu_mobile_csq"`

	//Mức tín hiệu Wifi, từ 0 – 100%
	McuWifiCsq int32 `json:"mcu_wifi_csq"`

	//Nhiệt độ bên trong hộp thiết bị
	McuTemp float32 `json:"mcu_temp"`

	//Trạng thái tín hiệu FM thu được (0 - no signal, 1 - good signal)
	McuFMStatus int32 `json:"mcu_fm_status"`
}

type McuReboot struct {

	// Theo nhóm hay multicast đến MCU
	ReceiptType int32 `json:"receipt_type"`

	// Id của nhóm hoặc MCU nhận lệnh reboot
	ReceiptIds []int64 `json:"receipt_ids"`
}

type Record struct {

	// ID của bản tin trong hệ thống CMS
	RecId int64 `json:"rec_id"`

	// Tóm tắt bằng text
	RecSummary string `json:"rec_summary"`

	// Index của file Header sẽ được phát khi bắt đầu phát bản tin này (range 0-255)
	RecHeader int32 `json:"rec_header"`

	// Đường dẫn tới file audio trên file server
	RecAudio string `json:"rec_audio"`

	// Kiểu codec của audio data (ogg speex)
	RecAudioCodec int32 `json:"rec_audio_codec"`

	// Chế độ phát bản tin (0 – không phát, 1 – phát theo lịch, 2 – phát ngay lập tức)
	RecPlayMode int32 `json:"rec_play_mode"`

	// Mức độ ưu tiên phát của bản tin
	RecPriority int32 `json:"rec_priority"`

	// Các thời điểm trong ngày phát bản tin (tính theo phút, một ngày có 1440 phút)
	RecPlayTime []int32 `json:"rec_play_time"`

	// Kiểu lặp lại việc phát bản tin (0 - theo tuần, 1 - theo tháng)
	RecPlayRepeatType int32 `json:"rec_play_repeat_type"`

	// Các ngày trong tuần (hoặc tháng) lặp lại bản tin
	RecPlayRepeatDays int32 `json:"rec_play_repeat_days"`

	//Mốc thời gian bắt đầu phát bản tin
	RecPlayStart time.Time `json:"rec_play_start"`

	// Thời gian bản tin tự hết hạn phát
	RecPlayExpire time.Time `json:"rec_play_expire"`

	// Kiểu phát (theo nhóm hay multicast đến MCU)
	RecReceiptType int32 `json:"rec_receipt_type"`

	// Id của nhóm hoặc MCU nhận bản tin
	RecReceiptIds []int64 `json:"rec_receipt_ids"`
}

type RecordUpdate struct {

	// ID của bản tin trong hệ thống CMS
	RecordId int64 `json:"record_id"`

	// ID của mcu đã phát bản tin
	McuId int64 `json:"mcu_id"`

	// Loại log của bản tin(1 - Đã tạo bản tin mới, 2 - Đã sửa bản tin, 3 - Bắt đầu phát phát bản tin, 4 - Đã xóa bản tin, 7 - Download bản tin bị lỗi, 11 - file bản tin bị lỗi - đọc ra từ thẻ nhớ khi chuẩn bị phát bị lỗi, 12 - dừng phát bản tin, 13 - chuẩn bị phát bản tin FM theo lịch - đưa vào playlist)
	LogType int32 `json:"log_type"`

	// Thời gian tạo log
	LogTime int64 `json:"log_time"`

	// Lý do phát (1 – phát theo lịch, 2 – phát ngay lập tức). Chỉ có ý nghĩa với kiểu log; 3 - Bắt đầu phát bản tin ( mặc định là 0)
	PlayReason int32 `json:"play_reason"`

	// Thời điểm trong ngày phát bản tin theo lịch. Chỉ có ý nghĩa nếu lý do phát là theo lịch (mặc định là -1)
	PlayReasonTime int32 `json:"play_reason_time"`
}

type RecordRes struct {

	// ID của bản tin
	RecId int64 `json:"rec_id"`
}

type RecordDel struct {

	// ID của bản tin cần xoá
	RecId int64 `json:"rec_id"`

	// Xoá theo nhóm hay MCU
	RecReceiptType int32 `json:"rec_receipt_type"`

	// Id của nhóm hoặc MCU cần xoá
	RecReceiptIds []int64 `json:"rec_receipt_ids"`
}

type RecordPlay struct {

	// ID của bản tin cần thay đổi thời gian, chế độ phát
	RecId int64 `json:"rec_id"`

	// Các thời điểm trong ngày phát bản tin đã thay đổi
	RecPlayTime []int32 `json:"rec_play_time"`

	// Chế độ phát đã thay đổi
	RecPlayMode int32 `json:"rec_play_mode"`

	// Thay đổi theo nhóm hay MCU
	RecReceiptType int32 `json:"rec_receipt_type"`

	// Id của nhóm hoặc MCU cần thay đổi
	RecReceiptIds []int64 `json:"rec_receipt_ids"`
}

type RecordHeader struct {

	// 0 - Xóa header; 1 - Cập nhật header
	HeaderCmd int32 `json:"header_cmd"`

	// Index của file header
	HeaderIndex int32 `json:"header_idx"`

	// Đường dẫn tới file audio trên file server
	HeaderAudioPath string `json:"header_audio_path"`

	// Kiểu codec của audio data (ogg speex)
	HeaderAudioCodec int32 `json:"header_audio_codec"`

	// Theo nhóm hay multicast đến MCU
	HeaderReceiptType int32 `json:"header_receipt_type"`

	// Id của nhóm hoặc MCU nhận file header
	HeaderReceiptIds []int64 `json:"header_receipt_ids"`
}

type PlayHeader struct {

	// Index của file header
	HeaderIndex int32 `json:"header_idx"`

	//Mức độ ưu tiên phát nhạc nền = -10
	HeaderPriority int32 `json:"header_priority"`

	// Theo nhóm hay multicast đến MCU
	HeaderReceiptType int32 `json:"header_receipt_type"`

	// Id của nhóm hoặc MCU nhận lệnh phát file header
	HeaderReceiptIds []int64 `json:"header_receipt_ids"`
}

type HeaderRes struct {

	// Index của file header
	HeaderIndex int32 `json:"header_idx"`
}

type RecLiveStream struct {

	// ID của bản tin trong hệ thống CMS
	RecId int64 `json:"rec_id"`

	// 0 - tiếp tục phát live, 1 - đây là đoạn cuối của bản tin tiếp sóng
	RecState int32 `json:"rec_state"`

	// Index của bản tin tiếp sóng, bắt đầu từ 1, tăng dần
	RecChunkIndex int32 `json:"rec_chunk_index"`

	// Số lượng chunk cần được cho vào bộ đệm trước khi phát.
	RecBufferSize int32 `json:"rec_buffer_size"`

	// Thời gian hết hiệu lực của chunk này (tính bằng giây). Vd - 30s
	RecLiveTime int32 `json:"rec_live_time"`

	// Kiểu codec của audio data (ogg speex)
	RecAudioCodec int32 `json:"rec_audio_codec"`

	// Đường dẫn tới file audio trên file server
	RecAudio string `json:"rec_audio"`

	// 	Mức độ ưu tiên phát của bản tin tiếp sóng
	RecPriority int32 `json:"rec_priority"`

	// Kiểu phát (theo địa bàn hay multicast đến MCU)
	RecReceiptType int32 `json:"rec_receipt_type"`

	// Id của địa bàn hoặc MCU nhận bản tin tiếp sóng
	RecReceiptIds []int64 `json:"rec_receipt_ids"`
}

type ChunkRes struct {

	// ID của tin tiếp sóng
	RecId int64 `json:"rec_id"`

	// Index của chunk
	RecChunIndex int64 `json:"rec_chunk_index"`
}

type RecFM struct {

	// ID của bản tin trong hệ thống CMS
	FMId int64 `json:"fm_id"`

	// index của file Header sẽ được phát khi bắt đầu phát bản tin FM này
	FMHeader int32 `json:"fm_header"`

	// Tần số của kênh FM
	FMFrequence float32 `json:"fm_frequency"`

	// Chế độ phát bản tin FM (0 – không phát, 1 – phát theo lịch, 2 – phát ngay lập tức)
	FMPlayMode int32 `json:"fm_play_mode"`

	// Thời gian mà Mgateway sẽ tự động chuyển sang chế độ tiếp sóng FM nếu mất kết nối đến Server thông qua 3G hoặc Ethernet (0 - giá trị = 0 sẽ ko tự động phát; đơn vị giây)
	FMAutoSwitchTime int32 `json:"fm_auto_switch_time"`

	// Mức độ ưu tiên phát của bản tin FM
	FMPriority int32 `json:"fm_priority"`

	// Các thời điểm trong ngày phát bản tin FM(tính theo phút, một ngày có 1440 phút)
	FMPlayTime []int32 `json:"fm_play_time"`

	// Kiểu lặp lại việc phát bản tin FM (0 - theo tuần, 1 - theo tháng)
	FMPlayRepeatType int32 `json:"fm_play_repeat_type"`

	// Các ngày trong tuần (hoặc tháng) lặp lại bản tin FM
	FMPlayRepeatDays int32 `json:"fm_play_repeat_days"`

	// Thời gian bắt đầu phát bản tin FM
	FMPlayStart time.Time `json:"fm_play_start"`

	//Thời gian bắt đầu phát bản tin FM
	FMPlayExpire time.Time `json:"fm_play_expire"`

	// Khoảng thời gian phát tin tiếp sòng FM
	FMPlayDuration int32 `json:"fm_play_duration"`

	// Kiểu phát (theo địa bàn hay multicast đến MCU)
	FMReceiptType int32 `json:"fm_receipt_type"`

	// Id của địa bàn hoặc MCU nhận bản tin tiếp sóng
	FMReceiptIds []int64 `json:"fm_receipt_ids"`
}

type RecordStat struct {
	// ID của MCU
	McuId int64 `json:"mcu_id"`

	// ID của bản tin trong hệ thống CMS
	RecId int64 `json:"rec_id"`

	// Trạng thái bản tin (0 - OFF, 1 - NORMAL, 2 - PREPARE,  3 - PLAYING, 4 - PLAYED)
	RecStatus int32 `json:"rec_status"`

	// Chế độ phát bản tin (0 – không phát, 1 – phát theo lịch, 2 – phát ngay lập tức)
	RecPlayMode int32 `json:"rec_play_mode"`

	// Mức độ ưu tiên phát của bản tin
	RecPriority int32 `json:"rec_priority"`

	// Các thời điểm trong ngày phát bản tin (tính theo giây, một ngày có 86400 giây)
	RecPlayTime []int32 `json:"rec_play_time"`

	// Kiểu lặp lại việc phát bản tin (0 - theo tuần, 1 - theo tháng)
	RecPlayRepeatType int32 `json:"rec_play_repeat_type"`

	// Các ngày trong tuần (hoặc tháng) lặp lại bản tin
	RecPlayRepeatDays int32 `json:"rec_play_repeat_days"`

	// Thời gian tạo
	RecCreateTime int64 `json:"rec_created"`

	// Mốc thời gian bắt đầu phát bản tin
	RecPlayStart int64 `json:"rec_play_start"`

	// Thời gian bản tin tự hết hạn phát
	RecPlayExpire int64 `json:"rec_play_expire"`

	// Kiểu codec của audio (0 - PCM/16bit/8kHz - 128kbps, 1 - G.711/A-law - 64kbps, 2 - mp3)
	RecAudioCodec int32 `json:"rec_audio_codec"`

	// Định dạng audio (0 - RAW, 1 - wav, 2 - mp3)
	RecAudioFormat int32 `json:"rec_audio_format"`

	// Kích thước file
	RecSize int32 `json:"rec_size"`

	// CRC32
	RecChecksum int64 `json:"rec_checksum"`

	// Tần số của kênh FM
	FMFrequence float32 `json:"fm_frequency"`

	// Khoảng thời gian phát tin tiếp sòng FM
	FMPlayDuration int32 `json:"fm_play_duration"`

	// Thời gian mà Mgateway sẽ tự động chuyển sang chế độ tiếp sóng FM nếu mất kết nối đến Server thông qua 3G hoặc Ethernet (0 - giá trị = 0 sẽ ko tự động phát; đơn vị giây)
	FMAutoSwitchTime int32 `json:"fm_auto_switch_time"`
}

type FMRes struct {

	// ID của tin FM
	FMId int64 `json:"fm_id"`
}

type NodeEmqttd struct {
	Name             string `json:"name"`
	OtpRelease       string `json:"otp_release"`
	TotalMemory      string `json:"total_memory"`
	UsedMemory       string `json:"used_memory"`
	ProcessAvailable int    `json:"process_available"`
	ProcessUsed      int    `json:"process_used"`
	MaxFds           int    `json:"max_fds"`
	ClusterStatus    string `json:"cluster_status"`
	NodeStatus       string `json:"node_status"`
}
type NodeVenmqttd struct {
	Nodes []NodeDetail `json:"table"`
	Type string `json:"type"`
}
type NodeDetail struct{
	NodeName string `json:"Node"`
	IsRunning bool `json:"Running"`
}

type V2Nodes struct {
	Code   int          `json:"code"`
	Result []NodeEmqttd `json:"result"`
}

type EmqttdListStatus struct {
	CurrentPage int                `json:"currentPage"`
	PageSize    int                `json:"pageSize"`
	TolalNum    int                `json:"totalNum"`
	TotalPage   int                `json:"totalPage"`
	Result      []MqttClientStatus `json:"result"`
}

type MqttClientStatus struct {
	ClientId    string `json:"client_id"`
	IsOnline	bool `json:"is_online"`
	MountPoint 	string `json:"mountpoint"`
	PeerHost 	string `json:"peer_host"`
	PeerPort 	int `json:"peer_port"`
	User 		string `json:"user"`
}
type MQTTClientList struct{
	MQTTClient []MqttClientStatus `json:"table"`
	Type string `json:"type"`
}
type MQTTClientTs struct{
	Queue_started_at           int64  `json:"queue_started_at"`
	Session_started_at           int64  `json:"session_started_at"`
}

type MQTTClientTsList struct{
	MQTTClientts []MQTTClientTs `json:"table"`
	Type string                 `json:"type"`
}

type McuCfgLog struct {

	// ID của MCU
	McuId string `json:"mcuId"`

	VCode string `json:"vCode"`

	// Tel num của sim gắn vào MCU
	McuTelNumber string `json:"telNumber"`

	// Giá trị âm lượng
	McuVolume string `json:"volume"`

	// Địa chỉ IP Internet hiện tại của MCU
	McuIp string `json:"ip"`

	// Firmware version của thiết bị
	McuFirmwareVersion string `json:"firmwareVersion"`

	// Giá trị âm lượng FM
	McuFMVolume string `json:"fmVolume"`

	// Ato-Detection hiện tại của khối FM
	McuFMAuto string `json:"fmAuto"`

	// Đường truyền dẫn về Server (0 - Ethernet, 1 - 3G, 2 - WIFI)
	McuTxType string `json:"txType"`

	// Danh sách cms_group_id
	CmsGroupIdList string `json:"groupIdList"`

	// Danh sách camera
	McuCfgCameraList string `json:"cfgCameraList"`

	// Danh sách cảm biến
	McuCfgSensorList string `json:"cfgSensorList"`

	// Danh sách cấu hình cảnh báo của MCU
	McuCfgEventList string `json:"cfgEventList"`
}

type RequestBody struct {
	// ID của MCU, địa bàn, nhóm MCU
	Id string `json:"id"`

	// ID của Channel Vernemq
	ChannelId string `json:"channelId"`
	// Thời gian bắt đầu
	StartTime string `json: "startTime"`
	//Thời gian kết thúc
	EndTime string `json: "endTime"`
	//Loại log: log_monitoring, log_OPU, log_control
	TypeLog string `json:"typeLog"`
	/*Mô tả cụ thể loại log
		log_monitoring: volume
						temp
		log_OPU: OPU_GENERIC,
				OPU_CAMERA,
				OPU_PHONE,
				OPU_SENSOR,
				OPU_ALARM,
				OPU_STATUS,
				OPU_MEDIA,
				OPU_LOG
		log_control: Control_Verify_Vcode,
					Control_AddRecord
					Control_LoadMCU_Data
					Control_Delete_MCU
					Control_Set_CMSGroup
					Control_MCU_AddRecord
					Control_MCU_Edit_Record
					Control_MCU_AddRecord

	*/
	DescLog string `json:"descLog"`
	//Số lượng bản ghi
	Limit int `json:"limit"`
}

type ValueQuerry struct {

	// Value của query database
	Value string `json:"value"`

	// time của Value của query database
	Time string `json:"time"`

}

type ResponseLogsData struct {
	// ID của MCU
	Id string `json:"id"`
	// ID của Channel Vernemq
	ChannelId string              `json:"channelId"`

	//Loại log: log_monitoring, log_OPU, log_control
	TypeLog string `json:"typeLog"`
	/*Mô tả cụ thể loại log
	log_monitoring: volume
					temp
	log_OPU: OPU_GENERIC,
			OPU_CAMERA,
			OPU_PHONE,
			OPU_SENSOR,
			OPU_ALARM,
			OPU_STATUS,
			OPU_MEDIA,
			OPU_LOG
	log_control: Control_Verify_Vcode,
				Control_AddRecord
				Control_LoadMCU_Data
				Control_Delete_MCU
				Control_Set_CMSGroup
				Control_MCU_AddRecord
				Control_MCU_Edit_Record
				Control_MCU_AddRecord

	*/
	DescLog string `json:"descLog"`
	//Số lượng bản ghi yêu cầu
	Limit int `json:"limit"`

	//Số lượng bản ghi thực tế
	Total int `json:"total"`
	//danh sách logs theo thời gian
	LogList      []ValueQuerry `json:"logs"`
}


const (
	mid       = 0
	status    = 1
	prio      = 2
	created   = 3
	expired   = 4
	start     = 5
	mode      = 6
	repeat    = 7
	days      = 8
	size      = 9
	codec     = 10
	format    = 11
	checksum  = 12
	frequency = 13
	duration  = 14
	swtime    = 15
	time_cnt  = 16
)

func ConvertOPUMediaToStruct(in [][]int64, mcuid int64) []RecordStat {
	items := []RecordStat{}
	for _, info := range in {
		if len(info) >= 17 {
			times := []int32{}
			cnt := int(info[time_cnt])
			for i := 1; i <= cnt; i++ {
				play := int32(info[i+time_cnt])
				times = append(times, play)
			}
			media := RecordStat{
				McuId:             mcuid,
				RecId:             info[mid],
				RecStatus:         int32(info[status]),
				RecPriority:       int32(info[prio]),
				RecCreateTime:     info[created],
				RecPlayExpire:     info[expired],
				RecPlayStart:      info[start],
				RecPlayMode:       int32(info[mode]),
				RecPlayRepeatType: int32(info[repeat]),
				RecPlayRepeatDays: int32(info[days]),
				RecPlayTime:       times,
				RecSize:           int32(info[size]),
				RecAudioCodec:     int32(info[codec]),
				RecAudioFormat:    int32(info[format]),
				RecChecksum:       info[checksum],
				FMFrequence:       float32(info[frequency] / 10),
				FMPlayDuration:    int32(info[duration]),
				FMAutoSwitchTime:  int32(info[swtime]),
			}
			items = append(items, media)
		}
	}
	return items
}
type MCUVerifyTTTM struct {
	McuId int64 `json:"mcuId"`
	VCode string `json:"vCode"`
	Name string `json:"name"`
	Address string `json:"address"`
	Lat string `json:"latitude"`
	Lon string `json:"longitude"`
	SimNumber string `json:"simNumber"`
	SimSerial string `json:"simSerial"`
}
type ResponseMSGTTTM struct {
	Status int `json:"status"`
	Code string `json:"code"`
	Message string `json:"message"`
}
type DisconnectClientMQTT struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
type McuIDReqTTTM struct {
	McuId int64 `json:"mcuId"`
}
