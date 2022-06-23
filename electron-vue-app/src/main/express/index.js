const express = require('express')


const router =  require('./router')

function Start(port) {

    const app = express()
    
    router(app);


    //app.use("/", router)
    
    app.listen(port, () => {
      console.log(`Example app listening on port ${port}`)
    })

}





export default {Start}