/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

// G804
var save_count = 0

type FactContract struct {
}

// 存证对象
type Fact struct {
	FileHash string `json:"fileHash"`
	FileName string `json:"fileName"`
	Time     int32  `json:"time"`
}

// 新建存证对象
func NewFact(fileHash string, fileName string, time int32) *Fact {
	fact := &Fact{
		FileHash: fileHash,
		FileName: fileName,
		Time:     time,
	}
	return fact
}

func (f *FactContract) InitContract() protogo.Response {
	// G805
	go func() {
		//something
	}()
	return sdk.Success([]byte("Init contract success"))
}

func (f *FactContract) UpgradeContract() protogo.Response {
	// G801
	rand.Int()
	return sdk.Success([]byte("Upgrade contract success"))
}

func (f *FactContract) InvokeContract(method string) protogo.Response {
	switch method {
	case "save":
		return f.save()
	case "findByFileHash":
		return f.findByFileHash()
	default:
		return sdk.Error("invalid method")
	}
}

func (f *FactContract) save() protogo.Response {
	params := sdk.Instance.GetArgs()

	// 获取参数
	fileHash := string(params["file_hash"])
	fileName := string(params["file_name"])
	// G802
	timeStr := time.Now().String()
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		msg := "time is [" + timeStr + "] not int"
		sdk.Instance.Errorf(msg)
		return sdk.Error(msg)
	}

	// 构建结构体
	fact := NewFact(fileHash, fileName, int32(time))

	// 序列化
	factBytes, err := json.Marshal(fact)
	if err != nil {
		return sdk.Error(fmt.Sprintf("marshal fact failed, err: %s", err))
	}
	// 发送事件
	sdk.Instance.EmitEvent("topic_vx", []string{fact.FileHash, fact.FileName})

	// 存储数据
	err = sdk.Instance.PutStateByte("fact_bytes", fact.FileHash, factBytes)
	if err != nil {
		return sdk.Error("fail to save fact bytes")
	}

	save_count++

	sdk.Instance.PutStateByte("fact_stat", fact.FileHash, []byte(fmt.Sprintf("%d", save_count)))

	// 返回结果
	return sdk.Success([]byte(fact.FileName + fact.FileHash))

}

func (f *FactContract) findByFileHash() protogo.Response {
	// 获取参数
	fileHash := string(sdk.Instance.GetArgs()["file_hash"])

	// 查询结果
	result, err := sdk.Instance.GetStateByte("fact_bytes", fileHash)
	if err != nil {
		return sdk.Error("failed to call get_state")
	}

	// 反序列化
	var fact Fact
	if err = json.Unmarshal(result, &fact); err != nil {
		return sdk.Error(fmt.Sprintf("unmarshal fact failed, err: %s", err))
	}

	// 返回结果
	return sdk.Success(result)
}

func main() {
	err := sandbox.Start(new(FactContract))
	if err != nil {
		log.Fatal(err)
	}
}
