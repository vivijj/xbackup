// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract Xbackup is ERC721URIStorage {
    // cortex data info hash => the json string of the url to get the data.
    mapping(string => string) public accessLinks;
    mapping(string => uint) public infohash2tokenid;

    // only operator can update the model url and mint the NFT
    address public operator;
    // the next token id to be mint, start from 1
    uint256 public freeTokenId;

    modifier onlyOperator() {
        require(msg.sender == operator);
        _;
    }

    constructor(string memory _name, string memory _symbol)
        ERC721(_name, _symbol)
    {
        operator = msg.sender;
        freeTokenId = 1;
    }

    function mintTokenForDataAuthor(
        string calldata _infoHash,
        string calldata _tokenUri,
        string calldata _link,
        address author
    ) public onlyOperator {
        super._safeMint(author, freeTokenId);
        super._setTokenURI(freeTokenId, _tokenUri);
        accessLinks[_infoHash] = _link;
        infohash2tokenid[_infoHash] = freeTokenId;
        freeTokenId = freeTokenId + 1;
    }

    // Note: the nft meta always have the access link too, so it should update together.
    function updateDataMeta(
        string memory _infoHash,
        string memory _newlinks,
        string calldata _newTokenUri
    ) external onlyOperator {
        accessLinks[_infoHash] = _newlinks;
        uint256 tokenid = infohash2tokenid[_infoHash];
        super._setTokenURI(tokenid, _newTokenUri);
    }
}
