/*
Learn how to use interface to implement OOPS concept
*/
package main 
import ( 
	"fmt" 
)
type Person struct{
	name string
}
type Animal struct{
	sound string
}
type Greeter interface{
	sayHi() string
}
//retrieve param by pointer's value
func (man *Person) sayHi() string{
	//update new name. since this is de-reference so the operation affects memory address
	man.name = "Jr. " + man.name;
	return "My name is " + man.name
}
func (dog Animal) sayHi() string{
	return dog.sound
}
func greet(gr Greeter) string {
	return gr.sayHi()
}
func main(){
	var man = Person{name:"Duy Tran"}
	var dog = Animal{sound:"Gaw Gaw!!!"}

	fmt.Println("People greet: ", greet(&man)) //pass reference to func param
	fmt.Println("After change: ", man.name)
	fmt.Println("Dogs greet: ", greet(dog))

}