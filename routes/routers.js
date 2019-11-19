var express = require('express');
var router = express.Router();

const userRouter = require('./userRouter');
const cargoRouter = require('./cargoRouter');

router.use(userRouter);
router.use(cargoRouter);

module.exports = router;
