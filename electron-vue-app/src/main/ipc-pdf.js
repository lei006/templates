const { ipcMain } = require('electron')
const axios = require('axios')

const Store = require('electron-store');
const store = new Store();




var webapi_url = "http://api.pdf2.eyoujs.com:8060/v1"
//var webapi_url = "http://127.0.0.1:8060/v1"

async function PdfLogin(username, password,hardsn) {
    let url = webapi_url + "/user/login"
    let ret = await axios({method: 'post', url: url, data: {username: username,password: password,hardsn:hardsn}});
    return ret.data;
}

async function PdfRegedit(username, password,hardsn) {
  let url = webapi_url + "/user/regedit"
  let ret = await axios({method: 'post', url: url, data: {username: username,password: password,hardsn:hardsn}});
  return ret.data;
}

async function PdfChangePass(username, password,new_password) {
  let url = webapi_url + "/user/changepass"
  let ret = await axios({method: 'post', url: url, data: {username: username,password: password,new_password:new_password}});
  return ret.data;
}



async function PdfAutoLogin(username, password,hardsn) {
  let url = webapi_url + "/user/autologin"
  let ret = await axios({method: 'post', url: url, data: {username: username,password: password,hardsn:hardsn}});
  return ret.data;
}


async function PdfSendSms(phone) {
    let url = webapi_url + "/user/send_sms/" + phone
    let ret = await axios({method: 'get', url: url});
    return ret.data;
}

async function PdfPriceList() {
    let url = webapi_url + "/prices"
    let ret = await axios({method: 'get', url: url});
    return ret.data;
}

async function PdfCreateOrder(pay_type, username, price_id) {
    let url = webapi_url + "/order"
    let ret = await axios({method: 'post', url: url, data: {username: username,price_id:price_id, pay_type:pay_type}});
    return ret.data;
}

async function PdfGetHelper(key) {
    let url = webapi_url + "/helper/" + key;
    let ret = await axios({method: 'get', url: url});
    return ret.data;
}





ipcMain.on('send-sms', async (event, arg) => {
  
  let ret = await PdfSendSms(arg)
  event.sender.send("send-sms-ack", ret);
})


ipcMain.on('login', async (event, username, checkcode) => {

  let hardsn = getHardsn()

  let ret = await PdfLogin(username, checkcode, hardsn)
  if (ret.code === 20000) {
    store.set("userinfo", ret.data);
    event.sender.send("userinfo", ret.data);
  }
  event.sender.send("login-ack", ret);
})


ipcMain.on('regedit', async (event, username, password) => {

  let hardsn = getHardsn()

  let ret = await PdfRegedit(username, password, hardsn)
  if (ret.code === 20000) {
    store.set("userinfo", ret.data);
    event.sender.send("userinfo", ret.data);
  }
  event.sender.send("regedit-ack", ret);
})


ipcMain.on('logout', async (event) => {
  store.set("userinfo", null);
  event.sender.send("userinfo", null);
})


ipcMain.on('user-changepass', async (event, password, new_password) => {

  let userinfo = store.get("userinfo")
  if (userinfo == null || userinfo == undefined) {
    return;
  }



  let ret = await PdfChangePass(userinfo.username, password, new_password)
  event.sender.send("user-changepass-ack", ret);
})



// 同步取硬件id
function getHardsn() {
  return _conver_sync("hardsn");
}


ipcMain.on('autologin', async (event) => {

  let userinfo = store.get("userinfo")
  if (userinfo == undefined || userinfo == null) {
    return;
  }


  let username = userinfo.username;
  let password = userinfo.autopassword;
  let hardsn = getHardsn()

  let ret = await PdfAutoLogin(username, password, hardsn)
  if (ret.code === 20000) {
    store.set("userinfo", ret.data);
    event.sender.send("userinfo", ret.data);
  }
})

ipcMain.on('price-list', async (event) => {
  let ret = await PdfPriceList()
  if (ret.code === 20000) {
    event.sender.send("price-list-ack", ret.data.items);
  }
})


ipcMain.on('create-order', async (event, pay_type, price_id) => {

  let userinfo = store.get("userinfo")

  let ret = await PdfCreateOrder(pay_type, userinfo.username, price_id)
  if (ret.code === 20000) {
    console.log("PdfCreateOrder ret ok:", ret)

    event.sender.send("create-order-ack", ret.data);
  }else{
    console.log("PdfCreateOrder ret error:", ret)
  }

})



ipcMain.on('menu-cmd', async (event, cmd) => {

  let ret = await PdfGetHelper(cmd)
  if (ret.code === 20000) {
    event.sender.send("menu-cmd-ack", cmd, ret.data);
  }else{
    console.log("menu-cmd ", cmd, " err:", ret)
  }

})



function ReplaceFileName(filename) {
  return filename.replace(/\s+/g,"ACmAp;");
}


//取提pdf页面数量
ipcMain.on('pdf2pagenum', (event, arg1, arg2, arg3, arg4) => {
  let fileNamePath = arg1;
  fileNamePath = ReplaceFileName(fileNamePath);

  let random_id = arg2;
  _conver_async("pdf2pagenum " + fileNamePath, function(stdout){
    let pagenum = stdout;
    event.sender.send('pdf2pagenum_reply', pagenum, random_id);
  })
});


//取提pdf转file
ipcMain.on('pdf2file', (event, random_id, infile, outfile, outFormat) => {

  console.log("pdf2file-------------->", random_id)

  infile = ReplaceFileName(infile);
  outfile = ReplaceFileName(outfile);


  _conver_async(`pdf2file ${infile} ${outfile} ${outFormat}`, function(result){
    event.sender.send('pdf2file_reply', random_id, result );
  })
});

//取提pdf转file
ipcMain.on('file2pdf', (event, random_id, infile, outfile) => {

  infile = ReplaceFileName(infile);
  outfile = ReplaceFileName(outfile);

  _conver_async(`file2pdf ${infile} ${outfile}`, function(result){
    event.sender.send('file2pdf_reply', random_id, result );
  })
});

//所有图像文件合并到一个pdf
ipcMain.on('img2pdf', (event, infile, outfile) => {

  infile = ReplaceFileName(infile);
  outfile = ReplaceFileName(outfile);


  _conver_async(`img2pdf ${infile} ${outfile}`, function(result){
    console.log(`img2pdf ${infile} ${outfile}`)
    event.sender.send('img2pdf_reply', result );
  })
});




ipcMain.on('pdf2merge', (event, infile, outfile) => {

  infile = ReplaceFileName(infile);
  outfile = ReplaceFileName(outfile);


  _conver_async(`pdf2merge ${infile} ${outfile}`, function(result){
    console.log("pdf2merge",result )
    event.sender.send('pdf2merge_reply', result );
  })
});

ipcMain.on('pdf2pickimage', (event, random_id, infile, outpath) => {

  infile = ReplaceFileName(infile);
  outpath = ReplaceFileName(outpath);

  _conver_async(`pdf2pickimage ${infile} ${outpath}`, function(result){
    event.sender.send('pdf2pickimage_reply', random_id, result );
  })
});

ipcMain.on('pdf2split', (event, random_id, infile, outpath) => {

  infile = ReplaceFileName(infile);
  outpath = ReplaceFileName(outpath);

  _conver_async(`pdf2split ${infile} ${outpath}`, function(result){
    event.sender.send('pdf2split_reply', random_id, result );
  })
});

ipcMain.on('pdf2compress', (event, random_id, infile, outfile) => {

  infile = ReplaceFileName(infile);
  outfile = ReplaceFileName(outfile);

  _conver_async(`pdf2compress ${infile} ${outfile} 30`, function(result){
    event.sender.send('pdf2compress_reply', random_id, result );
  })
});


//取硬件id(同步)
ipcMain.on('hardsn', (event) => {
  event.returnValue = getHardsn();
});




//是vip用户吗?

function user_is_vip() {
  //确保用户是vip
  let userinfo = store.get("userinfo")
  if (userinfo == undefined || userinfo == null) {
    return false;
  }

  if (userinfo.vip_type == undefined || userinfo.vip_type == null || userinfo.vip_type == "free") {
    return false;
  }

  return true;
}




function user_is_login() {

  //确保用户已经登录
  let userinfo = store.get("userinfo")
  if (userinfo == undefined || userinfo == null) {
    return false;
  }

  if (userinfo.vip_type == undefined || userinfo.vip_type == null) {
    return false;
  }

  return true;
}





ipcMain.on('is-vip', (event) => {
  event.returnValue = user_is_vip();
});


ipcMain.on('buy-vip', (event) => {

  if (user_is_login() == true) {
    event.sender.send("buy-vip-ack");
  }else{
    event.sender.send("show-user-login");
  }


});


ipcMain.on('user-to-regedit', (event) => {
    event.sender.send("show-user-regedit");
});

ipcMain.on('user-to-login', (event) => {
    event.sender.send("show-user-login");
});


function get_vip_flag() {
  if (user_is_vip() == false) {
    return "test";
  }

  return "eyoujs";
}



function _conver_async(cmd, stdout_callback) {
  const { exec } = require('child_process');


  let flag = get_vip_flag();

  let command = '.\\helper\\FFileConverter.exe ' + flag + " " + cmd;

  console.log("exec run -->", command);


  exec(command, (error, stdout, stderr) => {
    if (error) {
      console.error(`执行出错: ${error}`);
      return;
    }
    if (stdout == "success") {
      stdout = "转化完成";
    }
    if (stdout == "error") {
      stdout = "转化出错";
    }


    stdout_callback(stdout);
  });
}

function _conver_sync(cmd) {
  const { execSync } = require('child_process');
  var output = execSync('.\\helper\\FFileConverter.exe ' + cmd); 
  return output.toString('utf8');
}




