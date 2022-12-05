import { ethers } from "hardhat";
import { writeFile, readFile } from "fs/promises";
import { token } from "../typechain-types/@openzeppelin/contracts";

async function main() {
  const Xbackup = await ethers.getContractFactory("Xbackup");
  const xbackup = Xbackup.attach("0x60fA83Ea5dEC41f1D1a2724E06A40BE86aF325da");
  let res = await xbackup.accessLinks("b2f5b0036877be22c6101bdfa5f2c7927fc35ef8");
  console.log(`res is ${res}`);
  let tokenId = await xbackup.infohash2tokenid("b2f5b0036877be22c6101bdfa5f2c7927fc35ef8");
  console.log(`token id is ${tokenId}`);
  let owner = await xbackup.ownerOf(tokenId);
  console.log(`owner is ${owner}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
