package service

import (
	// 标准包
	"context"
	"fmt"
	"net"

	// 内部包
	"gotraining3/internal/week6/pkg"
)

type LengthFieldBasedFrameDecoder struct {
	hss                 *HomeworkSocketServer
	ctx                 context.Context
	cancelFunc          context.CancelFunc
	buf                 []byte
	maxFrameLength      int // 发送的数据帧最大长度
	lengthFieldOffset   int // 长度字段开始位置
	lengthFieldLength   int // 长度字段长度
	initialBytesToStrip int // 丢弃掉的数据长度
	failFast            bool
}

func NewLengthFieldBasedFrameDecoder(ctx context.Context, maxFrameLength int, lengthFieldOffset int, lengthFieldLength int, initialBytesToStrip int, failFast bool, port string) *LengthFieldBasedFrameDecoder {
	cancelContext, cancel := context.WithCancel(ctx)
	socketServer, err := NewSocketServer(cancelContext, port)
	if nil != err {
		panic(err)
	}

	var lfb LengthFieldBasedFrameDecoder
	lfb.hss = socketServer
	lfb.ctx = cancelContext
	lfb.cancelFunc = cancel
	lfb.buf = make([]byte, 1)
	lfb.maxFrameLength = maxFrameLength
	lfb.lengthFieldOffset = lengthFieldOffset
	lfb.lengthFieldLength = lengthFieldLength
	lfb.initialBytesToStrip = initialBytesToStrip

	return &lfb
}

func (l *LengthFieldBasedFrameDecoder) Start() error {
	l.hss.Start()

	for {
		select {
		case conn := <-l.hss.ConnChan:
			{
				l.decode(conn)
			}
		case <-l.ctx.Done():
			return nil
		}
	}

}

func (l *LengthFieldBasedFrameDecoder) Stop() {
	l.cancelFunc()
}

func (l *LengthFieldBasedFrameDecoder) decode(conn net.Conn) {
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
		l.buf = append(l.buf, tmpBuf...)
	}

	for {
		if len(l.buf) < l.maxFrameLength {
			break
		}

		// 取长度字段值
		lengthFieldLength_value, err := pkg.Bytes2IntS(l.buf[l.lengthFieldOffset:l.lengthFieldLength])
		if nil != err {
			println(err.Error())
			break
		}
		if l.failFast {
			if lengthFieldLength_value > l.maxFrameLength {
				fmt.Printf("长度字段长度大于数据帧长度")
				return
			}
		}
		// 长度字段到数据区的移移量。 数据包长度 - lengthFieldOffset - lengthFieldLength - 长度字段的值
		lengthAdjustment_value := l.maxFrameLength - l.lengthFieldOffset - l.lengthFieldLength - lengthFieldLength_value
		// 解码
		// 开始索引位置
		dataStartIndex := 0 + l.initialBytesToStrip
		dataEndIndex := l.lengthFieldOffset + l.lengthFieldLength + lengthAdjustment_value + lengthFieldLength_value
		decodeData := l.buf[dataStartIndex:dataEndIndex]
		fmt.Println(decodeData)

		l.buf = append(l.buf[dataEndIndex:])
	}
}
