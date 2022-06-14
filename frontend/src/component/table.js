import * as React from 'react';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Command from './command';

const createData = (day, morningSubject, afternoonSubject) =>{
    return {day, morningSubject, afternoonSubject}
}

const rows = [
    createData('Mon', 'Math', 'Eng'),
    createData('Mon', 'Math', 'Eng'),
    createData('Mon', 'Math', 'Eng'),
    createData('Mon', 'Math', 'Eng'),
    createData('Mon', 'Math', 'Eng')
]

export default function TableSubject(){
    return (
        <div className="grow">
            <div className='p-20 max-w-screen-lg mx-auto'>
        <TableContainer className='border-none' component={Paper}>
      <Table  aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Dessert (100g serving)</TableCell>
            <TableCell >Calories</TableCell>
            <TableCell >Fat&nbsp;(g)</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow
              key={row.day}
              sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                {row.day}
              </TableCell>
              <TableCell>{row.morningSubject}</TableCell>
              <TableCell>{row.afternoonSubject}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
            </div>
            <Command/>
        </div>
    )
}