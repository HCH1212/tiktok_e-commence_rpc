package rpc

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"log"
	"sync"
)

var (
	wg  sync.WaitGroup
	err error
	r   registry.Registry
)

func InitRpcServer(num int) {
	var errChan = make(chan error, num)

	wg.Add(num)

	serverRun(authRpc(), errChan)
	serverRun(userRpc(), errChan)
	serverRun(productRpc(), errChan)
	serverRun(cartRpc(), errChan)
	serverRun(orderRpc(), errChan)

	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Check for errors
	for err = range errChan {
		log.Fatalf("Service failed with error: %v", err)
	}

	log.Println("All services stopped")
}

// 并发同时跑多个服务
func serverRun(server server.Server, errChan chan error) {
	go func() {
		err = server.Run()
		if err != nil {
			log.Fatalf("Service failed with error: %v", err)
		}
		if err != nil {
			errChan <- err
		}
	}()
}

//func common() registry.Registry {
//	r, err = consul.NewConsulRegister(viper.GetString("consul.addr"))
//	if err != nil {
//		log.Fatal(err)
//	}
//	return r
//}
