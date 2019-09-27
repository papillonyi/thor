package exercise

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/papillonyi/thor/etcd"
	"log"
)

type Action struct {
	Type     string   `json:"type"`
	Metadata metadata `json:"metadata" binding:"required"`
}

func (action *Action) Add() (err error) {
	key, err := action.getKey()
	value, err := action.getValue()

	_, err = etcd.Client.Put(context.TODO(), key, value)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("put key %s value %s", key, value)
	return
}

func (action Action) GetListByNamespace(namespace string) (actionList []Action, err error) {
	storeList, err := getListByNamespace(&Action{}, namespace)
	for _, item := range storeList {
		//var action Action
		if err := mapstructure.Decode(item, &action); err != nil {
			log.Print(err)
		}
		log.Printf("get done item %s", item)
		actionList = append(actionList, action)
	}
	return
}

func GetListByNamespace(StoreManager, namespace string) (actionList []Action, err error) {
	storeList, err := getListByNamespace(&Action{}, namespace)
	for _, item := range storeList {
		var action Action
		if err := mapstructure.Decode(item, &action); err != nil {
			log.Print(err)
			//return
		}
		log.Printf("get done item %s", item)
		actionList = append(actionList, action)
	}
	return
}

//func GetActionListByNamespace(namespace string) (actionList []*Action, err error) {
//	key := "/action/" + namespace
//
//	log.Println(key)
//
//	resp, err := etcd.Client.KV.Get(context.TODO(), key, clientv3.WithPrefix())
//	if err != nil {
//		logging.Error(err)
//		return
//	}
//
//	for _,item := range resp.Kvs {
//		a := &Action{}
//		fmt.Println(string(item.Key),string(item.Value))
//		err = json.Unmarshal(item.Value,)
//
//		if err != nil {
//			logging.Error(err)
//			return
//		}
//
//		fmt.Println(a)
//		actionList = append(actionList, a)
//	}
//	return
//}

func (action *Action) getValue() (value string, err error) {
	data := make(map[string]interface{})
	key, err := action.getKey()
	if err != nil {
		log.Print(err)
		return
	}

	data["key"] = key
	data["value"] = action
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		return
	}

	return string(jsonData), nil
}

func (_ *Action) getPrefix() string {
	return "/action/"
}

func (action *Action) getKey() (key string, err error) {
	return fmt.Sprintf("/action/%s/%s", action.Metadata.Namespace, action.Metadata.Name), nil
}
