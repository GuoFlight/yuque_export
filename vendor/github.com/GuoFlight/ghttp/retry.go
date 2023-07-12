package ghttp

import "github.com/avast/retry-go"

type Retry struct {
	JudgeRetryFunc
	Options []retry.Option
}

type JudgeRetryFunc func(res *Res) error
