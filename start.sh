go mod tidy
cd /usr/src/code/app/usercenter/cmd/api
go run usercenter.go -f etc/usercenter.yaml
cd /usr/src/code/app/usercenter/cmd/rpc
go run usercenter.go -f etc/usercenter.yaml

