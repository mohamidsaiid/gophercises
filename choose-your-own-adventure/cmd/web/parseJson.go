package main

import (
	"encoding/json"
	"os"
)

func createObejct(key string, val interface{}) displayedText {
	tmp := new(displayedText)
	tmp.Arc = key
	for k, v := range val.(map[string]interface{}) {
		if k == "title" {
			tmp.Title = v.(string)
		} else if k == "story" {
			tmpArr := []string{}
			for _, vv := range v.([]interface{}) {
				tmpArr = append(tmpArr, vv.(string))
			}
			tmp.Story = tmpArr
		} else if k == "options" {
			tmpArr := []option{}
			for _, vv := range v.([]interface{}) {
				tmpStr := option{}
				for kk, vvv := range vv.(map[string]interface{}) {
					if kk == "text" {
						tmpStr.Text = vvv.(string)
					} else if kk == "arc" {
						tmpStr.Arc = vvv.(string)
					}
				}
				tmpArr = append(tmpArr, tmpStr)
			}
			tmp.Options = tmpArr
		}
	}
	return *tmp
}

func parsingJson() (map[string]int, []displayedText) {
	file, err := os.ReadFile("./cmd/web/gopher.json")
	if err != nil {
		panic(err)
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(file, &dat); err != nil {
		panic(err)
	}

	var data []displayedText
	for key, val := range dat {
		data = append(data, createObejct(key, val))
	}

	mpStoryToIdx := map[string]int{}
	for idx, val := range data {
		mpStoryToIdx[val.Arc] = idx
	}

	return mpStoryToIdx, data
}
