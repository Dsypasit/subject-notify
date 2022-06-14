import Cookies from 'js-cookie'
import axios from 'axios';

class Auth {
  constructor() {
    this.authenticated = false;
    this.t1 = ""
    this.t2 = ""
    this.checkCookies()
  }

  checkCookies(){
    this.t1 = Cookies.get('j')
    this.t2 = Cookies.get('a')
    if (this.t1 == undefined || this.t2 == undefined){
      this.authenticated = false
      return
    }
    this.authenticated = true
  }

  setToken(){
    this.t1 = Cookies.get('j')
    this.t2 = Cookies.get('a')
  }

  getToken(){
    const token = this.t1+"."+this.t2
    return token
  }

  async signup(name, password) {
    let res;
    let text
    let data = {
          username:name,
          password:password
      }
    try{
      res = await axios({
        url:'/OpenAccount',
        baseURL:'http://127.0.0.1:5000',
        method:'post',
        data:data,
        headers: {'Content-Type': 'application/json'}
      })
    }catch(err){
      return "server error"
    }
    return undefined
  }

  async login(username, password) {
    let res;
    let text;
    try{
      res = await fetch('http://localhost:5000/Login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                username,
                password
            })
        });
      let data = await res.text()
      console.log(data+"login1");
    }catch(err){
      let data = await res.text()
      console.log(data+"login2");
      return data
    }
    
    return text
  }

  logout() {
    this.authenticated = false;
    fetch('http://localhost:5000/Logout', {
      method: 'GET',
      credentials: 'include'
    })
    .catch((err)=> console.log(err))
  }

  isAuthenticated() {
    return this.authenticated;
  }

  getUser(setInfo){
    let success = false
    fetch('http://localhost:5000/GetAccount', {
      method: 'GET',
      credentials: 'include'
    }).then((res) => {
      return res.json()
    })
    .then((data) => {
      setInfo(data)
      console.log(data);
      success = true
    })
    .catch((err)=> {
      throw err
    })
    return success;
  }
}

export default new Auth();