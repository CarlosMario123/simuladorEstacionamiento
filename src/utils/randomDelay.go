package utils

import (
    "math/rand"
    "time"
)


func RandomDelay(min, max float64) time.Duration {
    rand.Seed(time.Now().UnixNano())
    delay := min + rand.Float64()*(max-min)
    return time.Duration(delay * float64(time.Second))
}
