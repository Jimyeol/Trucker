var myConnection = require('../dbConfig.js');


class  User {
    login(req) {
        var InsertedUser = req.body.user;
        var InsertedPassword = req.body.password;
        return new Promise(
            async (resolve, reject) => {
                const sql = `SELECT * FROM kyeonggidb WHERE user = ? AND password = ?`;
                try {
                    //geth 에서 지갑 Info load하는 코드
                    var result = await myConnection.query(sql, [InsertedUser, InsertedPassword]);
                    resolve(result[0][0]);
                } catch (err) {
                    reject(err);
                }
            }
        )
    }
}

module.exports = new User();