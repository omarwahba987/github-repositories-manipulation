	This is a cmd program that can retreive repositories data from github.com
		 based on some specifics will shown below
			you can control results with number limit, date and key word can be programing language or any thing else,
			you can choose to use one, two, three or none of these filters
	you can clone the project and run " go mod tidy " to import unexisting packeges
	
	for run 
		you can change directory to program directory
 		for linux and mac
			1- go build main.go
			2- ./main  -key=golang -limit=10 -date=2010-01-01
		for windows 
			1-  go run main.go -key=golang -limit=10 -date=2020-01-01
 	for testing
		there are two testing files you can run them by 
			go run test