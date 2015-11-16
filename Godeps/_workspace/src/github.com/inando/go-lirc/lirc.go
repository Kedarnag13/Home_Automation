//Go bindings for lirc
package lirc

type Client struct {
	Verbose bool
	Socket  int
}

//New initializes a Client struct
func New() (client *Client, err error) {
	client = new(Client)
	client.Verbose = false
	client.Socket, err = lircGetLocalSocket("", client.Verbose)
	return
}

//Send is for blasting ir codes.
//Here is more information: http://www.lirc.org/html/irsend.html
func (client *Client) Send(format string, v ...interface{}) (err error) {
	ctx, err := lircCommandInit(format, v...)
	if err != nil {
		return
	}

	err = lircCommandRun(ctx, client.Socket)
	if err != nil {
		return
	}

	return
}
