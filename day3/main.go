package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"sync"
)

var re = regexp.MustCompile(`(?m)#(\d*) @ (\d*),(\d*): (\d*)x(\d*)`)

func main() {
	challenge2()
}

func challenge2() {
	b, err := ioutil.ReadFile("C:\\Users\\dominik\\Documents\\go\\src\\github.com\\paradoxxl\\aoc\\day3\\input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := re.FindAllStringSubmatch(string(b), -1)
	w:= sync.WaitGroup{}

	m := sync.Map{}
	for _,v:= range input{
		idx,_ := strconv.Atoi(v[1])
		m.Store(idx,false)
	}


	for i:= range input{
		w.Add(1)
		go search(&w,input[i:],m)
	}

	w.Wait()

	m.Range(func(k, v interface{}) bool {
		if v == false {
			fmt.Println(k,v)
		}
		return true
	})
}

func search(w *sync.WaitGroup, s [][]string, m sync.Map){
	a := s[0]
	ai,_ := strconv.Atoi(a[1])
	if len(s) > 1{
		for _,b:= range s[1:]{
			if overlaps(a,b){
				m.Store(ai,true)
				bi,_ := strconv.Atoi(b[1])
				m.Store(bi,true)
			}
		}
	}
	w.Done()

}

func overlaps(a, b []string) bool {
	ax, _ := strconv.Atoi(a[2])
	ay, _ := strconv.Atoi(a[3])
	alx, _ := strconv.Atoi(a[4])
	aly, _ := strconv.Atoi(a[5])

	bx, _ := strconv.Atoi(b[2])
	by, _ := strconv.Atoi(b[3])
	blx, _ := strconv.Atoi(b[4])
	bly, _ := strconv.Atoi(b[5])

	if ax < bx+blx && ax+alx > bx &&
		ay < by+bly && ay+aly > by {
		return true
	}
	return false
}
func challenge1() {
	n := 1000
	m := 1000
	matrix := make([][]int, n)
	rows := make([]int, n*m)
	for i := 0; i < n; i++ {
		matrix[i] = rows[i*m : (i+1)*m]
	}

	b, err := ioutil.ReadFile("C:\\Users\\dominik\\Documents\\go\\src\\github.com\\paradoxxl\\aoc\\day3\\input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	input := re.FindAllStringSubmatch(string(b), -1)

	for _, line := range input {

		x, _ := strconv.Atoi(line[2])
		y, _ := strconv.Atoi(line[3])
		lx, _ := strconv.Atoi(line[4])
		ly, _ := strconv.Atoi(line[5])

		for dx := x; dx < x+lx; dx++ {
			for dy := y; dy < y+ly; dy++ {
				matrix[dx][dy]++
			}
		}
	}

	var res int
	for _, l := range matrix {
		for _, c := range l {
			if c >= 2 {
				res++
			}
		}
	}
	fmt.Println(res)
}
