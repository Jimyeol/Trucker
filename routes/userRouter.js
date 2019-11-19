var express = require('express');
var userRouter = express.Router();

//Middleware
var userModel = require('../model/userModel');



userRouter.get('/', (req, res, next) => {
    res.render('index.ejs')
});

userRouter.post('/login', async (req, res) => {
    try {
        console.log(req.body);
        await userModel.login(req);
        res.status(200).send(true);
    } catch(err) {
        res.status(500).send(false);
    }
});

userRouter.get('/main', (req, res) => {
    res.render('main.ejs')
});

userRouter.post('/register', async (req, res) => {
    try {
        console.log(req.body);
        await userModel.register(req);
        res.status(200).send(req.session.user)
    } catch(err) {
        console.log(err);
    }
});

userRouter.post('/regcargo', async (req, res) => {
    try {
        console.log(req.body);
        res.status(200).send('hi');
    } catch(err) {

    }
});

userRouter.post('/delcargo', async (req, res) => {
    try {
        //get Company index number -> select enrolled cargo list
    } catch (err) {

    }
});

module.exports = userRouter;