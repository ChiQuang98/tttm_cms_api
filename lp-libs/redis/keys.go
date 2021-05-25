package redis

const (
	DBMcuInfo    = 2 //m-gateway info
	DBMedia      = 3 //media info
	DBLiveStream = 1 //livestream info
)

//Key + Field
const (
	Mcu_Vcode            = "Vcode"
	Mcu_ConnectionStatus = "ConnectionStatus"
	Mcu_ConnectionTime   = "ConnectionTime"
	Mcu_PhoneNumber      = "PhoneNumber"
	Mcu_LocalIp          = "LocalIp"
	Mcu_WanIp            = "WanIp"
	Mcu_MediaIdLastest   = "MediaIdLastest"
	Mcu_Volume           = "Volume"
	Mcu_GroupList        = "GroupList"
	Mcu_SensorList       = "SensorList"
	Mcu_CameraList       = "CameraList"
	Mcu_AlarmList        = "EventList"
	Mcu_Version          = "Version"
	Mcu_FM_Volume        = "FMVolume"
	Mcu_FM_Auto          = "FMAuto"
	//status
	Mcu_TxType         = "TxType"
	Mcu_SpeakerErr     = "SpeakerErr"
	Mcu_SpeakerCurrent = "SpeakerCurrent"
	Mcu_MobileCsq      = "MobileCsq"
	Mcu_WifiCsq        = "WifiCsq"
	Mcu_Temp           = "Temp"
)

const (
	Record_Summary        = "Summary"
	Record_AudioLink      = "AudioLink"
	Record_AudioCodec     = "Codec"
	Record_PlayMode       = "Mode"
	Record_Priority       = "Priority"
	Record_PlayTime       = "PlayTime"
	Record_PlayRepeatType = "RepeatType"
	Record_PlayRepeatDays = "RepeatDay"
	Record_PlayStart      = "Start"
	Record_PlayExpire     = "Expire"
	Record_ReceiptType    = "ReceiptType"
	Record_ReceiptIds     = "ReceiptIds"
)

const (
	Record_Chunk_Index       = "ChunkIndex"
	Record_Chunk_BufferSize  = "ChunkBuffer"
	Record_Chunk_TimeLive    = "ChunkTimeLive"
	Record_Chunk_AudioLink   = "ChunkAudioLink"
	Record_Chunk_AudioCodec  = "ChunkCodec"
	Record_Chunk_State       = "ChunkState"
	Record_Chunk_Priority    = "ChunkPriority"
	Record_Chunk_ReceiptType = "ChunkReceiptType"
	Record_Chunk_ReceiptIds  = "ChunkReceiptIds"
)

const (
	FM_PlayMode       = "Mode"
	FM_Priority       = "Priority"
	FM_PlayTime       = "PlayTime"
	FM_PlayRepeatType = "RepeatType"
	FM_PlayRepeatDays = "RepeatDay"
	FM_PlayStart      = "Start"
	FM_PlayExpire     = "Expire"
	FM_ReceiptType    = "ReceiptType"
	FM_ReceiptIds     = "ReceiptIds"
	FM_Frequence      = "Frequence"
	FM_Duration       = "Duration"
	FM_AutoSwitchTime = "AutoSWTime"
)
