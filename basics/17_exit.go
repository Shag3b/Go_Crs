package basics

import "os"

func main() {
	defer println("defer exe") //will not run
	println("start of main")
	//exit
	os.Exit(1)

	//will not exe
	println("end of main")
}

/*
## Practical Use Cases
==>> Error Handling
==>> Termination Conditions
==>> Exit Codes
## Best Practices
==>> Avoid Deferred Actions
==>> Status Codes// 0->mean success else->error
==>> Avoid Abusive Use*/
