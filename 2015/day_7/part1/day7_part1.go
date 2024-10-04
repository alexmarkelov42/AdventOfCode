package day7_part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type wire struct {
	expr     string
	value    uint16
	isValSet bool
}

var wires = map[string]*wire{}

func RunLogicGates(filename string) (result uint16) {
	for key, val := range wires {
		fmt.Println(key, *val)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		tokens := strings.Split(data, " -> ")
		wires[tokens[1]] = &wire{}
		wires[tokens[1]].expr = tokens[0]
	}
	result = getValue("a")
	return result
}

func setValue(wireName string, value uint16) {
	wire1 := wires[wireName]
	wire1.value = value
	wire1.isValSet = true
}

func getValue(wireName string) (val uint16) {
	wire1 := wires[wireName]
	if wire1.isValSet {
		return wire1.value
	}
	val = evalExpr(wireName)
	return val
}

func evalExpr(wireName string) (result uint16) {
	wire1 := wires[wireName]
	tokens := strings.Split(wire1.expr, " ")
	if isNumber(wire1.expr) {
		tmp, _ := strconv.Atoi(string(wire1.expr))
		result = uint16(tmp)
		setValue(wireName, result)
		return result
	}
	if strings.Index(wire1.expr, "NOT") != -1 {
		in1 := tokens[1]
		result := getValue(in1)
		setValue(wireName, ^result)
		return ^result
	}
	if strings.Index(wire1.expr, "AND") != -1 {
		var (
			in1val uint16
			in2val uint16
		)
		in1 := tokens[0]
		in2 := tokens[2]
		if isNumber(in1) {
			tmp, _ := strconv.Atoi(string(in1))
			in1val = uint16(tmp)
		} else {
			in1val = getValue(in1)
		}
		if isNumber(in2) {
			tmp, _ := strconv.Atoi(string(in2))
			in2val = uint16(tmp)
		} else {
			in2val = getValue(in2)
		}
		setValue(wireName, in1val&in2val)
		return in1val & in2val
	}
	if strings.Index(wire1.expr, "OR") != -1 {
		var (
			in1val uint16
			in2val uint16
		)
		in1 := tokens[0]
		in2 := tokens[2]
		if isNumber(in1) {
			tmp, _ := strconv.Atoi(string(in1))
			in1val = uint16(tmp)
		} else {
			in1val = getValue(in1)
		}
		if isNumber(in2) {
			tmp, _ := strconv.Atoi(string(in2))
			in2val = uint16(tmp)
		} else {
			in2val = getValue(in2)
		}
		setValue(wireName, in1val|in2val)
		return in1val | in2val
	}
	if strings.Index(wire1.expr, "LSHIFT") != -1 {
		in1 := tokens[0]
		num, _ := strconv.Atoi(tokens[2])
		in1val := getValue(in1)
		result = in1val << num
		setValue(wireName, result)
		return result
	}
	if strings.Index(wire1.expr, "RSHIFT") != -1 {
		in1 := tokens[0]
		num, _ := strconv.Atoi(tokens[2])
		in1val := getValue(in1)
		result = in1val >> num
		setValue(wireName, result)
		return result
	}
	result = evalExpr(tokens[0])
	setValue(wireName, result)
	return result
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func parseNum(expr string) uint16 {
	regexprNumber, _ := regexp.Compile("[0-9]+")
	num := regexprNumber.Find([]byte(expr))
	tmp, _ := strconv.Atoi(string(num))
	return uint16(tmp)
}
