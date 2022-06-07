import MenuAppBar from "./component/navbar";
import Login from "./login";
import Signup from "./signup"
import Present from "./present";
import auth from "./auth";
import { BrowserRouter, Route, Routes, Switch } from "react-router-dom";
import { ProtectRoute } from "./protect.route";

function App() {
  return (
    <div className="w-full h-full">
      <MenuAppBar auth={auth}/>
      <Routes>
      <Route path='/signup' element={<Signup/>}/>
      <Route exact path='/' element={<Login/>}/>
      <Route path='/present' element={<ProtectRoute Component={Present}/>}/>
      <Route path="*" component={() => "404 NOT FOUND"} />
      </Routes>
    </div>
  );
}

export default App;
