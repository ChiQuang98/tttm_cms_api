package redis

import (
	"encoding/json"
	"strconv"
	"time"

	"tttm_cms_api/lp-libs/models"
)

var recordInfoFields = []string{
	Record_Summary,
	Record_AudioLink,
	Record_AudioCodec,
	Record_PlayMode,
	Record_Priority,
	Record_PlayTime,
	Record_PlayRepeatType,
	Record_PlayRepeatDays,
	Record_PlayStart,
	Record_PlayExpire,
	Record_ReceiptType,
	Record_ReceiptIds,
}

func UpdateRecordInfo(record *models.Record) error {
	recAudioCodec := string(record.RecAudioCodec)
	recPlayMode := string(record.RecPlayMode)
	recPriority := string(record.RecPriority)
	recPlayTime, err := json.Marshal(record.RecPlayTime)
	if err != nil {
		return err
	}
	recPlayRepeatType := string(record.RecPlayRepeatType)
	recPlayRepeatDays := string(record.RecPlayRepeatDays)

	recReceiptType := string(record.RecReceiptType)
	recPlayStart := record.RecPlayStart.Format("2006-01-02 15:04:05")
	recPlayExpire := record.RecPlayExpire.Format("2006-01-02 15:04:05")
	recReceiptIds, err := json.Marshal(record.RecReceiptIds)
	if err != nil {
		return err
	}

	return mClient.HMSet(
		strconv.FormatInt(record.RecId, 10),
		Record_Summary, record.RecSummary,
		Record_AudioLink, record.RecAudio,
		Record_AudioCodec, recAudioCodec,
		Record_PlayMode, recPlayMode,
		Record_Priority, recPriority,
		Record_PlayTime, string(recPlayTime),
		Record_PlayRepeatType, recPlayRepeatType,
		Record_PlayRepeatDays, recPlayRepeatDays,
		Record_PlayStart, recPlayStart,
		Record_PlayExpire, recPlayExpire,
		Record_ReceiptType, recReceiptType,
		Record_ReceiptIds, string(recReceiptIds),
	).Err()
}

func UpdateRecordPlayInfo(recPlay *models.RecordPlay) error {
	recPlayTime, err := json.Marshal(recPlay.RecPlayTime)
	if err != nil {
		return err
	}

	recReceiptIds, err := json.Marshal(recPlay.RecReceiptIds)
	if err != nil {
		return err
	}

	return mClient.HMSet(
		strconv.FormatInt(recPlay.RecId, 10),
		Record_PlayMode, string(recPlay.RecPlayMode),
		Record_PlayTime, string(recPlayTime),
		Record_ReceiptType, string(recPlay.RecReceiptType),
		Record_ReceiptIds, string(recReceiptIds),
	).Err()
}

func GetRecordInforById(recId int64) *models.Record {
	result, err := mClient.HMGet(strconv.FormatInt(recId, 10), recordInfoFields...).Result()
	if err != nil {
		return nil
	}
	if result == nil {
		return nil
	}
	recSummary := ""
	if result[0] != nil {
		recSummary = result[0].(string)
	}
	recAudioLink := ""
	if result[1] != nil {
		recAudioLink = result[1].(string)
	}
	recCodec := int32(0)
	if result[2] != nil {
		codec, _ := strconv.ParseInt(result[2].(string), 10, 32)
		recCodec = int32(codec)
	}
	recPlayMode := int32(0)
	if result[3] != nil {
		playMode, _ := strconv.ParseInt(result[3].(string), 10, 32)
		recPlayMode = int32(playMode)
	}
	recPriority := int32(0)
	if result[4] != nil {
		priority, _ := strconv.ParseInt(result[4].(string), 10, 32)
		recPriority = int32(priority)
	}
	recPlayTime := []int32{}
	if result[5] != nil {
		json.Unmarshal([]byte(result[5].(string)), &recPlayTime)
	}
	recPlayRepeatType := int32(0)
	if result[6] != nil {
		playRepeatType, _ := strconv.ParseInt(result[6].(string), 10, 32)
		recPlayRepeatType = int32(playRepeatType)
	}
	recPlayRepeatDay := int32(0)
	if result[7] != nil {
		playRepeatDay, _ := strconv.ParseInt(result[7].(string), 10, 32)
		recPlayRepeatDay = int32(playRepeatDay)
	}
	recPlayStart := time.Now()
	if result[8] != nil {
		layout := "2006-01-02T15:04:05.000Z"
		timeStr := result[8].(string)
		recPlayStart, _ = time.Parse(layout, timeStr)
	}
	recPlayExpire := time.Now()
	if result[9] != nil {
		layout := "2006-01-02T15:04:05.000Z"
		timeStr := result[9].(string)
		recPlayExpire, _ = time.Parse(layout, timeStr)
	}
	recReceiptType := int32(0)
	if result[10] != nil {
		receiptType, _ := strconv.ParseInt(result[10].(string), 10, 64)
		recReceiptType = int32(receiptType)
	}
	recReceiptIDs := []int64{}
	if result[11] != nil {
		json.Unmarshal([]byte(result[11].(string)), &recReceiptIDs)
	}

	record := &models.Record{
		RecId:             recId,
		RecSummary:        recSummary,
		RecAudio:          recAudioLink,
		RecAudioCodec:     recCodec,
		RecPlayMode:       recPlayMode,
		RecPriority:       recPriority,
		RecPlayTime:       recPlayTime,
		RecPlayRepeatType: recPlayRepeatType,
		RecPlayRepeatDays: recPlayRepeatDay,
		RecPlayStart:      recPlayStart,
		RecPlayExpire:     recPlayExpire,
		RecReceiptType:    recReceiptType,
		RecReceiptIds:     recReceiptIDs,
	}
	return record
}

func DeleteRecordById(recId int64) error {
	return mClient.Del(strconv.FormatInt(recId, 10)).Err()
}
