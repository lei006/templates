
const { app, BrowserWindow, dialog, globalShortcut, ipcMain, session } = require('electron')


const fs = require('fs')
const path = require('path')

import web_app from './express/index'



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
  //: `https://www.zhihu.com/`
  : `file://${__dirname}/index.html`

//app.commandLine.appendSwitch('ignore-certificate-errors')




let g_debug = false


let win_main = null;

let g_userinfo = {};



function createWindow () {  

  // 创建浏览器窗口
  win_main = new BrowserWindow({
    width: 1024,
    height: 720,
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
    console.log("show");
    win_main.show()
  })


  //main.js
  win_main.on('close', (event) => {
    
  });
  //win_main.loadFile('./src/views/main.html')
  win_main.loadURL(winURL)
  console.log("load");
  win_main.webContents.openDevTools()
  const options = {
    type: 'question',
    buttons: ['Cancel', 'Yes, please', 'No, thanks'],
    defaultId: 2,
    cancelId: 0,
    title: 'Question',
    message: 'my window?',
    detail: 'It does not really matter',
    checkboxLabel: 'remember',
    checkboxChecked: true,
  }; 
  
  //dialog.showMessageBoxSync(win_main, options);


}

app.whenReady().then(function(){

  web_app.Start(1234);
  createWindow();

})



ipcMain.on('app-mini', (event, arg) => {
  win_main.minimize();
})


ipcMain.on('app-exit', (event, arg) => {
  app.exit(0)
})


