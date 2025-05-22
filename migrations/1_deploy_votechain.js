const VotechainBase = artifacts.require("VotechainBase");
const KPUManager = artifacts.require("KPUManager");
const VoterManager = artifacts.require("VoterManager");
const ElectionManager = artifacts.require("ElectionManager");
const Votechain = artifacts.require("Votechain");
module.exports = async function(deployer, network, accounts) {
    const admin = accounts[0];

    // console.log("Deploying VotechainBase...");
    // await deployer.deploy(VotechainBase);
    // const baseInstance = await VotechainBase.deployed();
    // console.log(`VotechainBase deployed at: ${baseInstance.address}`);
    //
    // console.log("Deploying KPUManager...");
    // await deployer.deploy(KPUManager, baseInstance.address);
    // const kpuManagerInstance = await KPUManager.deployed();
    // console.log(`KPUManager deployed at: ${kpuManagerInstance.address}`);
    //
    // console.log("Deploying VoterManager...");
    // await deployer.deploy(VoterManager, baseInstance.address, kpuManagerInstance.address);
    // const voterManagerInstance = await VoterManager.deployed();
    // console.log(`VoterManager deployed at: ${voterManagerInstance.address}`);

    console.log("Deploying ElectionManager...");
    await deployer.deploy(ElectionManager, '0xBcDc1f66f7C4adEE4980567a29ba3687a7FCB294', '0xac10688B8C5ed19007CC6c1296E8155C3630db94');
    const electionManagerInstance = await ElectionManager.deployed();
    console.log(`ElectionManager deployed at: ${electionManagerInstance.address}`);

    // console.log("Deploying main Votechain contract...");
    // await deployer.deploy(
    //     Votechain,
    //     baseInstance.address,
    //     kpuManagerInstance.address,
    //     voterManagerInstance.address,
    //     electionManagerInstance.address
    // );
    // const votechainInstance = await Votechain.deployed();
    // console.log(`Votechain deployed at: ${votechainInstance.address}`);

    // console.log("\n======== Contract Addresses ========");
    // console.log(`VotechainBase: ${baseInstance.address}`);
    // console.log(`KPUManager: ${kpuManagerInstance.address}`);
    // console.log(`VoterManager: ${voterManagerInstance.address}`);
    // console.log(`ElectionManager: ${electionManagerInstance.address}`);
    // console.log(`Main Votechain: ${votechainInstance.address}`);
    // console.log("===================================\n");

};