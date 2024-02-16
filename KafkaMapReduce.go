package MapReducde

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	jsoniter "github.com/json-iterator/go"
)

type KafkaMapReduce struct {
	producer sarama.SyncProducer
	address  []string
}

func (mr *KafkaMapReduce) Map(taskList []*Task) error {
	for _, task := range taskList {
		//对task发送
		taskjson, err := jsoniter.MarshalToString(task)
		if err != nil {
			log.Panic(err)
			return err
		}
		msg := &sarama.ProducerMessage{
			Topic: task.Group,
			Value: sarama.ByteEncoder(taskjson),
		}
		_, _, err = mr.producer.SendMessage(msg)
		if err != nil {
			log.Panic(err)
			return err
		}
	}
	return nil
}

func (mr *KafkaMapReduce) Map2(mapFunc func() (*Task, error)) error {
	task, err := mapFunc()
	if err != nil {
		log.Panic(err)
		return err
	}
	//对task发送
	taskjson, err := jsoniter.MarshalToString(task)
	if err != nil {
		log.Panic(err)
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: task.Group,
		Value: sarama.ByteEncoder(taskjson),
	}
	_, _, err = mr.producer.SendMessage(msg)

	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

func (mr *KafkaMapReduce) Reduce(topic []string, handler func(task *Task) error) error {
	// 配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0                     // specify appropriat e version
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // 未找到组消费位移的时候从哪边开始消费

	//创建消费组
	consumergroup, err := sarama.NewConsumerGroup(mr.address, "consumer-group", config)
	if err != nil {
		panic(fmt.Sprintf("new consumergroup error: %s\n", err.Error()))
	}
	defer consumergroup.Close()

	q := &ReduceHandler{
		UserHandler: handler,
	}

	for {
		err = consumergroup.Consume(context.Background(), topic, q)
		if err != nil {
			log.Panic(err)
			return err
		}
	}
}
