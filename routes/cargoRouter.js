var express = require('express');
var cargoRouter = express.Router();

//Middleware
var cargoModel = require('../model/cargoModel');

cargoRouter.post('/regcargo', async (req, res) => {
    try {
        var cargodata = {
            companycode: req.session.user.userID,
            startpoint: req.body.startpoint,
            endpoint: req.body.endpoint,
            carweight: req.body.carweight,
            weight: req.body.weight,
            transport: req.body.transport,
            cost: req.body.cost,
        }
        await cargoModel.cargoRegistration(cargodata);
        res.status(200).send(true);
    } catch(err) {
        res.status(500).send(err);
    }
});

cargoRouter.post('/delcargo', async (req, res) => {
    try {
        //get Company index number -> select enrolled cargo list
        var result = await cargoModel.callAllCargo(req.session.user);
        console.log(result[0])
        res.status(200).send(result[0])
    } catch (err) {
        res.status(500).send(err);
    }
});

cargoRouter.get('/components/jusoPopup', (req, res) => {
    res.render('components/juso.ejs');
})

module.exports = cargoRouter;