all: dockerimg run togofile
	@echo all

dockerimg:
	cd ../../../ && docker build -t eip7210.8.0:v1 -f ./contracts/token/ERC721/Dockerfile .
	@echo dockerimg

run:
	docker run -v /home/usera/projects/sol/openzeppelin-contracts/contracts/token/ERC721/outputeip721:/outputeip721 eip7210.8.0:v1 --allow-paths /source -o /outputeip721 --abi --overwrite --bin /source/contracts/token/ERC721/ERC721.sol
	@echo dockerrun

togofile:
	abigen --abi=./outputeip721/ERC721.abi --bin=./outputeip721/ERC721.bin --pkg=eipnft --out=./eip721.go


