
const { app, BrowserWindow, dialog, globalShortcut, ipcMain, session } = require('electron')




const fs = require('fs')
const path = require('path')

import modeDB from './mode_db'

process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';


/**
 * Set `__static` path to static files in production
 * https://simulatedgreg.gitbooks.io/electron-vue/content/en/using-static-assets.html
 */
if (process.env.NODE_ENV !== 'development') {
  global.__static = require('path').join(__dirname, '/static').replace(/\\/g, '\\\\')
}

const winURL = process.env.NODE_ENV === 'development'
  ? `http://localhost:9080`
  : `file://${__dirname}/index.html`

//app.commandLine.appendSwitch('ignore-certificate-errors')




let g_debug = false


let win_main = null;

let g_userinfo = {};



function createWindow () {  

  // 创建浏览器窗口
  win_main = new BrowserWindow({
    width: 800,
    height: 600,
    show: false,
    frame:false,
    minWidth:640,
    minHeight:480,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false
    },
    titleBarStyle:"hidden",
    hasShadow:true,
    autoHideMenuBar:(g_debug==false),
  })

  win_main.once('ready-to-show', () => {
    win_main.show()
  })


  //main.js
  win_main.on('close', (event) => {
    
  });
  //win_main.loadFile('./src/views/main.html')
  win_main.loadURL(winURL)


}

app.whenReady().then(createWindow)



ipcMain.on('app-mini', (event, arg) => {
  win_main.minimize();
})


ipcMain.on('app-exit', (event, arg) => {
  app.exit(0)
})



//打开目录
ipcMain.on('open-directory-dialog', async function (event, filters) {
  let ret = await dialog.showOpenDialog({properties: ['openDirectory']});
  event.returnValue = ret;
})


ipcMain.on('app-path', async function (event) {
  //event.returnValue = app.getAppPath();

  let doc_path = app.getPath("documents") + "\\海豚PDF转化器";

  mkdirsSync(doc_path)

  event.returnValue = doc_path;
})

function mkdirsSync(dirname) {
  if (fs.existsSync(dirname)) {
    return true;
  } else {
    if (mkdirsSync(path.dirname(dirname))) {
      fs.mkdirSync(dirname);
      return true;
    }
  }
}











/*


//得到目录下的文件
ipcMain.on('get-directory-file', async function (event, filters) {

  let outfilelist = [];

  let ret = await dialog.showOpenDialog({properties: ['openDirectory', 'multiSelections']});
  console.log(ret)
  if (ret.canceled == false) {
    let filePaths = ret.filePaths;
    
    for (let i = 0; i < filePaths.length; i++) {
      const filepath = filePaths[i];
      let tmp = getFiles(filepath, filters)

      console.log("  get-directory-file = ", filepath, filters, tmp)


      outfilelist = outfilelist.concat(tmp);
    }
    
  }

  event.returnValue = outfilelist;
})




ipcMain.on('open-file-dialog', async function (event, filters) {
  //console.log(await dialog.showOpenDialog({properties: ['openFile', 'openDirectory', 'multiSelections']}))
  //event.returnValue = await dialog.showOpenDialog({properties: ['openFile', 'openDirectory']});
  let outfilelist = [];
  let ret = await dialog.showOpenDialog({properties: ['openFile', 'multiSelections'],filters:filters});
  if (ret.canceled == false) {
    outfilelist = ret.filePaths;
  }

  event.returnValue = outfilelist;
})


function getFiles(_path, filters) {
  let outfilelist = [];
  let ret_file_list = fs.readdirSync(_path);
  for (let i = 0; i < ret_file_list.length; i++) {
    const file = _path +"\\"+ ret_file_list[i];
    let ext_name = path.extname(file).replace(".","");
    let index_of = filters.indexOf(ext_name);
    if (index_of >= 0 ) {
      outfilelist.push(file)
    }
  }
  return outfilelist;
}
*/


modeDB.Start();
