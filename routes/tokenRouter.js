var express = require('express');
var tokenRouter = express.Router();

//Middleware
var tokenModel = require('../model/cargoModel');

tokenRouter.get('/token', (req, res) => {

});

module.exports = tokenRouter;