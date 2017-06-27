package main

import "fmt"

func main(){
	fmt.Println("-----Go in action #4------------------ ")

	fmt.Println("-----Arrays-----")
	var array1 [5]int
	array1[0]=10
	array2 :=[5]int{10,20,30,40,50}
	array2[2]=35
	array3 :=[5]int{1:10, 2:20}
	array3[1]=25
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	fmt.Println("-----Pointer Arrays-----")
	arrayPtr1 :=[2]*int{new(int),new(int)}
	fmt.Println(*arrayPtr1[0])
	*arrayPtr1[0]=100
	*arrayPtr1[1]=200
	fmt.Println(*arrayPtr1[0])

	fmt.Println("-----Assign one array to another-----")
	var array5 [2]string
	var array6 =[2]string{"red", "blue"}

	array5=array6
	fmt.Println(array5)
	fmt.Println(array6)

	fmt.Println("-----Assign one pointer array to another-----")

	var arrayPtr2A =[2]*string{new(string), new(string)}
	var arrayPtr2B =[2]*string{new(string), new(string)}
	*arrayPtr2A[0]="ppp"
	*arrayPtr2A[1]="qqq"
	fmt.Println(arrayPtr2A[0])
	fmt.Println(arrayPtr2B[0])
	fmt.Println("-----")
	fmt.Println(*arrayPtr2A[0])
	fmt.Println(*arrayPtr2B[0])
	fmt.Println("-----")
	arrayPtr2B=arrayPtr2A
	fmt.Println(*arrayPtr2A[0])
	fmt.Println(*arrayPtr2B[0])
	fmt.Println("-----")

	fmt.Println("-----Multidimensional array----------------")

	var arrayMulti_A =[2][3]int{{10,10},{20,20}}
	fmt.Println(arrayMulti_A)

}
