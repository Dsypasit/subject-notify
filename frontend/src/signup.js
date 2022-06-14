import React, { useEffect, useState } from 'react'
import { Avatar, Button, Grid, Link, Paper, TextField, Typography } from '@mui/material'
import PersonIcon from '@mui/icons-material/Person';
import { FormControlLabel, Checkbox } from '@mui/material';
import { Navigate, useNavigate } from 'react-router-dom';
import auth from './auth';

const Signup = () =>{
    const nevigate = useNavigate();
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPass, setConfirmPass] = useState("")
    const [errMessage, setErrMessage] = useState("")

    const [errPass, setErrPass] = useState(false)
    const [errName, setErrName] = useState(false)

    const submit = async(e) =>{
        e.preventDefault()
        if (name === "" || password ===""){
            setErrName(true)
            setErrPass(true)
            setErrMessage("Please Enter username and password")
            return
        }
        let err = await auth.signup(name, password)
        if (err != undefined){
            setErrName(true)
            setErrMessage(err)
            return
        }
        nevigate("/login")
    }
    useEffect(() => {
        if (confirmPass !== password){
            setErrPass(true)
        }else{
            setErrPass(false)
        }

        if (name.length > 0 || password.length>0){
            setErrName(false)
            setErrMessage("")
        }
    }, [confirmPass, name])
    return (
        <Grid className='py-20 h-100v'>
        <Paper elevation={10} className='p-20 w-72 mx-auto my-10'>
            <Grid align='center'>
                <Avatar className='bg-emerald-500'><PersonIcon/></Avatar>
                <h2>Sign up</h2>
            </Grid>
            <TextField error={errName} required value={name} onChange={(e)=> setName(e.target.value)} label='Username' placeholder='Enter username' fullWidth className='mt-5'/>
            <TextField error={errPass} required value={password} onChange={(e)=> setPassword(e.target.value)} label='Password' placeholder='Enter password' type='password' fullWidth className='mt-5'/>
            <TextField error={errPass} required value={confirmPass} onChange={(e)=> setConfirmPass(e.target.value)} label='Confirm Password' placeholder='Enter password' type='password' fullWidth className='mt-5'/>
            {errMessage ? <Typography display="block"  variant="caption" className='text-red-500 font-bold'>{errMessage}</Typography> : null}
            <FormControlLabel control={<Checkbox defaultChecked />} label={<Typography className='text-sm align-bottom'>Remember me</Typography>}/>
            <Button onClick={submit} type='submit' fullWidth className='bg-blue-50'>Sign up</Button>
        </Paper>
        </Grid>
    )
}

export default Signup