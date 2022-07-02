const { ipcMain } = require('electron')


const {spawn} = require('child_process')



// 子进程，参考 https://www.cnblogs.com/chyingp/p/node-learning-guide-child_process.html
let child_process = undefined;
function db_server_start(port, path, callback) {
    
    let command = "./extensions/mongo/mongod.exe";
    let args = [`--dbpath=${path}`,`--port=${port}`];
    child_process = spawn(command, args)
    child_process.unref();

    child_process.on('exit', (code) => {
        child_process = undefined;
        if(callback) {
            callback({state:"stoped", pid:child_process.pid, code, port, path, command, args});
        }
    });
    
    if(callback) {
        callback({state:"started", pid:child_process.pid , port, path, command, args});
    }
}

//var kill = require('tree-kill');

function server_stop() {

    try {
        if(child_process) {
            //不知道为什么，无法关闭....
            child_process.kill('SIGINT');
            child_process.kill();
            child_process = undefined;
        }        
    } catch (error) {
        console.error("helper stop error:", error)
    }

}

let port = 27017;
let data_path = './extensions/mongo/db'

function Start(callback){

    db_server_start(port, data_path, function(state){
        if(callback) {
            callback(state);
        }
    });
    
}



export default {Start}
