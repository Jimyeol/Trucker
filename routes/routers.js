var express = require('express');
var router = express.Router();

const userRouter = require('./userRouter');
const cargoRouter = require('./cargoRouter');
const mRouter = require('./mRouter');

router.use(userRouter);
router.use(cargoRouter);
router.use(mRouter);

module.exports = router;
