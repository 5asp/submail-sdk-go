module github.com/t0of/submail-sdk-go

go 1.15

replace (
	github.com/t0of/submail-sdk-go => ./
	github.com/t0of/submail-sdk-go/example => ./example
	github.com/t0of/submail-sdk-go/service => ./service
	github.com/t0of/submail-sdk-go/submail => ./submail
	github.com/t0of/submail-sdk-go/submail/client => ./submail/client
	github.com/t0of/submail-sdk-go/submail/signer => ./submail/signer
)
