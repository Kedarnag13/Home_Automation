package lirc_test

import (
	"testing"

	"github.com/inando/go-lirc"
)

func ExampleClient_Send() error {
	client, err := lirc.New()
	if err != nil {
		return err
	}
	return client.Send("%s %s %s", "SEND_ONCE", "denon", "vol-up")
}

func TestSend(t *testing.T) {
	err := ExampleClient_Send()
	if err != nil {
		t.Error(err)
		return
	}
}
