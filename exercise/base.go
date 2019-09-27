package exercise

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/papillonyi/thor/etcd"
	"go.etcd.io/etcd/clientv3"
	"log"
	"reflect"
)

type metadata struct {
	Namespace string            `json:"namespace" binding:"required"`
	Name      string            `json:"name" binding:"required"`
	Labels    map[string]string `json:"labels"`
}

type StoreManager interface {
	GetListByNamespace(namespace string) (actionList []Action, err error)
	Add() (err error)
	getKey() (key string, err error)
	getPrefix() (key string)
}

type storeData struct {
	Key   string       `json:"key" binding:"required"`
	Value StoreManager `json:"value" binding:"required"`
}

func getListByNamespace(manager StoreManager, namespace string) (storeList []interface{}, err error) {
	prefix := manager.getPrefix() + namespace

	resp, err := etcd.Client.KV.Get(context.TODO(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Print(err)
		return
	}

	for _, item := range resp.Kvs {
		value := make(map[string]interface{})
		fmt.Println(string(item.Key), string(item.Value))
		err = json.Unmarshal(item.Value, &value)

		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("done value %s", reflect.TypeOf(value))
		storeList = append(storeList, value["value"])
	}
	return
}
