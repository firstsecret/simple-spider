package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSave() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)

			itemCount++

			_, err := Save(item)

			if err != nil {
				log.Printf("es client insert err: %s", err)
			}
		}

	}()

	return out
}

func Save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://106.12.10.203:9200"))

	if err != nil {
		fmt.Println("客户端连接错误")
		//panic(err)
		return "", err
	}

	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())

	if err != nil {
		fmt.Println("客户端插入错误")
		//panic(err)

		return "", err
	}

	return resp.Id, nil

	//fmt.Printf("%+v", resp)
}
