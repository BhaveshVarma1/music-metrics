// noinspection JSUnresolvedVariable

import './stats.css';
import {BASE_URL_API, fetchInit, getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import {useEffect, useState} from "react";
import {Chart} from "react-google-charts";

const DEFAULT_SONG_COUNT_LIMIT = 100
const DEFAULT_ARTIST_COUNT_LIMIT = 50
const DEFAULT_ALBUM_COUNT_LIMIT = 50

export function Stats() {

    const selectedStyle = 'selector-selected'
    const unselectedStyle = 'selector-unselected'
    const [songStyle, setSongStyle] = useState(selectedStyle);
    const [artistStyle, setArtistStyle] = useState(unselectedStyle);
    const [albumStyle, setAlbumStyle] = useState(unselectedStyle);

    const [averageYear, setAverageYear] = useState('Calculating...');

    // TO PASS IN AS PROPS:
    // initialState, url, fixArtistNames, defaultCount, ddValues, thead, itemCallback
    const songTableProps = {
        initialState: [{"song": "Loading...", "artist": "Loading...", "count": 0}],
        url: '/api/v1/topSongs',
        fixArtistNames: true,
        defaultCount: DEFAULT_SONG_COUNT_LIMIT,
        ddValues: [25, 50, 100, 250],
        thead: () => {
            return (
                <thead>
                <tr className={"table-column-names"}>
                    <th>Rank</th>
                    <th>Song name</th>
                    <th>Artist</th>
                    <th style={{textAlign: 'right'}}>Count</th>
                </tr>
                </thead>
            )
        },
        itemCallback: (item) => {
            return (
                <tr className={"table-row"}>
                    <td>{item.rank}</td>
                    <td>{item.song}</td>
                    <td>{item.artist}</td>
                    <td style={{textAlign: 'right'}}>{item.count}</td>
                </tr>
            )
        }
    }
    const artistTableProps = {
        initialState: [{"artist": "Loading...", "count": 0}],
        url: '/api/v1/topArtists',
        fixArtistNames: false,
        defaultCount: DEFAULT_ARTIST_COUNT_LIMIT,
        ddValues: [10, 25, 50, 100],
        thead: () => {
            return (
                <thead>
                <tr className={"table-column-names"}>
                    <th>Rank</th>
                    <th>Artist name</th>
                    <th style={{textAlign: 'right'}}>Count</th>
                </tr>
                </thead>
            )
        },
        itemCallback: (item) => {
            return (
                <tr className={"table-row"}>
                    <td>{item.rank}</td>
                    <td>{item.artist}</td>
                    <td style={{textAlign: 'right'}}>{item.count}</td>
                </tr>
            )
        }
    }
    const albumTableProps = {
        initialState: [{"album": "Loading...", "artist": "Loading...", "count": 0}],
        url: '/api/v1/topAlbums',
        fixArtistNames: true,
        defaultCount: DEFAULT_ALBUM_COUNT_LIMIT,
        ddValues: [10, 25, 50, 100],
        thead: () => {
            return (
                <thead>
                <tr className={"table-column-names"}>
                    <th>Rank</th>
                    <th>Album name</th>
                    <th>Artist</th>
                    <th style={{textAlign: 'right'}}>Count</th>
                </tr>
                </thead>
            )
        },
        itemCallback: (item) => {
            return (
                <tr className={"table-row"}>
                    <td>{item.rank}</td>
                    <td>{item.album}</td>
                    <td>{item.artist}</td>
                    <td style={{textAlign: 'right'}}>{item.count}</td>
                </tr>
            )
        }
    }
    const [displayedTable, setDisplayedTable] = useState(<TopTable props={songTableProps}/>);

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
    }, [])

    if (getToken() == null || getToken() === 'undefined') {
        sessionStorage.setItem('route', 'stats')
        return (
            <div>
                <PrimaryInfo text="Log in to continue to stats..."/>
                <LoginButton text="LOGIN TO SPOTIFY"/>
            </div>
        )
    }

    function setToSong() {
        setSongStyle(selectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(unselectedStyle)

        setDisplayedTable(<TopTable props={songTableProps}/>)
    }

    function setToArtist() {
        setSongStyle(unselectedStyle)
        setArtistStyle(selectedStyle)
        setAlbumStyle(unselectedStyle)

        setDisplayedTable(<TopTable props={artistTableProps}/>)
    }

    function setToAlbum() {
        setSongStyle(unselectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(selectedStyle)

        setDisplayedTable(<TopTable props={albumTableProps}/>)
    }

    return (
        <div>
            <PrimaryInfo text="Stats central."/>
            <SecondaryInfo text={"Average release year: " + averageYear}/>
            <div className={'selector'}>
                <div className={songStyle + ' selector-option corner-rounded-left'} onClick={setToSong}>Top Songs</div>
                <div className={artistStyle + ' selector-option'} onClick={setToArtist}>Top Artists</div>
                <div className={albumStyle + ' selector-option corner-rounded-right'} onClick={setToAlbum}>Top Albums</div>
            </div>
            {displayedTable}
            <DecadePieChart/>
        </div>
    )

}

/*function SongsTable() {
    const [allSongs, setAllSongs] = useState([{"song": "Loading...", "artist": "Loading...", "count": 0}])
    const [displayedSongs, setDisplayedSongs] = useState([])

    useEffect(() => {
        fetch(BASE_URL_API + '/api/v1/topSongs/' + localStorage.getItem('username'), fetchInit('/api/v1/topSongs', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                fixArtistNames(data.topSongs)
                addRankColumn(data.topSongs)
                setAllSongs(data.topSongs)
                setDisplayedSongs(data.topSongs.slice(0, DEFAULT_SONG_COUNT_LIMIT))
            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [])

    function SongsDropdown() {
        const [isOpen, setIsOpen] = useState(false);
        const [dropdownValue, setDropdownValue] = useState(DEFAULT_SONG_COUNT_LIMIT)

        // Close dropdown when clicking outside of it
        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (isOpen && !event.target.classList.toString().includes('dropdown')) {
                    setIsOpen(false);
                }
            })
        }, [isOpen])

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setDropdownValue(size)
            setDisplayedSongs(allSongs.slice(0, size))
        }

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
                        Select table size... {dropdownValue}
                    </div>
                </div>
            </div>

        );
    }

    return (
        <div>
            <table className={"table-all"}>
                <thead>
                <tr className={"table-column-names"}>
                    <th>Rank</th>
                    <th>Song name</th>
                    <th>Artist</th>
                    <th style={{textAlign: 'right'}}>Count</th>
                </tr>
                </thead>
                <tbody>
                {displayedSongs.map(songCount => (
                    <tr className={"table-row"}>
                        <td>{songCount.rank}</td>
                        <td>{songCount.song}</td>
                        <td>{songCount.artist}</td>
                        <td style={{textAlign: 'right'}}>{songCount.count}</td>
                    </tr>
                ))}
                </tbody>
            </table>
            <SongsDropdown/>
        </div>
    )
}

function ArtistsTable() {
    const [allArtists, setAllArtists] = useState([{"artist": "Loading...", "count": 0}])
    const [displayedArtists, setDisplayedArtists] = useState([])

    useEffect(() => {
        fetch(BASE_URL_API + '/api/v1/topArtists/' + localStorage.getItem('username'), fetchInit('/api/v1/topArtists', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                addRankColumn(data.topArtists)
                setAllArtists(data.topArtists)
                setDisplayedArtists(data.topArtists.slice(0, DEFAULT_ARTIST_COUNT_LIMIT))
            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [])

    function ArtistsDropdown() {
        const [isOpen, setIsOpen] = useState(false);
        const [dropdownValue, setDropdownValue] = useState(DEFAULT_ARTIST_COUNT_LIMIT)

        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (isOpen && !event.target.classList.toString().includes('dropdown')) {
                    setIsOpen(false);
                }
            })
        }, [isOpen])

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setDropdownValue(size)
            setDisplayedArtists(allArtists.slice(0, size))
        }

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
                        Select table size... {dropdownValue}
                    </div>
                </div>
            </div>

        );
    }

    return (
        <div>
            <table className={"table-all table-all-artist"}>
                <thead>
                <tr className={"table-column-names"}>
                    <th>Rank</th>
                    <th>Artist name</th>
                    <th style={{textAlign: 'right'}}>Count</th>
                </tr>
                </thead>
                <tbody>
                {displayedArtists.map(artistCount => (
                    <tr className={"table-row"}>
                        <td>{artistCount.rank}</td>
                        <td>{artistCount.artist}</td>
                        <td style={{textAlign: 'right'}}>{artistCount.count}</td>
                    </tr>
                ))}
                </tbody>
            </table>
            <ArtistsDropdown/>
        </div>
    )
}

function AlbumsTable() {
    const [allAlbums, setAllAlbums] = useState([{"album": "Loading...", "artist": "Loading...", "count": 0}])
    const [displayedAlbums, setDisplayedAlbums] = useState([])

    useEffect(() => {
        fetch(BASE_URL_API + '/api/v1/topAlbums/' + localStorage.getItem('username'), fetchInit('/api/v1/topAlbums', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                fixArtistNames(data.topAlbums)
                addRankColumn(data.topAlbums)
                setAllAlbums(data.topAlbums)
                setDisplayedAlbums(data.topAlbums.slice(0, DEFAULT_ALBUM_COUNT_LIMIT))
            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [])

    function AlbumsDropdown() {
        const [isOpen, setIsOpen] = useState(false);
        const [dropdownValue, setDropdownValue] = useState(DEFAULT_ALBUM_COUNT_LIMIT)

        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (isOpen && !event.target.classList.toString().includes('dropdown')) {
                    setIsOpen(false);
                }
            })
        }, [isOpen])

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setDropdownValue(size)
            setDisplayedAlbums(allAlbums.slice(0, size))
        }

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
                        Select table size... {dropdownValue}
                    </div>
                </div>
            </div>

        );
    }

    return (
        <div>
            <table className={"table-all"}>
                <thead>
                <tr className={"table-column-names"}>
                    <th>Rank</th>
                    <th>Album name</th>
                    <th>Artist</th>
                    <th style={{textAlign: 'right'}}>Count</th>
                </tr>
                </thead>
                <tbody>
                {displayedAlbums.map(albumCount => (
                    <tr className={"table-row"}>
                        <td>{albumCount.rank}</td>
                        <td>{albumCount.album}</td>
                        <td>{albumCount.artist}</td>
                        <td style={{textAlign: 'right'}}>{albumCount.count}</td>
                    </tr>
                ))}
                </tbody>
            </table>
            <AlbumsDropdown/>
        </div>
    )
}*/

function TopTable(props) {
    props = props.props

    // TO PASS IN AS PROPS:
    // initialState, url, fixArtistNames, defaultCount, ddValues, thead, itemCallback

    const [allItems, setAllItems] = useState(props.initialState)
    const [displayedItems, setDisplayedItems] = useState(props.initialState)

    useEffect(() => {
        console.log("FETCHING: " + BASE_URL_API + props.url + '/' + localStorage.getItem('username'))
        fetch(BASE_URL_API + props.url + '/' + localStorage.getItem('username'), fetchInit(props.url, null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                if (props.fixArtistNames) fixArtistNames(data.items)
                addRankColumn(data.items)
                setAllItems(data.items)
                setDisplayedItems(data.items.slice(0, props.defaultCount))
            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [])

    function Dropdown() {
        const [isOpen, setIsOpen] = useState(false);
        const [dropdownValue, setDropdownValue] = useState(props.defaultCount)

        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (isOpen && !event.target.classList.toString().includes('dropdown')) {
                    setIsOpen(false);
                }
            })
        }, [isOpen])

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setDropdownValue(size)
            setDisplayedItems(allItems.slice(0, size))
        }

        return (
            <div className={'dd-wrapper'}>
                <div className='dropdown'>
                    {isOpen && (
                        <div className='dropdown-menu'>
                            <ul>
                                <li onClick={() => itemClicked(props.ddValues[0])}>props.ddValues[0]</li>
                                <li onClick={() => itemClicked(props.ddValues[1])}>props.ddValues[1]</li>
                                <li onClick={() => itemClicked(props.ddValues[2])}>props.ddValues[2]</li>
                                <li onClick={() => itemClicked(props.ddValues[3])}>props.ddValues[3]</li>
                            </ul>
                        </div>
                    )}
                    <div className='dropdown-button' onClick={toggle}>
                        Select table size... {dropdownValue}
                    </div>
                </div>
            </div>

        );
    }

    return (
        <div>
            <table className={"table-all"}>
                {props.thead}
                <tbody>
                {displayedItems != null && displayedItems.map(props.itemCallback)}
                </tbody>
            </table>
            <Dropdown/>
        </div>
    )
}

function DecadePieChart() {

    const [chartData, setChartData] = useState([["Decade", "Count"]])

    useEffect(() => {
        fetch(BASE_URL_API + '/api/v1/decadeBreakdown/' + localStorage.getItem('username'), fetchInit('/api/v1/decadeBreakdown', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setChartData(convertDecadesToPieChartData(data.decadeBreakdown))
            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [])

    return (
        <div className={'chart-wrapper'}>
            <Chart
                width={'100%'}
                height={'100%'}
                chartType="PieChart"
                data={chartData}
                options={{
                    backgroundColor: 'transparent',
                    fontColor: '#cce2e6',
                    legend: {
                        textStyle: {
                            color: '#cce2e6'
                        }
                    }
                }}
            />
        </div>
    )

}

// HELPER FUNCTIONS
function fixArtistNames(items) {
    items.forEach(item => {
        item.artist = item.artist.replaceAll(';;', ', ')
    })
}

function addRankColumn(items) {
    let rank = 1
    items.forEach(item => {
        item.rank = rank
        rank++
    })
}

function convertDecadesToPieChartData(data) {
    let result = [["Decade", "Count"]]
    data.forEach(item => {
        result.push([item.decade, item.count])
    })
    return result
}
