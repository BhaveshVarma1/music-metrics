import './stats.css';
import {BASE_URL_API, fetchInit, getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import {useEffect, useState} from "react";

export function Stats() {

    const [songCountsLimit, setSongCountsLimit] = useState(100);
    const [averageYear, setAverageYear] = useState('Calculating...');
    const [songCounts, setSongCounts] = useState([{"song": "Loading...", "artist": "Loading...", "count": 0}]);

    useEffect(() => {
        fetch(BASE_URL_API + '/averageYear/' + localStorage.getItem('username'), fetchInit('/averageYear', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setAverageYear(data.averageYear)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/songCounts/' + localStorage.getItem('username'), fetchInit('/songCounts', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setSongCounts(data.songCounts.slice(0, songCountsLimit))
            }).catch(error => {
            console.log("ERROR: " + error)
        })
    }, [])

    if (getToken() == null || getToken() === 'undefined') {
        sessionStorage.setItem('route', 'stats')
        return (
            <div>
                <PrimaryInfo text="Log in to continue to stats..."/>
                <LoginButton text="LOGIN TO SPOTIFY"/>
            </div>)
    }
    return (
        <div>
            <PrimaryInfo text="Stats central."/>
            <SecondaryInfo text={"Average release year: " + averageYear}/>
            <SecondaryInfo text={"Song counts:"}/>
            <DropdownMenu/>
            <CountsTable songCounts={songCounts}/>
        </div>
    )

}

function DropdownMenu() {
    const [isOpen, setIsOpen] = useState(false);
    const [value, setValue] = useState(100);

    function toggle() {
        setIsOpen(!isOpen);
    }

    function itemClicked(size) {
        toggle()
        setValue(size)
    }

    return (
        <div className='dropdown'>
            <div className='dropdown-button' onClick={toggle}>
                Select table size... {value}
            </div>
            {isOpen && (
                <div className='dropdown-menu'>
                    <ul>
                        <li onClick={() => itemClicked(25)}>25</li>
                        <li onClick={() => itemClicked(50)}>50</li>
                        <li onClick={() => itemClicked(100)}>100</li>
                        <li onClick={() => itemClicked(250)}>250</li>
                    </ul>
                </div>
            )}
        </div>
    );
}

function CountsTable({ songCounts }) {
    return (
        <table className={"table-all"}>
            <thead>
            <tr className={"table-column-names"}>
                <th>Song name</th>
                <th>Artist</th>
                <th>Count</th>
            </tr>
            </thead>
            <tbody>
            {songCounts.map(songCount => (
                <tr className={"table-row"}>
                    <td>{songCount.song}</td>
                    <td>{songCount.artist}</td>
                    <td>{songCount.count}</td>
                </tr>
            ))}
            </tbody>
        </table>
    )
}
