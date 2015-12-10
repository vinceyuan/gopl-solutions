// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/vinceyuan/gopl-solutions/ch02/ex2.2/tempconv"
)

func process(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	{
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	{
		f := tempconv.Feet(t)
		m := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToM(f), m, tempconv.MToF(m))
	}
	{
		p := tempconv.Pound(t)
		k := tempconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", p, tempconv.PToK(p), k, tempconv.KToP(k))
	}

}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			process(arg)
		}
		return
	}

	fmt.Println("Input number. Ctrl-C to quit.")

	for true {
		var arg string
		_, err := fmt.Scanf("%s", &arg)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		process(arg)
	}
}
