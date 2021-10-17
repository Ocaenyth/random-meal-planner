package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Dish string

type Dishes []Dish

func LoadDishes(filepath string) (dishes Dishes) {
	logrus.Infof("Loading dishes from [%s]", filepath)

	data, err := os.ReadFile(filepath)
	check(err)

	err = json.Unmarshal(data, &dishes)
	check(err)
	return
}

func (d Dishes) PickOne() Dish {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	length := len(d)
	index := r.Intn(length)

	return d[index]
}
