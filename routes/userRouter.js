var express = require('express');
var userRouter = express.Router();

//Middleware
var userModel = require('../model/userModel');
var registerUser = require('../model/registerUser');

userRouter.get('/', (req, res, next) => {
    res.render('index.ejs')
});

userRouter.post('/login', async (req, res) => {
    try {
        await userModel.login(req);
        console.log('login session', req.session.user)
        res.status(200).send(req.session.user);
    } catch(err) {
        res.status(500).send(false);
    }
});

userRouter.get('/main', (req, res) => {
    res.render('main.ejs')
});

userRouter.post('/register', async (req, res) => {
    try {
        console.log(req.body)
        var result = await userModel.register(req);
        console.log('register', req.session.user)
        await registerUser.registerUser(req.session.user.userID);
        res.status(200).send(result)
    } catch(err) {
        console.log(err);
    }
});

module.exports = userRouter;