
const { app, BrowserWindow, dialog, globalShortcut, ipcMain, session } = require('electron')


const fs = require('fs')
const path = require('path')


import {WebServer, DbServer, koa, Window} from './module'



process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';


/**
 * Set `__static` path to static files in production
 * https://simulatedgreg.gitbooks.io/electron-vue/content/en/using-static-assets.html
 */
if (process.env.NODE_ENV !== 'development') {
  global.__static = require('path').join(__dirname, '/static').replace(/\\/g, '\\\\')
}

app.whenReady().then(function(){
  console.log("app ready enter");
  
  WebServer.Start();
  DbServer.Start();
  Window.Start();

  console.log("app ready leave");
})

ipcMain.on('app-exit', (event, arg) => {
  app.exit(0)
})

