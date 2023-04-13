// noinspection JSUnresolvedVariable,JSCheckFunctionSignatures

import './stats.css';
import {BASE_URL_API, fetchInit, getToken, LoginButton, PrimaryInfo} from "../util/util";
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
    const [chartStyle, setChartStyle] = useState(unselectedStyle);
    const [showSelector2, setShowSelector2] = useState(true);
    const [countStyle, setCountStyle] = useState(selectedStyle);
    const [timeStyle, setTimeStyle] = useState(unselectedStyle);

    const songCountProps = {
        initialState: [{"song": "Loading...", "artist": "Loading...", "count": 0}],
        url: '/api/v1/topSongs',
        fixArtistNames: true,
        defaultCount: DEFAULT_SONG_COUNT_LIMIT,
        ddValues: [25, 50, 100, 250],
        tableStyle: 'table-all',
        thead: (
            <thead>
            <tr className={"table-column-names"}>
                <th>Rank</th>
                <th>Song name</th>
                <th>Artist</th>
                <th style={{textAlign: 'right'}}>Count</th>
            </tr>
            </thead>
        ),
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
    const songTimeProps = {
        initialState: [{"song": "Loading...", "artist": "Loading...", "count": 0}],
        url: '/api/v1/topSongsTime',
        fixArtistNames: true,
        defaultCount: DEFAULT_SONG_COUNT_LIMIT,
        ddValues: [25, 50, 100, 250],
        tableStyle: 'table-all',
        thead: (
            <thead>
            <tr className={"table-column-names"}>
                <th>Rank</th>
                <th>Song name</th>
                <th>Artist</th>
                <th style={{textAlign: 'right'}}>Minutes</th>
            </tr>
            </thead>
        ),
        itemCallback: (item) => {
            return (
                <tr className={"table-row"}>
                    <td>{item.rank}</td>
                    <td>{item.song}</td>
                    <td>{item.artist}</td>
                    <td style={{textAlign: 'right'}}>{Math.round(item.count/60)}</td>
                </tr>
            )
        }
    }
    const artistCountProps = {
        initialState: [{"artist": "Loading...", "count": 0}],
        url: '/api/v1/topArtists',
        fixArtistNames: false,
        defaultCount: DEFAULT_ARTIST_COUNT_LIMIT,
        ddValues: [10, 25, 50, 100],
        tableStyle: 'table-all table-all-artist',
        thead: (
            <thead>
            <tr className={"table-column-names"}>
                <th>Rank</th>
                <th>Artist name</th>
                <th style={{textAlign: 'right'}}>Count</th>
            </tr>
            </thead>
        ),
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
    const artistTimeProps = {
        initialState: [{"artist": "Loading...", "count": 0}],
        url: '/api/v1/topArtistsTime',
        fixArtistNames: false,
        defaultCount: DEFAULT_ARTIST_COUNT_LIMIT,
        ddValues: [10, 25, 50, 100],
        tableStyle: 'table-all table-all-artist',
        thead: (
            <thead>
            <tr className={"table-column-names"}>
                <th>Rank</th>
                <th>Artist name</th>
                <th style={{textAlign: 'right'}}>Minutes</th>
            </tr>
            </thead>
        ),
        itemCallback: (item) => {
            return (
                <tr className={"table-row"}>
                    <td>{item.rank}</td>
                    <td>{item.artist}</td>
                    <td style={{textAlign: 'right'}}>{Math.round(item.count/60)}</td>
                </tr>
            )
        }
    }
    const albumCountProps = {
        initialState: [{"album": "Loading...", "artist": "Loading...", "count": 0}],
        url: '/api/v1/topAlbums',
        fixArtistNames: true,
        defaultCount: DEFAULT_ALBUM_COUNT_LIMIT,
        ddValues: [10, 25, 50, 100],
        tableStyle: 'table-all',
        thead: (
            <thead>
            <tr className={"table-column-names"}>
                <th>Rank</th>
                <th>Album name</th>
                <th>Artist</th>
                <th style={{textAlign: 'right'}}>Count</th>
            </tr>
            </thead>
        ),
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
    const albumTimeProps = {
        initialState: [{"album": "Loading...", "artist": "Loading...", "count": 0}],
        url: '/api/v1/topAlbumsTime',
        fixArtistNames: true,
        defaultCount: DEFAULT_ALBUM_COUNT_LIMIT,
        ddValues: [10, 25, 50, 100],
        tableStyle: 'table-all',
        thead: (
            <thead>
            <tr className={"table-column-names"}>
                <th>Rank</th>
                <th>Album name</th>
                <th>Artist</th>
                <th style={{textAlign: 'right'}}>Minutes</th>
            </tr>
            </thead>
        ),
        itemCallback: (item) => {
            return (
                <tr className={"table-row"}>
                    <td>{item.rank}</td>
                    <td>{item.album}</td>
                    <td>{item.artist}</td>
                    <td style={{textAlign: 'right'}}>{Math.round(item.count/60)}</td>
                </tr>
            )
        }
    }
    const [currentData, setCurrentData] = useState(<TopTable props={songCountProps}/>);

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
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable props={songCountProps}/>)
    }

    function setToArtist() {
        setSongStyle(unselectedStyle)
        setArtistStyle(selectedStyle)
        setAlbumStyle(unselectedStyle)
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable props={artistCountProps}/>)
    }

    function setToAlbum() {
        setSongStyle(unselectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(selectedStyle)
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable props={albumCountProps}/>)
    }

    function setToChart() {
        setSongStyle(unselectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(unselectedStyle)
        setChartStyle(selectedStyle)
        setShowSelector2(false)

        setCurrentData(<AllCharts/>)
    }

    function setToCount() {
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        if (songStyle === selectedStyle) {
            setCurrentData(<TopTable props={songCountProps}/>)
        } else if (artistStyle === selectedStyle) {
            setCurrentData(<TopTable props={artistCountProps}/>)
        } else if (albumStyle === selectedStyle) {
            setCurrentData(<TopTable props={albumCountProps}/>)
        }
    }

    function setToTime() {
        setCountStyle(unselectedStyle)
        setTimeStyle(selectedStyle)

        if (songStyle === selectedStyle) {
            setCurrentData(<TopTable props={songTimeProps}/>)
        } else if (artistStyle === selectedStyle) {
            setCurrentData(<TopTable props={artistTimeProps}/>)
        } else if (albumStyle === selectedStyle) {
            setCurrentData(<TopTable props={albumTimeProps}/>)
        }
    }

    return (
        <div>
            <PrimaryInfo text="Stats central."/>
            <div className={'selector'}>
                <div className={songStyle + ' selector-option corner-rounded-left'} onClick={setToSong}>Top Songs</div>
                <div className={artistStyle + ' selector-option'} onClick={setToArtist}>Top Artists</div>
                <div className={albumStyle + ' selector-option'} onClick={setToAlbum}>Top Albums</div>
                <div className={chartStyle + ' selector-option corner-rounded-right'} onClick={setToChart}>Other</div>
            </div>
            {showSelector2 && (
                <div className={'selector'}>
                    <div className={countStyle + ' selector-option corner-rounded-left'} onClick={setToCount}>By Count</div>
                    <div className={timeStyle + ' selector-option corner-rounded-right'} onClick={setToTime}>By Time</div>
                </div>
            )}
            {currentData}
        </div>
    )

}

// SECONDARY COMPONENTS
function TopTable(props) {
    props = props.props

    // TO PASS IN AS PROPS:
    // initialState, url, fixArtistNames, defaultCount, ddValues, tableStyle, thead, itemCallback

    const [allItems, setAllItems] = useState(props.initialState)
    const [displayedItems, setDisplayedItems] = useState(props.initialState)
    const [dropdownValue, setDropdownValue] = useState(props.defaultCount)

    useEffect(() => {
        setDropdownValue(props.defaultCount)
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
    }, [props])

    function Dropdown() {
        const [isOpen, setIsOpen] = useState(false);
        //const [dropdownValue, setDropdownValue] = useState(props.defaultCount)

        // Close the dropdown if the user clicks outside of it
        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (!event.target.classList.toString().includes('dropdown')) {
                    setIsOpen(false);
                }
            })
        }, [])

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
                                <li onClick={() => itemClicked(props.ddValues[0])}>{props.ddValues[0]}</li>
                                <li onClick={() => itemClicked(props.ddValues[1])}>{props.ddValues[1]}</li>
                                <li onClick={() => itemClicked(props.ddValues[2])}>{props.ddValues[2]}</li>
                                <li onClick={() => itemClicked(props.ddValues[3])}>{props.ddValues[3]}</li>
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
            <table className={props.tableStyle}>
                {props.thead}
                <tbody>
                {displayedItems != null && displayedItems.map(props.itemCallback)}
                </tbody>
            </table>
            <Dropdown/>
        </div>
    )
}

function AllCharts() {
    const [averageLength, setAverageLength] = useState('Calculating...');
    const [averageYear, setAverageYear] = useState('Calculating...');
    const [medianYear, setMedianYear] = useState('Calculating...');
    const [percentExplicit, setPercentExplicit] = useState('Calculating...');
    const [totalSongs, setTotalSongs] = useState('Calculating...');
    const [uniqueAlbums, setUniqueAlbums] = useState('Calculating...');
    const [uniqueArtists, setUniqueArtists] = useState('Calculating...');
    const [uniqueSongs, setUniqueSongs] = useState('Calculating...');

    // Fetches ALL the data used in this component
    useEffect(() => {
        fetch(BASE_URL_API + '/api/v1/averageLength/' + localStorage.getItem('username'), fetchInit('/api/v1/averageLength', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                let minutes = Math.floor(data.value / 60)
                setAverageLength(minutes + " min " + (data.value - minutes * 60) + " sec")
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/averageYear/' + localStorage.getItem('username'), fetchInit('/api/v1/averageYear', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setAverageYear(data.value)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/medianYear/' + localStorage.getItem('username'), fetchInit('/api/v1/medianYear', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setMedianYear(data.value)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/percentExplicit/' + localStorage.getItem('username'), fetchInit('/api/v1/percentExplicit', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setPercentExplicit(data.value + "%")
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/totalSongs/' + localStorage.getItem('username'), fetchInit('/api/v1/totalSongs', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setTotalSongs(data.value)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/uniqueAlbums/' + localStorage.getItem('username'), fetchInit('/api/v1/uniqueAlbums', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setUniqueAlbums(data.value)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/uniqueArtists/' + localStorage.getItem('username'), fetchInit('/api/v1/uniqueArtists', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setUniqueArtists(data.value)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
        fetch(BASE_URL_API + '/api/v1/uniqueSongs/' + localStorage.getItem('username'), fetchInit('/api/v1/uniqueSongs', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setUniqueSongs(data.value)
            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [])

    return (
        <div className={'all-panels'}>
            <BasicPanel primary={"Average Year"} description={"Average release year of every song listened to"} data={averageYear} commentary={"That was a good year."}/>
            <BasicPanel primary={"Average Song Length"} data={averageLength} commentary={"That's not very long."}/>
            <BasicPanel primary={"Median Year"} description={"The middle year of every song listened to"} data={medianYear} commentary={"That was a better year."}/>
            <BasicPanel primary={"Percent Explicit"} description={"The percentage of all listens marked as explicit"} data={percentExplicit} commentary={"Good for you."}/>
            <BasicPanel primary={"Total Songs"} description={"The total number of listens"} data={totalSongs} commentary={"Baskemtball."}/>
            <BasicPanel primary={"Unique Album Count"} description={"The number of different albums you've listened to"} data={uniqueAlbums} commentary={"Wow, not a whole lot of diversity there."}/>
            <BasicPanel primary={"Unique Artist Count"} description={"The number of different artists you've listened to"} data={uniqueArtists} commentary={"Nice!"}/>
            <BasicPanel primary={"Unique Song Count"} description={"The number of different songs you've listened to"} data={uniqueSongs} commentary={"Looks like you repeat quite a lot of songs there bud!"}/>
            <BasicPanel primary={"Breakdown by Decade"} data={<DecadePieChart/>}/>
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
                setChartData(convertDecadesToPieChartData(data.items))
            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [])

    return (
        <div style={{padding: '1rem'}}>
            <Chart
                chartType="PieChart"
                data={chartData}
                options={{
                    backgroundColor: 'transparent',
                    fontColor: '#cce2e6',
                    legend: {
                        position: 'left',
                        alignment: 'center',
                        textStyle: {
                            color: '#cce2e6'
                        }
                    },
                    chartArea: {
                        left: 0,
                        top: 0,
                        width: '100%',
                        height: '100%',
                    },
                }}
            />
        </div>
    )

}

function BasicPanel(props) {
    return (
        <div className={'panel'}>
            <div className={'panel-sub panel-left'}>
                <div className={'panel-primary'}>{props.primary}</div>
                <div className={'panel-desc'}>{props.description}</div>
            </div>
            <div className={'panel-sub panel-right'}>
                <div className={'panel-data'}>{props.data}</div>
                <div className={'panel-desc'}>{props.commentary}</div>
            </div>
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
