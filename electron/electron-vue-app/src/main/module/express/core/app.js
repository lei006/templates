const express = require('express')
let app = express()

import router from '../router'


app.Router = function() {
    return express.Router();
}


router(app);



export default app;
