package api

import "time"

type setAble interface {
	SetTimeout(timeout time.Duration)
}

func SetTimeout[t setAble](a t, timeout time.Duration) t {
	a.SetTimeout(timeout)
	return a
}
