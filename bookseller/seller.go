package bookseller

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listCat) == 0 || len(listArt) == 0 {
		return ""
	}

	cat := map[string]int{}

	for _, v := range listCat {
		q := getQuantity(v, listArt)
		cat[v] = q
	}

	var result string
	var i int
	for _, v := range listCat {
		i += 1
		result = result + fmt.Sprintf("(%s : %s)",v,strconv.Itoa(cat[v]))
		if i < len(cat) {
			result = result + " - "
		}
	}

	return result
}

func getQuantity(category string, listArt []string) int {
	qArt := 0
	for _, art := range listArt {
		aArt := strings.Split(art, " ")
		if category == string(aArt[0][0]) {
			q, _ := strconv.Atoi(aArt[1])
			qArt = qArt + q
		}
	}
	return qArt
}

