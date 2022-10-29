package tushare

import (
	"testing"
)

func TestEventQueue(t *testing.T) {
	q := eventQueue{}
	q.add(&event{api: "1"})
	if q.size() != 1 {
		t.Errorf("error size after add first event into queue")
		return
	}
	if q.first().evt.api != "1" {
		t.Errorf("error last after add first event into queue")
		return
	}
	q.add(&event{api: "2"})
	q.add(&event{api: "3"})
	if q.first().evt.api != "1" {
		t.Errorf("error last after add three events into queue")
		return
	}
	if q.size() != 3 {
		t.Errorf("error size after add three events into queue")
		return
	}
	q.remove()
	if q.first().evt.api != "2" {
		t.Errorf("error last after remove one event from queue")
		return
	}
	q.remove()
	if q.first().evt.api != "3" {
		t.Errorf("error last after remove two events from queue")
		return
	}
	q.remove()
	if !q.isEmpty() && q.size() == 0 {
		t.Errorf("error is empty after all events from queue")
	}
}
