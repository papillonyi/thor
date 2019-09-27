package etcd

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"log"
	"time"
)

var (
	Client *clientv3.Client
	err    error
)

func Setup() {

	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"47.100.46.22:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	log.Println("connect to etcd")
}
