var myConnection = require('../dbConfig');
var web3 = require('web3');
// var Contract = require('./abi');
// var myContract = Contract.myContract;
// var web3 = Contract.web3;

var web3 = new web3(new web3.providers.HttpProvider('http://127.0.0.1:8545'));

class web3js {
    makeAccounts(data) {
        console.log('web3js로 넘어온 data', data);
        return new Promise(
            async (resolve, reject) => {
                try {
                    var newAccount = await web3.eth.personal.newAccount(data);
                    console.log('Ether Account', newAccount);
                    resolve(newAccount);
                } catch (err) {
                    console.log(err)
                    reject(err);
                }
            }
        );
    }
}

module.exports = new web3js();