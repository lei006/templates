



const fs = require('fs')

function read_config(filename, callback, err_callback) {

    let file = process.cwd() + '\\' + filename // 文件路径
    fs.readFile(file, 'utf-8', function (err, data) {
      if (err) {
        if (err_callback != undefined) {
          err_callback(err)
        }
      } else {

        let config = JSON.parse(data)
        callback(config)
      }
    })
  
  }

Date.prototype.Format = function (fmt) { //author: meizz 
    var o = {
        "M+": this.getMonth() + 1, //月份 
        "d+": this.getDate(), //日 
        "h+": this.getHours(), //小时 
        "m+": this.getMinutes(), //分 
        "s+": this.getSeconds(), //秒 
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度 
        "S": this.getMilliseconds() //毫秒 
    };
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
}

module.exports = {
    read_config,
}
