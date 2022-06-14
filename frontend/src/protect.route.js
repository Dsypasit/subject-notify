import auth from './auth'
import {Navigate} from 'react-router-dom'

export const ProtectRoute = ({info, Component}) => {
    console.log("kuy protect");
    return info?.Username ? <Component info={info}/> : <Navigate to="/login"/>
}