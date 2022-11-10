package repository

import (
	"context"
	"fmt"
	"time"
	"unsafe"

	"github.com/go-redis/redis/v8"
	"github.com/rs/xid"
)

const namePrefix = "EXENGINE_"

type Message redis.XMessage
type streamConnection struct {
	rdb *redis.Client
}

type StreamRepository interface {
	Ping(ctx context.Context) (string, error)
	PublishMessage(ctx context.Context, stream string, values map[string]interface{}) error
	Pendings(ctx context.Context, stream, gname string) ([]Message, error)
	Range(ctx context.Context, stream, start, stop string) ([]Message, error)
	CreateConsumerGroup(ctx context.Context, stream, gname, start string) error
	RemoveConsumerGroup(ctx context.Context, stream, gname string) error
	RemoveConsumer(ctx context.Context, stream, gname, consumer string) error
	ConsumeMessage(ctx context.Context, subject, consumersGroup, uniqueID string, t time.Duration, count int64) ([]Message, error)
	ConsumeAck(ctx context.Context, subject, consumersGroup, messageID string) error
	RemoveMessage(ctx context.Context, subject, id string) error
	Number(ctx context.Context, subject string) int64
	Info(ctx context.Context, stream string)
	LastDeliverdId(ctx context.Context, stream, gname string) string
	ReadAllMessages(ctx context.Context, subject, start string) ([]Message, error)
}

func NewStreamRepository(rdb *redis.Client) StreamRepository {
	return &streamConnection{
		rdb: rdb,
	}
}

func (c *streamConnection) Ping(ctx context.Context) (string, error) {
	return c.rdb.Ping(ctx).Result()
}

func (c *streamConnection) PublishMessage(ctx context.Context, stream string, values map[string]interface{}) error {

	if err := c.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream:       stream,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values:       values,
	}).Err(); err != nil {
		return err
	}

	fmt.Printf("\n--------- published message to %q stream. -----------------\n", stream)
	return nil
}

func (c *streamConnection) CreateConsumerGroup(ctx context.Context, stream, gname, start string) error {
	if err := c.PublishMessage(context.Background(), stream, map[string]interface{}{
		"init": "init", // This for initializing redis for reading from request stream -- This is needed for making consumer group
	}); err != nil {
		return fmt.Errorf("error while trying to run the app. Can not initialize redis stream %v", err)
	}
	return c.rdb.XGroupCreate(ctx, stream, gname, start).Err()
}

func (c *streamConnection) RemoveConsumerGroup(ctx context.Context, stream, gname string) error {
	return c.rdb.XGroupDestroy(ctx, stream, gname).Err()
}

func (c *streamConnection) RemoveConsumer(ctx context.Context, stream, gname, consumer string) error {
	return c.rdb.XGroupDelConsumer(ctx, stream, gname, consumer).Err()
}

func (c *streamConnection) ConsumeMessage(ctx context.Context, subject, consumersGroup, uniqueID string, t time.Duration, count int64) ([]Message, error) {

	entries, err := c.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    consumersGroup,
		Consumer: uniqueID,
		Streams:  []string{subject, ">"},
		Count:    count,
		Block:    t,
		NoAck:    false,
	}).Result()

	if err != nil {
		return nil, err
	}

	return *(*[]Message)(unsafe.Pointer(&(entries[0].Messages))), nil

}

func (c *streamConnection) ConsumeAck(ctx context.Context, subject, consumersGroup, messageID string) error {

	return c.rdb.XAck(ctx, subject, consumersGroup, messageID).Err()
}

func (c *streamConnection) Pendings(ctx context.Context, stream, gname string) ([]Message, error) {
	pnd, err := c.rdb.XPending(ctx, stream, gname).Result()

	if err != nil {
		fmt.Println("error")
		return nil, err
	}

	return c.Range(ctx, stream, pnd.Lower, pnd.Higher)
}

func (c *streamConnection) Range(ctx context.Context, stream, start, stop string) ([]Message, error) {
	entries, err := c.rdb.XRange(ctx, stream, start, stop).Result()
	if err != nil {
		return nil, err
	}

	return *(*[]Message)(unsafe.Pointer(&entries)), nil

}

func (c *streamConnection) Number(ctx context.Context, subject string) int64 {
	return c.rdb.XLen(ctx, subject).Val()
}

func (c *streamConnection) Info(ctx context.Context, stream string) {

	r := c.rdb.XInfoGroups(ctx, stream)

	res, err := r.Result()
	fmt.Printf("info %#v ,%v", res, err)

	// Output: "" redis: nil
	//cmd := redis.NewStringCmd(ctx, "xinfo", "stream", stream, "full")
	//err := c.rdb.Process(ctx, cmd)

	//fmt.Println("command value:", cmd.Val())
	//if err != nil {
	//	fmt.Println("Error happend in processing cmd", err, cmd)
	//} else {
	//	fmt.Println(cmd)
	//}

}

func (c *streamConnection) LastDeliverdId(ctx context.Context, stream, gname string) string {

	r := c.rdb.XInfoGroups(ctx, stream)

	groupsInfo, err := r.Result()
	//fmt.Printf("info %#v ,%v", res, err)

	if err != nil {
		return "0-0"
	}
	for _, v := range groupsInfo {
		if v.Name == gname {
			return v.LastDeliveredID
		}

	}

	return "0-0"
	// Output: "" redis: nil
	//cmd := redis.NewStringCmd(ctx, "xinfo", "stream", stream, "full")
	//err := c.rdb.Process(ctx, cmd)

	//fmt.Println("command value:", cmd.Val())
	//if err != nil {
	//	fmt.Println("Error happend in processing cmd", err, cmd)
	//} else {
	//	fmt.Println(cmd)
	//}

}

func (c *streamConnection) RemoveMessage(ctx context.Context, subject, id string) error {
	return c.rdb.XDel(ctx, subject, id).Err()
}

func (c *streamConnection) Close() error {

	return c.rdb.Close()
}

func (c *streamConnection) ReadAllMessages(ctx context.Context, subject, start string) ([]Message, error) {
	gn := namePrefix + time.Now().String()
	uid := xid.New().String()
	err := c.CreateConsumerGroup(ctx, subject, gn, start)

	if err != nil {
		return nil, err
	}
	entries, err := c.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    gn,
		Consumer: uid,
		Streams:  []string{subject, ">"},
		Count:    0,
		Block:    0,
	}).Result()

	_ = c.RemoveConsumerGroup(ctx, subject, gn)

	if err != nil {
		return nil, err
	}

	return *(*[]Message)(unsafe.Pointer(&(entries[0].Messages))), nil

}
