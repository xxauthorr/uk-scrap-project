package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type item struct {
	Reference     string `json:"refereance"`
	CreatedOn     string `json:"create_on"`
	EstimatedOn   string `json:"Estimated_on"`
	Status        string
	AffectedAreas string
}

var count int
var spEnergyNetworks []item

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.spenergynetworks.co.uk"),
	)

	c.OnHTML("div[class=Item]", func(h *colly.HTMLElement) {
		selector := h.DOM
		val := selector.Find("span.Value").Text()
		count = count + 1
		values := strings.Split(val, "  ")
		item := item{
			Reference:     values[0],
			CreatedOn:     values[1],
			EstimatedOn:   values[2],
			Status:        values[3],
			AffectedAreas: values[4],
		}
		if item.EstimatedOn == "" {
			item.EstimatedOn = "nil"
		}
		spEnergyNetworks = append(spEnergyNetworks, item)

	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println(e.Error())
	})

	err := c.Visit("https://www.spenergynetworks.co.uk/pages/power_cuts_list.aspx")
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := range spEnergyNetworks {
		fmt.Println(spEnergyNetworks[i])
	}

}

	// c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
	// 	next := h.Request.AbsoluteURL(h.Attr("href"))
	// 	c.Visit(next)
	// })