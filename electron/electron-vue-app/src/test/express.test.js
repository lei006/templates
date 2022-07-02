import web_srv from '../main/module/express/index'
const request = require('supertest');



function run(fn, times) {
    while (times > 0) {
      fn();
      times--;
    }
}

web_srv.Start(function({app}){


    describe('Web服务器简单回应', function () {

        it('测试一: responds with /test', function (done) {

            request(app)
            .get('/test')
            .set('Accept', 'application/json')
            //.expect('Content-Type', /json/)
            .expect(200, done);

        });


        /*
        it('测试二: responds with /file_list', function (done) {
            request(app)
            .get('/file_list')
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)
            .expect(200)
            .end(function(err, res) {
                
                if (err) throw done(err);
                assert(JSON.stringify(res.body) == "[]")
                
                done()
            })        
        });

        it('测试三: responds with /download', function (done) {

            request(app)
            .get('/download?aaa=bbb')
            .set('Accept', 'application/json')
            .expect('Content-Type', /json/)
            .expect(200)
            .end(function(err, res) {
                if (err) throw done(err);
                assert(res.body.aaa==='bbb');
                done()
            })
        });
        it('测试四', function () {

            let file_list = [`${__filename}`,`${__dirname}`];
            //fileSrvMod._set_file_list();
            let ret_list = fileSrvMod._get_all_file(file_list)
            assert(ret_list.length > 0);
            assert(ret_list[0].fs_path !== undefined && ret_list[0].srv_path !== undefined);
        });
        */
    });




});


