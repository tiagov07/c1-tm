package main

import "fmt"

var (
	tempeture uint8   = 23
	humidity  uint8   = 76
	presure   float32 = 1013.2
)

func main() {
	fmt.Printf(" the wather conditions in pereira are.. \n tempeture : %v celcius \n humidity : %v %% \n and presure : %v hPa \n", tempeture, humidity, presure)
}
