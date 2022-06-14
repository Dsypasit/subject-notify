import MenuAppBar from "./component/navbar";
import Login from "./login";
import Signup from "./signup"
import Present from "./present";
import auth from "./auth";
import { BrowserRouter, Navigate, Route, Routes, Switch, useNavigate } from "react-router-dom";
import { ProtectRoute } from "./protect.route";
import { useEffect, useState } from "react";

function App() {
  const [info, setInfo] = useState(undefined)
  const [pass, setPass] = useState(false)
  let nevigate = useNavigate()

  useEffect( () => {
    async function getUser () {
      try{
        let done = await auth.getUser(setInfo)
        if (done){
          nevigate("/")
          setPass(true)
        }
      }catch(e){
        console.log(e);
      }
    }
    getUser();
  }, [])

  return (
    <div className="w-full h-full">
      <MenuAppBar setPass={setPass} info={info} setInfo={setInfo}/>
      <Routes>
      <Route path='/signup' element={<Signup/>}/>
      <Route path='/login' element={<Login info={info} setPass={setPass} setInfo={setInfo}/>}/>
      <Route exact path='/' element={info ? <Present setInfo={setInfo} setPass={setPass} info={info}/> : <Login info={info} setPass={setPass} setInfo={setInfo}/>}/>
      <Route path="*" component={() => "404 NOT FOUND"} />
      </Routes>
    </div>
  );
}

export default App;
