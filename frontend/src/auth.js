import Cookies from 'js-cookie'

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

  async login(name, password, func) {
    const res = await axios.post('http://localhost:8000/login', {
        name,
        password
    })
    if (res.status <300){
        this.authenticated = true;
        func()
    }else{
        return res.data.message
    }
  }

  logout(cb) {
    this.authenticated = false;
    cb();
  }

  isAuthenticated() {
    return this.authenticated;
  }
}

export default new Auth();