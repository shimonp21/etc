package histogram

import (
	"fmt"
	"time"
)

const bucketSize time.Duration = 10 * time.Second

type Histogram struct {
	startTime time.Time
	buckets   []int64
}

func StartNewHistogram() *Histogram {
	return &Histogram{
		startTime: time.Now().UTC(),
		buckets:   make([]int64, 0),
	}
}

func (histogram *Histogram) Add() {
	currentTime := time.Now().UTC()
	bucket := int64(currentTime.Sub(histogram.startTime) / bucketSize)

	for bucket > int64(len(histogram.buckets))-1 {
		histogram.buckets = append(histogram.buckets, 0)
	}

	histogram.buckets[bucket]++

}

func (histogram *Histogram) String() string {
	return fmt.Sprintf("{%v, %v}", histogram.startTime, histogram.buckets)
}
