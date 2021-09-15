
const { app, BrowserWindow, dialog, globalShortcut, ipcMain, session } = require('electron')

const {read_config} = require('./readconfig.js');
var tcpPortUsed = require('tcp-port-used');


let g_config = undefined;

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

let win_main = null;

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
      nodeIntegration: true
    },
    titleBarStyle:"hidden",
    hasShadow:true,
    fullscreen: (g_config.runmode !== "debug"),
    autoHideMenuBar: true,
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

ipcMain.on('app-exit', (event, arg) => {
  app.exit(0)
})

read_config("config.json", function(config) {

  g_config = config;

  // user pro debug
  g_config.runmode = g_config.runmode || "user";
  g_config.port = g_config.port || 28083;

  app.whenReady().then(createWindow)
  /*
  tcpPortUsed.waitUntilFree(g_config.port, 500, 6000000).then(function() {
  }, function(err) {
      console.log('wait port free Error:', err.message);
  });
  */







})





