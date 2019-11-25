var express = require('express');
var userRouter = express.Router();

//sessions
const session = require('express-session');
const FileStore = require('session-file-store')(session);

userRouter.use(session({
    secret: 'keyboard cat',
    resave: false,
    saveUninitialized: true,
    store: new FileStore(),
}));

//Middleware
var userModel = require('../model/userModel');
var ccModel = require('../model/ccModel');

userRouter.get('/', (req, res, next) => {
    res.render('index.ejs');
})

userRouter.post('/login', (req, res) => {
        console.log(req.body);
        //DB 측 login 정보 Match
        // var result = await userModel.login(req);
        req.session.user = {
            user : 'hi',
        }
        console.log('hi');
        res.redirect('/main');
});

userRouter.get('/main', (req, res) => {
    console.log(req.session.user);
    res.render('main.ejs')
});

userRouter.get('/register', (req, res) => {
    userModel.register(req);
    //req.session.user 형성
    res.redirect('/main');
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

userRouter.get('/asd', async (req, res) => {
    try {
        var result = await ccModel.getTruckByValue();
        
    } catch (err) {
        
    }
})


module.exports = userRouter;