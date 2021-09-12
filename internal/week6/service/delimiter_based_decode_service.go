package service

import (
	// 标准包
	"context"
	"fmt"
	"net"

	// 内部包
	"gotraining3/internal/week6/pkg"
)

type DeLimiterBasedDecode struct {
	hss        *HomeworkSocketServer
	ctx        context.Context
	cancelFunc context.CancelFunc
	buf        []byte
	delimiter  string
}

func NewDeLimiterBasedDecode(ctx context.Context, delimiter string, port string) *DeLimiterBasedDecode {
	cancelContext, cancel := context.WithCancel(ctx)
	socketServer, err := NewSocketServer(cancelContext, port)
	if nil != err {
		panic(err)
	}

	var dbd DeLimiterBasedDecode
	dbd.hss = socketServer
	dbd.ctx = cancelContext
	dbd.cancelFunc = cancel
	dbd.buf = make([]byte, 1)
	dbd.delimiter = delimiter

	return &dbd
}

func (f *DeLimiterBasedDecode) Start() error {
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

func (f *DeLimiterBasedDecode) Stop() {
	f.cancelFunc()
}

func (f *DeLimiterBasedDecode) decode(conn net.Conn) {
	defer conn.Close()
	tmpBuf := make([]byte, 1024)
	_, err := conn.Read(tmpBuf)
	if err != nil {
		println(err.Error())
		return
	}

	// 读出所有数据
	for {
		readLength, err := conn.Read(tmpBuf)
		if err != nil {
			println("Error reading:", err.Error())
			return
		}
		if readLength <= 0 {
			break
		}
		f.buf = append(f.buf, tmpBuf...)
	}

	// 找到分隔符索引并截取出数据
	for {
		delimiterIndex := pkg.IndexOf(f.buf, []byte(f.delimiter))
		if delimiterIndex < 0 {
			return
		} else {
			frame := f.buf[:delimiterIndex]
			fmt.Println(frame)
			f.buf = append(f.buf[delimiterIndex:])
		}
	}
}
