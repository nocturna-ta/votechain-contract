generate-abi:
	@echo  ">> Generating ABI"
	@solcjs --abi --bin ./contracts/VotechainBase.sol -o ./binding/abi-bin
	@solcjs --abi --bin ./contracts/KPUManager.sol -o ./binding/abi-bin
	@solcjs --abi --bin ./contracts/VoterManager.sol -o ./binding/abi-bin
	@solcjs --abi --bin ./contracts/ElectionManager.sol -o ./binding/abi-bin
	@solcjs --abi --bin ./contracts/Votechain.sol -o ./binding/abi-bin

generate-gobind:
	@echo ">> Generating Go Bindings"
	@abigen --bin=./binding/abi-bin/contracts_VotechainBase_sol_VotechainBase.bin --abi ./binding/abi-bin/contracts_VotechainBase_sol_VotechainBase.abi --pkg votechainBase --out ./binding/votechainBase/VotechainBase.go
	@abigen --bin=./binding/abi-bin/contracts_KPUManager_sol_KPUManager.bin --abi ./binding/abi-bin/contracts_KPUManager_sol_KPUManager.abi --pkg kpuManager --out ./binding/kpuManager/KPUManager.go
	@abigen --bin=./binding/abi-bin/contracts_VoterManager_sol_VoterManager.bin --abi ./binding/abi-bin/contracts_VoterManager_sol_VoterManager.abi --pkg voterManager --out ./binding/voterManager/VoterManager.go
	@abigen --bin=./binding/abi-bin/contracts_ElectionManager_sol_ElectionManager.bin --abi ./binding/abi-bin/contracts_ElectionManager_sol_ElectionManager.abi --pkg electionManager --out ./binding/electionManager/ElectionManager.go
	@abigen --bin=./binding/abi-bin/contracts_Votechain_sol_Votechain.bin --abi ./binding/abi-bin/contracts_Votechain_sol_Votechain.abi --pkg votechain --out ./binding/votechain/Votechain.go

change-abi:
	@mv ./binding/contracts_Votechain_sol_Votechain.abi ./binding/Votechain.abi

generate-mock:
	@echo ">> Generating Mocks"
		mockery --dir=./interfaces --name=ElectionManagerInterface --output=./mocks --outpkg=mocks
		mockery --dir=./interfaces --name=VotechainBaseInterface --output=./mocks --outpkg=mocks
		mockery --dir=./interfaces --name=KpuManagerInterface --output=./mocks --outpkg=mocks
		mockery --dir=./interfaces --name=VoterManagerInterface --output=./mocks --outpkg=mocks
		mockery --dir=./interfaces --name=VotechainInterface --output=./mocks --outpkg=mocks


