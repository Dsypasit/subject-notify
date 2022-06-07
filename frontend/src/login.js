import React, { useState } from 'react'
import { Avatar, Button, Grid, Link, Paper, TextField, Typography } from '@mui/material'
import PersonIcon from '@mui/icons-material/Person';
import { FormControlLabel, Checkbox } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import auth from './auth';

const Login = () =>{
    let err = ""
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    const nevigate = useNavigate()
    const submit = async () =>{
        console.log('before');
        err = await auth.login(name, password, nevigate)
    }
    return (
        <Grid className='py-20 h-100v'>
        <Paper elevation={10} className='p-20 w-72 mx-auto my-10'>
            <Grid align='center'>
                <Avatar className='bg-emerald-500'><PersonIcon/></Avatar>
                <h2>Sign in</h2>
            </Grid>
            {err ? <Typography variant="caption" className='text-red-500 font-bold'>test ku hello hahaha watafak</Typography> : null}
            <TextField error={err} value={name} onChange={(e)=>setName(e.target.value)} required label='Username' placeholder='Enter username' fullWidth className='mt-5'/>
            <TextField error={err} value={password} onChange={e => setPassword(e.target.value)} required label='Password' placeholder='Enter password' type='password' fullWidth className='mt-5'/>
            <FormControlLabel control={<Checkbox defaultChecked />} label={<Typography className='text-sm align-bottom'>Remember me</Typography>}/>
            <Button onClick={submit} type='submit' fullWidth className='bg-blue-50'>Sign in</Button>
            <Typography className='pt-5'>
                <Link href='#' className='no-underline text-sm '>
                    Forgot password
                </Link>
            </Typography>
            <Typography align='' className='text-sm pt-1'>Do you have an account? 
                <Link href='/signup' className='text-sm no-underline'>
                    &nbsp;Sign up
                </Link>
            </Typography>
        </Paper>
        </Grid>
    )
}

export default Login