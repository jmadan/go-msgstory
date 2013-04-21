all:
	go install

main:
	go build main.go
	go run main.go

clean:
	go clean

user:
	cd user/
	go build user.go
	go test -v
	cd ..
	go install user

circle:
	cd circle/
	go build circle.go
	go test -v
	cd ..
	go install circle

message:
	cd message/
	go build message.go
	go test -v
	cd ..
	go install message

geolocation:
	cd geolocation/
	go build geolocation.go
	go test -v
	cd ..
	go install geolocation

register:
	cd register/
	go build register.go
	go test -v
	cd ..
	go install register

authentication:
	cd authentication/
	go build authentication.go
	go test -v
	cd ..
	go install authentication


test:
	go test -v ./message
