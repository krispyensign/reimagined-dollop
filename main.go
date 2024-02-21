package main

import (
	"fmt"
	"time"
	"bytes"
	"compress/gzip"
	"log"

	"github.com/google/uuid"
)

func main() {
	var (
		b1 bytes.Buffer
		b2 bytes.Buffer
	)

	gz1 := gzip.NewWriter(&b1)
	for i := 0; i < 6000000; i++ {
		id := uuid.New().String()
		if _, err := gz1.Write([]byte(id)); err != nil {
			log.Fatal(err)
		}
		if i % 1000 == 0 {
			fmt.Println(i)
		}
	}
	if err := gz1.Close(); err != nil {
		log.Fatal(err)
	}

	gz2 := gzip.NewWriter(&b2)
	for i := 0; i < 6000000; i++ {
		timenow := time.Now().Format(time.RFC3339Nano)
		timeid := uuid.NewSHA1(uuid.Nil, []byte(timenow)).String()
		if _, err := gz2.Write([]byte(timeid)); err != nil {
			log.Fatal(err)
		}
		if i % 1000 == 0 {
			fmt.Println(i)
		}
	}
	if err := gz2.Close(); err != nil {
		log.Fatal(err)
	}


	fmt.Println(len(b1.Bytes()))
	fmt.Println(len(b2.Bytes()))
}
