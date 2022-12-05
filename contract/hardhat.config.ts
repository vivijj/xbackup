import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  networks: {
    hardhat: {
        chainId: 43,
        blockGasLimit: 8000000,
        accounts: {
            accountsBalance: "10000000000000000000000000",
        },
    },
    dev: {
        // this is the dev cortex network
        url: "http://0.0.0.0:8545/",
        chainId: 43,
    },
    ctxc: {
        url: "https://cortex.logistic.ml",
        chainId: 21, 
    },
},
  solidity: {
    version: "0.8.17",
    settings: {
        optimizer: {
            enabled: true,
            runs: 200,
        },
    },
},
};

export default config;
