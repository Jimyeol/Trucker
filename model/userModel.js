var myConnection = require('../dbConfig.js');


class  User {
    login(req) {
        var ca = req.body.ca;
        var InsertedUser = req.body.user;
        var InsertedPassword = req.body.password;
        return new Promise(
            async (resolve, reject) => {
                try {
                    if(ca == 'web') {
                        const sql = `SELECT * FROM companydb WHERE user = ? AND password = ?`;
                        var result = await myConnection.query(sql, [InsertedUser, InsertedPassword]);
                        resolve(result[0][0]);
                    } else {
                        const sql = `SELECT * FROM driverdb WHERE user = ? AND password = ?`;
                        var result = await myConnection.query(sql, [InsertedUser, InsertedPassword]);
                        resolve(result[0][0]);
                    }
                } catch (err) {
                    reject(err);
                }
            }
        )
    }

    register(req) {
        var ca = req.body.ca;
        var InsertedUser = req.body.user;
        var InsertedPassword = req.body.password;
        return new Promise(
            async (resolve, reject) => {
                try {

                    if(ca == 'web') {
                        const sql = 'INSERT INTO companydb (user, password) values (?, ?)';
                        await myConnection.query(sql, [InsertedUser, InsertedPassword]);
                        req.session.user = {
                            userID: InsertedUser,
                        }
                    resolve(req.session.user);
                    } else {
                        const sql = 'INSERT INTO driverdb (user, password) values (?, ?)';
                        await myConnection.query(sql, [InsertedUser, InsertedPassword]);
                    }
                    
                } catch (err) {
                    reject(err);
                }
            }
        )
    }
}

module.exports = new User();