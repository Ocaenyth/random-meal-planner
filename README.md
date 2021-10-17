# random-meal-planner

Do you find it difficult to decide what you should prepare for your meals ?

Did you develop an addiction for delivering apps because of the pandemic ? 

Then this little program is perfect for you, by specifying main and side dishes
that you know how to prepare, it will randomly pick one of each and decide your
next meal for you !

## Setup

### Golang

This script is golang, so you're gonna need that, if you don't then here's a 
neat link for you : https://golang.org/doc/install

### Dependencies

In order to download the dependencies, please run the following command inside
the repository you just cloned

`go get`

### Configuration

This script allows some runtime arguments, for which you can get more 
information by running the following command

`go run . -help`

This will print the following message

```
  -main string
        Path to the json file containing your main dishes (default "main.json")
  -sides string
        Path to the json file containing your side dishes (default "sides.json")
  -v    Whether the program will display debug logs
```

### Dishes

For this program to run, you need to provide two JSON files (`main.json` and 
`sides.json` in your current folder by default), each containing what main / 
side dishes you know how to cook respectively, in a JSON array.

Two example files are provided in this repository if you need any help, feel 
free to remove / update them to your will.

### Generating a meal (running the program)

At this point all you need to do is run the following command : 

`go run .`

![normal_run](https://i.imgur.com/njkOkHT.png)

You can also run this program in debug mode and it will display a little more
information, like so :

`go run . -v`

![verbose_run](https://i.imgur.com/zO2i4va.png)