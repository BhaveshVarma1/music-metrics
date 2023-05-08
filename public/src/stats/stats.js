// noinspection JSUnresolvedVariable,JSCheckFunctionSignatures

import './stats.css';
import {BASE_URL_API, fetchInit, getToken, LoginButton, PrimaryInfo} from "../util/util";
import React, {useEffect, useState} from "react";
import {Chart} from "react-google-charts";

// Default values for the dropdowns (must be in the array specified in the props)
const DEFAULT_SONG_COUNT_LIMIT = 100
const DEFAULT_ARTIST_COUNT_LIMIT = 50
const DEFAULT_ALBUM_COUNT_LIMIT = 50

const OPEN_SPOTIFY = 'https://open.spotify.com'

export function Stats() {

    // STYLING VARIABLES
    const selectedStyle = 'selector-selected'
    const unselectedStyle = 'selector-unselected'
    const [songStyle, setSongStyle] = useState(selectedStyle);
    const [artistStyle, setArtistStyle] = useState(unselectedStyle);
    const [albumStyle, setAlbumStyle] = useState(unselectedStyle);
    const [chartStyle, setChartStyle] = useState(unselectedStyle);
    const [showSelector2, setShowSelector2] = useState(true);
    const [countStyle, setCountStyle] = useState(selectedStyle);
    const [timeStyle, setTimeStyle] = useState(unselectedStyle);

    // DATA VARIABLES (only fetched once, when the Stats component loads)
    const [averageLength, setAverageLength] = useState(0);
    const [averagePopularity, setAveragePopularity] = useState([]);
    const [averageYear, setAverageYear] = useState(0);
    const [decadeBreakdown, setDecadeBreakdown] = useState([]);
    const [hourBreakdown, setHourBreakdown] = useState([]);
    const [medianYear, setMedianYear] = useState(0);
    const [modeYear, setModeYear] = useState(0);
    const [percentExplicit, setPercentExplicit] = useState(0);
    const [topAlbums, setTopAlbums] = useState([]);
    const [topAlbumsTime, setTopAlbumsTime] = useState([]);
    const [topArtists, setTopArtists] = useState([]);
    const [topArtistsTime, setTopArtistsTime] = useState([]);
    const [topSongs, setTopSongs] = useState([]);
    const [topSongsTime, setTopSongsTime] = useState([]);
    const [totalSongs, setTotalSongs] = useState(0);
    const [uniqueAlbums, setUniqueAlbums] = useState(0);
    const [uniqueArtists, setUniqueArtists] = useState(0);
    const [uniqueSongs, setUniqueSongs] = useState(0);
    const [weekDayBreakdown, setWeekDayBreakdown] = useState([]);

    // OTHER
    const [dataOrLoading, setDataOrLoading] = useState(<Loading/>)
    const songCountProps = {
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
                    <td><a href={OPEN_SPOTIFY + '/track/' + item.songId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.song}</a></td>
                    <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                    <td style={{textAlign: 'right'}}>{item.count}</td>
                </tr>
            )
        }
    }
    const songTimeProps = {
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
                    <td><a href={OPEN_SPOTIFY + '/track/' + item.songId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.song}</a></td>
                    <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                    <td style={{textAlign: 'right'}}>{Math.round(item.count/60)}</td>
                </tr>
            )
        }
    }
    const artistCountProps = {
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
                    <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.artist}</a></td>
                    <td style={{textAlign: 'right'}}>{item.count}</td>
                </tr>
            )
        }
    }
    const artistTimeProps = {
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
                    <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.artist}</a></td>
                    <td style={{textAlign: 'right'}}>{Math.round(item.count/60)}</td>
                </tr>
            )
        }
    }
    const albumCountProps = {
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
                    <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.album}</a></td>
                    <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                    <td style={{textAlign: 'right'}}>{item.count}</td>
                </tr>
            )
        }
    }
    const albumTimeProps = {
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
                    <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.album}</a></td>
                    <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                    <td style={{textAlign: 'right'}}>{Math.round(item.count/60)}</td>
                </tr>
            )
        }
    }
    const [currentData, setCurrentData] = useState();

    const toggleLoading = () => {
        setCurrentData(<TopTable items={topSongs} props={songCountProps}/>)
        setDataOrLoading(
            <>
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
            </>
        )
    }

    useEffect(() => {
        fetch(BASE_URL_API + '/api/v1/allStats/' + localStorage.getItem('username'), fetchInit('/api/v1/allStats', null, getToken()))
            .then(response => response.json())
            .then(data => {

                console.log(data)

                // ADD RANK COLUMN FOR RELEVANT ARRAYS
                addRankColumn(data.topAlbums.items)
                addRankColumn(data.topAlbumsTime.items)
                addRankColumn(data.topArtists.items)
                addRankColumn(data.topArtistsTime.items)
                addRankColumn(data.topSongs.items)
                addRankColumn(data.topSongsTime.items)

                // DO CALCULATIONS FOR OTHER RELEVANT DATA
                let minutes = Math.floor(data.value / 60)

                // ASSIGN DATA TO RESPECTIVE STATES
                setAverageLength(minutes + ":" + (data.averageLength.value - minutes * 60))
                setAveragePopularity(data.averagePopularity.items)
                setAverageYear(data.averageYear.value)
                setDecadeBreakdown(data.decadeBreakdown.items)
                setHourBreakdown(data.hourBreakdown.items)
                setMedianYear(data.medianYear.value)
                setModeYear(data.modeYear.items)
                setPercentExplicit(data.percentExplicit.value + "%")
                setTopAlbums(data.topAlbums.items)
                setTopAlbumsTime(data.topAlbumsTime.items)
                setTopArtists(data.topArtists.items)
                setTopArtistsTime(data.topArtistsTime.items)
                setTopSongs(data.topSongs.items)
                setTopSongsTime(data.topSongsTime.items)
                setTotalSongs(data.totalSongs.value)
                setUniqueAlbums(data.uniqueAlbums.value)
                setUniqueArtists(data.uniqueArtists.value)
                setUniqueSongs(data.uniqueSongs.value)
                setWeekDayBreakdown(data.weekDayBreakdown.items)

                // REMOVING LOADING SCREEN AND SHOWS STATS
                toggleLoading()

            }).catch(error => {
                console.log("ERROR: " + error)
            })
    }, [toggleLoading])

    // LOGIN SCREEN
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

        setCurrentData(<TopTable items={topSongs} props={songCountProps}/>)
    }

    function setToArtist() {
        setSongStyle(unselectedStyle)
        setArtistStyle(selectedStyle)
        setAlbumStyle(unselectedStyle)
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable items={topArtists} props={artistCountProps}/>)
    }

    function setToAlbum() {
        setSongStyle(unselectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(selectedStyle)
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable items={topAlbums} props={albumCountProps}/>)
    }

    function setToChart() {
        setSongStyle(unselectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(unselectedStyle)
        setChartStyle(selectedStyle)
        setShowSelector2(false)

        setCurrentData(<AllCharts
            averageLength={averageLength}
            averagePopularity={averagePopularity}
            averageYear={averageYear}
            decadeBreakdown={decadeBreakdown}
            hourBreakdown={hourBreakdown}
            medianYear={medianYear}
            modeYear={modeYear}
            percentExplicit={percentExplicit}
            totalSongs={totalSongs}
            uniqueAlbums={uniqueAlbums}
            uniqueArtists={uniqueArtists}
            uniqueSongs={uniqueSongs}
            weekDayBreakdown={weekDayBreakdown}
        />)
    }

    function setToCount() {
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        if (songStyle === selectedStyle) {
            setCurrentData(<TopTable items={topSongs} props={songCountProps}/>)
        } else if (artistStyle === selectedStyle) {
            setCurrentData(<TopTable items={topArtists} props={artistCountProps}/>)
        } else if (albumStyle === selectedStyle) {
            setCurrentData(<TopTable items={topAlbums} props={albumCountProps}/>)
        }
    }

    function setToTime() {
        setCountStyle(unselectedStyle)
        setTimeStyle(selectedStyle)

        if (songStyle === selectedStyle) {
            setCurrentData(<TopTable items={topSongsTime} props={songTimeProps}/>)
        } else if (artistStyle === selectedStyle) {
            setCurrentData(<TopTable items={topArtistsTime} props={artistTimeProps}/>)
        } else if (albumStyle === selectedStyle) {
            setCurrentData(<TopTable items={topAlbumsTime} props={albumTimeProps}/>)
        }
    }

    return (
        <div>
            <PrimaryInfo text="Stats central."/>
            {dataOrLoading}
        </div>
    )

}

// SECONDARY COMPONENTS
function TopTable(props) {
    let allItems = props.items
    props = props.props

    // TO PASS IN AS PROPS:
    // array of items
    // defaultCount, ddValues, tableStyle, thead, itemCallback

    const [displayedItems, setDisplayedItems] = useState([])
    const [dropdownValue, setDropdownValue] = useState(props.defaultCount)

    useEffect(() => {
        setDropdownValue(props.defaultCount)
        setDisplayedItems(allItems.slice(0, props.defaultCount))
    }, [props, allItems])

    function Dropdown() {
        const [isOpen, setIsOpen] = useState(false);

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

function Loading() {
    return (
        <div>Loading...</div>
    )
}

function AllCharts(props) {

    return (
        <div className={'all-panels'}>
            <BasicPanel primary={"Average Year"} data={props.averageYear} commentary={"That was a good year."}/>
            <BasicPanel primary={"Average Song Length"} data={props.averageLength} commentary={"That's not very long."}/>
            <BasicPanel primary={"Median Year"} data={props.medianYear} commentary={"That was a better year."}/>
            <BasicPanel primary={"Percent Explicit"} data={props.percentExplicit} commentary={"That's too high."}/>
            <BasicPanel primary={"Total Songs"} data={props.totalSongs} commentary={"Looks like you spend too much time on Spotify."}/>
            <BasicPanel primary={"Unique Album Count"} data={props.uniqueAlbums} commentary={"Wow, not a whole lot of diversity there."}/>
            <BasicPanel primary={"Unique Artist Count"} data={props.uniqueArtists} commentary={"Nice!"}/>
            <BasicPanel primary={"Unique Song Count"} data={props.uniqueSongs} commentary={"That's pretty ok."}/>
            <BasicPanel primary={"Breakdown by Decade"} data={<DecadePieChart data={props.decadeBreakdown}/>} commentary={"Looks like you need more diversity."}/>
            <BasicPanel primary={"Breakdown by Hour"} data={<HourChart data={props.hourBreakdown}/>} last={true}/>
        </div>
    )
}

function DecadePieChart(props) {

    return (
        <div className={'decade-wrapper'}>
            <Chart
                chartType="PieChart"
                data={convertDecadesToPieChartData(props.data)}
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
                    enableInteractivity: false,
                }}
            />
        </div>
    )

}

function HourChart(props) {

        return (
            <div className={'hour-wrapper'}>
                <Chart
                    chartType="BarChart"
                    data={convertHoursToChartData(props.data)}
                    options={{
                        backgroundColor: 'transparent',
                        fontColor: '#cce2e6',
                        enableInteractivity: false,
                        orientation: 'horizontal',
                        hAxis: {
                            title: 'Hour',
                            textStyle: {
                                color: '#cce2e6',
                            }
                        },
                        vAxis: {
                            title: 'Count',
                            textStyle: {
                                color: '#cce2e6',
                            },
                            viewWindow: {
                                min: 0,
                                max: 150
                            }
                        },
                        chartArea: {
                            left: 50, // adjust the left margin to make space for the y-axis labels
                            top: 20, // adjust the top margin to make space for the x-axis labels
                            width: '100%', // adjust the width to make the chart larger
                            height: '60%' // adjust the height to make the chart larger
                        }
                    }}
                />
            </div>
        )
}

function BasicPanel(props) {

    let style = 'panel'
    if (props.last) {
        style += ' panel-last'
    }

    return (
        <div className={style}>
            <div className={'panel-primary'}>{props.primary}</div>
            <div className={'panel-right'}>
                <div className={'panel-data'}>{props.data}</div>
                <div className={'panel-commentary'}>{props.commentary}</div>
            </div>
        </div>
    )
}

function LinkedArtistList(props) {
    if (props === undefined || props.nameString === undefined || props.idString === undefined) {
        return null
    }
    const names = props.nameString.split(';;');
    const ids = props.idString.split(';;');

    return (
        <div>
            {names.map((name, index) => (
                <React.Fragment key={index}>
                    {index !== 0 && ', '}
                    <a key={index} href={`https://open.spotify.com/artist/${ids[index]}`} target={"_blank"} rel={"noreferrer"} className={'table-link'}>
                        {name}
                    </a>
                </React.Fragment>
            ))}
        </div>
    )
}

// HELPER FUNCTIONS
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

function convertHoursToChartData(data) {
    let result = [["Hour", "Count"]]
    let hours = ["12AM", "1AM", "2AM", "3AM", "4AM", "5AM", "6AM", "7AM", "8AM", "9AM", "10AM", "11AM", "12PM", "1PM", "2PM", "3PM", "4PM", "5PM", "6PM", "7PM", "8PM", "9PM", "10PM", "11PM"]
    let i = 0;
    data.forEach(item => {
        result.push([hours[i], item])
        i++
    })
    return result
}
