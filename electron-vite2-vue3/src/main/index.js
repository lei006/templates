
const {app, BrowserWindow, Notification, Menu, MenuItem, Tray, dialog } = require('electron')

const path = require('path')

let mainWin = undefined;

const winURL = process.env.NODE_ENV === 'devlopment'
  ? `http://localhost:3000`
  : `file://${__dirname}/../../dist/index.html`



function createWindow(){

    const win = new BrowserWindow({
        width: 800,
        height: 600,
        webPreferences: {
          contextIsolation:false,
          nodeIntegration:true,
          //preload: path.join(__dirname, 'preload.js')
        }
      })
    
      win.loadURL(winURL);

      const menu = new Menu();
      menu.append(new MenuItem({label:'复制', role:'copy'}));
      win.webContents.on('context-menu', (e,params)=>{
        menu.popup({window:win, x:params.x, y:params.y});
      });

      //mainWin.webContents.openDevTools();

    if(process.env.NODE_ENV === 'devlopment'){
        //mainWin.webContents.openDevTools();
    }


      mainWin = win;

}



function tray(){
    let tray = new Tray('./tray.png');
    const contextMenu = Menu.buildFromTemplate([
        {
            role:'minimize',
            label:"最小化",
            click:()=>{
                mainWin.minimize();
            }
        },{
            label:"新窗口1",
            click:()=>{
                const {shell} = require('electron')
                shell.openExternal('https://electronjs.org')
            }
        },{
            role:'togglefullscreen',
            label:"全屏",
            click:()=>{
                mainWin.setFullScreen(mainWin.isFullScreen()!==true);
            }
        },{
            role:'toggleshowdialog',
            label:"对话框",
            click:()=>{
                dialog.showOpenDialog(mainWin, { properties: ['openDirectory'] }).then(result=>{
                    console.log(result.canceled)
                    console.log(result.filePaths)                    
                })
            }
        },{
            label:'退出',
            //role:'quit',
            click:()=>{

                dialog.showMessageBox(mainWin, {title:"警告", message:"是否退出?", cancelId:1, buttons:['退出','取消']}).then(function(result){
                    if(result.response === 0) {
                        console.log("退出");

                        app.quit();
                    }else{
                        console.log("退出被取消");
                        ShowNotification('退出被取消');
                    }
                })

            }
        }
    ])

    tray.setToolTip('我是一个托盘程序..');
    tray.on('right-click', ()=>{
        tray.popUpContextMenu(contextMenu);
    })
    tray.on('click', (e,bounds)=>{
        mainWin.show();
    })


}





function ShowNotification(message, title){
    title = title|| "";
    new  Notification({title,body:message}).show();
}

app.whenReady().then(() => {
    createWindow()
    tray();
    
    app.on('activate', () => {
      if (BrowserWindow.getAllWindows().length === 0) createWindow()
    })
})


app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
      app.quit()
    }
})
