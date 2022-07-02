
function testRouter(router){

  router.get("/",function(req,res,next){
    res.send("test");
  })
  
  return router;
}


function homeRouter(router){


  router.get("/",function(req,res,next){
    res.send("home");
  })
  

  return router;
}




module.exports = app => {

    app.use('/test', testRouter(app.Router()))
    app.use('/home', homeRouter(app.Router()))

};
