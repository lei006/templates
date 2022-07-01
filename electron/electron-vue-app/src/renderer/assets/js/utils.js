





function getFilePath(fileNamePath) {
    let tmp = fileNamePath.substr(0, fileNamePath.lastIndexOf('\\')+1);
    return tmp;
  }
  
  
  function getFileName(fileNamePath) {
    let tmp = fileNamePath.substr(fileNamePath.lastIndexOf('\\') + 1);
    return tmp;
  }
  
  
  function getOutFileName(fileNamePath, outPath, outFormat) {
    //let tmp = fileNamePath.substr(fileNamePath.lastIndexOf('\\') + 1);
  
    let filename = getFileName(fileNamePath);
  
    let tmp = filename.substr(0, filename.lastIndexOf('.'));
  
    if (outFormat == "png") {
        tmp = outPath + tmp + "_" + outFormat + "\\";
    }else if (outFormat == "jpg") {
        tmp = outPath + tmp + "_" + outFormat + "\\";
    }else if (outFormat == "bmp") {
        tmp = outPath + tmp + "_" + outFormat + "\\";
    }else{
        tmp = outPath + tmp + "." + outFormat;
    }
  
  
    return tmp;
  }
  
  
  export default {getFilePath,getFileName,getOutFileName}

  
  