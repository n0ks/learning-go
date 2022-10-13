package main

import "time"

const write = "write"
const sleep = "sleep"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)

}

type SpyCount struct {
	Calls []string
}

func (s *SpyCount) Sleep() {
	s.Calls = append(s.Calls, sleep)

}

func (s *SpyCount) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
