package rpc

import (
	"github.com/cloudwego/kitex/server"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	wg  sync.WaitGroup
	err error
)

func InitRpcServer(num int) {
	var errChan = make(chan error, num)

	wg.Add(num)

	serverRun(authRpc(), errChan)
	serverRun(userRpc(), errChan)
	serverRun(productRpc(), errChan)
	serverRun(cartRpc(), errChan)
	serverRun(orderRpc(), errChan)
	serverRun(paymentRpc(), errChan)

	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Check for errors
	for err = range errChan {
		logrus.Fatalf("Service failed with error: %v", err)
	}

	logrus.Println("All services stopped")
}

// 并发同时跑多个服务
func serverRun(server server.Server, errChan chan error) {
	go func() {
		err = server.Run()
		if err != nil {
			logrus.Fatalf("Service failed with error: %v", err)
		}
		if err != nil {
			errChan <- err
		}
	}()
}
