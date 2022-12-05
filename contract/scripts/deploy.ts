import { ethers } from "hardhat";
import { writeFile, readFile } from "fs/promises";

async function main() {
  let deployAddr: Record<string, string> = {};
  const Xbackup = await ethers.getContractFactory("Xbackup");
  const xbackup = await Xbackup.deploy("CortexAIModelToken", "CTXCDATAToken");
  await xbackup.deployed();
  console.log(`xbackup contract deploy on ${xbackup.address}`);
  deployAddr.xbackup = xbackup.address
  const data = JSON.stringify(deployAddr, null, 2);
  await writeFile(`./address_deploy.json`, data)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
