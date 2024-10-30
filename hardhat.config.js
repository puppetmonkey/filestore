require("@nomiclabs/hardhat-waffle");

module.exports = {
  solidity: "0.8.0", // 根据你的合约版本进行调整
  networks: {
    hardhat: {
      // 设定自定义的持久化数据目录
      dataDir: "/home/xusir/project/filestore/pkg/eth/hardhat-data",
      chainId: 1337,
    },
  },
};
