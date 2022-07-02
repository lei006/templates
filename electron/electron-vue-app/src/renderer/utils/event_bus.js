

function EventEmitter(){
    this._events = Object.create(null);
  }
  
  EventEmitter.prototype.on = function(eventName, callback){
    if(!this._events){
      this._events = Object.create(null);
    }
    if(eventName !== 'newListener'){
      this.emit('newListener', callback);
    }
    if(this._events[eventName]){
      this._events[eventName].push(callback);
    }else{
      this._events[eventName] = [callback];
    }
  }
  
  EventEmitter.prototype.once = function(eventName, callback){
    let one = ()=>{
      callback();
      this.off(eventName, one);
    }
    one.l = callback;
    this.on(eventName, one);
  }
  
  EventEmitter.prototype.off = function(eventName,callback){
    if(this._events[eventName]){
      this._events[eventName].filter(fn=>{
        return fn != callback && fn.l != callback;
      })
    }
  }
  
  EventEmitter.prototype.emit = function(eventName, ...args){
    if(this._events[eventName]){
      this._events[eventName].forEach(fn => fn(...args));
    }
  }


let install = function (Vue) {
	// 设置eventBus
	Vue.prototype.$bus = new EventEmitter();
}

export default { install };



