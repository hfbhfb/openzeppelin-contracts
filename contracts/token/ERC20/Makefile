all: dockerimg run togofile
	@echo all

dockerimg:
	cd ../../../ && docker build -t eip200.8.0:v1 -f ./contracts/token/ERC20/Dockerfile .
	@echo dockerimg

run:
	docker run -v /home/usera/projects/sol/openzeppelin-contracts/contracts/token/ERC20/outputeip20:/outputeip20 eip200.8.0:v1 --allow-paths /source -o /outputeip20 --abi --overwrite --bin /source/contracts/token/ERC20/ERC20.sol
	@echo dockerrun

togofile:
	abigen --abi=./outputeip20/ERC20.abi --bin=./outputeip20/ERC20.bin --pkg=eiptwostore --out=./Eip20.go


