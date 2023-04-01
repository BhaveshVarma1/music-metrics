import './stats.css';
import {BASE_URL_API, fetchInit, getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import {useEffect, useState} from "react";

export function Stats() {

    const [songCountsLimit, setSongCountsLimit] = useState(100);
    const [averageYear, setAverageYear] = useState('Calculating...');
    const [songCounts, setSongCounts] = useState([{"song": "Loading...", "artist": "Loading...", "count": 0}]);
    const [displayedCounts, setDisplayedCounts] = useState([{"song": "Loading...", "artist": "Loading...", "count": 0}]);

    useEffect(() => {
        fetch(BASE_URL_API + '/api/v1/averageYear/' + localStorage.getItem('username'), fetchInit('/api/v1/averageYear', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setAverageYear(data.averageYear)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/songCounts/' + localStorage.getItem('username'), fetchInit('/api/v1/songCounts', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                fixArtistNames(data.songCounts)
                setSongCounts(data.songCounts)
                setDisplayedCounts(data.songCounts.slice(0, songCountsLimit))
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

    function DropdownMenu() {
        const [isOpen, setIsOpen] = useState(false);

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setSongCountsLimit(size)
            setDisplayedCounts(songCounts.slice(0, size))
        }

        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (isOpen && !event.target.classList.toString().includes('dropdown')) {
                    setIsOpen(false);
                }
            })
        }, [isOpen])

        return (
            <div className={'dd-wrapper'}>
                <div className='dropdown'>
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
                    <div className='dropdown-button' onClick={toggle}>
                        Select table size... {songCountsLimit}
                    </div>
                </div>
            </div>

        );
    }

    return (
        <div>
            <PrimaryInfo text="Stats central."/>
            <SecondaryInfo text={"Average release year: " + averageYear}/>
            <SecondaryInfo text={"Song counts:"}/>
            <CountsTable displayedCounts={displayedCounts}/>
            <DropdownMenu/>
        </div>
    )

}

function CountsTable({ displayedCounts }) {
    return (
        <table className={"table-all"}>
            <thead>
            <tr className={"table-column-names"}>
                <th>Song name</th>
                <th>Artist</th>
                <th style={{textAlign: 'right'}}>Count</th>
            </tr>
            </thead>
            <tbody>
            {displayedCounts.map(songCount => (
                <tr className={"table-row"}>
                    <td>{songCount.song}</td>
                    <td>{songCount.artist}</td>
                    <td style={{textAlign: 'right'}}>{songCount.count}</td>
                </tr>
            ))}
            </tbody>
        </table>
    )
}

function fixArtistNames(songCounts) {
    songCounts.forEach(songCount => {
        songCount.artist = songCount.artist.replaceAll(';;', ', ')
    })
}
