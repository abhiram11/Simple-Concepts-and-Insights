pragma solidity >=0.7.0 <0.9.0;

contract Kudos {
    int startingNumber;
    
    function setStartingNumber(int startNum) public {
        startingNumber = startNum;
    }
    
    function getSum(int numToAdd) public view returns (int) {
        return startingNumber + numToAdd;
    }
    
}
