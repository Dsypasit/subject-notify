import BasicList from "./component/listItem"
import TableSubject from "./component/table"
export default function Present () {
    return (
        <div className="flex">
            <BasicList/>
            <TableSubject/>
        </div>
    )
}