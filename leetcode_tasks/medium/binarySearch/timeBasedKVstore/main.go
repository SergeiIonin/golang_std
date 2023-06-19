package main

import "fmt"

// https://leetcode.com/problems/time-based-key-value-store/

// Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

// Implement the TimeMap class:

// TimeMap() Initializes the object of the data structure.
// void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
// String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp.
// If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".

func main() {
	tm := Constructor()

	tm.Set("foo", "bar", 1)
	r0 := tm.Get("foo", 1)
	r1 := tm.Get("foo", 3)
	fmt.Println("r0 = ", r0)
	fmt.Println("r1 = ", r1)

	tm.Set("foo", "bar2", 4)
	r2 := tm.Get("foo", 4)
	r3 := tm.Get("foo", 5)
	fmt.Println("r2 = ", r2)
	fmt.Println("r3 = ", r3)
}

type TimeMap struct {
	underlying map[string][]VwithTS
}

func Constructor() TimeMap {
	underlying := make(map[string][]VwithTS, 100)
	return TimeMap{underlying: underlying}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	arr, ok := this.underlying[key]
	if ok {
		arr = append(arr, VwithTS{value, timestamp})
		this.underlying[key] = arr
	} else {
		arr = make([]VwithTS, 0, 10)
		arr = append(arr, VwithTS{value, timestamp})
		this.underlying[key] = arr
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr, ok := this.underlying[key]
	if ok {
		return findValueForTimestamp(arr, timestamp)
	}
	return ""
}

func findValueForTimestamp(arr []VwithTS, ts int) string {
	v := arr[0]
	if ts < v.ts {
		return ""
	}
	l := 0
	r := len(arr) - 1
	diff := ts
	for l <= r {
		ind := (l + r) / 2
		med := arr[ind]
		diffMed := abs(med.ts - ts)
		if diffMed < diff {
			diff = diffMed
			v = med
		}
		if ts > med.ts {
			l = ind + 1
		} else if ts < med.ts {
			r = ind - 1
		} else {
			return med.value
		}
	}
	return v.value
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type VwithTS struct {
	value string
	ts    int
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
