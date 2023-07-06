package misc

import (
	"fmt"
	"testing"
	"time"
)

var msg *message

func TestPow(t *testing.T) {
	var start, end int
	msg = NewMessage("Hi, Alice. I have a proof that I did some work")

	// Benchmark Bob's work
	start = time.Now().Nanosecond()
	msg.doWork()
	end = time.Now().Nanosecond()
	bobTime := end - start
	fmt.Printf("took %d for Bob to do his work\n", bobTime)

	// Benchmark Alice' work
	start = time.Now().Nanosecond()
	if !msg.verifyWork() {
		t.Fatalf("Alice failed to verify Bob's work; rand num: %d, powHash %x", msg.randNum, msg.powHash)
	}
	end = time.Now().Nanosecond()
	aliceTime := end - start
	fmt.Printf("took %d for Alice to verify Bob's work; POW hash: %x\n", aliceTime, msg.powHash)
	fmt.Printf("Work time ratio %.2f\n", float64(bobTime)/float64(aliceTime))
}
