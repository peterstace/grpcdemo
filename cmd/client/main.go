package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/peterstace/grpcdemo/calc"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not dial: %v", err)
	}
	defer conn.Close()
	c := calc.NewCalculatorClient(conn)

	for {
		start := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		n1 := decimal.New(int64(rand.Intn(10000)), -2)
		n2 := decimal.New(int64(rand.Intn(10000)), -2)
		sum, err := c.Add(ctx, &calc.Numbers{
			Num1: n1.String(),
			Num2: n1.String(),
		})
		if err != nil {
			log.Fatalf("Could not add: %v", err)
		}
		log.Printf("Result %v + %v = %v", n1, n2, sum.GetNum())
		log.Println(time.Since(start))
		cancel()
	}
}
