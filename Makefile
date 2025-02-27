generate-abi:
	@echo  ">> Generating ABI"
	@solcjs --abi ./contracts/Votechain.sol -o ./binding

generate-gobind:
	@echo ">> Generating Go Bindings"
	@abigen --abi ./binding/Votechain.abi --pkg binding --type Votechain --out ./binding/Votechain.go

change-abi:
	@mv ./binding/contracts_Votechain_sol_Votechain.abi ./binding/Votechain.abi