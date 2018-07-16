package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"time"
)

func main()  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll  // 等待服务器所有副本都保存成功后响应
	config.Producer.Partitioner  = sarama.NewRandomPartitioner // 随机的分区策略

	// 是否等待成功和失败后的响应。注意，只有在上面 RequireAcks 设置不是 NoResponse 时才能起作用
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors    = true

	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if nil != err {
		fmt.Println("sarama.NewSyncProducer failed")
		return
	}
	defer client.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "tinytiger-20180714"
	msg.Value = sarama.StringEncoder(fmt.Sprintf("[%s] a good message", time.Now()))

	// pid 是分区号，offset 是该队列的偏移量
	pid, offset, err := client.SendMessage(msg)
	if nil != err {
		fmt.Println("client.SendMessage failed")
		return
	}

	fmt.Printf("pid:%v, offset:%v\n", pid, offset)
}
