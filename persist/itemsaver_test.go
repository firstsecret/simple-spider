package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"learnGo/crawier/model"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name:       "安静的雪",
		Gender:     "女",
		Age:        "23",
		Height:     "165cm",
		Weight:     "56",
		Income:     "4234324",
		Marriage:   "未婚",
		Education:  "本科",
		Occupation: "人事/行政",
		Hokou:      "上海",
		Xinzuo:     "白羊座",
		Hourse:     "已购房",
		Car:        "未购车",
	}

	id, err := Save(profile)

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://106.12.10.203:9200"))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(
		*resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	if actual != profile{
		t.Errorf("got %v; expected %v", actual, profile)
	}
}
