var myConnection = require('../dbConfig');

class Cargo {
    cargoRegistration(data) {

        return new Promise(
            async (resolve, reject) => {
                try {
                    var sql = 'INSERT INTO cargodb (company, startpoint, endpoint, carweight, weight, transport, cost) values (?, ?, ?, ?, ?, ?, ?)';
                    await myConnection.query(sql,[data.companycode, data.startpoint, data.endpoint, data.carweight, data.weight, data.transport, data.cost]);
                } catch (err) {
                    reject(err);
                }
            }
        )
    }

    callAllCargo(data) {
        return new Promise (
            async (resolve, reject) => {
                try {
                    console.log(data);
                    var sql = 'SELECT * FROM cargodb WHERE company = ?';
                    var result = myConnection.query(sql, [data.userID]);
                    resolve(result);
                } catch(err) {
                    reject(err);
                }
            }
        )
    }

}

module.exports = new Cargo();