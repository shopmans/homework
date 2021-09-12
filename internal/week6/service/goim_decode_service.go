package service

import (
	// 标准包
	"context"
	"fmt"
	"net"
	"strconv"

	// 第三方包

	// 内部包
	"gotraining3/internal/week6/pkg"
)

// Proto proto.
type Proto struct {
	PackLen   int32  // package length
	HeaderLen int16  // header length
	Ver       int16  // protocol version
	Operation int32  // operation for request
	Seq       int32  // sequence number chosen by client
	Body      []byte // body
}

type GoImDecode struct {
	hss        *HomeworkSocketServer
	ctx        context.Context
	cancelFunc context.CancelFunc
	buf        []byte
}

func NewGoImDecode(ctx context.Context, port string) *GoImDecode {
	cancelContext, cancel := context.WithCancel(ctx)
	socketServer, err := NewSocketServer(cancelContext, port)
	if nil != err {
		panic(err)
	}

	var imd GoImDecode
	imd.hss = socketServer
	imd.ctx = cancelContext
	imd.cancelFunc = cancel
	imd.buf = make([]byte, 1)

	return &imd
}

func (g *GoImDecode) Start() error {
	g.hss.Start()

	for {
		select {
		case conn := <-g.hss.ConnChan:
			{
				g.decode(conn)
			}
		case <-g.ctx.Done():
			return nil
		}
	}

}

func (g *GoImDecode) Stop() {
	g.cancelFunc()
}

func (g *GoImDecode) decode(conn net.Conn) {
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
		g.buf = append(g.buf, tmpBuf...)
	}

	pkgLength := 4
	hdrLength := 2
	verLength := 2
	optLength := 4
	seqLength := 4
	var proto Proto
	index := 0

	for {

		decodePkgLength, err := pkg.Bytes2IntS(g.buf[index : index+pkgLength])
		if nil != err {
			fmt.Println(err.Error())
			return
		}
		index = pkgLength
		proto.PackLen = int32(decodePkgLength)

		if decodePkgLength > len(g.buf) {
			break
		}

		decodeHdrLength, err := pkg.Bytes2Int16(g.buf[index : index+hdrLength])
		if nil != err {
			fmt.Println(err.Error())
			return
		}
		index = index + hdrLength
		proto.HeaderLen = decodeHdrLength

		decodeVer, err := pkg.Bytes2Int16(g.buf[index : index+verLength])
		if nil != err {
			fmt.Println(err.Error())
			return
		}
		index = index + verLength
		proto.Ver = decodeVer

		decodeOperation, err := pkg.Bytes2IntS(g.buf[index : index+optLength])
		if nil != err {
			fmt.Println(err.Error())
			return
		}
		index = index + optLength
		proto.Operation = int32(decodeOperation)

		decodeSeq, err := pkg.Bytes2IntS(g.buf[index : index+seqLength])
		if nil != err {
			fmt.Println(err.Error())
			return
		}
		index = index + seqLength
		proto.Seq = int32(decodeSeq)

		decodeData := g.buf[index : proto.PackLen-int32(proto.HeaderLen)]
		proto.Body = decodeData
		index = int(int32(index) + (proto.PackLen - int32(proto.HeaderLen)))

		fmt.Println("package length: " + strconv.Itoa(int(proto.PackLen)))
		fmt.Println("header length: " + strconv.Itoa(int(proto.HeaderLen)))
		fmt.Println("protocol version: " + strconv.Itoa(int(proto.Ver)))
		fmt.Println("operation for request: " + strconv.Itoa(int(proto.Operation)))
		fmt.Println("sequence number chosen by client: " + strconv.Itoa(int(proto.Seq)))
		fmt.Println("body: " + string(proto.Body))
	}
}
