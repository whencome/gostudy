package main

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Etcdx struct {
	endpoints []string
	timeout   int64
	client    *clientv3.Client
}

func NewEtcdx(endpoints []string, timeout int64) *Etcdx {
	if endpoints == nil || len(endpoints) == 0 {
		log.Panic("empty etcd endpoints")
	}
	if timeout <= 0 {
		timeout = int64(3)
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second * time.Duration(timeout),
	})
	if err != nil {
		log.Panicf("connect etcd failed: %s", err)
	}
	return &Etcdx{
		endpoints: endpoints,
		timeout:   timeout,
		client:    cli,
	}
}

func (x *Etcdx) Close() error {
	return x.client.Close()
}

// action对应Event_EventType， 0-PUT, 1-DELETE
func (x *Etcdx) Watch(key string, handleFunc func(action int32, key, newValue, oldValue string) error) {
	ch := x.client.Watch(context.Background(), key)
	for resp := range ch {
		for _, ev := range resp.Events {
			if handleFunc == nil {
				continue
			}
			// err := handleFunc(int32(ev.Type), key, string(ev.Kv.Value), string(ev.PrevKv.Value))
			newValue := ""
			oldValue := ""
			if ev.Kv != nil {
				newValue = string(ev.Kv.Value)
			}
			if ev.PrevKv != nil {
				oldValue = string(ev.PrevKv.Value)
			}
			err := handleFunc(int32(ev.Type), key, newValue, oldValue)
			if err != nil {
				log.Printf("handle %s change [type=%d, new value = %s, old value = %s] failed: %s\n", key, ev.Type, ev.Kv.Value, ev.PrevKv.Value, err)
			}
		}
	}
	log.Printf("------------- WATCH %s OVER ----------------\n", key)
}

func (x *Etcdx) Put(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(x.timeout))
	_, err := x.client.Put(ctx, key, value)
	cancel()
	return err
}

func (x *Etcdx) Get(key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(x.timeout))
	resp, err := x.client.Get(ctx, key)
	cancel()
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, kv := range resp.Kvs {
		result = append(result, string(kv.Value))
	}
	return result, nil
}

func (x *Etcdx) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(x.timeout))
	_, err := x.client.Delete(ctx, key)
	cancel()
	return err
}
