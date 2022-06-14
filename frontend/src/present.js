import BasicList from "./component/listItem"
import TableSubject from "./component/table"
import React, { useEffect, useState } from 'react'
import { useNavigate } from "react-router-dom"
import auth from "./auth"

export default function Present ({info, setInfo, setPass}) {

    return (
        <div>
        <div className="flex">
            <BasicList info={info}/>
            <TableSubject info={info}/>
        </div>
        </div>
    )
}