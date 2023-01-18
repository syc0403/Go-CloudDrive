package mq

import "Go-CloudDrive/common"

/**
* @Description
* @Author YuCheng
* @Date 2023/1/18 22:55
**/

// TransferDate 将要写到rabbitmq的数据的结构体
type TransferDate struct {
	FileHash      string
	CurLocation   string
	DestLocation  string
	DestStoreType common.StoreType
}
