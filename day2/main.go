package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

func main() {
	challenge2()
}

func challenge2(){
	b,err :=ioutil.ReadFile("C:\\Users\\dominik\\Documents\\go\\src\\github.com\\paradoxxl\\aoc\\day2\\input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := strings.Split(string(b),"\n")
	ctx := context.Background()

	for i:=range input{
		go findClosest(ctx,input[i:])
	}

	for ctx.Err() == nil{

	}
}

func findClosest(ctx context.Context,remainers[]string){
	start := remainers[0]
	if len(remainers) >= 2{
		for _,line:= range remainers[1:]{
			if ctx.Err() != nil{
				return
			}
			if len(line) == 0 {
				continue
			}

			if matchStrings(start,line) == 1{
				ctx.Done()
				fmt.Println(start,line)
			}
		}
	}

}

func matchStrings (a,b string) int{
	var r int
	if len(a) != len(b){
		fmt.Println("ojdo","a:",a,"b",b)
		return 0
	}
	for i := 0;i<len(a);i++{
		if a[i] != b[i]{
			r++
		}
	}
	return r
}

func challenge1(){
	b,err :=ioutil.ReadFile("C:\\Users\\dominik\\Documents\\go\\src\\github.com\\paradoxxl\\aoc\\day2\\input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := strings.Split(string(b),"\n")
	println(len(input))
	r := &res{}

	w := sync.WaitGroup{}

	for _,s := range input{
		w.Add(1)
		go func(text string) {
			m := mapText(text)
			evaluateMap(r,m)
			w.Done()
		}(s)
	}

	w.Wait()
	fmt.Println(r.Result())
	fmt.Println(r.two, r.three)
}

func evaluateMap(r *res, m map[rune]int){
	var two,three bool
	for _,v:= range m{
		switch v{
		case 2:
			if !two {
				r.AddTwo()
				two = true
			}
		case 3:
			if !three{
				r.AddThree()
				three = true
			}
		}
	}
}
func challenge1Demo(){
	input := []string{"abcdef","bababc","abbcde","abcccd","aabcdd","abcdee","ababab"}
	r := &res{}

	w := sync.WaitGroup{}

	for _,s := range input{
		w.Add(1)
		go func(text string) {
			m := mapText(text)
			evaluateMap(r,m)
			w.Done()
		}(s)
	}

	w.Wait()
	fmt.Println(r.Result())
	fmt.Println(r.two, r.three)

}

func mapText(s string) map[rune]int{
	r := make(map[rune]int)

	for _,c := range s{
		r[c]++
	}
	return r
}

type res struct{
	sync.RWMutex

	two int
	three int
}

func (r * res) Result() int{
	return r.two*r.three
}
func (r *res) AddTwo(){
	r.Lock()
	defer r.Unlock()

	r.two++
}

func (r *res) AddThree(){
	r.Lock()
	defer r.Unlock()

	r.three++
}