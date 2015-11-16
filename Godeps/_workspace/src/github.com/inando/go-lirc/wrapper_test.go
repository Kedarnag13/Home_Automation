package lirc

import "testing"

func Test0001(t *testing.T) {
	err := lircInit("Test0001", true)
	if err != nil {
		t.Error(err)
		return
	}
	lircDeinit()
}

func Test0002(t *testing.T) {
	ctx, err := lircCommandInit("%s %s %s", "SEND_ONCE", "denon", "vol-up")
	if err != nil {
		t.Error(err)
		return
	}

	fd, err := lircGetLocalSocket("", true)
	if err != nil {
		t.Error(err)
		return
	}

	err = lircCommandRun(ctx, fd)
	if err != nil {
		t.Error(err)
		return
	}
}
