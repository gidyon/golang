package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type numbers struct {
	Nums []int
}

func (nums *numbers) mean() string {
	sum := 0
	for _, num := range nums.Nums {
		sum += num
	}
	if len(nums.Nums) == 0 {
		return "0"
	}
	return fmt.Sprint(sum / len(nums.Nums))
}

func (nums *numbers) median() string {
	return "NA"
}

func (nums *numbers) add(val int) {
	nums.Nums = append(nums.Nums, val)
}

func (nums *numbers) count() string {
	return fmt.Sprint(len(nums.Nums))
}

func (nums *numbers) print() string {
	return fmt.Sprint((*nums).Nums)
}

var Numbers = new(numbers)

var form = `
	<h2>Statistics</h2>
	<div>Compute basic statistics for a given list of numbers</div>
	<div>Numbers(comma or space-separated)</div>
	<form action="#" method="post" name="bar">
		<input type="text" name="num"></input><br>
		<input type="submit" name="submit" value="Calculate"></input>
	</form>
	<table style="border:1px;">
		<thead>Results</thead>
		<tr>
			<td>Numbers</td><td>` + Numbers.print() + `</td>
		</tr>
		<tr>
			<td>Count</td><td>` + Numbers.count() + `</td>
		</tr>
		<tr>
			<td>Mean</td><td>` + Numbers.mean() + `</td>
		</tr>
		<tr>
			<td>Median</td><td>` + Numbers.median() + `</td>
		</tr>
	</table>
`

func Statistics(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		fmt.Println(Numbers)
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data*/
		//request.ParseForm();
		//io.WriteString(w, request.Form["in"][0])
		num, _ := strconv.ParseInt(request.FormValue("num"), 10, 64)
		Numbers.add(int(num))
		io.WriteString(w, form)
	}
}
func main() {
	http.HandleFunc("/", Statistics)
	http.ListenAndServe(":9001", nil)
}
