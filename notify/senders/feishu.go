package senders

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/kercylan98/minotaur/notify"
	"net/http"
)

// NewFeiShu 根据特定的 webhook 地址创建飞书发送器
func NewFeiShu(webhook string) *FeiShu {
	return &FeiShu{
		client:  resty.New(),
		webhook: webhook,
	}
}

// FeiShu 飞书发送器
type FeiShu struct {
	client  *resty.Client
	webhook string
}

// Push 推送通知
func (slf *FeiShu) Push(notify notify.Notify) error {
	content, err := notify.Format()
	if err != nil {
		return err
	}
	resp, err := slf.client.R().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetBody(content).
		Post(slf.webhook)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("FeiShu notify push failed, err: %s", resp.String())
	}

	var respStruct = struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}{}
	if err := json.Unmarshal(resp.Body(), &respStruct); err != nil {
		return fmt.Errorf("FeiShu notify response unmarshal failed, err: %s", err)
	}
	if respStruct.Code != 0 {
		return fmt.Errorf("FeiShu notify push failed, err: [%d] %s", respStruct.Code, respStruct.Msg)
	}
	return err
}
