// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.

let oracleAddress = "0x2F894A010318637b8310210A990C42fa00b00a4a";
let swapRouter = "0x2fF7642CD6eD40CD7551ef217B3185E1e0389e35";
let feeAddress = "0xbceF988D62b043C59914509c02Aa8487f6487FF0";
let multiSignatureAddress = "0xa75d4642CeE17d908B79528B5f0Cb8f5647870D2";

const {ethers} = require("hardhat");

async function main() {

  // const [deployerMax,,,,deployerMin] = await ethers.getSigners();
  const [deployerMin,,,,deployerMax] = await ethers.getSigners();

  console.log(
    "Deploying contracts with the account:",
    deployerMin.address
  );

  console.log("Account balance:", (await deployerMin.getBalance()).toString());

  const pledgePoolToken = await ethers.getContractFactory("PledgePool");
  const pledgeAddress = await pledgePoolToken.connect(deployerMin).deploy(oracleAddress,swapRouter,feeAddress, multiSignatureAddress);


  console.log("pledgeAddress address:", pledgeAddress.address);
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });