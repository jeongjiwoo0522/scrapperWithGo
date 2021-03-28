package something

import "fmt"

func sayBye() { // private function
	fmt.Println("Bye")
}

func SayHello() { // public function
	fmt.Println("Hello")
}
