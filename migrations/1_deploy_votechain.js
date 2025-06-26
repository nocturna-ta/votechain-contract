const MinimalForwarder = artifacts.require("MinimalForwarder");
const VotechainBase = artifacts.require("VotechainBase");
const KPUManager = artifacts.require("KPUManager");
const VoterManager = artifacts.require("VoterManager");
const ElectionManager = artifacts.require("ElectionManager");
const Votechain = artifacts.require("Votechain");

module.exports = async function(deployer, network, accounts) {
    const admin = accounts[0];

    console.log(`Deploying to network: ${network}`);
    console.log(`Admin account: ${admin}`);
    console.log(`Admin balance: ${web3.utils.fromWei(await web3.eth.getBalance(admin), 'ether')} ETH`);

    // Deploy MinimalForwarder first
    console.log("0. Deploying MinimalForwarder...");
    await deployer.deploy(MinimalForwarder);
    const forwarderInstance = await MinimalForwarder.deployed();
    console.log(`   ✅ MinimalForwarder deployed at: ${forwarderInstance.address}`);

    // Define trusted forwarders, starting with the newly deployed MinimalForwarder
    let trustedForwarders = [forwarderInstance.address];

    // Add additional forwarders based on network (if needed)
    if (network === 'development' || network === 'ganache') {
        console.log("Local development - Using deployed MinimalForwarder");
    } else if (network === 'mumbai' || network === 'polygon-mumbai') {
        trustedForwarders.push("0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b"); // Biconomy Mumbai forwarder
    } else if (network === 'polygon' || network === 'matic') {
        trustedForwarders.push("0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8"); // Biconomy Polygon forwarder
    } else {
        console.log("Unknown network - Using only deployed MinimalForwarder");
    }

    console.log("Trusted forwarders:", trustedForwarders);
    console.log("\n======== Starting Deployment ========");

    // 1. Deploy VotechainBase with trusted forwarders
    console.log("1. Deploying VotechainBase...");
    await deployer.deploy(VotechainBase, trustedForwarders);
    const baseInstance = await VotechainBase.deployed();
    console.log(`   ✅ VotechainBase deployed at: ${baseInstance.address}`);

    // 2. Deploy KPUManager
    console.log("2. Deploying KPUManager...");
    await deployer.deploy(KPUManager, baseInstance.address, trustedForwarders);
    const kpuManagerInstance = await KPUManager.deployed();
    console.log(`   ✅ KPUManager deployed at: ${kpuManagerInstance.address}`);

    // 3. Deploy VoterManager
    console.log("3. Deploying VoterManager...");
    await deployer.deploy(VoterManager, baseInstance.address, kpuManagerInstance.address, trustedForwarders);
    const voterManagerInstance = await VoterManager.deployed();
    console.log(`   ✅ VoterManager deployed at: ${voterManagerInstance.address}`);

    // 4. Deploy ElectionManager
    console.log("4. Deploying ElectionManager...");
    await deployer.deploy(ElectionManager, baseInstance.address, voterManagerInstance.address, trustedForwarders);
    const electionManagerInstance = await ElectionManager.deployed();
    console.log(`   ✅ ElectionManager deployed at: ${electionManagerInstance.address}`);

    // 5. Deploy main Votechain contract
    console.log("5. Deploying main Votechain contract...");
    await deployer.deploy(
        Votechain,
        baseInstance.address,
        kpuManagerInstance.address,
        voterManagerInstance.address,
        electionManagerInstance.address,
        trustedForwarders
    );
    const votechainInstance = await Votechain.deployed();
    console.log(`   ✅ Votechain deployed at: ${votechainInstance.address}`);

    // Verification and summary steps remain the same
    console.log("\n======== Verifying Trusted Forwarders ========");
    for (const forwarder of trustedForwarders) {
        const isBaseTrusted = await baseInstance.isTrustedForwarder(forwarder);
        const isKpuTrusted = await kpuManagerInstance.isTrustedForwarder(forwarder);
        const isVoterTrusted = await voterManagerInstance.isTrustedForwarder(forwarder);
        const isElectionTrusted = await electionManagerInstance.isTrustedForwarder(forwarder);
        const isVotechainTrusted = await votechainInstance.isTrustedForwarder(forwarder);

        console.log(`Forwarder ${forwarder}:`);
        console.log(`   - Base: ${isBaseTrusted ? '✅' : '❌'}`);
        console.log(`   - KPUManager: ${isKpuTrusted ? '✅' : '❌'}`);
        console.log(`   - VoterManager: ${isVoterTrusted ? '✅' : '❌'}`);
        console.log(`   - ElectionManager: ${isElectionTrusted ? '✅' : '❌'}`);
        console.log(`   - Votechain: ${isVotechainTrusted ? '✅' : '❌'}`);
    }

    // Deployment summary
    const deploymentSummary = {
        network: network,
        deployer: admin,
        timestamp: new Date().toISOString(),
        contracts: {
            MinimalForwarder: forwarderInstance.address,
            VotechainBase: baseInstance.address,
            KPUManager: kpuManagerInstance.address,
            VoterManager: voterManagerInstance.address,
            ElectionManager: electionManagerInstance.address,
            Votechain: votechainInstance.address
        },
        trustedForwarders: trustedForwarders
    };

    console.log("\n======== Deployment Summary ========");
    console.log(`Network: ${deploymentSummary.network}`);
    console.log(`Deployer: ${deploymentSummary.deployer}`);
    console.log(`Timestamp: ${deploymentSummary.timestamp}`);
    console.log("\n📋 Contract Addresses:");
    console.log(`MinimalForwarder: ${deploymentSummary.contracts.MinimalForwarder}`);
    console.log(`VotechainBase: ${deploymentSummary.contracts.VotechainBase}`);
    console.log(`KPUManager: ${deploymentSummary.contracts.KPUManager}`);
    console.log(`VoterManager: ${deploymentSummary.contracts.VoterManager}`);
    console.log(`ElectionManager: ${deploymentSummary.contracts.ElectionManager}`);
    console.log(`Main Votechain: ${deploymentSummary.contracts.Votechain}`);
    console.log(`\n🔗 Trusted Forwarders: ${trustedForwarders.join(', ')}`);
    console.log("=====================================");
};