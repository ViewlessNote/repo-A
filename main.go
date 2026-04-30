package main

import (
	"fmt"

	greeterv1 "github.com/ViewlessNote/repo-A/gen/go/greeter/v1"
)

// greet is a simple handler demonstrating the generated SDK types.
func greet(req *greeterv1.GreetRequest) *greeterv1.GreetResponse {
	return &greeterv1.GreetResponse{Message: fmt.Sprintf("Marco, %s!", req.GetVorName())}
}

func main() {
	req := &greeterv1.GreetRequest{VorName: "Polo"}
	resp := greet(req)
	fmt.Println(resp.GetMessage())
}
