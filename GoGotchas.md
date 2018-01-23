##### Go Gotchas
* __Opening Brace Can't Be Placed on a Separate Line__
    - syntax error: unexpected semicolon or newline before {
* __Unused variables__
    - Declared and unused local variables in function body -> Compile error (even if variable is initialized or assigned some value)
    - Exceptions: Global variables, Unused function arguments
* __Unused Imports__
    - Compile error when unused package remains in import statements
* __Short Variable Declarations Can Be Used Only Inside Functions__
    - Compile error -> non-declaration statement outside function body
    ```
    package main
    myvar := 1 //error
    func main() {
    }
    ```
* __Redeclaring Variables Using Short Variable Declarations__
    ```
    package main
    func main() {  
    one := 0
    one := 1 //error
    }
    ```
    Corect declaration
    ```
    package main
    func main() {  
    one := 0
    one, two := 1,2
    one,two = two,one
    }
    ```
* __Can't Use Short Variable Declarations to Set Field Values__
    - Use temporary variables or predeclare all your variables and use the standard assignment operator.
    ```
    package main
    import (  
    "fmt"
    )
    type info struct {  
      result int
    }
    func work() (int,error) {  
        return 13,nil  
    }
    func main() {  
      var data info
      data.result, err := work() //error
      var err
      data.result, err = work() // ok, works!
      fmt.Printf("info: %+v\n",data)
    }
    ```

* __Accidental Variable Shadowing__
    ```
    package main
    import "fmt"
    func main() {  
        x := 1
        fmt.Println(x)     //prints 1
        {
            fmt.Println(x) //prints 1
            x := 2
            fmt.Println(x) //prints 2
        }
        fmt.Println(x)     //prints 1 (bad if you need 2)
    }
    ```
    use : go tool vet -shadow main.go // Note that this will not report all shadowed variables.

* __Can't Use "nil" to Initialize a Variable Without an Explicit Type__
    ```
    package main
    func main() {  
        var x = nil //error
        var y interface{} = nil //works
        _ = x
        _ = y
    }
    ```

* __Using "nil" Slices and Maps__
    ```
    package main
    func main() {  
        var s []int
        s = append(s,1) //works
        var m map[string]int
        m["one"] = 1             //error --> panic: assignment to entry in nil map
        m = make(map[string]int) // initialize map first
        m["one"] = 1             // then it works
    }
    ```
