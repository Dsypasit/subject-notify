import React from 'react'
import { Avatar, Button, Grid, Link, Paper, TextField, Typography } from '@mui/material'
import PersonIcon from '@mui/icons-material/Person';
import { FormControlLabel, Checkbox } from '@mui/material';
import { useNavigate } from 'react-router-dom';

const Signup = ({setAuth}) =>{
    const nevigate = useNavigate();
    const submit = () =>{
        setAuth(true)
        nevigate('/login')
    }
    return (
        <Grid className='py-20 h-100v'>
        <Paper elevation={10} className='p-20 w-72 mx-auto my-10'>
            <Grid align='center'>
                <Avatar className='bg-emerald-500'><PersonIcon/></Avatar>
                <h2>Sign up</h2>
            </Grid>
            <TextField required label='Username' placeholder='Enter username' fullWidth className='mt-10'/>
            <TextField required label='Password' placeholder='Enter password' type='password' fullWidth className='mt-10'/>
            <TextField required label='Confirm Password' placeholder='Enter password' type='password' fullWidth className='mt-10'/>
            <FormControlLabel control={<Checkbox defaultChecked />} label={<Typography className='text-sm align-bottom'>Remember me</Typography>}/>
            <Button onClick={submit} type='submit' fullWidth className='bg-blue-50'>Sign up</Button>
        </Paper>
        </Grid>
    )
}

export default Signup