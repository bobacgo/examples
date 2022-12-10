package bigo_test

import (
	"log"
	"testing"
)

// 时间复杂度 O(1)
func TestO1(t *testing.T) {
	n := 1000
	log.Println(n + 1)
	log.Println(n * 2)
	log.Println(n * n)
}

// 时间复杂度 O(n)
func TestOn(t *testing.T) {
	n := 1000
	for i := 0; i < n; i++ {
		log.Println(n)
	}
}

// 时间复杂度 O(n^2)
func TestOnpow2(t *testing.T) {
	n := 1000
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			log.Println(i, j)
		}
	}
}

// 时间复杂度 O(logn)
func TestOlogn(t *testing.T) {
	n := 1000
	for i := 0; i < n; i = i * 2 {
		log.Println(n)
	}
}

// 时间复杂度 O(k^n)
func TestOkpown(t *testing.T) {
	var n uint = 1000
	fib(n)
}

func fib(n uint) uint {
	if n <= 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
