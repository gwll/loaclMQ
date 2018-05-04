package loaclMQ

type Mq struct {
	mq map[string]chan interface{}
}

var m Mq
var ch = make(chan int, 1)
var mqLen = 10000 //队列长度
var mqFull = 9000

func init() {
	//初始化消息队列列表
	m.mq = make(map[string]chan interface{})

}

func create(topic string) {
	//	创建新的消息队列
	ch <- 1
	m.mq[topic] = make(chan interface{}, mqLen)

	<-ch

}
func Pub(topic string, contents interface{}) { //发布消息
	_, ok := m.mq[topic]
	if !ok {
		create(topic)
	}

	for len(m.mq[topic]) > mqFull {
		<-m.mq[topic]
	}

	m.mq[topic] <- contents
}
func Sub(topic string) interface{} {
	//订阅消息
	_, ok := m.mq[topic]
	if !ok {
		create(topic)
	}
	contents := <-m.mq[topic]
	return contents
}
