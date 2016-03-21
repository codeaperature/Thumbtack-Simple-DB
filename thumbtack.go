

// https://www.thumbtack.com/challenges/simple-database
package main

import (
"bufio"
"fmt"
"strings"
"os"
)


//func getMax(mm0 map[int]string) (int, string) {
//	return 1, ""
//}


func main() {

	mm := make(map[string]map[int]string)
	level := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
//		fmt.Println(text)
		all := strings.Fields(text)

        switch all[0] {

		case "SET":
			_, ok := mm[all[1]]
			if(!ok) {
				mm[all[1]] = make(map[int]string)
			}
			mm[all[1]][level] = all[2]
			break
		case "GET":
			v, ok := mm[all[1]]
			maxs := "NULL"
			if(ok) {
				max := -1
				for k0, v0 := range v {
					if(k0 > max) {
						max = k0
						maxs = v0
						if(maxs == "") {
							maxs = "NULL"
						}
					}
				}
			}
			fmt.Println(maxs)
			break

		case "UNSET":
			mm[all[1]][level] = ""
			break

		case "NUMEQUALTO":
			count := 0
			for _, v := range mm {
				max := -1;
				for k0, _ := range v {
					if(k0 > max) {
						max = k0
					}
				}
				if(v[max] == all[1]) {
					count++
				}
		    }
			fmt.Println(count)
			break

		case "END":
			os.Exit(0)
			break;

		case "COMMIT":
			for _, v := range mm {
				max := -1
				maxs := ""
				for k0, v0 := range v {
					if(k0 > max) {
						max = k0
						maxs = v0
					}
					delete(v, max)
				}
				v[0] = maxs
			}
			level = 0;
	 		break

		case "BEGIN":
			level++
			break

		case "ROLLBACK":
			xacted := false
			if(level > 0) {
				for _, v := range mm {
					_, ok := v[level]
					if(ok) {
//						fmt.Printf("v = %v, ok = %v, level = %v\n", v, ok, level)
						delete(v, level)
						xacted = true
					}
				}
				level--
			}
			if(!xacted) {
				fmt.Println("NO TRANSACTION")
			}
			break
		}
	}
}