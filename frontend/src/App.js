import MenuAppBar from "./component/navbar";
import Login from "./login";
import Signup from "./signup"
import Present from "./present";
import auth from "./auth";
import { BrowserRouter, Route, Routes, Switch } from "react-router-dom";
import { ProtectRoute } from "./protect.route";
import { useEffect } from "react";

function App() {

  useEffect(()=>{
    auth.checkCookies()
  }, [auth.authenticated])

  return (
    <div className="w-full h-full">
      <MenuAppBar auth={auth}/>
      <Routes>
      <Route path='/signup' element={<Signup/>}/>
      <Route path='/login' element={<Login/>}/>
      <Route path='/' element={<ProtectRoute Component={Present}/>}/>
      <Route path="*" component={() => "404 NOT FOUND"} />
      </Routes>
    </div>
  );
}

export default App;
