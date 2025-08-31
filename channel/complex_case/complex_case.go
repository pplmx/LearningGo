package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	ID    int
	Name  string
	Stage int
}

type Control struct {
	Cmd string
}

// ---------------------------
// 生产者
// ---------------------------
func producer(ctx context.Context, id int, out chan<- Task, ctrl chan<- Control) {
	defer close(out) // 生产者完成后关闭输出通道

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}

		task := Task{ID: id*10 + i, Name: fmt.Sprintf("Task-%d-%d", id, i)}
		fmt.Printf("[Producer %d] Send %v\n", id, task)

		select {
		case out <- task:
		case <-ctx.Done():
			return
		}

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	}

	// 发送控制信号
	select {
	case ctrl <- Control{Cmd: fmt.Sprintf("Producer %d done", id)}:
	case <-ctx.Done():
	}
}

// ---------------------------
// 流水线阶段 - 改进版
// ---------------------------
func stage(ctx context.Context, name string, in <-chan Task, out chan<- Task, async chan<- string) {
	defer func() {
		if out != nil {
			close(out)
		}
	}()

	for {
		select {
		case task, ok := <-in:
			if !ok {
				return // 输入通道关闭，退出
			}

			task.Stage++
			fmt.Printf("[%s] Processing %v\n", name, task)

			// 模拟处理时间
			select {
			case <-time.After(150 * time.Millisecond):
			case <-ctx.Done():
				return
			}

			// 发送异步消息
			select {
			case async <- fmt.Sprintf("%s completed stage %d for task %d", name, task.Stage, task.ID):
			default: // 非阻塞发送
			}

			// 传递到下一阶段
			if out != nil {
				select {
				case out <- task:
				case <-ctx.Done():
					return
				}
			}

		case <-ctx.Done():
			return
		}
	}
}

// ---------------------------
// 消费者 - 改进版
// ---------------------------
func consumer(ctx context.Context, id int, taskCh <-chan Task, ctrlCh <-chan Control, async chan<- string, completedTasks *int64, totalTasks int64, cancel context.CancelFunc) {
	for {
		select {
		case task, ok := <-taskCh:
			if !ok {
				fmt.Printf("[Consumer %d] Task channel closed\n", id)
				return
			}
			fmt.Printf("[Consumer %d] Received %v\n", id, task)

			// 非阻塞发送异步消息
			select {
			case async <- fmt.Sprintf("[Consumer %d] processed task %d", id, task.ID):
			default:
			}

			// 增加完成任务计数
			completed := atomic.AddInt64(completedTasks, 1)
			fmt.Printf("[Consumer %d] Progress: %d/%d tasks completed\n", id, completed, totalTasks)

			// 检查是否所有任务都已完成
			if completed >= totalTasks {
				fmt.Printf("[Consumer %d] All tasks completed! Initiating graceful shutdown...\n", id)
				// 延迟一点时间让异步消息处理完
				go func() {
					time.Sleep(500 * time.Millisecond)
					cancel()
				}()
			}

		case ctrl, ok := <-ctrlCh:
			if !ok {
				fmt.Printf("[Consumer %d] Control channel closed\n", id)
				continue
			}
			fmt.Printf("[Consumer %d] Control received: %v\n", id, ctrl)

		case <-ctx.Done():
			fmt.Printf("[Consumer %d] Context cancelled\n", id)
			return
		}
	}
}

// ---------------------------
// 异步处理 worker - 改进版
// ---------------------------
func asyncWorker(ctx context.Context, id int, asyncCh <-chan string) {
	for {
		select {
		case msg, ok := <-asyncCh:
			if !ok {
				fmt.Printf("[AsyncWorker %d] Channel closed\n", id)
				return
			}
			fmt.Printf("[AsyncWorker %d] %s\n", id, msg)

			select {
			case <-time.After(50 * time.Millisecond):
			case <-ctx.Done():
				return
			}

		case <-ctx.Done():
			fmt.Printf("[AsyncWorker %d] Context cancelled\n", id)
			return
		}
	}
}

// ---------------------------
// 主函数 - 改进版
// ---------------------------
func main() {
	rand.Seed(time.Now().UnixNano())

	// 使用 context 进行优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 计算总任务数
	const producerCount = 3
	const tasksPerProducer = 5
	const totalTasks = producerCount * tasksPerProducer
	var completedTasks int64 // 完成任务计数器

	fmt.Printf("🚀 Starting pipeline with %d producers, %d tasks each (total: %d tasks)\n",
		producerCount, tasksPerProducer, totalTasks)

	// 创建 channels
	taskCh := make(chan Task, 10)
	ctrlCh := make(chan Control, 5)
	asyncCh := make(chan string, 50) // 增大缓冲区

	var wg sync.WaitGroup

	// 启动多个生产者 (每个生产者有自己的输出通道)
	producerChannels := make([]chan Task, producerCount)
	for i := 0; i < producerCount; i++ {
		producerChannels[i] = make(chan Task, 5)
		wg.Add(1)
		go func(id int, ch chan Task) {
			defer wg.Done()
			producer(ctx, id+1, ch, ctrlCh)
		}(i, producerChannels[i])
	}

	// 合并多个生产者的输出
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(taskCh)

		var mergeWg sync.WaitGroup
		for _, pCh := range producerChannels {
			mergeWg.Add(1)
			go func(ch <-chan Task) {
				defer mergeWg.Done()
				for task := range ch {
					select {
					case taskCh <- task:
					case <-ctx.Done():
						return
					}
				}
			}(pCh)
		}
		mergeWg.Wait()
	}()

	// 创建流水线阶段的 channels
	stage1Ch := make(chan Task, 10)
	stage2Ch := make(chan Task, 10)
	stage3Ch := make(chan Task, 10)

	// 启动流水线阶段
	wg.Add(1)
	go func() {
		defer wg.Done()
		stage(ctx, "Stage1-Load", taskCh, stage1Ch, asyncCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stage(ctx, "Stage2-Compress", stage1Ch, stage2Ch, asyncCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stage(ctx, "Stage3-Upload", stage2Ch, stage3Ch, asyncCh)
	}()

	// 启动消费者
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			consumer(ctx, id, stage3Ch, ctrlCh, asyncCh, &completedTasks, totalTasks, cancel)
		}(i)
	}

	// 启动异步 workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			asyncWorker(ctx, id, asyncCh)
		}(i)
	}

	// 等待所有任务完成或超时
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("✅ All tasks completed successfully!")
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Println("⏰ Execution timeout")
		} else {
			fmt.Printf("🎯 Graceful shutdown completed! Final count: %d/%d tasks\n",
				atomic.LoadInt64(&completedTasks), totalTasks)
		}
	}

	// 关闭剩余 channels
	close(ctrlCh)
	close(asyncCh)

	// 给一点时间让最后的消息处理完
	time.Sleep(200 * time.Millisecond)
}
