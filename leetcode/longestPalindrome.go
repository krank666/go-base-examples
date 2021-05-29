package main

import (
	"fmt"
)
func longestPalindrome(s string) string {

	slice := make([]string,0,4)

	for _ , char := range s {
		slice = append(slice,"#")
		slice = append(slice,string(char))
	}

	slice = append(slice,"#") //先把字符串加上“#”

	maxR := 0    //记录最长的半径
	maxIndex := 0  //记录最长的 index
	sliceLen := len(slice)

	for index , _ := range slice{

		if(index >=1){

			r := 0
			i := index - 1
			j := index + 1

			for {       //每一个字符 计算最长的半径

				if i<0 || j >= sliceLen{
					break
				}

				if(slice[i] == slice[j]){
					r++
					i--
					j++
				}else{
					break
				}

			}

			if r > maxR{
				maxR = r
				maxIndex = index
			}

		}
	}

	res := ""

	for index,str := range slice{  //maxIndex-maxR到maxIndex+maxR是最长串

		if index >= (maxIndex-maxR) && index <= (maxIndex+maxR) && str != "#"{
			res += str
		}
	}

	return res

}

func main(){
	str :="12312他师傅日日日十多个确认他认为而非委任为"
	fmt.Println(longestPalindrome(str))
}