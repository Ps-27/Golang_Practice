package main
import (
	"fmt"
)
type contactInfo struct{
	email string
	zipcode int
}
type person struct{
	Fname string
	Lname string
	contactInfo
}

func main(){
	p := person{}
	fmt.Println(p)
	s := "prity"
	fmt.Println(&s)
	fmt.Println(s)
	s +="sinha"
	fmt.Println(&s)
	fmt.Println(s)

	jim:=person{
		Fname:"jim",
		Lname:"party",
		contactInfo:contactInfo{
			email:"jim@email.com",
			zipcode:1234,
		},
	}
	jim.print()
	jim.updateFname("jimmy")
	jim.print()

	jimPointer:=&jim
	jimPointer.updateFname("jimmmii")
	jimPointer.print()


	sl:=[]string{"hi","there","how","are","you"}
	updateSlice(sl)
	fmt.Println(sl) 


	name := "bill"
 
	namePointer := &name
	
	fmt.Println(&name)
	fmt.Println(&namePointer)
	fmt.Println(*namePointer)
	printPointer(namePointer)
	
}
func printPointer(namePointer *string) {
	fmt.Println(*namePointer)
	fmt.Println(namePointer)
	fmt.Println(&namePointer)
}
func updateSlice(s[]string){
	s[0]="Bye"
	fmt.Println(p)
}

func (p *person)updateFname(NewName  string){
	p.Fname=NewName
}
func (p person) print(){
	fmt.Printf(" %+v\n",p)
}