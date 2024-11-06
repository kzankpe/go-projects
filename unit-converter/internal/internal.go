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

func Convert(req Request) Response {
	fmt.Println("Call Convert function")

	switch req.ConvertFrom {
	case "m":
		switch req.ConvertTo {
		case "km":
			return Response{req.Input / 1000}
		case "cm":
			return Response{req.Input * 100}
		case "mm":
			return Response{req.Input * 1000}
		case "inch":
			return Response{req.Input * 39.37}
		case "ft":
			return Response{req.Input / 0.3048}
		case "yard":
			return Response{req.Input / 0.9144}
		case "miles":
			return Response{req.Input / 1609.344}
		}
	}
	return Response{}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	value, err := strconv.ParseFloat(r.FormValue("value"), 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}
	convertfrom := r.FormValue("from")
	convertto := r.FormValue("to")
	result := Convert(Request{value, convertfrom, convertto})
	fmt.Fprintf(w, "Value = %f\n", value)
	fmt.Fprintf(w, "Conversion From = %s\n", convertfrom)
	fmt.Fprintf(w, "Final result : %f\n", result)
}
