var express = require('express');
var mRouter = express.Router();

var cargoModel = require('../model/cargoModel.js');


mRouter.get('/api/auth', async (req, res) => {
    try {
        console.log('api/auth', req.session.user)
        res.status(200).send(req.session.user);
    } catch (err) {
        console.log(err)
    }
});

mRouter.get('/api/cargo', async (req, res) => {
    try {
        console.log('hi')
        var result = await cargoModel.callAllCargo();
        console.log(result[0]);
        res.status(200).send(result[0])
    } catch(err) {
        console.log(err)
    }   
})

module.exports = mRouter;