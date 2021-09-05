import { HardhatNodeRunner } from "../src/devenv/hardhatNode";
import { GolangBuilder, GolangResults } from "../src/devenv/golangBuilder";
import { SifnodedRunner, ValidatorValues } from "../src/devenv/sifnoded";
import { DeployedContractAddresses } from "../scripts/deploy_contracts";
import { SmartContractDeployer } from "../src/devenv/smartcontractDeployer";
import { EbrelayerRunner } from "../src/devenv/ebrelayer";
import { EthereumAddressAndKey } from "../src/devenv/devEnv";
import { notify } from "node-notifier";

async function startHardhat() {
  const node = new HardhatNodeRunner()
  const resultsPromise = node.go()
  const results = await resultsPromise
  console.log(`rsltis: ${JSON.stringify(results, undefined, 2)}`)
  return { process, results }
}

async function golangBuilder() {
  const node = new GolangBuilder()
  const resultsPromise = node.go()
  const results = await resultsPromise
  console.log(`golangBuilder: ${JSON.stringify(results, undefined, 2)}`)
  const output = await Promise.all([process, results])
  return {
    process: output[0],
    results: output[1]
  }
}

async function sifnodedBuilder(golangResults: GolangResults) {
  console.log('in sifnodedBuilder')
  const node = new SifnodedRunner(golangResults)
  const resultsPromise = node.go()
  const results = await resultsPromise
  console.log(`golangBuilder: ${JSON.stringify(results, undefined, 2)}`)
  return {
    process,
    results
  }
}

async function smartContractDeployer() {
  const node: SmartContractDeployer = new SmartContractDeployer();
  const resultsPromise = node.go();
  const result = await resultsPromise;
  console.log(`Contracts deployed: ${JSON.stringify(result.contractAddresses, undefined, 2)}`)
  return { process, result };
}

async function ebrelayerBuilder(
  contractAddresses: DeployedContractAddresses,
  ethereumAccount: EthereumAddressAndKey,
  validater: ValidatorValues,
  golangResults: GolangResults
) {
  const node: EbrelayerRunner = new EbrelayerRunner({
    smartContract: contractAddresses,
    golangResults: golangResults,
    account: ethereumAccount,
    validatorValues: validater,
  });
  const resultsPromise = node.go();
  const result = await resultsPromise;
  return { process, result };
}

async function main() {
  try {
    const sigterm = new Promise((res, _) => {
      console.log("We are in sigterm, res:" + res);
      process.on('SIGINT', res);
      process.on('SIGTERM', res);
    });
    // cmd: golangBuilder: "make -C .. install"
    const golang = await golangBuilder();
    // golang.results: GolangResults {
    //   completed: true,
    //   error: undefined,
    //   output: "make: Entering directory '/mnt/shared/work/projects/sif/sifnode/local'\n" +
    //     `go install -ldflags '-X github.com/cosmos/cosmos-sdk/version.Name=sifchain -X github.com/cosmos/cosmos-sdk/version.ServerName=sifnoded -X github.com/cosmos/cosmos-sdk/version.ClientName=sifnoded -X github.com/cosmos/cosmos-sdk/version.Version=0.9.3-ibc -X github.com/cosmos/cosmos-sdk/version.Commit=9a293c0fe20dc06cf8144c2a92a0406c0f36c430 -X "github.com/cosmos/cosmos-sdk/version.BuildTags=testnet,"' -tags testnet ./cmd/sifnoded ./cmd/sifgen ./cmd/ebrelayer\n` +
    //     "make: Leaving directory '/mnt/shared/work/projects/sif/sifnode/local'\n",
    //   goBin: '/home/jurez/go/bin'
    // }
    
    // cmd: "node_modules/.bin/hardhat node --hostname localhost --port 8545"
    const hardhat = await startHardhat();
    // hardhat.results: {
    //   process: <ref *1> ChildProcess { ... },
    //   accounts: {
    //     proxyAdmin: {
    //       address: '0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266',
    //       privateKey: '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'
    //     },
    //     operator: {
    //       address: '0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266',
    //       privateKey: '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'
    //     },
    //     owner: {
    //       address: '0x70997970c51812dc3a010c7d01b50e0d17dc79c8',
    //       privateKey: '0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d'
    //     },
    //     pauser: {
    //       address: '0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc',
    //       privateKey: '0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a'
    //     },
    //     validators: [
    //       {
    //         address: '0x90f79bf6eb2c4f870365e785982e1f101e93b906',
    //         privateKey: '0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6'
    //       }
    //     ],
    //     available: [
    //       {
    //         address: '0x15d34aaf54267db7d7c367839aaf71a00a2c6a65',
    //         privateKey: '0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a'
    //       },
    //       {
    //         address: '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc',
    //         privateKey: '0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba'
    //       },
    //       {
    //         address: '0x976ea74026e726554db657fa54763abd0c3a0aa9',
    //         privateKey: '0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e'
    //       },
    //       {
    //         address: '0x14dc79964da2c08b23698b3d3cc7ca32193d9955',
    //         privateKey: '0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356'
    //       },
    //       {
    //         address: '0x23618e81e3f5cdf7f54c3d65f7fbc0abf5b21e8f',
    //         privateKey: '0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97'
    //       },
    //       {
    //         address: '0xa0ee7a142d267c1f36714e4a8f75612f20a79720',
    //         privateKey: '0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6'
    //       }
    //     ]
    //   },
    //   httpHost: 'localhost',
    //   httpPort: 8545,
    //   chainId: 1
    // }

    // const [hardhat, golang] = (await Promise.all([startHardhat(), golangBuilder()]))    
    // const hardhat = await startHardhat()
    // console.log(hardhat)


    const sifnode = await sifnodedBuilder(golang.results);
    console.log("golang.results: %o", golang.results);
    console.log("sifnode.results: %o", sifnode.results);
/*
    // sifnode.results: {
    //   validatorValues: [
    //     {
    //       chain_id: 'localnet',
    //       node_id: 'c4d27f85839701d7626cc286f8af2f1513ba3059',
    //       ipv4_address: '10.10.1.2',
    //       moniker: 'icy-sun',
    //       password: 'qYiUlbueB6rKMZnOVvykH2mt3jN9Ifa4',
    //       address: 'sif1uyxgl5cl0walc6pr4nrk4ce5qel30n23xms797',
    //       pub_key: 'sifpub1addwnpepqfa69p52kzj2ydkhrf25xgc302nk0kgxsttm6z6xjxvggxc6dtvdj3l7gvq',
    //       mnemonic: 'there matter glory wild damage viable sunset token father father crouch round friend seek potato primary affair teach onion shadow food capable pepper combine',
    //       validator_address: 'sifvalconspub1zcjduepqlqe3nd3vzxcq6z726u5vygdltt83e6zt39fcqdsycxrl3u3qptzqd7haxs',
    //       validator_consensus_address: 'sifvalcons1eyh97qcs9x2cc7x6f5uwfd5znxjyygzcs8wdxh',
    //       is_seed: true
    //     },
    //     [length]: 1
    //   ],
    //   process: <ref *1> ChildProcess {
    //     [domain]: null,
    //     _events: [Object: null prototype] {
    //       exit: [Function (anonymous)] { [length]: 1, [name]: '' }
    //     },
    //     _eventsCount: 1,
    //     _maxListeners: undefined,
    //     _closesNeeded: 1,
    //     _closesGot: 0,
    //     connected: false,
    //     signalCode: null,
    //     exitCode: null,
    //     killed: false,
    //     spawnfile: '/bin/sh',
    //     _handle: Process {
    //       onexit: [Function (anonymous)] { [length]: 2, [name]: '' },
    //       pid: 170551,
    //       [Symbol(owner_symbol)]: [Circular *1]
    //     },
    //     spawnargs: [
    //       '/bin/sh',
    //       '-c',
    //       '/home/jurez/go/bin/sifnoded start --minimum-gas-prices 0.5rowan --rpc.laddr tcp://0.0.0.0:26657 --home /tmp/sifnodedNetwork/validators/localnet/icy-sun/.sifnoded',
    //       [length]: 3
    //     ],
    //     pid: 170551,
    //     stdin: null,
    //     stdout: null,
    //     stderr: null,
    //     stdio: [ null, null, null, [length]: 3 ],
    //     [Symbol(kCapture)]: false
    //   },
    //   tcpurl: 'tcp://0.0.0.0:26657'
    // }

*/

    // cmd: "npx hardhat run scripts/deploy_contracts.ts --network localhost"    
    const smartcontract = await smartContractDeployer();
    // smartcontract.result: SmartContractDeployResult {
    // completed: true,
    // error: undefined,
    // output: 'No need to generate any newer typings.\n{"bridgeBank":"0x5FC8d32690cc91D4c39d9d3abcBD16989F875707","bridgeRegistry":"0x610178dA211FEF7D417bC0e6FeD39F05609AD788","rowanContract":"0x5FbDB2315678afecb367f032d93F642f64180aa3"}\n',
    // contractAddresses: {
    //     bridgeBank: '0x5FC8d32690cc91D4c39d9d3abcBD16989F875707',
    //     bridgeRegistry: '0x610178dA211FEF7D417bC0e6FeD39F05609AD788',
    //     rowanContract: '0x5FbDB2315678afecb367f032d93F642f64180aa3'
    // }

    // ebrelayer init $TCP_URL[tcp://0.0.0.0:26657] $WEBSOCKET_ADDRESS[ws://localhost:8545]
    //      $BRIDGE_REGISTRY $VALIDATOR_MONIKER "$VALIDATOR_MNEMONIC"
    //     --chain-id $CHAINNET[localnet] --keyring-backend test
    //     --from $VALIDATOR_MONIKER --symbol-translator-file
    //     ../test/integration/config/symbol_translator.json
    // Then ./src/devenv/wait_for_sif_account.py $VALIDATOR_ADDRESS
    //   result: SmartContractDeployResult {

    // const ebrelayer = await ebrelayerBuilder(
      // smartcontract.result.contractAddresses,
      // hardhat.results.accounts.validators[0],
      // sifnode.results.validatorValues[0],
      // golang.results
    // )


    console.log("Before await sigterm");
    await sigterm
    console.log("After await sigterm");
    console.log("Caught interrupt signal, cleaning up.");
    // sifnode.process.kill(sifnode.process.pid);
    // hardhat.process.kill(hardhat.process.pid);
    // ebrelayer.process.kill(ebrelayer.process.pid);
    console.log("All child process terminated, goodbye.");
    notify({
      title: "Sifchain DevEnvironment Notice",
      message: `Dev Environment has recieved either a SIGINT or SIGTERM signal, all process have exited.`
    })
  } catch (error) {
    console.log("Deployment failed. Lets log where it broke: ", error);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    if (typeof error == "number")
      process.exit(error)
    else {
      console.error(error);
      process.exit(1)
    }
  });
