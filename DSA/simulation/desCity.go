package simulation

import "fmt"


func destCity(paths [][]string) string {
	startC := paths[0][0]
	endC := paths[0][1]
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths); j++ {
			if paths[j][0] == endC {
				startC = paths[j][0]
				endC = paths[j][1]
				fmt.Println("Start City : ", startC, " End City : ", endC)
			}
		}
	}
	return endC
}
