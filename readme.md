# Learning Go

## Objects in Go
- Go does not use the term class.
- A struct is a collection of data.
- Go uses structs with associated methods.
- Go uses a simplified object orientation:
	- no inheritance
	- no constructors
	- no generics

## Concurrency
..is the management of multiple tasks that (could rheoretically) run at the same time. Concurrent programming enables parallelism: 
- management of task execution
- communication between tasks
- synchronisation between tasks 
Gorutines represent concurrent tasks
Channels are used to communicate between tasks
Select enables task synchronisation 


## Pointer
- .. is an address to data in memory
- & operator returns  the address of some data
- * operator returns the data at some address
```
var x int = 1 	//declaration and initialisation of an int
var y int 		//declaration of an int
var ip *int		//declaration of a pointer to an int

ip = &x 			//setting ip to the address of x > ip now points to x
y = *ip			//setting y to the data stored at the address ip > y now is x > y = 1
----------------------------------------------------
var ptr *int = new(int) //new() creates a var and returns pointer to that var, therfore *int
*ptr = 3
fmt.Println(*ptr)		 //3
fmt.Println(ptr) 		//some address
```

## Variable Scope
A variable is accessible from block Bj if:
	- the variable is declared in Bi and 
	- Bi >= Bj

## Memory: Stack vs Heap
Stack: variables defined within a function 
	- Memory will be deallocated when function call completed (in C for example)
Heap: Global variables 
	- Needs to be explicitly deallocated ( in C for example), is fast but error-prone

## Type conversion
Use T() to convert, where T stands for the type you want to convert into
```
var x int32 = 1
var x int16 = 2
x = y //error
x = int32(y)
```

## ASCII
Each char is associated with an 8-bit number. 
8 bits —> 0-255, not enough for different alphabets and symbols

## Unicode
Each char is associated with an 32-bit number. 
Unicode is variable in length. From 8-bits to 32-bits.
8-bit UTF codes are the same as ASCII

## Strings
— a sequence of bytes, each byte is a “rune” (Go specific, = UTF-8 code point)
Good packages: Unicode (for individual runes), String (for strings)
```
var hello = "hello"
var test = hello[0]
fmt.Println(test) //104
fmt.Println(string(test)) //h
```

## Switch/case statement
If x is 1, enter case 1, if x is 2, enter case 2……..
Compared to C, no break statements are needed.  Go will not fall through, it automatically breaks.
```
switch x {
case 1: 
	fmt.printf("case 1")
case 1: 
	fmt.printf("case 2")
default:
	fmt.printf("default case")
```

## Exported names
When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package. A name is exported if it begins with a capital letter.

## Variables
```
var i, j int = 1, 2
var i, j = 1, 2 //int not needed, var will take the type of the initialiser

var c, python, java = true, false, "no!"
```

## Short variable declarations (only within functions)
```
func main() {
	var i, j int = 1, 2
	k := 3    					//within fucntions we can omit the var keyword and use := 
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

## Functions
```
func add(x, y int) int {
	return x + y
} //adds x and y, returns integer
```

## Function with multiple results
```
func swap(x, y string) (string, string) {
	return y, x
} // parameter x and y get swapped, two strings get returned

```

## Named return values
```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
} //a return statement without arguments, returns the named return values, here (x, y int)
```

## Arrays
- fixed length in Go
- Elements are accessed via [ ]
- Index starts at 0
```
var x [5]int		 //declares an integer array with 5 zeroed ints in it
x[0] = 2 			
________________________
//Array literals
var x [5] int = [5]{1,2,3,4,5}
x := [...]{1,2,3,4} 					//... is subsituted with  the length of the literal
___________________________
//Iterating through arry
for i, v range x {
	fmt.Printf("index %d, value %d", i, v) 
--> range returns two values, the current index and the value at the index
```

## Slices
- A window on an underlying array
- The pointer indicates the start of the array
- The length is the number of elements in the array
- The capacity is the max number of elements (from start of slice to end of array)
```
array := [...]int{1,2,3,4,5}
slice1 := array[1:3] 				//2,3
slice2 := array[2:5] 				//3,4,5
fmt.Printf(cap(slice2)) 			// 3, because starting from index 2, the slice can have 3 elements 										at most
```

- changing a slice changes the underlying array
- overlapping arrays refer to the same elements, changing it in one slice, also changes it for the other element

## Slice literal
- creates the underlying array and then the slice that refers to it
- that means the slice and the arrat are the same.
- slice points to the start of the array, length is capacity
```
//Sliche literal
sli := []{1,2,3} 			//slices don't need ... in the [ ]
```

## Variable slices
- you can also create a slice by using make()
- make either takes two to three arguments:
- Two: type and length 
 `sli = make([]int, `10) 			// then length and capacity ate equal
- Three arguments : : type and length and capacity
`sli = make(int, 10,15)
- Append()  lets you append elements. If there is space left in the underlying array, it will use that until it is full. Once full and you keep appending, it copies the content and creates a new, bigger underlying array


## Hash table
- contains key value pairs
- each value is associated with a unique key
- hash function is used to compute the slot for the key in the hash table. Input=key, output=slot to put value

## Maps (Go’s implementation of hash tables)
```
var idMap map[string]int 
idMap = make(map[string]int)
_____________________________
//Map literal
idMap := map[string]int {
	"joe": 23}

_____________________________
//Accessing maps
fmt.Printf(idMap["joe"]) 				//prints 23

_____________________________
//Adding a key value pair
idMap["Jane"] = 34					//there was no key "Jane" so the key/val pair gets added

______________________________
//Editing a key value pair
idMap["Jane"] = 98					//overrides the old key/val pair

____________________________
//Deleting a key/val pair
delete(idMap, "joe"

_____________________________
//Two value assignment to test whether key is in map
id, b := idMap["joe"]					//id=23, b=true, because joe is a key in the map

_____________________________	
//Number of key/val pairs
len(idMap)

_____________________________
//Iterating through map
for key, val := range idMap {
	fmt.Prinf(key, val) 
}
```

## Struct
- Aggregate data type
- Groups together objects of arbitrary types
```
type struct Person {
	name string
	addr string
	phone string
}
var p1 Person
p2 := new(Person) 		//creates new Person struct with all fields zeroed 
p1 := Person(name: "joe", addr: "a str", phone: "0123")			//struct literal
p1.name = "joe"


```
- 