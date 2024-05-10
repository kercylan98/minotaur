package brokers_test

import (
	"context"
	"github.com/kercylan98/minotaur/toolkit/nexus"
	"github.com/kercylan98/minotaur/toolkit/nexus/brokers"
	"github.com/kercylan98/minotaur/toolkit/nexus/events"
	"github.com/kercylan98/minotaur/toolkit/nexus/queues"
	"os"
	"testing"
	"time"
)

func TestIndependentGoroutine_NotBindAndUnBind(t *testing.T) {
	ig, _, _ := brokers.NewIndependentGoroutine[int, string](
		func(index int) nexus.Queue[int, string] {
			return queues.NewNonBlockingRW[int, string](index, 1024, 1024)
		}, func(handler nexus.EventExecutor) {
			handler.Exec()
		},
		brokers.NewIndependentGoroutineOptions[int, string]().
			WithQueueCreatedHook(func(topic string, queue nexus.Queue[int, string], queueNum int) {
				t.Log("queue created", topic, "queue num", queueNum)
			}).
			WithQueueClosedHook(func(topic string, queue nexus.Queue[int, string], queueNum int) {
				t.Log("queue closed", topic, "queue num", queueNum)
			}).
			WithQueueBindCounterChangedHook(func(topic string, count int) {
				t.Log("queue bind counter changed", topic, count)
			}),
	)

	// 没有使用 bind 和 unBind 的情况下，每一个消息执行完成后且没有消息时，队列会自动关闭

	go func() {
		time.Sleep(time.Second)
		if err := ig.Publish("test", events.Synchronous[int, string](func(ctx context.Context) {
			t.Log("test")
		})); err != nil {
			panic(err)
		}

		ig.Close()
	}()

	ig.Run()

	t.Log("done")
}

func TestIndependentGoroutine_BindButNotUnBind(t *testing.T) {
	ig, bind, _ := brokers.NewIndependentGoroutine[int, string](
		func(index int) nexus.Queue[int, string] {
			return queues.NewNonBlockingRW[int, string](index, 1024, 1024)
		}, func(handler nexus.EventExecutor) {
			handler.Exec()
		},
		brokers.NewIndependentGoroutineOptions[int, string]().
			WithQueueCreatedHook(func(topic string, queue nexus.Queue[int, string], queueNum int) {
				t.Log("queue created", topic, "queue num", queueNum)
			}).
			WithQueueClosedHook(func(topic string, queue nexus.Queue[int, string], queueNum int) {
				t.Log("queue closed", topic, "queue num", queueNum)
			}).
			WithQueueBindCounterChangedHook(func(topic string, count int) {
				t.Log("queue bind counter changed", topic, count)
			}),
	)

	// 使用 bind 但是没有使用 unBind 的情况下，每一个消息执行完成后且没有消息时，队列不会自动关闭，由于关闭 broker 时会等待所有队列关闭，所以会一直阻塞
	bind("test")
	go func() {
		time.Sleep(time.Second)
		if err := ig.Publish("test", events.Synchronous[int, string](func(ctx context.Context) {
			t.Log("test")
		})); err != nil {
			panic(err)
		}

		go func() {
			time.Sleep(time.Second * 3)
			t.Log("timeout, exit")
			os.Exit(0)
		}()
		ig.Close()
	}()

	ig.Run()

	t.Fatal("not should be here")
}

func TestIndependentGoroutine_BindAndUnBind(t *testing.T) {
	ig, bind, unBind := brokers.NewIndependentGoroutine[int, string](
		func(index int) nexus.Queue[int, string] {
			return queues.NewNonBlockingRW[int, string](index, 1024, 1024)
		}, func(handler nexus.EventExecutor) {
			handler.Exec()
		},
		brokers.NewIndependentGoroutineOptions[int, string]().
			WithQueueCreatedHook(func(topic string, queue nexus.Queue[int, string], queueNum int) {
				t.Log("queue created", topic, "queue num", queueNum, queue.GetId())
			}).
			WithQueueClosedHook(func(topic string, queue nexus.Queue[int, string], queueNum int) {
				t.Log("queue closed", topic, "queue num", queueNum, queue.GetId())
			}).
			WithQueueBindCounterChangedHook(func(topic string, count int) {
				t.Log("queue bind counter changed", topic, count)
			}),
	)

	// 使用 bind 和 unBind 的情况下，每一个消息执行完成后且没有消息时，队列会自动关闭
	bind("test")
	go func() {
		time.Sleep(time.Second)
		if err := ig.Publish("test", events.Synchronous[int, string](func(ctx context.Context) {
		})); err != nil {
			panic(err)
		}

		unBind("test")

		time.Sleep(time.Second)
		if err := ig.Publish("test", events.Synchronous[int, string](func(ctx context.Context) {
		})); err != nil {
			panic(err)
		}

		ig.Close()
	}()

	ig.Run()

	t.Log("done")
}
