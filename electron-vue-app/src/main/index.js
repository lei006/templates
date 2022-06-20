
const { app, BrowserWindow, dialog, globalShortcut, ipcMain, session } = require('electron')




const fs = require('fs')
const path = require('path')


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
  : `www.baidu.com`
  //: `file://${__dirname}/index.html`

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

  win_main.webContents.openDevTools()


}

app.whenReady().then(createWindow)



ipcMain.on('app-mini', (event, arg) => {
  win_main.minimize();
})


ipcMain.on('app-exit', (event, arg) => {
  app.exit(0)
})


