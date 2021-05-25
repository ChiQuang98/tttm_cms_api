package redis

import (
	"errors"
	"fmt"
	"net"
	"time"

	"gopkg.in/redis.v3"
)

var (
	client  *redis.Client
	iClient *redis.Client
	mClient *redis.Client
	lClient *redis.Client
)

var ExpireKeyChannelName string

const (
	REDIS_RECEIVE_MESSAGE_TIMEOUT             = 5
	REDIS_RECEIVE_MESSAGE_TIMEOUT_RESUBSCRIBE = 150
	JOB_QUEUE_MAX                             = 10000000
	JOB_PRIORITY                              = 1000
)

func IsInit() bool {
	return client != nil
}

//var redisPubChan chan PublishPacket
//go-redis's code
func Init(address string, password string) {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     address,
			Password: password,
			//WriteTimeout: time.Second * 2,
			IdleTimeout: time.Minute * 3,
			DB:          int64(0),
			PoolSize:    10000,
			PoolTimeout: time.Minute * 5,
			//MaxRetries:   3,
		})
		lClient = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     address,
			Password: password,
			//WriteTimeout: time.Second * 2,
			IdleTimeout: time.Minute * 4,
			DB:          int64(DBLiveStream),
			PoolSize:    10000,
			PoolTimeout: time.Minute * 5,
			//MaxRetries:   3,
		})
		iClient = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     address,
			Password: password,
			//WriteTimeout: time.Second * 2,
			IdleTimeout: time.Minute * 4,
			DB:          int64(DBMcuInfo),
			PoolSize:    10000,
			PoolTimeout: time.Minute * 5,
			//MaxRetries:   3,
		})
		mClient = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     address,
			Password: password,
			//WriteTimeout: time.Second * 2,
			IdleTimeout: time.Minute * 4,
			DB:          int64(DBMedia),
			PoolSize:    10000,
			PoolTimeout: time.Minute * 5,
			//MaxRetries:   3,
		})
	}
}
func Init1(address ... string) {
	if client == nil {
		//StreamClient = RedisStream.NewClient(&RedisStream.Options{
		//	Addr:     "10.4.200.61:6379",
		//	Password: "",
		//	DB:       0, // use default DB
		//})
		//createConsumerGroup()
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName: "mymaster",
			SentinelAddrs: []string{
				address[0],
				address[1],
				address[2],
			},
		})
		lClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName: "mymaster",
			SentinelAddrs: []string{
				address[0],
				address[1],
				address[2],
			},
		})
		iClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName: "mymaster",
			SentinelAddrs: []string{
				address[0],
				address[1],
				address[2],
			},
		})
		mClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName: "mymaster",
			SentinelAddrs: []string{
				address[0],
				address[1],
				address[2],
			},
		})

	}
}
func Ping() bool {
	_, err := client.Ping().Result()
	return err == nil
}

/* PUB-SUB */

func Subscribe(recFunc func(*redis.Message), unSubscribe chan bool, keyExpire bool, isPattern bool, channels ...string) {

	if keyExpire {
		client.ConfigSet("notify-keyspace-events", "Ex")
		channels = append(channels, ExpireKeyChannelName)
	}

	var c *redis.PubSub
	var err error
	if !isPattern {
		c, err = client.Subscribe(channels...)
	} else {
		c, err = client.PSubscribe(channels...)
	}

	defer func() {
		//unsubscribe & close current pubsub
		if !isPattern {
			c.Unsubscribe(channels...)
		} else {
			c.PUnsubscribe(channels...)
		}
		c.Close()

		//will recover panic error when connection lost (redis server restart, down, ...)
		err := recover()
		if err != nil {
			fmt.Errorf("Pubsub error %v", err)
			if Ping() {
				go Subscribe(recFunc, unSubscribe, keyExpire, isPattern, channels...) //must use other routine
				return
			}
		}
	}()

	if err != nil {
		panic(err)
	}

	//must receive in same routine with subscribe caller
	timeOut := 0
	for {

		select {
		case <-unSubscribe:
			//glog.Info("unsubscribe")
			//c.Close() // already close in defer
			return

		default:
			msgi, err := c.ReceiveTimeout(REDIS_RECEIVE_MESSAGE_TIMEOUT * time.Second)
			if err == nil {
				switch msg := msgi.(type) {
				case *redis.PMessage:
					recMsg := &redis.Message{
						Channel: msg.Channel,
						Pattern: msg.Pattern,
						Payload: msg.Payload,
					}
					recFunc(recMsg)
					timeOut = 0
				}
			} else {
				//panic error when need to re-subscribe

				if err.Error() == "EOF" {
					panic(err)
				}

				if e, ok := err.(net.Error); !ok || !e.Timeout() {
					panic(err)
				}

				if !Ping() {
					panic(errors.New("could not ping to redis"))
				}

				//ok this is timeout error
				timeOut += REDIS_RECEIVE_MESSAGE_TIMEOUT
				if timeOut > REDIS_RECEIVE_MESSAGE_TIMEOUT_RESUBSCRIBE {
					panic(errors.New(fmt.Sprintf("Subscribe time out for %v on channel: %v", timeOut, channels)))
				}
			}
		}
	}

}

//Publish value to channel
func Publish(channel string, value string) error {
	return client.Publish(channel, value).Err()
}

func ClearAll() {
	client.FlushAll()
}

func ClearDB(index int64) {
	client.Select(index)
	client.FlushDb()
	client.Select(0)
}
