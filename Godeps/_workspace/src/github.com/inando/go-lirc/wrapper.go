package lirc

// #cgo pkg-config: lirc
// #cgo LDFLAGS: -llirc_client
// #include <stdlib.h>
// #include <lirc_client.h>
//
// int wrap_lirc_command_init(lirc_cmd_ctx* ctx, const char* cmd)
// {
// 	   return lirc_command_init(ctx, cmd);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

//wrapper for lirc_init
func lircInit(prog string, verbose bool) (err error) {
	p := C.CString(prog)
	defer C.free(unsafe.Pointer(p))
	v := bool2Cint(verbose)

	fd, err := C.lirc_init(p, v)
	if fd == C.LIRC_RET_ERROR {
		msg := fmt.Sprint(fd, err)
		err = fmt.Errorf("lirc_init returns %s", msg)
	}

	return
}

//wrapper for lirc_deinit
func lircDeinit() (err error) {
	ok, err := C.lirc_deinit()
	if ok < 0 || err != nil {
		msg := fmt.Sprint(ok, err)
		err = fmt.Errorf("lirc_deinit returns %s", msg)
	}
	return
}

//wrapper for lirc_command_init
func lircCommandInit(format string, v ...interface{}) (ctx C.lirc_cmd_ctx, err error) {

	cmd := C.CString(fmt.Sprintf(format+"\n", v...))
	defer C.free(unsafe.Pointer(cmd))

	ok, err := C.wrap_lirc_command_init(&ctx, cmd)
	if ok != 0 || err != nil {
		msg := fmt.Sprint(ok, err)
		err = fmt.Errorf("lirc_command_init returns %s", msg)
	}

	return
}

//wrapper for lirc_command_run
func lircCommandRun(ctx C.lirc_cmd_ctx, fd int) (err error) {

	for repeat := false; repeat; {
		ok, err := C.lirc_command_run(&ctx, C.int(fd))
		if ok == C.EAGAIN {
			repeat = true
		} else {
			if ok != 0 || err != nil {
				msg := fmt.Sprint(ok, err)
				err = fmt.Errorf("lirc_command_run returns %s", msg)
				return err
			}
		}
	}

	return
}

//wrapper for get_local_socket
//use empty path for default location
func lircGetLocalSocket(path string, verbose bool) (fd int, err error) {
	var p *C.char
	if path != "" {
		p := C.CString(path)
		defer C.free(unsafe.Pointer(p))
	}

	quiet := bool2Cint(!verbose)

	fdC, err := C.lirc_get_local_socket(p, quiet)
	fd = int(fdC)
	if fd < 0 || err != nil {
		msg := fmt.Sprint(fd, err)
		err = fmt.Errorf("lirc_command_run returns %s", msg)
	}
	return
}

//convert bool to C.int
func bool2Cint(v bool) C.int {
	b := C.int(0)
	if v {
		b = C.int(1)
	}
	return b
}
