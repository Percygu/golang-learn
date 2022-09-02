package main

import (
	"fmt"
	"time"
)

// 实现一个简单的协程池(task, worker，channel,pool)

// 1. Task 定义task结构，成员为一个执行函数
type Task struct {
	f func() error
}

func NewTask(function func() error) *Task {
	return &Task{
		f: function,
	}
}

func (t *Task) Exec() {
	t.f()
}

// 2. 定义一个poll结构
type Pool struct {
	//JobChannel  chan *Task
	TaskChannel chan *Task
	workerNum   int
}

func NewPool(taskLen int, cap int) *Pool {
	return &Pool{
		// JobChannel:  make(chan *Task, jobLen),
		TaskChannel: make(chan *Task, taskLen),
		workerNum:   cap,
	}
}

// 协程池内每个woker处理任务
func (p *Pool) Worker(workID int) {
	for task := range p.TaskChannel {
		task.Exec()
		fmt.Printf("worker%d 执行完了一个任务!\n", workID)
	}
}

// 启动协程池
func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.Worker(i)
	}
}

func (p *Pool) AddTask(t *Task) {
	p.TaskChannel <- t
}

func main() {
	t := NewTask(func() error {
		fmt.Printf("当前时间是：%v\n", time.Now())
		return nil
	})
	p := NewPool(10, 5)

	go p.Run()

	for {
		p.AddTask(t)
		time.Sleep(time.Second)
	}
}
