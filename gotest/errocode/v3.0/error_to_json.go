/**
 * @Author: xueyanghan
 * @File: error_to_json.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/7/6 18:22
 */

package errocode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

//{
//"key": "FailedOperation.InvalidUserStatus",
//"value": "用户状态异常",
//"describe": "user-center",
//"featureTypeId": 1
//},

type Item struct {
	Key           string `json:"key"`
	Value         string `json:"value"`
	FeatureTypeId int    `json:"featureTypeId"`
	Describe      string `json:"describe"`
}

var errorRangeToDesc = map[int]string{
	1:  "bcc-common",
	2:  "bcc-consortium",
	3:  "bcc-calcmodel",
	4:  "bcc-datasource",
	5:  "bcc-calctask",
	6:  "bcc-eventcenter",
	7:  "bcc-resource",
	8:  "bcc-dataconnect",
	9:  "bcc-account",
	10: "bcc-chain",
	20: "bcc-other",
}

func errorCodeToJson() error {
	var errCodes []int
	for key, _ := range ErrMessage {
		errCodes = append(errCodes, int(key))
	}
	sort.Ints(errCodes)

	var itemsMap = make(map[int][]Item)
	for _, code := range errCodes {
		item := Item{
			Key:           ErrMessage[ErrCode(code)][0],
			Value:         ErrMessage[ErrCode(code)][2],
			FeatureTypeId: 1,
			Describe:      errorRangeToDesc[code/100],
		}
		//fmt.Printf("code: %d, descIndex: %d,  item: %+v\n", code, code/100, item)
		itemsMap[code/100] = append(itemsMap[code/100], item)
	}

	for code, items := range itemsMap {
		itemsJson, err := json.Marshal(items)
		if err != nil {
			return err
		}

		fmt.Println(string(itemsJson))
		err = ioutil.WriteFile(fmt.Sprintf("%s.json", errorRangeToDesc[code]), itemsJson, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
