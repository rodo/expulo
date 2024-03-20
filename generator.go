package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

type genericFake struct{}

func (genericFake) FakeEmail() string     { return faker.Email() }
func (genericFake) FakeName() string      { return faker.Name() }
func (genericFake) FakeFirstName() string { return faker.FirstName() }

func mask() string { return "********" }

func randomFloat() float32   { return rand.Float32() }
func randomFloat32() float32 { return rand.Float32() }
func randomFloat64() float64 { return rand.Float64() }

func randomInt() int32   { return rand.Int31() }
func randomInt32() int32 { return rand.Int31() }
func randomInt64() int64 { return rand.Int63() }

func randomIntMinMax(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

type someStruct struct {
	String string
}

func randomString() string {
	a := someStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	return a.String
}

func md5signature(alpha string) string {
	data := []byte(alpha)
	return fmt.Sprintf("%x", md5.Sum(data))
}

/*
 * Time functions
 *
 */
func randomTimeTZ(timezone string) time.Time {

	location, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("Error:", err)
		loc, _ := time.LoadLocation("UTC")
		return time.Date(1970, 1, 1, 0, 0, 0, 0, loc)
	}

	min := time.Date(1973, 1, 1, 0, 0, 0, 0, location).Unix()
	max := time.Date(2024, 1, 1, 0, 0, 0, 0, location).Unix()
	randomUnix := rand.Int63n(max-min) + min
	return time.Unix(randomUnix, 0)
}
