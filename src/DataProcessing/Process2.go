package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * 分类汇总计算平均值
 */
func main() {
	f, err := os.Open("D:\\Data\\yyh\\groups.csv")
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
	groupMap := make(map[string]string)
	for _, v := range result {
		groupMap[v[0]] = v[1]
	}
	basePath := "D:\\Data\\yyh\\"
	for year := 2019; year <= 2021; year++ {
		for month := 1; month <= 12; month++ {
			path := basePath + strconv.Itoa(year) + "\\moral_time_filter" + strconv.Itoa(month) + ".csv"
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
			avgDataArray := make([][]float64, 6)
			countMap := make(map[int]int)
			for i := 0; i < 6; i++ {
				countMap[i] = 0
			}
			title := []string{}
			for _, v := range lines {
				groupIdStr, exist := groupMap[v[0]]
				if exist {
					groupId, _ := strconv.Atoi(groupIdStr)
					groupId = groupId - 1
					if countMap[groupId] == 0 {
						avgDataArray[groupId] = make([]float64, len(v)-1)
					}
					for i := 0; i < len(v)-1; i++ {
						data, _ := strconv.Atoi(strings.TrimSpace(v[i+1]))
						avgDataArray[groupId][i] = (avgDataArray[groupId][i]*float64(countMap[groupId]) + float64(data)) / float64(countMap[groupId]+1)
					}
					countMap[groupId] = countMap[groupId] + 1
				} else {
					title = v
				}
			}
			itemSize := len(title) - 1
			avgStrArray := make([][]string, 7)
			avgStrArray[0] = title
			for i := 1; i < 7; i++ {
				avgStrArray[i] = make([]string, itemSize+1)
				avgStrArray[i][0] = strconv.Itoa(i)
				for j := 0; j < itemSize; j++ {
					avgStrArray[i][j+1] = fmt.Sprintf("%f", avgDataArray[i-1][j])
				}
			}
			path2 := basePath + strconv.Itoa(year) + "\\moral_time_avg" + strconv.Itoa(month) + ".csv"
			nf, err := os.Create(path2)
			if err != nil {
				panic(err)
			}
			defer nf.Close()
			writer := csv.NewWriter(nf)
			writer.WriteAll(avgStrArray)
			writer.Flush()
			if err = writer.Error(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
