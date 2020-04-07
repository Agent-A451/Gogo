package main

import (
	"fmt"
	"math"
)	


//Converter sets a new Converter value
type Converter struct{}

//Feet should be represented
type Feet float64
//Centimeters should be represented
type Centimeters float64

//Minutes should be represented
type Minutes float64
//Seconds should be represented
type Seconds float64

//Kilograms should be represented
type Kilograms float64
//Pounds should be represented
type Pounds float64

//Celsius should be represented
type Celsius float64
//Fahrenheit should be represented
type Fahrenheit float64

//Radians should be represented
type Radians float64
//Degrees should be represented
type Degrees float64

//CentimetersToFeet should be represented
func (cvr Converter) CentimetersToFeet(c Centimeters) Feet {
	var res = Feet(c / 30.48)
	return res
}
//FeetToCentimeters should be represented
func (cvr Converter) FeetToCentimeters(f Feet) Centimeters {
	var des = Centimeters(f * 30.48)
	return des
}
//MinutesToSeconds should be represented
func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	var mes = Seconds(m * 60)
	return mes
}
//SecondsToMinutes should be represented
func (cvr Converter) SecondsToMinutes(s Seconds) Minutes {
	var bes = Minutes(s / 60)
	return bes
}
//KilogramsToPounds should be represented
func (cvr Converter) KilogramsToPounds(k Kilograms) Pounds {
	var ces = Pounds(k * 2.205)
	return ces 
}
//PoundsToKilograms should be represented
func (cvr Converter) PoundsToKilograms(p Pounds) Kilograms {
	var pes = Kilograms(p / 2.205)
	return pes
}
//CelsiusToFahrenheit should be represented
func (cvr Converter) CelsiusToFahrenheit(z Celsius) Fahrenheit {
	var ues = Fahrenheit(z * 1.8) + 32
	return ues
}
//FahrenheitToCelsius should be represented
func (cvr Converter) FahrenheitToCelsius(f Fahrenheit) Celsius {
	var zes = Celsius(f * 0.5556) - 17.7778
	return zes
}
//RadiansToDegrees should be reprensented
func (cvr Converter) RadiansToDegrees(r Radians) Degrees {
	const π = math.Pi
	var ves = Degrees(r * 180/π)   
	return ves
}
//DegreesToRadians should be represented
func (cvr Converter) DegreesToRadians(d Degrees) Radians {
	const π = math.Pi
	var wes = Radians(d * π/180)
	return wes 
}
func main() {
	cvr := Converter{}
	fmt.Println(cvr.CentimetersToFeet(50))
	fmt.Println(cvr.FeetToCentimeters(24))
	fmt.Println(cvr.MinutesToSeconds(5))
	fmt.Println(cvr.SecondsToMinutes(600))
	fmt.Println(cvr.KilogramsToPounds(20))
	fmt.Println(cvr.PoundsToKilograms(200))
	fmt.Println(cvr.CelsiusToFahrenheit(100))
	fmt.Println(cvr.FahrenheitToCelsius(212))
	fmt.Println(cvr.RadiansToDegrees(1))
	fmt.Println(cvr.DegreesToRadians(1))
}
