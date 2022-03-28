all: dockerimg run togofile
	@echo all

dockerimg:
	docker build -t eip200.8.0:v1 .
	@echo dockerimg

run:
	docker run -v /home/usera/projects/sol/openzeppelin-contracts:/sources eip200.8.0:v1 -o /sources/contracts/token/ERC20/outputeip20 --abi --overwrite --bin ERC20.sol
	@echo dockerrun

togofile:
	abigen --abi=./outputeip20/ERC20.abi --bin=./outputeip20/ERC20.bin --pkg=eiptwostore --out=./Eip20.go


