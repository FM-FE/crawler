package parser

import (
	"io/ioutil"
	"testing"
)

func TestCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := CityList(contents)
	// verify results
	//for i, item := range result.Items {
	//	fmt.Printf("%d : %v\n", i, item)
	//}

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expect url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	for i, item := range expectedCities {
		if result.Items[i] != item {
			t.Errorf("expect url #%d: %s; but was %s", i, item, result.Items[i])
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Items))
	}

}
