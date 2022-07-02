
const {app} =  require('./core')


function web_server_start(port, callback) {

    

    //app.use("/", router)
    /*
    app.get('/test', function (req, res) {
      res.send('test');
    })
    */
 
    app.listen(port, () => {
      if(callback){
        callback({state:"started", port});
      }
      console.log(`web server listening at:${port}`)
    })

}



function Start(callback){

  let port =  8080
  web_server_start(port, function(state){
    if(callback){
      
      state.app = app;
      callback(state);
    }
  })


}





export default {Start}