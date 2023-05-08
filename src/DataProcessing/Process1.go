package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/**
 * 筛选有效数据行
 */
func main() {
	f, err := os.Open("D:\\Data\\yyh\\ll.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	reader := csv.NewReader(f)
	result, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	keySet := make(map[string]bool)
	for _, v := range result {
		keySet[v[0]] = true
	}
	basePath := "D:\\Data\\yyh\\"
	for year := 2019; year <= 2021; year++ {
		for i := 1; i <= 12; i++ {
			path := basePath + strconv.Itoa(year) + "\\moral_time" + strconv.Itoa(i) + ".csv"
			f, err := os.Open(path)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			reader := csv.NewReader(f)
			lines, err := reader.ReadAll()
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			newStrArr := [][]string{}
			for _, v := range lines {
				_, exist := keySet[v[0]]
				if exist {
					newStrArr = append(newStrArr, v)
				}
			}
			path2 := basePath + strconv.Itoa(year) + "\\moral_time_filter" + strconv.Itoa(i) + ".csv"
			nf, err := os.Create(path2)
			if err != nil {
				panic(err)
			}
			defer nf.Close()
			writer := csv.NewWriter(nf)
			writer.WriteAll(newStrArr)
			writer.Flush()
			if err = writer.Error(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
