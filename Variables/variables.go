// Variables

bool  // True or False


string 


int int8 int16 int32 int64  // 8,16,32,64 represents bits

uint uint8 uint16 uint32 uint64 uintptr // uint means Unsigned Integer. uint represents non-negative whole numbers

byte // Its an alias for uint8. Alias means an alternative name for a thing.

rune // Its an alias for int32
     // Represents a Unicode code point


float32 float64 // Decimal values

complex64 complex128 // complex number 
                  	// Eg:- 5 + 6i


// Different methods to declare a variables
var number int 

var number int = 1 

number := 1 // Short assignment


// Same line Declaration
car, model := "BMW",1916

// is same as
car :="BMW"
model :=1916

// Constants
const pi=3.14