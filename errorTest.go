package main

import (
	"errors"
	"fmt"
	"strings"
)

type SudokuError []error

const rows,columns = 9,9

type Grid [rows] [columns] int8
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit = errors.New("invalid digit")
)
func main(){
	var g Grid
	err := g.Set(-1,0,15)
	if (err != nil) {
		switch err {
		case ErrBounds,ErrDigit:
			fmt.Println( "哈，出错了，错误是",err)
		default:
			fmt.Println("未知错误",err)

		}
	}

}
func (g *Grid) Set(row,column int,digit int8) error {
	var errs SudokuError
	if !inBounds(row,column) {
		errs = append (errs,ErrDigit)
	}
	if !validDigit(digit) {
		errs = append(errs,ErrDigit)
	}
	if (len(errs) > 0) {
		return errs
	}
	g [row] [column] = digit
	return nil
}
func validDigit(digit int8) bool {
	return digit >= 1 && digit <=9
}
func inBounds(row,column int) bool {
	if row < 0 ||row >= rows {
		return false
	}
	if column < 0 || column>=columns  {
		return false
	}
	return true
}


func (se SudokuError) Error() string{
	var s []string
	for _,err := range se {
		s = append(s,err.Error())
	}
	return strings.Join(s,",")
}