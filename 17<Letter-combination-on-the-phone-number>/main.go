package main

import "fmt"

// str 2-9

func letterCombinations(digits string) []string {
  
  if digits == ""{
    return []string{}
  }

  type DigToLet struct {
    Key int
    Vals []string
  } 

  hash := func (a int)int{return a%8}
  var table [8]DigToLet

  var futureJ int = 0 
  var temp int
  for i := 0; i<=7; i++{
    var digValues []string
    if (i == 7) || (i == 5) {
      for j:=futureJ; j<4+futureJ; j++ {
        digValues = append(digValues, string(int('a')+j))
        temp = j
      }
      futureJ = temp+1
    }else{
      for j:=futureJ; j<3+futureJ; j++ {
        digValues = append(digValues, string(int('a')+j))
        temp = j
      }
      futureJ = temp+1
    }

    table[hash(i+2)] = DigToLet{i+2,digValues }
  }

  //fmt.Println(table)
  
  var result []string = []string{""}

  for _,digit := range digits{
    //table[hash(int(val))].Key 2,3
    var new_combo []string;
    sub_string := table[hash(int(digit))].Vals
    for _,combo := range result{
      for _,val2 := range sub_string {
        new_combo = append(new_combo, combo+val2)
      }
    } 
    result = new_combo
  }
  //fmt.Println(result)

  return result 
}


func main(){

  fmt.Println(letterCombinations(""))
}
