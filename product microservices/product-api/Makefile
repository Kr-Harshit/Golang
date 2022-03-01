check_install_swagger:
	which swagger || GO111MODULE="off" go get github.com/go-swagger/go-swagger/cmd/swagger

generate_swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	swagger generate client -f ./swagger.yaml -A product-api -t ./testing

start_server:
	go run main.go

start_client:
	go test main_test.go

