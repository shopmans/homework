package service

import (
	// 标准包
	"context"
	"fmt"
	"net"
)

type FixLengthDecode struct {
	hss        *HomeworkSocketServer
	ctx        context.Context
	cancelFunc context.CancelFunc
	buf        []byte
	fixLength  int
}

func NewFixLengthDecode(ctx context.Context, fix_length int, port string) *FixLengthDecode {
	cancelContext, cancel := context.WithCancel(ctx)
	socketServer, err := NewSocketServer(cancelContext, port)
	if nil != err {
		panic(err)
	}

	var fld FixLengthDecode
	fld.hss = socketServer
	fld.ctx = cancelContext
	fld.cancelFunc = cancel
	fld.buf = make([]byte, 1)
	fld.fixLength = fix_length

	return &fld
}

func (f *FixLengthDecode) Start() error {
	f.hss.Start()

	for {
		select {
		case conn := <-f.hss.ConnChan:
			{
				f.decode(conn)
			}
		case <-f.ctx.Done():
			return nil
		}
	}

}

func (f *FixLengthDecode) Stop() {
	f.cancelFunc()
}

func (f *FixLengthDecode) decode(conn net.Conn) {
	defer conn.Close()
	tmpBuf := make([]byte, 1024)

	// 读出所有数据
	for {
		readLength, err := conn.Read(tmpBuf)
		if err != nil {
			println(err.Error())
			return
		}
		if readLength <= 0 {
			break
		}
		f.buf = append(f.buf, tmpBuf...)
	}

	for {
		if len(f.buf) < f.fixLength {
			return
		}

		decodeContext := string(f.buf[:f.fixLength])
		fmt.Println(decodeContext)
		f.buf = append(f.buf[f.fixLength:])
	}
}
