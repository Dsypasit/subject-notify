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

  async signup(name, password, navigate) {
    let res;
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
      console.log(err);
      return res?.data ? res.data : "server error"
    }
    console.log(res.status);
    if (res.status === 200){
        navigate("/")
    }else if(res.status == 417){
        return "server error"
    }else{
      return res.data
    }
  }

  async login(username, password, navigate) {
    let res;
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
    }catch(err){
      return res?.data ? res.data : "server error"
    }
    if (res.status <300){
        this.authenticated = true;
        navigate("/")
    }else if(res.status == 417){
        return "server error"
    }else{
      return res.data
    }
  }

  logout(nevigate) {
    this.authenticated = false;
    fetch('http://localhost:5000/Logout', {
      method: 'GET',
      credentials: 'include'
    }).then(() => nevigate('/login'))
    .catch((err)=> console.log(err))
  }

  isAuthenticated() {
    return this.authenticated;
  }
}

export default new Auth();