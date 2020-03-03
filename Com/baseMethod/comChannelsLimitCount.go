/*
* Domain 例子，主要思想是，设置执行最高次数。
 */
package baseMethod

import (
	"fmt"
	"sync"
	"time"
)

//LimitRate 限数量
type LimitRate struct {
	rate  int
	begin time.Time //开始时间
	count int
	lock  sync.Mutex //互斥锁
}

//LimitChannel() 根据
func (l *LimitRate) LimitChannel() bool {
	result := true
	l.lock.Lock()
	if l.count >= l.rate {
		if time.Now().Sub(l.begin) >= time.Second {
			l.begin = time.Now()
			l.count = 0
		} else {
			result = false
		}
	} else {
		l.count++
	}
	l.lock.Unlock()

	return result
}

//SetRate 设置每秒允许的请求数
func (l *LimitRate) SetRate(r int) {
	l.rate = r
	l.begin = time.Now()
}

func (l *LimitRate) GetRate() int {
	return l.rate
}

func Domain() {
	var wg sync.WaitGroup
	var lr LimitRate
	lr.SetRate(3)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			if lr.LimitChannel() {
				fmt.Println("Go git", j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
