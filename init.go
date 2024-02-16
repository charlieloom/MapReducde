package MapReducde

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func InitKMr(address []string) (KafkaMapReduce, error) {
	// 配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(address, config)

	mr := KafkaMapReduce{
		producer: p,
		address:  address,
	}
	if err != nil {
		log.Printf("new sync producer error : %s \n", err.Error())
		return mr, err
	}
	return mr, nil
}
