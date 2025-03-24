package vscale_api_go

import (
	"io"
	"log"
	"testing"
)

// Test data
var BillingLimit int64 = 10000

func TestNotificationService_BillingSettings(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, _, err = client.Notification.BillingSettings()
	if err != nil {
		t.Error(err)
		return
	}

}

func TestNotificationService_BillingSettingsUpdate(t *testing.T) {

	token, err := GetToken()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(token)
	_, res, err := client.Notification.BillingSettingsUpdate(BillingLimit)
	data, _ := io.ReadAll(res.Body)
	log.Println(string(data))
	if err != nil {
		t.Error(err)
		return
	}

}
