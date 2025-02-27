const Votechain = artifacts.require("Votechain")

module.exports = function (deployer) {
    deployer.deploy(Votechain).then(async (instance) =>{
        const accounts = await web3.eth.getAccounts()
        await instance.setKpuAdmin(accounts[1])
    })
}