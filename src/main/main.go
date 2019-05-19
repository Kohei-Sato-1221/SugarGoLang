package main

import (
	"context"
	"fmt"
	"time"
	
	"golang.org/x/sync/semaphore"
	"gopkg.in/ini.v1"
)


func main(){
	lesson76()
}

// go ini → configファイルをよむライブラリ
func lesson77(){
	
}

type ConfigList struct{
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init(){
	cfg, _ := ConfigList{
		Port
	}
}

// Semaphore →　同時に実行されるgoroutineの数を制御する
func lesson76(){
	ctx := context.TODO()
	go longProcess(ctx)	
	go longProcess(ctx)	
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context){
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock....")
		return
	}
	
//	if err := s.Acquire(ctx, 1); err != nil{
//		fmt.Print(err)
//		return
//	}
	defer s.Release(1)
	fmt.Println("Wait......")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")	
}





