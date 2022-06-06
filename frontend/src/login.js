import React from 'react'
import { Avatar, Button, Grid, Link, Paper, TextField, Typography } from '@mui/material'
import PersonIcon from '@mui/icons-material/Person';
import { FormControlLabel, Checkbox } from '@mui/material';

const Login = () =>{
    return (
        <Grid className='py-20'>
        <Paper elevation={10} className='p-20 w-72 mx-auto my-20'>
            <Grid align='center'>
                <Avatar className='bg-emerald-500'><PersonIcon/></Avatar>
                <h2>Sign in</h2>
            </Grid>
            <TextField required label='Username' placeholder='Enter username' fullWidth className='mt-10'/>
            <TextField required label='Password' placeholder='Enter password' type='password' fullWidth className='mt-10'/>
            <FormControlLabel control={<Checkbox defaultChecked />} label={<Typography className='text-sm align-bottom'>Remember me</Typography>}/>
            <Button type='submit' fullWidth className='bg-blue-50'>Sign in</Button>
            <Typography className='pt-5'>
                <Link href='#' className='no-underline text-sm '>
                    Forgot password
                </Link>
            </Typography>
            <Typography align='' className='text-sm pt-1'>Do you have an account? 
                <Link href='#' className='text-sm no-underline'>
                    &nbsp;Sign up
                </Link>
            </Typography>
        </Paper>
        </Grid>
    )
}

export default Login