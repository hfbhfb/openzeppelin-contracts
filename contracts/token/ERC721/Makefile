all: dockerimg run togofile
	@echo all

dockerimg:
	cd ../../../ && docker build -t Eip7210.8.0:v1 -f ./contracts/token/ERC721/Dockerfile .
	@echo dockerimg

run:
	docker run -v /home/usera/projects/sol/openzeppelin-contracts/contracts/token/ERC721/outputEip721:/outputEip721 Eip7210.8.0:v1 --allow-paths /source -o /outputEip721 --abi --overwrite --bin /source/contracts/token/ERC721/ERC721.sol
	@echo dockerrun

togofile:
	abigen --abi=./outputEip721/ERC721.abi --bin=./outputEip721/ERC721.bin --pkg=eiptwostore --out=./Eip721.go


