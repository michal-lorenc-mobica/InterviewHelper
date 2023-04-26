package main

// to control correctness in element amount for tables
const itemAmount int = 23

//each line is a single question
var questions = [itemAmount]string{
	"What are diferences between table and slice?",
	"What are modules?",
	"How to convert string to []byte?",
	"When pointers are equal?",
	"What is \"defer function\"?",
	"Can two maps can be compared directly?",
	"Is go a strong type language?",
	"Is there a pointer arithmetics in golang?",
	"Is there an exponentiation operator in golang?",
	"Is there a function overloading in golang?",
	"What is iota?",
	"What is RWMutex?",
	"How to invoke Go Race Detector for main.go file?",
	"What are GOOS and GOARCH flags?",
	"When two tables are equal?",
	"What is an anonymous struct? Where can you use them?",
	"Why to use Interfaces?",
	"What is interface{} type?",
	"What does \"concurrency\" is?",
	"What does \"parallelism\" is?",
	"goroutine vs thread",
	"When can panic occur in Go?",
	"Write an example of two-dimensional array"}

//each line is a single answer
var answers = [itemAmount]string{
	"- slice dimension is dynamic\n- uninitialized slice is NIL\n- uninitialized table has default values for used type",
	"- improved dependency management system\n- a collection of related go packages in a directory with a common go.mod file",
	"byteSlice := []byte(\"text\")",
	"Pointers are equal if they point to the same variable or when they are NIL",
	"- function will be called just before exiting the code block where it is defined (after the body of the code block has been executed)\n- called in reverse order in the code (from the end)",
	"You cannot compare maps with each other directly, but you can convert them to a string form and then compare them OR you can use reflect.DeepEqual",
	"Yes",
	"No",
	"No, there is special function POW(x,y) for that",
	"No",
	"- number generator for const counting from 0\n- resets on the next const word encountered",
	"- multiple goroutines can read data\n- ony one goroutin can write data\n- reading will not be perform during writing",
	"go run -race main.go",
	"They enable code compilation for selected platforms (example: \"linux amd64\")",
	"Arrays are equal if they have the same number of arguments, the same values, and the same order of elements",
	"An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code. If a struct is only meant to be used once (example: parsing specific JSON), then it makes no sense to declare it in a form usable for other developers.",
	"- interfaces are named collections of method signatures\n- data encapsulation method\n- to use an interface, you must implement all its methods\n- better to use more smaller interfaces than one big one",
	"- interface that has no methods\n- all types satisfy the empty interface\n",
	"- we only do one process at a time, but we have several to do at the same time\n- sometimes one process has to wait for another's completion",
	"- running several processes simultaneously\n- requires several CPUs",
	"- goroutines are executed by the Go runtime program which runs in user space above the kernel known as the Go scheduler\n- kernel is responsible for scheduling threads\n- goroutines are faster that threads in general\n- goroutines requires less space on stack than threads",
	"Go panics when:\n- developer used panic() function\n- there is logical error in code (like sending data to closed channel)",
	"tableExample := [2][3] int {\n		{5,6,7},\n		[3]int{8,9,10},\n		}\n"}
