package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"
	"unicode"

	"github.com/Shopify/sarama"
	"github.com/patrickmn/go-cache"
)

func consumer_test() {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"172.16.248.136:9092", "172.16.248.37:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()
	p, _ := consumer.Partitions("plutus_biz_data")
	fmt.Println(p)
	partition_consumer, err := consumer.ConsumePartition("plutus_biz_data", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}

func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			fmt.Println(r)
			return true
		}
	}
	return false
}

var chinesePattern = cache.New(72*time.Hour, 72*time.Hour)
var i = 0

func operateMap() {
	for ; i < 1000; i++ {
		chinesePattern.SetDefault("sadasdasdasdasdasdasdasdasdasdasdasdasdasdasdasd"+strconv.Itoa(i), 1)
	}
}

type hehe struct {
	data string
	sync.Mutex
}

func (h *hehe) funcA() {
	h.data = "aa"
	fmt.Println("a")

}
func (h hehe) funcB() {
	h.data = "bb"
	fmt.Println("b")

}

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	//s.data = str
}

type haha struct {
	a int
}

func x_y2_XY_stub(s string) string {
	length := len(s)
	var result []byte
	var inSpace bool
	var flag = false
	for i := 0; i < length; i++ {
		if s[i] == '{' {
			inSpace = true
			result = append(result, s[i])
			continue
		}
		if s[i] == '}' {
			inSpace = false
			result = append(result, s[i])
			continue
		}
		if inSpace {
			if s[i] == '_' {
				flag = true
				continue
			}
			if flag == true {
				result = append(result, s[i]+'A'-'a')
				flag = false
				continue
			}
			result = append(result, s[i])
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}

func XY2_x_y_stub(s string) string {
	length := len(s)
	inSpace := false
	if length == 0 {
		return ""
	}
	result := make([]byte, 0, length)
	for index := 0; index < length; index++ {
		if s[index] == '{' {
			inSpace = true
			result = append(result, s[index])
			continue
		}
		if s[index] == '}' {
			inSpace = false
			result = append(result, s[index])
			continue
		}
		if inSpace {
			if (s[index] >= 'A' && s[index] <= 'Z') && (s[index-1] >= 'a' && s[index-1] <= 'z') &&
				(s[index+1] >= 'a' && s[index+1] <= 'z') {
				result = append(result, '_')
				result = append(result, s[index]-'A'+'a')
				continue
			}
		}
		result = append(result, s[index])
	}
	return string(result)
}

func main() {
	fmt.Println(x_y2_XY_stub("/v1.0/get_contry/{contry_codA_aaa}"))
	fmt.Println(XY2_x_y_stub("/v1.0/stuDent/{enAbc}/{sdsd_sdsdAbc}/{enAac}/aaBB/{aAaaaBB}"))
}
