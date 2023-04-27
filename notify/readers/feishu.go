package readers

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"minotaur/notify"
	"net/http"
)

func NewFeiShu(webhook, receiveId string) *FeiShu {
	return &FeiShu{
		client:    resty.New(),
		webhook:   webhook,
		receiveId: receiveId,
	}
}

type FeiShu struct {
	client    *resty.Client
	webhook   string
	receiveId string
}

func (slf *FeiShu) Push(notify notify.Notify) error {
	content, err := notify.Format()
	if err != nil {
		return err
	}
	resp, err := slf.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(content).
		Post(slf.webhook)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("FeiShu notify reader push failed, err: %s", resp.String())
	}
	return err
}
