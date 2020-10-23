package kafka

import (
	"errors"
	"fmt"
	conf "work_report/config"
	"work_report/libraries/elog"

	"github.com/Shopify/sarama"
)

type Kafka struct {
	Topic   string
	Key     string
	StrVal  string
	ByteVal []byte
}

//创建kafka连接
func Connection() sarama.SyncProducer {
	//获取kafka url
	url := conf.GetApolloString(conf.KafkaUrl, "127.0.0.1") + ":" + conf.GetApolloString(conf.KafKaProt, "9092")
	// 新建一个arama配置实例
	config := sarama.NewConfig()

	// WaitForAll waits for all in-sync replicas to commit before responding.
	config.Producer.RequiredAcks = sarama.WaitForAll

	// NewRandomPartitioner returns a Partitioner which chooses a random partition each time.
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	config.Producer.Return.Successes = true

	// 新建一个同步生产者
	client, err := sarama.NewSyncProducer([]string{url}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		elog.Newf("producer close, err:%s", err.Error())
		return nil
	}
	return client
}

//生产者
func Producer(producer Kafka) error {
	client := Connection()
	defer client.Close()
	if producer.Topic == "" {
		return errors.New("Topic 不能为空")
	}
	// 定义一个生产消息，包括Topic、消息内容、
	msg := &sarama.ProducerMessage{}
	msg.Topic = producer.Topic
	if producer.Key != "" {
		msg.Key = sarama.StringEncoder(producer.Key)
	}
	if len(producer.ByteVal) > 0 {
		msg.Value = sarama.ByteEncoder(producer.ByteVal)
	} else {
		msg.Value = sarama.StringEncoder(producer.StrVal)
	}
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		elog.New("send message failed,"+err.Error(), elog.FileMsg{})
	} else {
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		elog.Newf("pid:%v offset:%v\n", pid, offset)
	}
	return err
}
