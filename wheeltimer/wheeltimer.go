package wheeltimer

import (
	"log"
	"time"
)

const (
	WORKER_STATE_INIT     = 0                      // 定时任务[补始]状态
	WORKER_STATE_STARTED  = 1                      // 定时任务[已经]状态
	WORKER_STATE_SHUTDOWN = 2                      // 定时任务[关闭]状态
	MIN_DELAY_UNIT        = 100 * time.Millisecond // 最小延迟毫秒数单位
)

var (
	tick             *time.Ticker      // 时钟
	tickDuration     time.Duration     // 每tick一次的时间间隔, 每tick一次就会到达下一个槽位
	ticksPerWheel    int64             // 轮中的slot数
	currentTickIndex int64             // 当前槽位
	rounds           int64             // 已转的轮数
	ticks            int64             // 目前已经 tick 过的次数
	startTime        int64             // 定时任务启动时间
	status           int               // 定时器状态
	tasks            map[int64][]*Task // 任务库
	stop             chan bool         // 停止WORKER
)

func init() {
	stop = make(chan bool)
	tickDuration = 100 * time.Millisecond
	tick = time.NewTicker(tickDuration)
	ticksPerWheel = 512
	status = WORKER_STATE_INIT
	tasks = make(map[int64][]*Task, ticksPerWheel)
}

func StartTimer() {
	startTime = time.Now().UnixNano()
	status = WORKER_STATE_STARTED
	go run()
}

func StopTimer() {
	status = WORKER_STATE_SHUTDOWN
	stop <- true
}

func AddTask(taskName string, doFunc TaskDoFunc, args []string, after time.Duration) {
	if after.Nanoseconds() < MIN_DELAY_UNIT.Nanoseconds() {
		after = MIN_DELAY_UNIT
	}
	// 任务需要经过的tick数为
	ticks := int64(after) / int64(tickDuration)
	// 任务需要经过的轮数为
	remainingRounds := ticks / ticksPerWheel
	// 任务存放的wheel索引为
	stopIndex := ticks & (ticksPerWheel - 1)
	var task = &Task{
		name:            taskName,
		doFunc:          doFunc,
		args:            args,
		delay:           after,
		stopIndex:       stopIndex,
		remainingRounds: remainingRounds,
	}
	tasks[stopIndex] = append(tasks[stopIndex], task)
	log.Printf("[%s]任务已经添加成功, %s后执行, 所在卡位:%d", taskName, after, stopIndex)
}

func run() {
	for {
		select {
		case <-tick.C:
			tasks := tasks[currentTickIndex]
			for _, t := range tasks {
				if nil != t {
					isRun := t.remainingRounds == 0 && t.stopIndex == currentTickIndex && !t.done
					if isRun {
						go t.run()
					}
					t.remainingRounds--
				}
			}
			ticks++
			currentTickIndex++
			if currentTickIndex == ticksPerWheel {
				rounds++
				currentTickIndex = 0
			}
		case <-stop:
			return
		}
	}
}

type TaskDoFunc func(args ...string) error

type Task struct {
	name            string        // 任务的名字
	doFunc          TaskDoFunc    // 任务要执行的方法
	args            []string      // 方法参数
	delay           time.Duration // 任务延迟
	stopIndex       int64         // 在哪个槽位
	remainingRounds int64         // 剩余的轮数
	done            bool          // 是否执行完成
}

func (t *Task) run() {
	var startTime = time.Now()
	err := t.doFunc(t.args...)
	if nil != err {
		log.Printf("[%s]执行失败: %s\n", t.name, err)
		return
	}
	t.done = true
	log.Printf("[%s]执行完毕, 用时:%s\n", t.name, time.Now().Sub(startTime))
}
