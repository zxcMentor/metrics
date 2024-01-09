package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func TimeUpdate() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:

			//fmt.Println("writing to file...")
			err := os.WriteFile("/app/static/tasks/_index.md", []byte(fmt.Sprint(content1, time.Now().Format(time.RFC1123), content2, b, content3)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}

func TreeUpdate() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:

			//fmt.Println("writing to file...")
			err := os.WriteFile("/app/static/tasks/_index.md", []byte(fmt.Sprint(content1, time.Now().Format(time.RFC1123), content2, b, content3)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func BinTreeBuilt() {
	output := content4

	arr := make([]int, 0)

	arr = sortIntArray(arr)

	t := time.NewTicker(5 * time.Second)

	for i := len(arr); i < 100; i++ {

	link1:
		newElement := rand.Intn(999)

		for i := 0; i < len(arr); i++ {
			if newElement == arr[i] {
				goto link1
			}
		}
		arr = append(arr, newElement)
		arr = sortIntArray(arr)

		balancedTree := recurseSplit(arr)

		output = binTreePrintRecurse(balancedTree)

		output = content4 + output + content5 + output + "\n{{< /mermaid >}}"

		err := os.WriteFile("/app/static/tasks/binary.md", []byte(fmt.Sprint(output)), 0644)
		if err != nil {
			log.Println(err)
		}

		<-t.C

	}

}

func recurseSplit(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	newNode := &Node{}

	if len(arr) == 1 {
		newNode.value = arr[0]
		//fmt.Println(newNode.value)
		return newNode
	}

	var max, ind int

	for i, element := range arr {
		if element > max {
			max = element
			ind = i
		}

		//fmt.Println("max element, index, arr:", max, ind, arr)

	}

	ind /= 2

	newNode.value = arr[ind]
	//fmt.Println(arr[ind])

	newNode.left = recurseSplit(arr[:ind])

	newNode.right = recurseSplit(arr[ind+1:])

	//fmt.Println(newNode.value, newNode.left, newNode.right)
	return newNode

}

func binTreePrintRecurse(bt *Node) string {
	var output string

	if bt.left != nil {
		output += strconv.Itoa(bt.value) + "-->" + strconv.Itoa(bt.left.value) + "\n"
		output += binTreePrintRecurse(bt.left)
	}

	if bt.right != nil {
		output += strconv.Itoa(bt.value) + "-->" + strconv.Itoa(bt.right.value) + "\n"
		output += binTreePrintRecurse(bt.right)
	}

	return output

}

func sortIntArray(arr []int) []int {
	var tmp int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp

			}

		}
	}
	return arr
}

type GraphElement struct {
	ID    int
	Name  string
	Form  string
	Links []*GraphElement
}

func graphRandomBuilt() {
	var output string

	form := []string{"circle", "rect", "square", "ellipse", "round-rect", "rhombus"}

	t := time.NewTicker(5 * time.Second)

	for {
		graphArr := make([]*GraphElement, 0)
		for i := 0; i < 15; i++ {
			newElement := &GraphElement{}
			newElement.ID = i
			newElement.Name = string(byte(i + 65))
			newElement.Form = form[rand.Intn(len(form)-1)]
			graphArr = append(graphArr, newElement)

			if i > 3 {

				for j := i - 3; j < i; j++ {

					if graphArr[j].Links != nil && len(graphArr[j].Links) > 1 {
						continue
					}

					rnd := rand.Intn(j)
					rnd1 := rand.Intn(i-j) + j + 1

					if graphArr[rnd].Links == nil {

						graphArr[rnd].Links = append(graphArr[rnd].Links, graphArr[j])

					} else {
						if len(graphArr[rnd].Links) < 2 && graphArr[rnd].Links[0] != graphArr[j] {
							graphArr[rnd].Links = append(graphArr[rnd].Links, graphArr[j])
						}
					}

					if graphArr[j].Links == nil {

						graphArr[j].Links = append(graphArr[j].Links, graphArr[rnd1])

					} else {

						if len(graphArr[j].Links) < 2 && graphArr[j].Links[0] != graphArr[rnd1] {

							graphArr[j].Links = append(graphArr[j].Links, graphArr[rnd1])

						}

					}

				}

			}

			if len(graphArr) > 4 {

				output = graphPrint(&graphArr)

				output = content6 + output + content7 + output + content8

				err := os.WriteFile("/app/static/tasks/graph.md", []byte(fmt.Sprint(output)), 0644)
				if err != nil {
					log.Println(err)
				}

				<-t.C

				//fmt.Println(output)

			}

		}

	}

}

func graphPrint(grArr *[]*GraphElement) string {
	var output, br1, br2 string

	mapForm := make(map[string]int)

	for i, gr := range *grArr {
		if i == len(*grArr)-1 {
			break
		}
		//fmt.Println("i", i)
		for j := 0; j < len(gr.Links); j++ {
			br1, br2 = "", ""

			if _, ok := mapForm[gr.Name]; !ok {
				output += gr.Name

				mapForm[gr.Name] = 1

				switch gr.Form {
				case "circle":
					br1 = "(("
					br2 = "))"
				case "rect":
					br1 = "["
					br2 = "]"
				case "square":
					br1 = "["
					br2 = "]"

				case "ellipse":
					br1 = "(["
					br2 = "])"

				case "round-rect":
					br1 = "("
					br2 = ")"

				case "rhombus":

					br1 = "{"
					br2 = "}"

				}

			}

			output += br1 + gr.Name + br2 + " --> "

			br1, br2 = "", ""

			if _, ok := mapForm[gr.Links[j].Name]; !ok {

				output += gr.Links[j].Name

				mapForm[gr.Links[j].Name] = 1

				switch gr.Form {
				case "circle":
					br1 = "(("
					br2 = "))"
				case "rect":
					br1 = "["
					br2 = "]"
				case "square":
					br1 = "["
					br2 = "]"

				case "ellipse":
					br1 = "(["
					br2 = "])"

				case "round-rect":
					br1 = "("
					br2 = ")"

				case "rhombus":

					br1 = "{"
					br2 = "}"

				}

			}

			output += br1 + gr.Links[j].Name + br2 + "\n"

		}

	}

	return output
}
