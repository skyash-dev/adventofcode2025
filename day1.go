package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func eval(expr string) int {
	expr = strings.ReplaceAll(expr, " ", "") 
	expr = strings.ReplaceAll(expr, "\n", "")
	n := len(expr)

	count:=0
	total:=50
	sign:=1
	i:=0

	for i<n{
		if expr[i]=='+'{
			sign=1
			i++
		}else if expr[i] == '-'{
			sign=-1
			i++
		}

		j:=i
		for j<n && expr[j]>='0'&& expr[j]<='9'{
			j++
		}



		num, _ := strconv.Atoi((expr[i:j]))
		if(num>100||num<(-100)){
			num %= 100;
		}
		total += sign*num
		if(total<0){
			total+=100
		}else if total>99{
			total-=100
		}
		i=j
		if total==0{
			count++
		}
	}
	return  count
}

func main() {
	file, err := os.ReadFile("day1_input.txt")

	if err != nil {
		log.Fatal(err)
	}
    
	input := string(file)
	input = strings.ReplaceAll(input, "L", "-")
	input = strings.ReplaceAll(input, "R", "+")

	res := eval(input)

	fmt.Println(res)

}
