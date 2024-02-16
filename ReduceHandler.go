package MapReducde

import (
	"log"

	"github.com/IBM/sarama"
	jsoniter "github.com/json-iterator/go"
)

type ReduceHandler struct {
	UserHandler func(task *Task) error
}

func (r ReduceHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (r ReduceHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (r ReduceHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {

		task := &Task{}
		err := jsoniter.UnmarshalFromString(string(msg.Value), task)
		if err != nil {
			log.Panic(err)
			return err
		}
		// log.Println("task:", task)
		//用户定义UserHandler处理逻辑
		err = r.UserHandler(task)
		if err != nil {
			log.Panic(err)
			return err
		}
		session.MarkMessage(msg, "")
	}
	return nil
}
