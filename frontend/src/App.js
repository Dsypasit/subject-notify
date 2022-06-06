import MenuAppBar from "./component/navbar";
import Login from "./login";
import Signup from "./signup"
import Present from "./present";
import { useState } from "react";
import { BrowserRouter, Route, Routes, Switch } from "react-router-dom";

function App() {
  const [auth, setAuth] = useState(false)
  return (
    <div className="w-full h-full">
      <MenuAppBar auth={auth}/>
      <Routes>
      <Route path='/' element={auth ?<Present/> :<Signup setAuth={setAuth}/>}/>
      <Route path='/login' element={<Login setAuth={setAuth}/>}/>
      </Routes>
    </div>
  );
}

export default App;
