import React, {useEffect, useState} from "react";
import {makeStyles} from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import axios from "axios";

const useStyle = makeStyles({
    table: {
        minWidth: 650
    }
})



export default function BasicTable() {
    const classes = useStyle()
    const [rows, setRows] = useState([]);
    useEffect(() => {
        getAllTodo();
    },[])
    const url = "http://localhost:8080/todo"
    const getAllTodo = async () => {
        try {
            const response = await axios.get(url);
            const arrayPost = response.data.tasks;
            setRows(arrayPost);
            console.log(arrayPost);
        } catch (error){
            console.log(error);
        }
    }
    return (
        <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell align="left">id</TableCell>
                        <TableCell align="left">title</TableCell>
                        <TableCell align="left">text</TableCell>
                        <TableCell align="left">created at</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {rows.map((row) => (
                        <TableRow key={row.id}>
                            <TableCell align="left">{row.id}</TableCell>
                            <TableCell align="left">{row.title}</TableCell>
                            <TableCell align="left">{row.text}</TableCell>
                            <TableCell align="left">{row.created_at}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
}