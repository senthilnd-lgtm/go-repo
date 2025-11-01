package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now()
	ctx := context.WithValue(context.Background(), "userid", "sedurais")
	val, err := fetchUserData(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value from response", val)
	fmt.Println("whole process took ", time.Since(start))
}

type Response struct {
	val int
	err error
}

func fetchUserData(ctxp context.Context) (int, error) {
	convalue := ctxp.Value("userid")
	fmt.Println("Value from context", convalue)
	ctxc, cancel := context.WithTimeout(ctxp, time.Millisecond*200)
	defer cancel() // going to close the context

	rech := make(chan Response)

	go func() {
		val, err := fetch3rdParty()
		rech <- Response{
			val: val,
			err: err,
		}

	}()

	for {
		select {
		case <-ctxc.Done(): // Dont is getting called due to specified time out
			return 0, fmt.Errorf("Response took too long")
		case resp := <-rech:
			return resp.val, resp.err

		}
	}

}

func fetch3rdParty() (int, error) {
	time.Sleep(time.Millisecond * 150)
	return 666, nil
}
