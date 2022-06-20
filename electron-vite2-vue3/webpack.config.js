
const path = require('path')

var UglifyJsPlugin = require('uglifyjs-webpack-plugin');


module.exports = {
    
    entry: './src/main/index.js',
    output: {
      filename: 'main.js'
    },
    mode:'production', //  development production
    plugins: [
      new UglifyJsPlugin()
    ]    
  };
  

  