// noinspection JSUnresolvedVariable

import './stats.css';
import {BASE_URL_API, fetchInit, getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import {useEffect, useState} from "react";

export function Stats() {

    const DEFAULT_SONG_COUNT_LIMIT = 100
    const DEFAULT_ALBUM_COUNT_LIMIT = 100

    const selectedStyle = 'selector-selected'
    const unselectedStyle = 'selector-unselected'
    const [songStyle, setSongStyle] = useState(selectedStyle);
    const [albumStyle, setAlbumStyle] = useState(unselectedStyle);

    const [songCountsLimit, setSongCountsLimit] = useState(DEFAULT_SONG_COUNT_LIMIT);
    const [albumCountsLimit, setAlbumCountsLimit] = useState(DEFAULT_ALBUM_COUNT_LIMIT);
    const [averageYear, setAverageYear] = useState('Calculating...');
    const [songCounts, setSongCounts] = useState([{"song": "Loading...", "artist": "Loading...", "count": 0}]);
    const [displayedCounts, setDisplayedCounts] = useState([{"song": "Loading...", "artist": "Loading...", "count": 0}]);
    const [topAlbums, setTopAlbums] = useState([{"album": "Loading...", "artist": "Loading...", "count": 0}]);
    const [displayedAlbums, setDisplayedAlbums] = useState([{"album": "Loading...", "artist": "Loading...", "count": 0}]);

    const songCountsTable = <CountsTable displayedCounts={displayedCounts}/>
    const topAlbumsTable = <AlbumsTable displayedAlbums={displayedAlbums}/>
    const [displayedTable, setDisplayedTable] = useState(songCountsTable);

    const countsDropdown = <CountsDropdown/>
    const albumsDropdown = <AlbumsDropdown/>
    const [currentDropdown, setCurrentDropdown] = useState(<CountsDropdown/>)

    // Call MusicMetrics APIs
    useEffect(() => {
        if (getToken() == null || getToken() === 'undefined') {
            return
        }
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
                setDisplayedCounts(data.songCounts.slice(0, DEFAULT_SONG_COUNT_LIMIT))
                // This line is needed because React's state update is asynchronous
                if (songStyle === selectedStyle) {
                    setDisplayedTable(<CountsTable displayedCounts={data.songCounts.slice(0, DEFAULT_SONG_COUNT_LIMIT)}/>)
                }
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/topAlbums/' + localStorage.getItem('username'), fetchInit('/api/v1/topAlbums', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                fixArtistNames(data.topAlbums)
                setTopAlbums(data.topAlbums)
                setDisplayedAlbums(data.topAlbums.slice(0, DEFAULT_ALBUM_COUNT_LIMIT))
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

    function setToSong() {
        setSongStyle(selectedStyle)
        setAlbumStyle(unselectedStyle)
        setDisplayedTable(songCountsTable)
        setCurrentDropdown(countsDropdown)
    }

    function setToAlbum() {
        setSongStyle(unselectedStyle)
        setAlbumStyle(selectedStyle)
        setDisplayedTable(topAlbumsTable)
        setCurrentDropdown(albumsDropdown)
    }

    function tempClick() {
        //setDisplayedTable(songCountsTable)
    }

    function CountsDropdown() {
        const [isOpen, setIsOpen] = useState(false);

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setSongCountsLimit(size)
            //setDisplayedCounts(songCounts.slice(0, size))
            console.log(songCounts.slice(0, size))
            setDisplayedTable(<CountsTable displayedCounts={songCounts.slice(0, size)}/>)
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

    function AlbumsDropdown() {
        const [isOpen, setIsOpen] = useState(false);

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setAlbumCountsLimit(size)
            setDisplayedAlbums(topAlbums.slice(0, size))
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
                                <li onClick={() => itemClicked(10)}>10</li>
                                <li onClick={() => itemClicked(25)}>25</li>
                                <li onClick={() => itemClicked(50)}>50</li>
                                <li onClick={() => itemClicked(100)}>100</li>
                            </ul>
                        </div>
                    )}
                    <div className='dropdown-button' onClick={toggle}>
                        Select table size... {albumCountsLimit}
                    </div>
                </div>
            </div>

        );
    }

    return (
        <div>
            <div onClick={tempClick}><PrimaryInfo text="Stats central."/></div>
            <SecondaryInfo text={"Average release year: " + averageYear}/>
            <div className={'selector'}>
                <div className={songStyle + ' selector-option corner-rounded-left'} onClick={setToSong}>Top Songs</div>
                <div className={albumStyle + ' selector-option corner-rounded-right'} onClick={setToAlbum}>Top Albums</div>
            </div>
            {displayedTable}
            {currentDropdown}
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

function AlbumsTable({ displayedAlbums }) {
    return (
        <table className={"table-all"}>
            <thead>
            <tr className={"table-column-names"}>
                <th>Album name</th>
                <th>Artist</th>
                <th style={{textAlign: 'right'}}>Count</th>
            </tr>
            </thead>
            <tbody>
            {displayedAlbums.map(albumCount => (
                <tr className={"table-row"}>
                    <td>{albumCount.album}</td>
                    <td>{albumCount.artist}</td>
                    <td style={{textAlign: 'right'}}>{albumCount.count}</td>
                </tr>
            ))}
            </tbody>
        </table>
    )
}

function fixArtistNames(items) {
    items.forEach(item => {
        item.artist = item.artist.replaceAll(';;', ', ')
    })
}
