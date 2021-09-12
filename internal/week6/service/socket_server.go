package service

import (
	// 标准包
	"context"
	"fmt"
	"net"
)

type HomeworkSocketServer struct {
	listener   net.Listener
	ConnChan   chan net.Conn
	cancelFunc context.CancelFunc
}

func NewSocketServer(ctx context.Context, port string) (*HomeworkSocketServer, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("error listening:", err.Error())
		return nil, err
	}
	cancelContext, cancel := context.WithCancel(ctx)
	var hss HomeworkSocketServer
	hss.listener = listener
	hss.ConnChan = make(chan net.Conn)
	hss.cancelFunc = cancel

	go func() {
		select {
		case <-cancelContext.Done():
			{
				hss.listener.Close()
				close(hss.ConnChan)
			}
		}
	}()

	return &hss, nil
}

func (h *HomeworkSocketServer) Start() {
	go func() {
		for {
			conn, err := h.listener.Accept()
			if err != nil {
				println("Error accept:", err.Error())
				return
			}
			h.ConnChan <- conn
		}
	}()
}

func (h *HomeworkSocketServer) Stop() {
	h.cancelFunc()
}
