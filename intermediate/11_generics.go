package intermediate

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }
type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s stack[T]) print() {
	if len(s.elements) == 0 {
		fmt.Println("stack is empty.")
		return
	}
	fmt.Print("stack elemnts: ")
	for _, v := range s.elements {
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	// x, y := 1, 2
	// fmt.Println(swap(x, y))
	// x1, y1 := "dog", "cat"
	// fmt.Println(swap(x1, y1))
	list1 := stack[int]{}
	list1.push(1)
	list1.push(2)
	list1.push(3)
	list1.push(4)
	list1.print()

	fmt.Println(list1.pop())
	fmt.Println(list1.pop())
	fmt.Println(list1.pop())
	fmt.Println(list1.pop())
	list1.print()

	fmt.Println(list1.isEmpty())

	stringStack := stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.push("John")
	stringStack.print()
	fmt.Println(stringStack.pop())
	fmt.Println("Is stringStack empty:", stringStack.isEmpty())
	stringStack.print()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stringStack empty:", stringStack.isEmpty())
}

/*
• Benefits of Generics
==>>• Code Reusability
==>>• Type Safety
==>>• Performance
• Considerations
==>>• Type Constraints
==>>• Documentation
==>>• Testing
*/
