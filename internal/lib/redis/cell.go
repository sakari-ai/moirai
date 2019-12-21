package redis

import (
	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
	redis "gopkg.in/redis.v5"
)

// redisCell implements https://github.com/brandur/redis-cell
type Cell struct {
	client         *redis.ClusterClient
	maxBust        int
	countPerPeriod int
	periodInSecond int
	quantity       int
}

func NewRedisCell(client *redis.ClusterClient, opts ...int) *Cell {
	cell := &Cell{client: client}
	for k, v := range opts {
		if k == 0 {
			cell.maxBust = v
		} else if k == 1 {
			cell.countPerPeriod = v
		} else if k == 2 {
			cell.periodInSecond = v
		} else {
			cell.quantity = v
		}
	}
	if cell.quantity == 0 {
		cell.quantity = 1
	}
	return cell
}

func (cell *Cell) Rate(key string) (allowed bool, remain int) {
	allowed = true
	var args = []interface{}{"CL.THROTTLE", key, cell.maxBust, cell.countPerPeriod, cell.periodInSecond, cell.quantity}
	cmd := redis.NewSliceCmd(args...)
	if cell.client != nil {
		err := cell.client.Process(cmd)
		if err != nil {
			log.Error("redis-cell command is failed", field.Error(err))
		}
		value := cmd.Val()
		if value[0].(int64) > 0 {
			allowed = false
		}
		remain = int(value[2].(int64))
	}
	return
}
