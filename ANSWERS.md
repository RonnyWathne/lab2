##Important Instructions

Please preserve the structure of this file, as it will subjected to *partial*
automatic analysis. **Only insert your answers by replacing the text `YOUR
ANSWER HERE`; do not delete anything else.**

Please use [markdown](https://help.github.com/articles/markdown-basics)
formating to typeset code and Unix commands with the backtick character, for
example, `ls -la`, or if you want to write code blocks, each line should be
indented with four spaces, as done in the code below:

    #include <stdio.h>
    
    int main(void) {
    	printf("Hello, world!\n");
    	return 0;
    }


##Exercises from the Go tour.golang.org

###Exercise 25

Answer: *

     pakage main
    
     func Sqrt(x float64) float64 { 
     y := float64(2) 
      for i:= float64(1); i < 11 ;i++ { 
        y = y - ((math.Pow(y,2) - x)/(2*y)) 
     } 
      return y 
    }*
Exercise 38

    Answer: *
    
    package main
    
    import "code.google.com/p/go-tour/pic"
    
    func Pic(dx, dy int) [][]uint8 {
    
    array := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
       array[i] = make([]uint8, dx)
    }
    
    var counteri uint8
    var counterj uint8
    
    for i := 0; i < dy; i++ {
       counteri ++
      for j := 0; j < dx; j++ {
    
              array[i][j] = counteri * counterj
    
        counterj ++
     }
    }   
    return array
    }
    
    func main() {
       pic.Show(Pic)
    }
*

Exercise 43
    
    Answer: *
    
    Package main
    
    import (
       "code.google.com/p/go-tour/wc"
      "strings"
    
    )

    func WordCount(s string) map[string]int {
     result := make(map[string]int) 
      field := strings.Fields(s)
     numberofwords := len(field)
    
      for i := 0; i < numberofwords; i++ {
         (result[field[i]]) ++
      }
      return result
    
    }
    func main() {
    
    wc.Test(WordCount)
    }
*
Exercise 58
    
    Answer: *
    
    package main
    
    import (
     "fmt"
    )
    
    type ErrNegativeSqrt float64
    
    func (e ErrNegativeSqrt) Error() string {
       return fmt.Sprintf("Error, can not be negative, input: %g", float64(e))
    }
    
    func Sqrt(f float64) (float64, error) {
    
      if f < 0 {
          return 0, ErrNegativeSqrt(f)
       }
      return 0, nil
    
    
      }

    func main() {
       fmt.Println(Sqrt(2))
       fmt.Println(Sqrt(-2))
       }*
       
Exercise 60

    Answer: * package main
    
    import (
     "net/http"
      "fmt"
    )
    
    type String string
    
    func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
       fmt.Fprint(w, s)
    }
    
    
    type Struct struct {
        Greeting string
        Punct    string
        Who      string
    }
    
    func (s Struct) ServeHTTP(
       w http.ResponseWriter,
      r *http.Request) {
      fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
    }
    
    type Hello struct{}
    
    func main() {
    
     http.Handle("/string", String("I'm a frayed knot."))
    
     http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    
      http.ListenAndServe("localhost:4000", nil)
    }
*

Exercise 63
    
    Answer: *
    
    package main
    
    import (
       "io"
      "os"
     "strings"
    )
    
    type rot13Reader struct {
    
       r io.Reader 
    }
    
    func (rot *rot13Reader) Read(bytearray []byte) (n int, err error) {
        n,err = rot.r.Read(bytearray)
        for i := 0; i < len(bytearray); i++ {
           if (bytearray[i] >= 'A' && bytearray[i] < 'N') || (bytearray[i] >='a' && bytearray[i] < 'n') {
               bytearray[i] += 13
           } else if (bytearray[i] > 'M' && bytearray[i] <= 'Z') || (bytearray[i] > 'm' && bytearray[i] <= 'z'){
               bytearray[i] -= 13
           }
       }
       return
    }

    func main() {
       s := strings.NewReader(
           "Lbh penpxrq gur pbqr!")
       r := rot13Reader{s}
      io.Copy(os.Stdout, &r)
    }
*
Go Language Questions

Write a loop that repeats exactly n times. (Please inline your code below, as long as it is just a few line code snippet).
Answer: i:=0 ; i<n ; i++{ }
What is the value of ok in the following example:
Answer: the value is String
Object-oriented programming: How can you define a “class” Message with two attributes Sender and Content in Go?
Answer: type Message struct { sender string content string}
How can a method CheckForError be defined on the above Message “class”, that returns an error? (The function body can be empty.)
Answer: func (e CheckForError) Error() { }
