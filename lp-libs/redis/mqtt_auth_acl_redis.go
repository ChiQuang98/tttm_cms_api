package redis

import (
	"fmt"
)

//NOTE: 1: subscribe, 2: publish, 3: pubsub

const (
	ACL_TOPIC_SUB    = "1"
	ACL_TOPIC_PUB    = "2"
	ACL_TOPIC_PUBSUB = "3"
)

func SetAclSuperUser(username, password string) error {
	key := fmt.Sprintf("mqtt_user:%s", username)
	err := client.HMSet(key, "is_superuser ", "1").Err()
	err = client.HMSet(key, "password", password).Err()
	return err
}

func SetAclMcuUser(mcuId string) error {
	key := fmt.Sprintf("mqtt_user:%s", mcuId)
	mcuSubTopic := fmt.Sprintf("h/d/%s", mcuId)
	mcuPubTopic := fmt.Sprintf("h/s/%s", mcuId)
	err := client.HMSet(key, "password", mcuId).Err()
	err = client.HMSet(key, mcuSubTopic, ACL_TOPIC_SUB).Err()
	err = client.HMSet(key, mcuPubTopic, ACL_TOPIC_PUB).Err()
	return err
}

func SetAclMcuTopic(mcuId int64, groups []int64) error {
	key := fmt.Sprintf("mqtt_acl:%d", mcuId)
	fields := map[string]string{}
	//acl subscriber topic
	mcuSubTopic := fmt.Sprintf("h/d/%d", mcuId)
	fields[mcuSubTopic] = ACL_TOPIC_SUB
	//acl publish topic
	mcuPubTopic := fmt.Sprintf("h/s/%d", mcuId)
	fields[mcuPubTopic] = ACL_TOPIC_PUB

	for _, g := range groups {
		groupSubTopic := fmt.Sprintf("h/g/%d", g)
		fields[groupSubTopic] = ACL_TOPIC_SUB
	}

	return client.HMSetMap(key, fields).Err()
}

func DeleteAclMcuTopic(mcuId int64) error {
	key := fmt.Sprintf("mqtt_acl:%d", mcuId)
	return client.Del(key).Err()
}
