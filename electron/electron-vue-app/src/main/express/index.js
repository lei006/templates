const express = require('express')


const router =  require('./router')

function Start(port) {

    const app = express()
    
    router(app);


    //app.use("/", router)
    
    app.listen(port, () => {
      console.log(`web server listening at:${port}`)
    })

}





export default {Start}