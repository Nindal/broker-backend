package utils

import (
	"github.com/sony/gobreaker"
)

var CircuitBreaker *gobreaker.CircuitBreaker

func InitCircuitBreaker() {
	CircuitBreaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{Name: "ServiceCircuitBreaker"})
}
