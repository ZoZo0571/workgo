package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func task_1_13_8() {
	var x int
	fmt.Scan(&x)
	for x > 9 {
		sum := 0
		for i := x; i > 0; i /= 10 {
			sum += i % 10
		}
		x = sum
		// x = x/1000000 + x/100000%10 + x/10000%10 + x/1000%10 + x/100%10 + x/10%10 + x%10
	}
	fmt.Println(x)
}

func task_1_13_9() {
	var a, b, max7 int
	fmt.Scan(&a, &b)

	max7 = (b / 7) * 7

	if max7 < a {
		fmt.Print("NO")
	} else {
		fmt.Print(max7)
	}

}

// func task_1_13_11() {
// 	var N, x, result int
// 	fmt.Scan(&N)
// 	result = 0
// 	for x = 0; result < N; x++ {
// 		result = int(math.Pow(float64(2), float64(x)))
// 		if result > N {
// 			break
// 			}
// 		fmt.Print(result, " ")
// 	}

// }

func task_1_13_11() {
	var N, x int
	fmt.Scan(&N)

	for x = 1; x < N; x *= 2 {
		fmt.Print(x, " ")
	}

}

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: ":8080"}

	wg.Add(1)
	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}

func assing_handler(s *string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(*s))
	})
}

func webs() {
	httpServerExitDone := &sync.WaitGroup{}
	srv := startHttpServer(httpServerExitDone)
	var s string = "Пока ничего на задано"
	var sptr *string = &s
	assing_handler(sptr)
	for {
		fmt.Print("Введи что писать в браузер: ")
		in := bufio.NewReader(os.Stdin)
		*sptr, _ = in.ReadString('\n')
		if s == "q" {
			if err := srv.Shutdown(context.TODO()); err != nil {
				panic(err) // failure/timeout shutting down the server gracefully
			}
			break
		}
	}

	httpServerExitDone.Wait()
	fmt.Println("Server stopped.")
}
