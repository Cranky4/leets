package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV")) // LVIII
}

func romanToInt(s string) int {
    

    tt := map[string]int {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,

        "IV": 4,
        "IX": 9,

        "XL": 40,
        "XC": 90,

        "CD": 400,
        "CM": 900,
    }


    res := 0

    for i := 0; i< len(s); i++ {
        if i+1 < len(s) {
            // pair
            ss := s[i:i+2]

            val, ok := tt[ss]
            if ok {
            res += val
            i += 1
            continue
            }
        }
        

        // by one
        ss := string(s[i])
        val, ok := tt[ss]
        if ok {
           res += val
        }
    }

    return res
}