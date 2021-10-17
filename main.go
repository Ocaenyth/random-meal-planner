package main

import (
	"log"
	"os"

	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
}

func setLogLevel(config Config) {
	if config.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	config := NewConfig()
	prettyPrint("Config:", config)
	setLogLevel(config)

	mainDishes := LoadDishes(config.MainFilepath)
	prettyPrint("Main", mainDishes)
	sideDishes := LoadDishes(config.SidesFilepath)
	prettyPrint("Sides", sideDishes)

	GenerateMeal(mainDishes, sideDishes)
}

func GenerateMeal(mainDishes Dishes, sidesDishes Dishes) {
	logrus.Info("Generating a meal ...")

	m := mainDishes.PickOne()
	s := sidesDishes.PickOne()

	logrus.Infof("You should eat some [%s] with [%s] on the side :]", m, s)
}
