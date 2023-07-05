/*
Copyright 2022 The deepauto-io LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package redisdb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	single      *redis.Client
	cluster     *redis.ClusterClient
	mux         *sync.RWMutex
	clusterMode bool
}

func New(opt Config) *Redis {
	redisClient := &Redis{}
	if len(opt.Address) == 1 {
		redisClient.single = redis.NewClient(&redis.Options{
			Addr:         opt.Address[0],
			DB:           opt.DB,
			Password:     opt.Password,
			PoolSize:     opt.PoolSize,
			DialTimeout:  opt.DailTimeOut,
			ReadTimeout:  opt.ReadTimeOut,
			WriteTimeout: opt.WriteTimeout,
		})
		_, err := redisClient.single.Ping(context.Background()).Result()
		if err != nil {
			panic(fmt.Errorf("open redis service: %s", err.Error()))
		}
		redisClient.mux = new(sync.RWMutex)
		redisClient.clusterMode = false
		return redisClient
	}
	redisClient.cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        opt.Address,
		Password:     opt.Password,
		PoolSize:     opt.PoolSize,
		DialTimeout:  opt.DailTimeOut,
		ReadTimeout:  opt.ReadTimeOut,
		WriteTimeout: opt.WriteTimeout,
	})
	redisClient.clusterMode = true
	redisClient.mux = new(sync.RWMutex)
	_, err := redisClient.cluster.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("open redis service: %s", err.Error()))
	}
	return redisClient
}

func (r *Redis) Set(k, v string, t time.Duration) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Set(context.Background(), k, v, t).Err()
	}
	return r.single.Set(context.Background(), k, v, t).Err()
}

func (r *Redis) Get(k string) string {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Get(context.Background(), k).Val()
	}
	return r.single.Get(context.Background(), k).Val()
}

func (r *Redis) HSet(k, f string, v interface{}) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.HSet(context.Background(), k, f, v).Err()
	}
	return r.single.HSet(context.Background(), k, f, v).Err()
}

func (r *Redis) HGet(k, f string) (string, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		strObj := r.cluster.HGet(context.Background(), k, f)
		if strObj.Err() != nil && strObj.Err() != redis.Nil {
			return "", strObj.Err()
		}
		return strObj.Val(), nil
	}
	strObj := r.single.HGet(context.Background(), k, f)
	if strObj.Err() != nil && strObj.Err() != redis.Nil {
		return "", strObj.Err()
	}
	return strObj.Val(), nil
}

func (r *Redis) HGetAll(k string) (map[string]string, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.HGetAll(context.Background(), k).Result()
	}
	return r.single.HGetAll(context.Background(), k).Result()
}

func (r *Redis) HDel(k, f string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.HDel(context.Background(), k, f).Err()
	}
	return r.single.HDel(context.Background(), k, f).Err()
}

func (r *Redis) Expire(k string, t time.Duration) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Expire(context.Background(), k, t).Err()
	}
	return r.single.Expire(context.Background(), k, t).Err()
}

func (r *Redis) HSetTTL(k, field string, value interface{}, t time.Duration) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		if err := r.cluster.HSet(context.Background(), k, field, value).Err(); err != nil {
			return err
		}
		return r.cluster.Expire(context.Background(), k, t).Err()
	}
	if err := r.single.HSet(context.Background(), k, field, value).Err(); err != nil {
		return err
	}
	return r.single.Expire(context.Background(), k, t).Err()
}

func (r *Redis) Keys(k string) []string {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Keys(context.Background(), k).Val()
	}
	return r.single.Keys(context.Background(), k).Val()
}

func (r *Redis) Del(k string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Del(context.Background(), k).Err()
	}
	return r.single.Del(context.Background(), k).Err()
}

func (r *Redis) ZRemRangeByScore(k, min, max string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.ZRemRangeByScore(context.Background(), k, min, max).Err()
	}
	return r.single.ZRemRangeByScore(context.Background(), k, min, max).Err()
}

func (r *Redis) ZRange(k string, start, end int64) []string {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.ZRange(context.Background(), k, start, end).Val()
	}
	return r.single.ZRange(context.Background(), k, start, end).Val()
}

func (r *Redis) ZAddNx(k string, score, member float64) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.ZAddNX(context.Background(), k, &redis.Z{Score: float64(score), Member: float64(member)}).Err()
	}
	return r.single.ZAddNX(context.Background(), k, &redis.Z{Score: float64(score), Member: float64(member)}).Err()
}

func (r *Redis) Incr(k string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Incr(context.Background(), k).Err()
	}
	return r.single.Incr(context.Background(), k).Err()
}

func (r *Redis) Watch(ctx context.Context, key string) <-chan interface{} {
	var pubsub *redis.PubSub
	if r.clusterMode {
		pubsub = r.cluster.PSubscribe(context.Background(), fmt.Sprintf("__key*__:%s", key))
	} else {
		pubsub = r.single.PSubscribe(context.Background(), fmt.Sprintf("__key*__:%s", key))
	}

	res := make(chan interface{})
	go func() {
		for {
			select {
			case msg := <-pubsub.Channel():
				op := msg.Payload
				res <- op
			case <-ctx.Done():
				_ = pubsub.Close()
				close(res)
				return
			}
		}
	}()
	return res
}

func (r *Redis) LPush(k, f string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.LPush(context.Background(), k, f).Err()
	}
	return r.single.LPush(context.Background(), k, f).Err()
}

func (r *Redis) BRPop(timeout time.Duration, k string) ([]string, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.BRPop(context.Background(), timeout, k).Result()
	}
	return r.single.BRPop(context.Background(), timeout, k).Result()
}

func (r *Redis) LockNx(k, v string, t time.Duration) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.SetNX(context.Background(), k, v, t).Err()
	}
	return r.single.SetNX(context.Background(), k, v, t).Err()
}

func (r *Redis) UnLockNx(k string) error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Del(context.Background(), k).Err()
	}
	return r.single.Del(context.Background(), k).Err()
}

func (r *Redis) Close() error {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.clusterMode {
		return r.cluster.Close()
	}
	return r.single.Close()
}
