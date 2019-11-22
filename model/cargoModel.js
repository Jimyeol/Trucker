var myConnection = require('../dbConfig');

class Cargo {
    cargoRegistration(data) {
        return new Promise(
            async (resolve, reject) => {
                try {
                    console.log('cargoRegi', data)
                    var sql = 'INSERT INTO cargodb (company, startpoint, endpoint, carweight, weight, transport, cost) values (?, ?, ?, ?, ?, ?, ?)';
                    await myConnection.query(sql,[data.companycode, data.startpoint, data.endpoint, data.carweight, data.weight, data.transport, data.cost]);
                    resolve(true);
                } catch (err) {
                    reject(err);
                }
            }
        )
    }

    callAllCargo() {
        return new Promise (
            async (resolve, reject) => {
                try {
                    var sql = 'SELECT * FROM cargodb';
                    var result = myConnection.query(sql);
                    console.log(result[0]);
                    resolve(result);
                } catch(err) {
                    reject(err);
                }
            }
        )
    }

}

module.exports = new Cargo();