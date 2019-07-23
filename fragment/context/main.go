package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func timeoutHandler() {
	// 超时后会自动调用 cancel(), 取消上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second)) // 这句等同于上面那句

	//go doTimeOutStuff(ctx)
	go doTimeOutStuff(ctx)

	time.Sleep(10 * time.Second)

	cancel() // 取消上下文, 会释放与其关联的资源
}

func someHandler() {
	// 这个不设置超时, 在资源释放后需要手动调用 cancel()
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

// 每秒 work 一下，同时会判断 ctx 是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logger.Printf("done")
			return
		default:
			logger.Printf("work")
		}
	}
}

// 每秒获取 deadline, 如果 deadline 已过, 返回, 如果没过等同于 doStuff
func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok { //获取deadline
			logger.Printf("deadline set")
			if time.Now().After(deadline) {
				logger.Printf(ctx.Err().Error())
				return
			}
		}

		select {
		case <-ctx.Done():
			logger.Printf("done")
			return
		default:
			logger.Printf("work")
		}
	}
}

func main() {
	logger = log.New(os.Stdout, "[ctx] ", log.Ltime)
	timeoutHandler()
}
