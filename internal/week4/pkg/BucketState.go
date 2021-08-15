package roll

import (
	"sync"
	"time"
)

type BucketState struct {
	Buckets []*Bucket
	// bucket总数
	NumBucket int
	// 每个bucket占用时长（毫秒）
	InWdindowsTime int
	//  窗口时长（毫秒）
	WindowsMilliseconds int
	// 最后一个bucket索引
	LastBucketIndex int
	lock            sync.Mutex
}

func NewBucketState(numBucket int, inWdindowsTime int) *BucketState {
	if numBucket <= 0 {
		return nil
	}
	if inWdindowsTime <= 0 {
		return nil
	}
	return &BucketState{
		Buckets:             make([]*Bucket, numBucket),
		NumBucket:           numBucket,
		InWdindowsTime:      inWdindowsTime,
		WindowsMilliseconds: numBucket * inWdindowsTime,
		LastBucketIndex:     0,
	}
}

func (b *BucketState) GetCurrentBucket() (error, *Bucket) {
	currentTime := time.Nanosecond.Microseconds()

	// 最后的bucket
	currentBucket := b.Buckets[b.LastBucketIndex]

	// 如果当前时间是在currentBucket对应的时间窗口内，直接返回currentBucket
	if nil != currentBucket && currentTime < (currentBucket.WindowStart+int64(b.InWdindowsTime)) {
		return nil, currentBucket
	}

	b.lock.Lock()
	defer b.lock.Unlock()

	// 首次建bucket
	if nil == currentBucket {
		newBucket := Bucket{
			WindowStart: currentTime,
			Count:       0,
		}
		b.Buckets = append(b.Buckets, &newBucket)
		return nil, &newBucket
	}

	// 将创建一个或者多个Bucket，直到Bucket代表的时间窗口赶上当前时间
	for i := 0; i < b.NumBucket; i++ {
		lastBucket := b.Buckets[b.LastBucketIndex]

		if lastBucket != nil && currentTime < (lastBucket.WindowStart+int64(b.InWdindowsTime)) {
			// 最后一个bucket在当前时间窗口内
			return nil, currentBucket
		} else if currentTime-(lastBucket.WindowStart+int64(b.InWdindowsTime)) > int64(b.WindowsMilliseconds) {
			// 整过滑动窗口时间过期，清空窗口内的所有bucket
			b.Buckets = make([]*Bucket, 0)
			b.LastBucketIndex = 0
			return b.GetCurrentBucket()
		} else {
			// 紧随最后一个bucket新建一个
			b.LastBucketIndex++
			if b.LastBucketIndex >= b.NumBucket {
				b.LastBucketIndex = 0
			}
			b.Buckets[b.LastBucketIndex] = &Bucket{
				WindowStart: lastBucket.WindowStart + int64(b.InWdindowsTime),
				Count:       0,
			}
		}
	}

	return nil, b.Buckets[b.LastBucketIndex]
}

func (b *BucketState) GetWindowsCount() int {
	b.lock.Lock()
	defer b.lock.Unlock()

	sum := 0
	for _, v := range b.Buckets {
		sum = sum + v.Count
	}

	return sum
}
