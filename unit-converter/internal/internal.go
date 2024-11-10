package internal

import (
	"fmt"
	"net/http"
	"strconv"
)

type Request struct {
	Input       float64 `json:"input"`
	ConvertFrom string  `json:"convertfrom"`
	ConvertTo   string  `json:"convertto"`
}

type Response struct {
	Result float64 `json:"result"`
}

func ConvertLength(req Request) Response {
	fmt.Println("Call Convert Length function")

	//Convert to meter
	var res Response
	switch req.ConvertFrom {
	case "m":
		res = Response{req.Input}
	case "km":
		res = Response{req.Input * 1000}
	case "cm":
		res = Response{req.Input / 100}
	case "mm":
		res = Response{req.Input / 1000}
	case "ft":
		res = Response{req.Input * 0.3048}
	case "yard":
		res = Response{req.Input * 1.0936133}
	case "miles":
		res = Response{req.Input * 1609.34}
	case "inch":
		res = Response{req.Input * 39.3700787}
	}

	//Convert to the desired units

	switch req.ConvertTo {
	case "m":
		return res
	case "km":
		return Response{res.Result / 1000}
	case "cm":
		return Response{res.Result * 100}
	case "mm":
		return Response{res.Result * 1000}
	case "ft":
		return Response{res.Result / 0.3048}
	case "yard":
		return Response{res.Result / 1.0936133}
	case "miles":
		return Response{res.Result / 1609.34}
	case "inch":
		return Response{res.Result / 39.3700787}
	}

	return Response{}
}

func ConvertWeigth(r Request) Response {
	fmt.Println("Call Convert Length function")

	var response Response

	//Convert all input to gram
	switch r.ConvertFrom {
	case "g":
		response = Response{r.Input}
	case "mg":
		response = Response{r.Input / 1000}
	case "kg":
		response = Response{r.Input * 1000}
	case "oz":
		response = Response{r.Input * 28.3495231}
	case "pd":
		response = Response{r.Input * 453.59237}
	}

	// Convert to final unit
	switch r.ConvertTo {
	case "g":
		return response
	case "mg":
		return Response{response.Result / 1000}
	case "kg":
		return Response{response.Result * 1000}
	case "oz":
		return Response{response.Result / 28.3495231}
	case "pd":
		return Response{response.Result / 453.59237}
	}
	return Response{}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Fprintf(w, "POST request successful\n")
	value, err := strconv.ParseFloat(r.FormValue("value"), 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}
	conversionType := r.FormValue("conversionType")
	fmt.Println(conversionType)
	convertfrom := r.FormValue("from")
	convertto := r.FormValue("to")
	result := ConvertLength(Request{value, convertfrom, convertto})
	fmt.Println(result)
	fmt.Fprintf(w, "<h1>Length Conversion Result</h1>")
	fmt.Fprintf(w, "<h2>Final result : %.3f %s\n</h2>", result.Result, convertto)

	// Add a button to go back to the home page
	fmt.Fprintf(w, `<br><a href="/"> <button>Go Back to Home</button> </a>`)

}
