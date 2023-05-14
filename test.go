package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func main() {
	log.SetFlags(0)

	var (
		r map[string]interface{}
		//wg sync.WaitGroup
	)

	es, _ := elasticsearch.NewClient(
		elasticsearch.Config{},
	)
	//// 1. Get cluster info
	////
	//res, _ := es.Info()
	//defer res.Body.Close()
	//
	//_ = json.NewDecoder(res.Body).Decode(&r)
	//
	//for i, title := range []string{"《西游记》，又称《西游释厄传》，简称《西游》，是中国古代明朝的第一部浪漫主义章回体长篇神魔小说，全书58.5万字（世德堂本），共100回，中国四大名著之一、四大奇书之一。成书于16世纪明朝中叶，一般认为作者是明朝的吴承恩。书中讲述唐三藏与徒弟孙悟空、猪八戒和沙悟净等师徒四人前往西天取经的故事，表现了惩恶扬善的古老主题，也有观点认为《西游记》是暗讽权力官场的讽刺小说。 ", "因《西游记》的传颂，明清之际吴元泰、吴政泰、余象斗等因而又据佛、道两教之有关戏曲杂剧和神话传说，撰写《东游记》、《南游记》和《北游记》，再加上杨志和之另本《西游记》，合称《四游记》。\n\n《西游记》自问世以来，在中国及世界各地广为流传，被翻译成多种语言。在中国，乃至亚洲部分地区《西游记》家喻户晓，其中孙悟空、唐僧、猪八戒、沙僧等人物和“大闹天宫”、“三打白骨精”、“孙悟空三借芭蕉扇”等故事尤其为人熟悉。几百年来，西游记被改编成各种地方戏曲、电影、电视剧、动画片、漫画等，版本繁多。 "} {
	//	wg.Add(1)
	//
	//	go func(i int, title string) {
	//		defer wg.Done()
	//
	//		// Build the request body.
	//		data, _ := json.Marshal(struct {
	//			Title string `json:"title"`
	//		}{Title: title})
	//		req := esapi.IndexRequest{
	//			Index:      "upload",
	//			DocumentID: strconv.Itoa(i + 1),
	//			Body:       bytes.NewReader(data),
	//			Refresh:    "true",
	//		}
	//
	//		// Perform the request with the client.
	//		res, _ := req.Do(context.Background(), es)
	//		defer res.Body.Close()
	//	}(i, title)
	//}
	//wg.Wait()
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "西游记",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, _ := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("upload"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Println("内容太多了！")
		} else {
			fmt.Println("Index Not found")
		}
	}

	_ = json.NewDecoder(res.Body).Decode(&r)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	fmt.Println(int64(len(r["hits"].(map[string]interface{})["hits"].([]interface{}))))

}
