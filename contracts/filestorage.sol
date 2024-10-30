// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract FileStorage {
    struct File {
        string fileHash; // 外部文件哈希
        string ipfsHash; // IPFS hash
        string fileName; // 文件名称
        uint256 timestamp; // 上传时间
    }

    // 用户地址到文件数组的映射
    mapping(address => File[]) public userFiles;

    // 文件哈希到文件信息的映射（fileHash -> File）
    mapping(string => File) public fileHashToFile;

    // 用户上传文件
    function storeFile(
        string memory _fileHash,
        string memory _ipfsHash,
        string memory _fileName
    ) public {
        File memory newFile = File({
            fileHash: _fileHash,
            ipfsHash: _ipfsHash,
            fileName: _fileName,
            timestamp: block.timestamp
        });

        // 将文件添加到用户的文件数组
        userFiles[msg.sender].push(newFile);

        // 将文件哈希映射到文件信息
        fileHashToFile[_fileHash] = newFile;
    }

    // 通过文件索引获取文件信息
    function getFile(
        uint256 _index
    )
        public
        view
        returns (string memory, string memory, string memory, uint256)
    {
        require(_index < userFiles[msg.sender].length, "Invalid file index.");
        File memory file = userFiles[msg.sender][_index];
        return (file.fileHash, file.ipfsHash, file.fileName, file.timestamp);
    }

    // 通过外部文件哈希获取文件信息
    function getFileByHash(
        string memory _fileHash
    )
        public
        view
        returns (string memory, string memory, string memory, uint256)
    {
        File memory file = fileHashToFile[_fileHash];
        require(bytes(file.fileHash).length != 0, "File not found."); // 确保文件存在
        return (file.fileHash, file.ipfsHash, file.fileName, file.timestamp);
    }

    // 获取用户上传的文件数量
    function getUserFileCount() public view returns (uint256) {
        return userFiles[msg.sender].length;
    }
}
