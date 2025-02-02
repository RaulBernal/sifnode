import {BigNumber, BigNumberish} from "ethers";
import {HardhatRuntimeEnvironment} from "hardhat/types";
import {DependencyContainer} from "tsyringe";
import {
    BridgeBankMainnetUpgradeAdmin,
    CosmosBridgeMainnetUpgradeAdmin,
    DeploymentChainId,
    DeploymentDirectory,
    DeploymentName
} from "./tsyringe/injectionTokens";
import {BridgeToken, BridgeToken__factory, ERC20} from "../build";
import {SignerWithAddress} from "@nomiclabs/hardhat-ethers/signers";
import {NotNativeCurrencyAddress} from "./ethereumAddress";
import {DeployedBridgeBank, DeployedBridgeToken} from "./contractSupport";

export const eRowanMainnetAddress = "0x07bac35846e5ed502aa91adf6a9e7aa210f2dcbe"

export async function impersonateAccount<T>(
    hre: HardhatRuntimeEnvironment,
    address: string,
    newBalance: BigNumberish | undefined,
    fn: (s: SignerWithAddress) => Promise<T>
) {
    await hre.network.provider.request({
        method: "hardhat_impersonateAccount",
        params: [address]
    });
    if (newBalance) {
        await setNewEthBalance(hre, address, newBalance)
    }
    const signer = await hre.ethers.getSigner(address)
    const result = await fn(signer)
    await hre.network.provider.request({
        method: "hardhat_stopImpersonatingAccount",
        params: [address],
    });
    return result
}

export async function startImpersonateAccount<T>(
    hre: HardhatRuntimeEnvironment,
    address: string,
    newBalance?: BigNumberish
): Promise<SignerWithAddress> {
    await hre.network.provider.request({
        method: "hardhat_impersonateAccount",
        params: [address]
    });
    if (newBalance) {
        await setNewEthBalance(hre, address, newBalance)
    }
    return await hre.ethers.getSigner(address)
}

export async function stopImpersonateAccount<T>(
    hre: HardhatRuntimeEnvironment,
    address: string,
) {
    await hre.network.provider.request({
        method: "hardhat_stopImpersonatingAccount",
        params: [address],
    });
}

export async function setNewEthBalance(
    hre: HardhatRuntimeEnvironment,
    address: string,
    newBalance: BigNumberish | undefined,
) {
    const newValue = BigNumber.from(newBalance)
    await hre.network.provider.send("hardhat_setBalance", [
        address,
        newValue.toHexString().replace(/0x0+/, "0x")
    ]);
}

export async function setupSifchainMainnetDeployment(c: DependencyContainer, hre: HardhatRuntimeEnvironment, deploymentName: string) {
    c.register(DeploymentDirectory, {useValue: "./deployments"})
    c.register(DeploymentName, {useValue: deploymentName})
    // We'd like to be able to use chainId from the provider,
    // but it doesn't actually work.  It returns 1 even when
    // you're looking at forked ropsten.
    // const chainId = (await hre.ethers.provider.getNetwork()).chainId
    let chainId = 0;
    switch (deploymentName) {
        case "sifchain":
        case "sifchain-1":
            chainId = 1
            break
        default:
            chainId = 3
            break
    }
    c.register(DeploymentChainId, {useValue: chainId})
    const bridgeTokenFactory = await hre.ethers.getContractFactory("BridgeToken") as BridgeToken__factory

    // BridgeToken for Rowan doesn't have a json file in deployments, so we need to build DeployedBridgeToken by hand
    // instead of using
    const existingRowanToken: BridgeToken = await bridgeTokenFactory.attach(eRowanMainnetAddress)
    const syntheticBridgeToken = {
        contract: Promise.resolve(existingRowanToken),
        contractName: () => "BridgeToken"
    }
    c.register(BridgeBankMainnetUpgradeAdmin, {useValue: "0x7c6c6ea036e56efad829af5070c8fb59dc163d88"})
    c.register(CosmosBridgeMainnetUpgradeAdmin, {useValue: "0x7c6c6ea036e56efad829af5070c8fb59dc163d88"})
    c.register(DeployedBridgeToken, {useValue: syntheticBridgeToken as DeployedBridgeToken})
}

export async function impersonateBridgeBankAccounts(c: DependencyContainer, hre: HardhatRuntimeEnvironment, deploymentName: string) {
    const bridgeBank = await c.resolve(DeployedBridgeBank).contract
    const operator = await bridgeBank.operator()
    const owner = await bridgeBank.owner()
    const pauser = owner // TODO you can't look up pausers, so probably need to add a pauser
    startImpersonateAccount(hre, operator)
    startImpersonateAccount(hre, owner)
    startImpersonateAccount(hre, pauser)

    await setNewEthBalance(hre, owner, BigNumber.from("100000000000000000000"))
}

export async function setupRopstenDeployment(c: DependencyContainer, hre: HardhatRuntimeEnvironment, deploymentName: string) {
    c.register(DeploymentDirectory, {useValue: "./deployments"})
    c.register(DeploymentName, {useValue: deploymentName})
    c.register(DeploymentChainId, {useValue: 3})
}

export async function approveThenDo<T>(
    token: ERC20,
    agent: NotNativeCurrencyAddress,
    amount: BigNumberish,
    account: SignerWithAddress,
    fn: () => Promise<T>
) {
    await token.connect(account).approve(agent.address, amount)
    return fn()
}
