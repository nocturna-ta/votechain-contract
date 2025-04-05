generate-abi:
	@echo  ">> Generating ABI"
	@solcjs --abi ./contracts/Votechain.sol -o ./binding

generate-gobind:
	@echo ">> Generating Go Bindings"
	@abigen --abi ./binding/Votechain.abi --pkg binding --type Votechain --out ./binding/Votechain.go

change-abi:
	@mv ./binding/contracts_Votechain_sol_Votechain.abi ./binding/Votechain.abi

generate-mock:
	@echo ">> Generating Mocks"
		mockery --dir=./interfaces --name=IVotechainCaller --output=./mocks --outpkg=mocks --filename=votechain_caller_mock.go
		mockery --dir=./interfaces --name=IVotechainTransactor --output=./mocks --outpkg=mocks --filename=votechain_transactor_mock.go
		mockery --dir=./interfaces --name=IVotechainFilterer --output=./mocks --outpkg=mocks --filename=votechain_filterer_mock.go
		mockery --dir=./interfaces --name=IVotechain --output=./mocks --outpkg=mocks --filename=votechain_mock.go


