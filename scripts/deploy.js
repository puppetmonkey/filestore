async function main() {
  // 获取合约工厂
  const FileStorage = await ethers.getContractFactory("FileStorage");

  // 部署合约
  const fileStorage = await FileStorage.deploy();

  console.log("FileStorage deployed to:", fileStorage.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
