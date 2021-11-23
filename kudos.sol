pragma solidity >=0.7.0 <0.9.0;

contract Kudos {
    //create a structure
    mapping (address => Kudo[]) allKudos; //similar to state variables

    //     function giveKudos(string memory what, address who, address giver, string memory comments) public {

    function giveKudos(string memory what, address who, string memory comments) public {
         
        
        Kudo memory kudo = Kudo(what, msg.sender, comments); //new Kudo(wont be used)
        
        //in the allKudos Database/Structure, search for Abhiram, and for him we get array,
        //push this new kudo in that array
        allKudos[who].push(kudo);
    }
    
    function getKudosLength(address who) public view returns(uint) {
        Kudo[] memory allKudosForWho = allKudos[who];
        return allKudosForWho.length;
    }
    
    function getKudosAtIndex(address who, uint idx) public view returns (string memory, address, string memory) { //(what, giver, comments)
    Kudo memory kudo = allKudos[who][idx];
    return (kudo.what, kudo.giver, kudo.comments);
    }
    
}

///define the structure
struct Kudo {
    string what;
    address giver;
    string comments;
}


//giving kudos: 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2

// 0xd9145CCE52D386f254917e481eB44e9943F39138 --> Kudos Deploy Memory
