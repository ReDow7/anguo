package tushare

import "time"

type throttle struct {
	limitPerMinutes int
	queue           eventQueue
	apiCount        map[string]int
	allCount        int
	byApi           bool
}

type eventQueue struct {
	head *eventNode
	tail *eventNode
}

func newThrottle(byApi bool, limitPerMinutes int) *throttle {
	return &throttle{
		limitPerMinutes,
		eventQueue{nil, nil},
		make(map[string]int),
		0, byApi,
	}
}

func (t *throttle) tryByApi(api string) bool {
	t.updateCountByNow()
	if t.checkCountLimit(api) {
		t.newEvent(api)
		return true
	}
	return false
}

func (t *throttle) newEvent(apiName string) {
	t.queue.add(&event{
		api: apiName,
		t:   time.Now(),
	})
	t.updateCount(apiName, 1)
}

func (t *throttle) updateCountByNow() {
	oneMinAgo := time.Now().Add(-1 * time.Minute)
	for ptr := t.queue.first(); !t.queue.isEmpty(); ptr = t.queue.first() {
		if ptr.evt.t.Before(oneMinAgo) {
			t.updateCount(ptr.evt.api, -1)
			t.queue.remove()
		} else {
			break
		}
	}
}

func (t *throttle) checkCountLimit(api string) bool {
	if t.byApi {
		return t.apiCount[api] < t.limitPerMinutes
	} else {
		return t.allCount < t.limitPerMinutes
	}
}

func (t *throttle) updateCount(api string, val int) {
	if t.byApi {
		t.apiCount[api] += val
		if t.apiCount[api] < 0 {
			t.apiCount[api] = 0
		}
	} else {
		t.allCount += val
		if t.allCount < 0 {
			t.allCount = 0
		}
	}
}

func (e *eventQueue) add(event *event) {
	toAdd := &eventNode{evt: event}
	if e.head == nil {
		e.head, e.tail = toAdd, toAdd
	} else {
		e.tail.next = toAdd
		toAdd.pre = e.tail
		e.tail = toAdd
	}
}

func (e *eventQueue) size() int {
	if e.head == nil {
		return 0
	}
	cnt := 0
	for ptr := e.head; ptr != nil; ptr = ptr.next {
		cnt++
	}
	return cnt
}

func (e *eventQueue) isEmpty() bool {
	return e.head == nil
}

func (e *eventQueue) first() *eventNode {
	return e.head
}

func (e *eventQueue) remove() {
	if e.head.next != nil {
		e.head.next.pre = nil
		e.head = e.head.next
	} else {
		e.head, e.tail = nil, nil
	}
}

type eventNode struct {
	evt  *event
	next *eventNode
	pre  *eventNode
}

type event struct {
	api string
	t   time.Time
}
