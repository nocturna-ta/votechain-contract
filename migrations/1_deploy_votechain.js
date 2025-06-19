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

    // Define trusted forwarders based on network
    let trustedForwarders = [];

    if (network === 'development' || network === 'ganache') {
        // For local development, you can use empty array or add test forwarders
        trustedForwarders = [
            "0xF63B38349538202bB2444BD6928e11F2ff68f239"
        ];

        console.log("Local development - No trusted forwarders initially");
    } else if (network === 'mumbai' || network === 'polygon-mumbai') {
        // Biconomy forwarder for Mumbai testnet
        trustedForwarders = [
            "0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b" // Biconomy Mumbai forwarder
        ];
    } else if (network === 'polygon' || network === 'matic') {
        // Biconomy forwarder for Polygon mainnet
        trustedForwarders = [
            "0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8" // Biconomy Polygon forwarder
        ];
    } else {
        // For other networks, start with empty array
        trustedForwarders = [];
        console.log("Unknown network - No trusted forwarders initially");
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
        baseInstance.address,           // Use newly deployed base
        kpuManagerInstance.address,     // Use newly deployed kpuManager
        voterManagerInstance.address,   // Use newly deployed voterManager
        electionManagerInstance.address, // Use newly deployed electionManager
        trustedForwarders               // Add trusted forwarders array
    );
    const votechainInstance = await Votechain.deployed();
    console.log(`   ✅ Votechain deployed at: ${votechainInstance.address}`);

    // 6. Verify trusted forwarders are set correctly
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

    // 7. Verify admin permissions
    console.log("\n======== Verifying Admin Setup ========");
    const kpuAdmin = await baseInstance.kpuAdmin();
    const votingActive = await baseInstance.votingActive();
    console.log(`KPU Admin: ${kpuAdmin}`);
    console.log(`Admin matches deployer: ${kpuAdmin.toLowerCase() === admin.toLowerCase() ? '✅' : '❌'}`);
    console.log(`Voting Active: ${votingActive}`);

    // 8. Create deployment summary
    const deploymentSummary = {
        network: network,
        deployer: admin,
        timestamp: new Date().toISOString(),
        contracts: {
            VotechainBase: baseInstance.address,
            KPUManager: kpuManagerInstance.address,
            VoterManager: voterManagerInstance.address,
            ElectionManager: electionManagerInstance.address,
            Votechain: votechainInstance.address
        },
        trustedForwarders: trustedForwarders,
        gasUsed: {
            // Note: You can track gas usage if needed
        }
    };

    console.log("\n======== Deployment Summary ========");
    console.log(`Network: ${deploymentSummary.network}`);
    console.log(`Deployer: ${deploymentSummary.deployer}`);
    console.log(`Timestamp: ${deploymentSummary.timestamp}`);
    console.log("\n📋 Contract Addresses:");
    console.log(`VotechainBase: ${deploymentSummary.contracts.VotechainBase}`);
    console.log(`KPUManager: ${deploymentSummary.contracts.KPUManager}`);
    console.log(`VoterManager: ${deploymentSummary.contracts.VoterManager}`);
    console.log(`ElectionManager: ${deploymentSummary.contracts.ElectionManager}`);
    console.log(`Main Votechain: ${deploymentSummary.contracts.Votechain}`);
    console.log(`\n🔗 Trusted Forwarders: ${trustedForwarders.length > 0 ? trustedForwarders.join(', ') : 'None initially set'}`);
    console.log("=====================================");

    // 9. Save deployment info to file (optional)
    const fs = require('fs');
    const path = require('path');

    try {
        const deploymentDir = path.join(__dirname, '..', 'deployments');
        if (!fs.existsSync(deploymentDir)) {
            fs.mkdirSync(deploymentDir, { recursive: true });
        }

        const filename = `${network}_deployment_${Date.now()}.json`;
        const filepath = path.join(deploymentDir, filename);

        fs.writeFileSync(filepath, JSON.stringify(deploymentSummary, null, 2));
        console.log(`\n💾 Deployment info saved to: deployments/${filename}`);
    } catch (error) {
        console.log(`\n⚠️  Could not save deployment info: ${error.message}`);
    }

    // 10. Next steps guidance
    console.log("\n======== Next Steps ========");
    console.log("1. 🔧 Configure meta-transaction service (Biconomy, OpenGSN, etc.)");
    if (trustedForwarders.length === 0) {
        console.log("2. ➕ Add trusted forwarder addresses:");
        console.log("   await baseInstance.addTrustedForwarder(forwarderAddress)");
    }
    console.log("3. 👥 Register KPU Provinsi accounts");
    console.log("4. 🗳️  Set up frontend for gasless transactions");
    console.log("5. ✅ Test the complete voting flow");
    console.log("=============================");

    return {
        base: baseInstance,
        kpuManager: kpuManagerInstance,
        voterManager: voterManagerInstance,
        electionManager: electionManagerInstance,
        votechain: votechainInstance
    };
};