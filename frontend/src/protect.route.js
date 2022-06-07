import auth from './auth'
import {Navigate} from 'react-router-dom'

export const ProtectRoute = ({Component}) => {
    console.log(auth.isAuthenticated());
    return auth.isAuthenticated() ? <Component/> : <Navigate to="/"/>
}