package redis

import (
	"encoding/json"
	"strconv"
	"time"

	"tttm_cms_api/lp-libs/models"
)

var fmInfoFields = []string{
	FM_PlayMode,
	FM_Priority,
	FM_PlayTime,
	FM_PlayRepeatType,
	FM_PlayRepeatDays,
	FM_PlayStart,
	FM_PlayExpire,
	FM_ReceiptType,
	FM_ReceiptIds,
	FM_Frequence,
	FM_Duration,
	FM_AutoSwitchTime,
}

func UpdateFmInfo(fm *models.RecFM) error {
	fmPlayMode := string(fm.FMPlayMode)
	fmPriority := string(fm.FMPriority)
	fmFrequence := strconv.FormatFloat(float64(fm.FMFrequence), 'f', 5, 64)
	fmDuration := string(fm.FMPlayDuration)
	fmAutoSwitchTime := string(fm.FMAutoSwitchTime)
	fmPlayTime, err := json.Marshal(fm.FMPlayTime)
	if err != nil {
		return err
	}
	fmPlayRepeatType := string(fm.FMPlayRepeatType)
	fmPlayRepeatDays := string(fm.FMPlayRepeatDays)

	fmReceiptType := string(fm.FMReceiptType)
	fmPlayStart := fm.FMPlayStart.Format("2006-01-02 15:04:05")
	fmPlayExpire := fm.FMPlayExpire.Format("2006-01-02 15:04:05")
	fmReceiptIds, err := json.Marshal(fm.FMReceiptIds)
	if err != nil {
		return err
	}

	return mClient.HMSet(
		strconv.FormatInt(fm.FMId, 10),
		FM_PlayMode, fmPlayMode,
		FM_Priority, fmPriority,
		FM_Frequence, fmFrequence,
		FM_Duration, fmDuration,
		FM_AutoSwitchTime, fmAutoSwitchTime,
		FM_PlayTime, string(fmPlayTime),
		FM_PlayRepeatType, fmPlayRepeatType,
		FM_PlayRepeatDays, fmPlayRepeatDays,
		FM_PlayStart, fmPlayStart,
		FM_PlayExpire, fmPlayExpire,
		FM_ReceiptType, fmReceiptType,
		FM_ReceiptIds, string(fmReceiptIds),
	).Err()
}

func GetFMInforById(fmId int64) *models.RecFM {
	result, err := mClient.HMGet(strconv.FormatInt(fmId, 10), fmInfoFields...).Result()
	if err != nil {
		return nil
	}
	if result == nil {
		return nil
	}
	fmPlayMode := int32(0)
	if result[0] != nil {
		mode, _ := strconv.ParseInt(result[0].(string), 10, 32)
		fmPlayMode = int32(mode)
	}
	fmPriority := int32(0)
	if result[1] != nil {
		priority, _ := strconv.ParseInt(result[1].(string), 10, 32)
		fmPriority = int32(priority)
	}
	fmFrequence := float32(0)
	if result[2] != nil {
		frequence, _ := strconv.ParseFloat(result[2].(string), 32)
		fmFrequence = float32(frequence)
	}
	fmDuration := int32(0)
	if result[3] != nil {
		duration, _ := strconv.ParseInt(result[3].(string), 10, 32)
		fmDuration = int32(duration)
	}
	fmAutoSwitchTime := int32(0)
	if result[4] != nil {
		swtime, _ := strconv.ParseInt(result[4].(string), 10, 32)
		fmAutoSwitchTime = int32(swtime)
	}
	fmPlayTime := []int32{}
	if result[5] != nil {
		json.Unmarshal([]byte(result[5].(string)), &fmPlayTime)
	}
	fmPlayRepeatType := int32(0)
	if result[6] != nil {
		playRepeatType, _ := strconv.ParseInt(result[6].(string), 10, 32)
		fmPlayRepeatType = int32(playRepeatType)
	}
	fmPlayRepeatDay := int32(0)
	if result[7] != nil {
		playRepeatDay, _ := strconv.ParseInt(result[7].(string), 10, 32)
		fmPlayRepeatDay = int32(playRepeatDay)
	}
	fmPlayStart := time.Now()
	if result[8] != nil {
		layout := "2006-01-02T15:04:05.000Z"
		timeStr := result[8].(string)
		fmPlayStart, _ = time.Parse(layout, timeStr)
	}
	fmPlayExpire := time.Now()
	if result[9] != nil {
		layout := "2006-01-02T15:04:05.000Z"
		timeStr := result[9].(string)
		fmPlayExpire, _ = time.Parse(layout, timeStr)
	}
	fmReceiptType := int32(0)
	if result[10] != nil {
		receiptType, _ := strconv.ParseInt(result[10].(string), 10, 64)
		fmReceiptType = int32(receiptType)
	}
	fmReceiptIDs := []int64{}
	if result[11] != nil {
		json.Unmarshal([]byte(result[11].(string)), &fmReceiptIDs)
	}

	fm := &models.RecFM{
		FMId:             fmId,
		FMPlayMode:       fmPlayMode,
		FMPriority:       fmPriority,
		FMFrequence:      fmFrequence,
		FMPlayDuration:   fmDuration,
		FMAutoSwitchTime: fmAutoSwitchTime,
		FMPlayTime:       fmPlayTime,
		FMPlayRepeatType: fmPlayRepeatType,
		FMPlayRepeatDays: fmPlayRepeatDay,
		FMPlayStart:      fmPlayStart,
		FMPlayExpire:     fmPlayExpire,
		FMReceiptType:    fmReceiptType,
		FMReceiptIds:     fmReceiptIDs,
	}
	return fm
}
