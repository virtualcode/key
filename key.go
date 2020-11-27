package key

import "math"

// TypeableCharacters is number of different
// characters which can be used for the keyspace
const TypeableCharacters = 100

// Space is how big the keyspace will be
func Space(k string) uint64 {
	return uint64(math.Pow(TypeableCharacters, float64(len(k))))
}

// MedianCrackTime calculates the average time required to
// guess a password based on keyspace size and processor power
func MedianCrackTime(k string, tryRate uint64) uint64 {
	return Space(k) / 2 / tryRate
}

// StrongEnough for government work ...but not my bank account :)
// a minimum time required to guess the password
// based on MedianCrackTime and processing power
func StrongEnough(k string, tryRate uint64, minCrackTime uint64) bool {
	return MedianCrackTime(k, tryRate) >= minCrackTime
}
