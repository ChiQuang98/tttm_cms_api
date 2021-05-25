package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"tttm_cms_api/lp-libs/models"
	"github.com/golang/glog"
)

func CmsRevertRecordById(recordId int) (int, []byte) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsRevertRecordById err: ", err)
		}
	}()
	client := &http.Client{}
	req, err := http.NewRequest("GET", serverBasePath+"/record-revert/"+strconv.Itoa(recordId), nil)
	if err != nil {
		return http.StatusBadGateway, nil
	}
	if len(cmsInfo.UserName) > 0 {
		req.SetBasicAuth(cmsInfo.UserName, cmsInfo.Password)
	}
	res, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil
	}
	defer res.Body.Close()
	if res != nil {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return http.StatusInternalServerError, nil
		}
		status := res.StatusCode
		return status, body
	}
	return http.StatusInternalServerError, nil
}

func CmsUpdateRecord(record models.RecordUpdate) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsUpdateRecord err: ", err)
		}
	}()
	client := &http.Client{}
	buf, err := json.Marshal(record)
	if err != nil {
		return false, "", err
	}

	req, err := http.NewRequest(http.MethodPut, serverBasePath+"/record-update", bytes.NewBuffer(buf))
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

func CmsRecordStat(medias []models.RecordStat) (bool, string, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("RECOVER/CmsUpdateRecord err: ", err)
		}
	}()
	client := &http.Client{}
	buf, err := json.Marshal(medias)
	if err != nil {
		return false, "", err
	}

	req, err := http.NewRequest(http.MethodPost, serverBasePath+"/record-stat", bytes.NewBuffer(buf))
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
