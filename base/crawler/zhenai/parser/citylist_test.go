package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err:= ioutil.ReadFile("zhenai_data.html")

	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s\n",contents)
	result := ParseCityList(contents)
	const resultSize = 470

	//expectedUrls := []string{
	//	"","",""
	//}

	if len(result.Requests) != resultSize{
		t.Errorf("result shoud have %d "+ "requests; but had %d" ,resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize{
		t.Errorf("result shoud have %d "+ "requests; but had %d" ,resultSize, len(result.Items))
	}
}
