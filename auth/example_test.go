package auth

import "fmt"

func ExampleAuthorization() {
	if err := Authorization(request); err != nil {
		panic(err)
	}
	// Output:
	//
}

func ExampleParseHeaderAuthorization() {
	authorizationInHeader := "SHA-256 Access=123, SignedHeaders=host;x-request-date, Signature=7b74a0ede6c7e3e89d307da6dcb41ce75ba7e29fa145dac7de39751c869f3a2a"
	headerAuthorization, err := ParseHeaderAuthorization(authorizationInHeader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Algorithm: %s\n", headerAuthorization.Algorithm)
	fmt.Printf("Access: %s\n", headerAuthorization.Access)
	fmt.Printf("SignedHeaders: %s\n", headerAuthorization.SignedHeaders)
	fmt.Printf("Signature: %s\n", headerAuthorization.Signature)
	// Output:
	// Algorithm: SHA-256
	// Access: 123
	// SignedHeaders: host;x-request-date
	// Signature: 7b74a0ede6c7e3e89d307da6dcb41ce75ba7e29fa145dac7de39751c869f3a2a
}
