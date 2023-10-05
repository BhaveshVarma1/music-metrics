// noinspection JSUnresolvedVariable,JSCheckFunctionSignatures,JSUnusedGlobalSymbols

import './stats.css';
import {ExtendedStreamingInfo, fetchInit, getToken, LoginButton, PrimaryInfo} from "../util/util";
import React, {useEffect, useState} from "react";
import {Chart} from "react-google-charts";
import DatePicker from "react-datepicker";
import 'react-datepicker/dist/react-datepicker.css'
import spotifyIcon from './spotify-icon.png'

// Default values for the dropdowns (must be in the array specified in the props)
const DEFAULT_TRACK_COUNT_LIMIT = 100
const DEFAULT_ARTIST_COUNT_LIMIT = 50
const DEFAULT_ALBUM_COUNT_LIMIT = 50

const DEFAULT_TIME_RANGES = ['All time', 'Last 7 days', 'Last 30 days', 'This year so far', 'Custom range...']

const DEFAULT_START_TIME = 0
const DEFAULT_END_TIME = Date.now()

const OPEN_SPOTIFY = 'https://open.spotify.com'

export function Stats() {

    // STYLING VARIABLES
    const selectedStyle = 'selector-selected'
    const unselectedStyle = 'selector-unselected'
    const [trackStyle, setTrackStyle] = useState(selectedStyle);
    const [artistStyle, setArtistStyle] = useState(unselectedStyle);
    const [albumStyle, setAlbumStyle] = useState(unselectedStyle);
    const [chartStyle, setChartStyle] = useState(unselectedStyle);
    const [showSelector2, setShowSelector2] = useState(true);
    const [countStyle, setCountStyle] = useState(selectedStyle);
    const [timeStyle, setTimeStyle] = useState(unselectedStyle);
    const [showDataSelectors, setShowDataSelectors] = useState(false);
    const [showTimeSelector, setShowTimeSelector] = useState(false);

    // TIME VARIABLES
    const [displayedTimeRange, setDisplayedTimeRange] = useState('All time');
    const [usingCustomTimeRange, setUsingCustomTimeRange] = useState(false);
    const [startTime, setStartTime] = useState(DEFAULT_START_TIME);
    const [endTime, setEndTime] = useState(DEFAULT_END_TIME);
    const [selectedStartDate, setSelectedStartDate] = useState(new Date());
    const [selectedEndDate, setSelectedEndDate] = useState(new Date());

    // DATA VARIABLES (only fetched once, when the Stats component loads)
    const [averageLength, setAverageLength] = useState(0);
    const [averagePopularity, setAveragePopularity] = useState([]);
    const [averageYear, setAverageYear] = useState(0);
    const [decadeBreakdown, setDecadeBreakdown] = useState([]);
    const [firstTrack, setFirstTrack] = useState({});
    const [hourBreakdown, setHourBreakdown] = useState([]);
    const [medianYear, setMedianYear] = useState(0);
    const [modeYear, setModeYear] = useState(0);
    const [percentExplicit, setPercentExplicit] = useState(0);
    const [topAlbums, setTopAlbums] = useState([]);
    const [topAlbumsTime, setTopAlbumsTime] = useState([]);
    const [topArtists, setTopArtists] = useState([]);
    const [topArtistsTime, setTopArtistsTime] = useState([]);
    const [topTracks, setTopTracks] = useState([]);
    const [topTracksTime, setTopTracksTime] = useState([]);
    const [totalMinutes, setTotalMinutes] = useState(0);
    const [totalTracks, setTotalTracks] = useState(0);
    const [uniqueAlbums, setUniqueAlbums] = useState(0);
    const [uniqueArtists, setUniqueArtists] = useState(0);
    const [uniqueTracks, setUniqueTracks] = useState(0);
    const [weekDayBreakdown, setWeekDayBreakdown] = useState([]);

    // OTHER
    const [currentData, setCurrentData] = useState(<StatsInfo text={"Loading..."}/>);
    const [popupVisible, setPopupVisible] = useState(true)

    useEffect(() => {
        if (getToken() == null || getToken() === 'undefined') return
        fetch('/stats/' + localStorage.getItem('username') + '/' + startTime + '-' + endTime, fetchInit('/stats', null, getToken()))
            .then(response => response.json())
            .then(data => {
                console.log(data)

                if (data === "No tracks found for this time period.") {
                    setShowTimeSelector(true)
                    setShowDataSelectors(false)
                    setCurrentData(<StatsInfo text="No listening history found for this time period."/>)
                    return
                }

                // ADD RANK COLUMN FOR RELEVANT ARRAYS
                addRankColumn(data.topAlbums.items)
                addRankColumn(data.topAlbumsTime.items)
                addRankColumn(data.topArtists.items)
                addRankColumn(data.topArtistsTime.items)
                addRankColumn(data.topTracks.items)
                addRankColumn(data.topTracksTime.items)

                // DO CALCULATIONS FOR OTHER RELEVANT DATA
                let minutes = Math.floor(data.averageLength.value / 60)

                // ASSIGN DATA TO RESPECTIVE STATES
                setAverageLength(minutes + ":" + makeIntDoubleDigit(data.averageLength.value - minutes * 60))
                setAveragePopularity(data.averagePopularity.items)
                setAverageYear(data.averageYear.value)
                setDecadeBreakdown(convertDecadesToPieChartData(data.decadeBreakdown.items))
                setFirstTrack(data.firstTrack)
                setHourBreakdown(data.hourBreakdown.items)
                setMedianYear(data.medianYear.value)
                setModeYear(data.modeYear.items)
                setPercentExplicit(data.percentExplicit.value + "%")
                setTopAlbums(data.topAlbums.items)
                setTopAlbumsTime(data.topAlbumsTime.items)
                setTopArtists(data.topArtists.items)
                setTopArtistsTime(data.topArtistsTime.items)
                setTopTracks(data.topTracks.items)
                setTopTracksTime(data.topTracksTime.items)
                setTotalMinutes(addCommaToNumber(data.totalMinutes.value))
                setTotalTracks(addCommaToNumber(data.totalTracks.value))
                setUniqueAlbums(addCommaToNumber(data.uniqueAlbums.value))
                setUniqueArtists(addCommaToNumber(data.uniqueArtists.value))
                setUniqueTracks(addCommaToNumber(data.uniqueTracks.value))
                setWeekDayBreakdown(addWeekDayToData(data.weekDayBreakdown.items))

                // REMOVE LOADING MESSAGE
                setShowTimeSelector(true)
                setShowDataSelectors(true)
                setCurrentData(<TopTable items={data.topTracks.items} type={'trackCount'}/>)

                // RESET THE SELECTORS
                setTrackStyle(selectedStyle)
                setArtistStyle(unselectedStyle)
                setAlbumStyle(unselectedStyle)
                setChartStyle(unselectedStyle)
                setShowSelector2(true)
                setCountStyle(selectedStyle)
                setTimeStyle(unselectedStyle)
                setCountStyle(selectedStyle)
                setTimeStyle(unselectedStyle)

            }).catch(error => {
                console.error(error)
            })
    }, [startTime, endTime])

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

    function setToTrack() {
        setTrackStyle(selectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(unselectedStyle)
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable items={topTracks} type={'trackCount'}/>)
    }
    function setToArtist() {
        setTrackStyle(unselectedStyle)
        setArtistStyle(selectedStyle)
        setAlbumStyle(unselectedStyle)
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable items={topArtists} type={'artistCount'}/>)
    }
    function setToAlbum() {
        setTrackStyle(unselectedStyle)
        setArtistStyle(unselectedStyle)
        setAlbumStyle(selectedStyle)
        setChartStyle(unselectedStyle)
        setShowSelector2(true)
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        setCurrentData(<TopTable items={topAlbums} type={'albumCount'}/>)
    }
    function setToChart() {
        setTrackStyle(unselectedStyle)
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
            totalMinutes={totalMinutes}
            totalTracks={totalTracks}
            uniqueAlbums={uniqueAlbums}
            uniqueArtists={uniqueArtists}
            uniqueTracks={uniqueTracks}
            weekDayBreakdown={weekDayBreakdown}
        />)
    }
    function setToCount() {
        setCountStyle(selectedStyle)
        setTimeStyle(unselectedStyle)

        if (trackStyle === selectedStyle) {
            setCurrentData(<TopTable items={topTracks} type={'trackCount'}/>)
        } else if (artistStyle === selectedStyle) {
            setCurrentData(<TopTable items={topArtists} type={'artistCount'}/>)
        } else if (albumStyle === selectedStyle) {
            setCurrentData(<TopTable items={topAlbums} type={'albumCount'}/>)
        }
    }
    function setToTime() {
        setCountStyle(unselectedStyle)
        setTimeStyle(selectedStyle)

        if (trackStyle === selectedStyle) {
            setCurrentData(<TopTable items={topTracksTime} type={'trackTime'}/>)
        } else if (artistStyle === selectedStyle) {
            setCurrentData(<TopTable items={topArtistsTime} type={'artistTime'}/>)
        } else if (albumStyle === selectedStyle) {
            setCurrentData(<TopTable items={topAlbumsTime} type={'albumTime'}/>)
        }
    }

    function submitTimes(potStartTime, potEndTime) {
        if (validateTimes(potStartTime, potEndTime)) {
            if (potStartTime === startTime && potEndTime === endTime) return
            setStartTime(potStartTime)
            setEndTime(potEndTime)
            setShowTimeSelector(false)
            setShowDataSelectors(false)
            setCurrentData(<StatsInfo text="Loading..."/>)
        } else {
            console.log("ERROR: Invalid times: " + potStartTime + " " + potEndTime)
            setShowTimeSelector(true)
            setShowDataSelectors(false)
            setCurrentData(<StatsInfo text="Invalid time range, try again."/>)
        }
        // useEffect triggered when startTime / endTime change
    }

    function Dropdown() {

        const [isOpen, setIsOpen] = useState(false);

        // Close the dropdown if the user clicks outside of it
        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (!event.target.classList.toString().includes('dropdown-time')) setIsOpen(false);
            })
        }, [])

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(startTime, endTime, index) {
            // Assumes that custom time range is the last item in the array
            if (index !== DEFAULT_TIME_RANGES.length - 1) {
                submitTimes(startTime, endTime)
                setUsingCustomTimeRange(false)
            } else {
                setUsingCustomTimeRange(true)
            }
            toggle()
            setDisplayedTimeRange(DEFAULT_TIME_RANGES[index])
        }

        return (
            <div className={'dd-wrapper'}>
                <div className='dropdown-time'>
                    {isOpen && (
                        <div className='dropdown-time-menu'>
                            <ul>
                                <li onClick={() => {
                                    itemClicked(0, Date.now(), 0)
                                }}>{DEFAULT_TIME_RANGES[0]}</li>
                                <li onClick={() => {
                                    const now = Date.now()
                                    itemClicked(now - (7 * 24 * 60 * 60 * 1000), now, 1)
                                }}>{DEFAULT_TIME_RANGES[1]}</li>
                                <li onClick={() => {
                                    const now = Date.now()
                                    itemClicked(now - (30 * 24 * 60 * 60 * 1000), now, 2)
                                }}>{DEFAULT_TIME_RANGES[2]}</li>
                                <li onClick={() => {
                                    const now = new Date()
                                    const yearEpoch = new Date(now.getFullYear(), 0, 1);
                                    itemClicked(yearEpoch.getTime(), Date.now(), 3)
                                }}>{DEFAULT_TIME_RANGES[3]}</li>
                                <li onClick={() => {
                                    itemClicked(0, Date.now(), 4)
                                }}>{DEFAULT_TIME_RANGES[4]}</li>
                            </ul>
                        </div>
                    )}
                    <div className='dropdown-time-button' onClick={toggle}>
                        {displayedTimeRange}
                    </div>
                </div>
            </div>
        )
    }

    return (
        <div>
            <PrimaryInfo text="Stats central."/>
            {showTimeSelector && (
                <>
                    <div className={'small-description'}>Showing stats from:</div>
                    <div className={'extra-bottom-margin'}>
                        <Dropdown/>
                        {usingCustomTimeRange && (
                            <div className={'custom-time-wrapper'}>
                                <div className={'custom-time'}>
                                    <div className={'time-input-wrapper'}>
                                        <DatePicker className={'time-input'} selected={selectedStartDate} onChange={(date) => {
                                            setSelectedStartDate(date)
                                        }}/>
                                    </div>
                                    <div className={'time-input-wrapper'}>
                                        <DatePicker className={'time-input'} selected={selectedEndDate} onChange={(date) => {
                                            setSelectedEndDate(date)
                                        }}/>
                                    </div>
                                    <div className={'submit-times'} onClick={() => submitTimes(dateToMillis(selectedStartDate), dateToMillis(selectedEndDate) + 86399999)}><b>GO</b></div>
                                </div>
                            </div>
                        )}
                    </div>
                </>
            )}
            {showDataSelectors && (
                <>
                    <div className={'selector'}>
                        <div className={trackStyle + ' selector-option corner-rounded-left'} onClick={setToTrack}>Top Tracks</div>
                        <div className={artistStyle + ' selector-option'} onClick={setToArtist}>Top Artists</div>
                        <div className={albumStyle + ' selector-option'} onClick={setToAlbum}>Top Albums</div>
                        <div className={chartStyle + ' selector-option corner-rounded-right'} onClick={setToChart}>Other</div>
                    </div>
                    {showSelector2 && (
                        <div className={'selector extra-bottom-margin'}>
                            <div className={countStyle + ' selector-option corner-rounded-left'} onClick={setToCount}>By Count</div>
                            <div className={timeStyle + ' selector-option corner-rounded-right'} onClick={setToTime}>By Time</div>
                        </div>
                    )}
                </>
            )}
            {currentData}
            {popupVisible && (
                <ExtendedStreamingInfo callback={() => setPopupVisible(false)} />
            )}
        </div>
    )
}

// SECONDARY COMPONENTS
function TopTable(props) {

    // TO PASS IN AS PROPS:
    // array of items
    // string representing type of item (trackCount, artistTime, ...)

    const [currentHeader, setCurrentHeader] = useState(null)
    const [currentRow, setCurrentRow] = useState(null)
    const [displayedItems, setDisplayedItems] = useState([])
    const [tableStyle, setTableStyle] = useState('table-all')
    const [dropdown, setDropdown] = useState(null)
    const allItems = props.items

    useEffect(() => {

        let type = props.type

        let currentProps = null
        const trackCountProps = {
            defaultCount: DEFAULT_TRACK_COUNT_LIMIT,
            ddValues: [25, 50, 100, 250],
            tableStyle: 'table-all',
            head_horiz: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "5%"}}>#</th>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "40%"}}>Track</th>
                    <th style={{width: "40%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Listens</th>
                </tr>
                </thead>
            ),
            head_vert: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{}}></th>
                    <th style={{}}>#</th>
                    <th style={{}}></th>
                    <th style={{}}>Track</th>
                    <th style={{textAlign: 'right'}}>Listens</th>
                </tr>
                </thead>
            ),
            row_horiz: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td><a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.track}</a></td>
                        <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(item.count)}</td>
                    </tr>
                )
            },
            row_vert: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td>
                            <div style={{display: "flex", flexDirection: "column"}}>
                                <a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"} className={'table-link'}><b>{item.track}</b></a>
                                <LinkedArtistList nameString={item.artist} idString={item.artistId}/>
                            </div>
                        </td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(item.count)}</td>
                    </tr>
                )
            }
        }
        const trackTimeProps = {
            defaultCount: DEFAULT_TRACK_COUNT_LIMIT,
            ddValues: [25, 50, 100, 250],
            tableStyle: 'table-all',
            head_horiz: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "5%"}}>#</th>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "40%"}}>Track</th>
                    <th style={{width: "40%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Minutes</th>
                </tr>
                </thead>
            ),
            head_vert: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{}}></th>
                    <th style={{}}>#</th>
                    <th style={{}}></th>
                    <th style={{}}>Track</th>
                    <th style={{textAlign: 'right'}}>Minutes</th>
                </tr>
                </thead>
            ),
            row_horiz: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td><a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.track}</a></td>
                        <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(Math.round(item.count/60))}</td>
                    </tr>
                )
            },
            row_vert: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td>
                            <div>
                                <a href={OPEN_SPOTIFY + '/track/' + item.trackId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.track}</a>
                                <LinkedArtistList nameString={item.artist} idString={item.artistId}/>
                            </div>
                        </td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(Math.round(item.count/60))}</td>
                    </tr>
                )
            }
        }
        const artistCountProps = {
            defaultCount: DEFAULT_ARTIST_COUNT_LIMIT,
            ddValues: [10, 25, 50, 100],
            tableStyle: 'table-all table-all-artist',
            head_horiz: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "15%"}}>#</th>
                    <th style={{width: "75%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Listens</th>
                </tr>
                </thead>
            ),
            head_vert: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "15%"}}>#</th>
                    <th style={{width: "75%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Listens</th>
                </tr>
                </thead>
            ),
            row_horiz: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.artist}</a></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(item.count)}</td>
                    </tr>
                )
            },
            row_vert: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.artist}</a></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(item.count)}</td>
                    </tr>
                )
            }
        }
        const artistTimeProps = {
            defaultCount: DEFAULT_ARTIST_COUNT_LIMIT,
            ddValues: [10, 25, 50, 100],
            tableStyle: 'table-all table-all-artist',
            head_horiz: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "15%"}}>#</th>
                    <th style={{width: "75%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Minutes</th>
                </tr>
                </thead>
            ),
            head_vert: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "15%"}}>#</th>
                    <th style={{width: "75%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Minutes</th>
                </tr>
                </thead>
            ),
            row_horiz: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.artist}</a></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(Math.round(item.count/60))}</td>
                    </tr>
                )
            },
            row_vert: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><a href={OPEN_SPOTIFY + '/artist/' + item.artistId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.artist}</a></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(Math.round(item.count/60))}</td>
                    </tr>
                )
            }
        }
        const albumCountProps = {
            defaultCount: DEFAULT_ALBUM_COUNT_LIMIT,
            ddValues: [10, 25, 50, 100],
            tableStyle: 'table-all',
            head_horiz: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "5%"}}>#</th>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "40%"}}>Album</th>
                    <th style={{width: "40%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Listens</th>
                </tr>
                </thead>
            ),
            head_vert: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{}}></th>
                    <th style={{}}>#</th>
                    <th style={{}}></th>
                    <th style={{}}>Album</th>
                    <th style={{textAlign: 'right'}}>Listens</th>
                </tr>
                </thead>
            ),
            row_horiz: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.album}</a></td>
                        <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(item.count)}</td>
                    </tr>
                )
            },
            row_vert: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td>
                            <div>
                                <a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.album}</a>
                                <LinkedArtistList nameString={item.artist} idString={item.artistId}/>
                            </div>
                        </td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(item.count)}</td>
                    </tr>
                )
            }
        }
        const albumTimeProps = {
            defaultCount: DEFAULT_ALBUM_COUNT_LIMIT,
            ddValues: [10, 25, 50, 100],
            tableStyle: 'table-all',
            head_horiz: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "5%"}}>#</th>
                    <th style={{width: "5%"}}></th>
                    <th style={{width: "40%"}}>Album</th>
                    <th style={{width: "40%"}}>Artist</th>
                    <th style={{textAlign: 'right', width: "5%"}}>Minutes</th>
                </tr>
                </thead>
            ),
            head_vert: (
                <thead>
                <tr className={"table-column-names"}>
                    <th style={{}}></th>
                    <th style={{}}>#</th>
                    <th style={{}}></th>
                    <th style={{}}>Album</th>
                    <th style={{textAlign: 'right'}}>Minutes</th>
                </tr>
                </thead>
            ),
            row_horiz: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.album}</a></td>
                        <td><LinkedArtistList nameString={item.artist} idString={item.artistId}/></td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(Math.round(item.count/60))}</td>
                    </tr>
                )
            },
            row_vert: (item) => {
                return (
                    <tr className={"table-row"}>
                        <td><a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"}><img src={spotifyIcon} alt={"Unavailable"} style={{width: "1.5rem"}}/></a></td>
                        <td>{item.rank}</td>
                        <td><img src={item.image} style={{width: "3rem"}} alt={"Unavailable"}/></td>
                        <td>
                            <div>
                                <a href={OPEN_SPOTIFY + '/album/' + item.albumId} target={"_blank"} rel={"noreferrer"} className={'table-link'}>{item.album}</a>
                                <LinkedArtistList nameString={item.artist} idString={item.artistId}/>
                            </div>
                        </td>
                        <td style={{textAlign: 'right'}}>{addCommaToNumber(Math.round(item.count/60))}</td>
                    </tr>
                )
            }
        }

        switch (type) {
            case 'trackCount':
                currentProps = trackCountProps
                break
            case 'trackTime':
                currentProps = trackTimeProps
                break
            case 'artistCount':
                currentProps = artistCountProps
                break
            case 'artistTime':
                currentProps = artistTimeProps
                break
            case 'albumCount':
                currentProps = albumCountProps
                break
            case 'albumTime':
                currentProps = albumTimeProps
                break
            default:
                currentProps = trackCountProps
        }
        setDisplayedItems(allItems.slice(0, currentProps.defaultCount))
        setTableStyle(currentProps.tableStyle)
        setDropdown(<Dropdown values={currentProps.ddValues} initial={currentProps.defaultCount} />)

        const mediaQuery = window.matchMedia('(orientation: portrait)');
        setCurrentHeader(mediaQuery.matches ? currentProps.head_vert : currentProps.head_horiz)
        setCurrentRow(mediaQuery.matches ? () => currentProps.row_vert : () => currentProps.row_horiz)

        const handleOrientationChange = (event) => {
            setCurrentHeader(event.matches ? currentProps.head_vert : currentProps.head_horiz)
            setCurrentRow(event.matches ? () => currentProps.row_vert : () => currentProps.row_horiz)
        };

        mediaQuery.addEventListener('change', handleOrientationChange);
        return () => {
            mediaQuery.removeEventListener('change', handleOrientationChange);
        };
    }, [props, allItems])

    function Dropdown(props) {

        const ddValues = props.values
        const [isOpen, setIsOpen] = useState(false);
        const [currentValue, setCurrentValue] = useState(props.initial)

        // Close the dropdown if the user clicks outside of it
        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (!event.target.classList.toString().includes('dropdown-stats')) {
                    setIsOpen(false);
                }
            })
        }, [])

        function toggle() {
            setIsOpen(!isOpen);
        }

        function itemClicked(size) {
            toggle()
            setCurrentValue(size)
            setDisplayedItems(allItems.slice(0, size))
        }

        return (
            <div className={'dd-wrapper'}>
                <div className='dropdown-stats'>
                    {isOpen && (
                        <div className='dropdown-stats-menu'>
                            <ul>
                                <li onClick={() => itemClicked(ddValues[0])}>{ddValues[0]}</li>
                                <li onClick={() => itemClicked(ddValues[1])}>{ddValues[1]}</li>
                                <li onClick={() => itemClicked(ddValues[2])}>{ddValues[2]}</li>
                                <li onClick={() => itemClicked(ddValues[3])}>{ddValues[3]}</li>
                            </ul>
                        </div>
                    )}
                    <div className='dropdown-stats-button' onClick={toggle}>
                        Select table size... {currentValue}
                    </div>
                </div>
            </div>
        );
    }

    return (
        <div>
            <table className={tableStyle}>
                {currentHeader}
                <tbody>
                {(displayedItems != null && currentRow != null) && displayedItems.map(currentRow)}
                </tbody>
            </table>
            {dropdown}
        </div>
    )
}

function StatsInfo(props) {
    return (
        <div className={'default-text-color loading'}>{props.text}</div>
    )
}

function AllCharts(props) {

    return (
        <>
            <div className={'all-panels'}>
                <BasicPanel primary={"Total Minutes"} data={props.totalMinutes} commentary={getCommentary(props.totalMinutes, 'totalMinutes', props.totalMinutes)}/>
                <BasicPanel primary={"Average Year"} data={props.averageYear} commentary={getCommentary(props.totalMinutes, 'averageYear', props.averageYear)}/>
                <BasicPanel primary={"Average Track Length"} data={props.averageLength} commentary={getCommentary(props.totalMinutes, 'averageLength', props.averageLength)}/>
                <BasicPanel primary={"Median Year"} data={props.medianYear} commentary={getCommentary(props.totalMinutes, 'medianYear', props.medianYear)}/>
                {/*<BasicPanel primary={"Percent Explicit"} data={props.percentExplicit} commentary={getCommentary(props.totalMinutes, 'percentExplicit', props.percentExplicit)}/>*/}
                <BasicPanel primary={"Total Tracks"} data={props.totalTracks} commentary={getCommentary(props.totalMinutes, 'totalTracks', props.totalTracks)}/>
                <BasicPanel primary={"Unique Album Count"} data={props.uniqueAlbums} commentary={getCommentary(props.totalMinutes, 'uniqueAlbums', props.uniqueAlbums)}/>
                <BasicPanel primary={"Unique Artist Count"} data={props.uniqueArtists} commentary={getCommentary(props.totalMinutes, 'uniqueArtists', props.uniqueArtists)}/>
                <BasicPanel primary={"Unique Track Count"} data={props.uniqueTracks} commentary={getCommentary(props.totalMinutes, 'uniqueTracks', props.uniqueTracks)}/>
                <BasicPanel primary={"Breakdown by Decade"} data={<DecadePieChart data={props.decadeBreakdown}/>} commentary={getCommentary(props.totalMinutes, 'decadeBreakdown', props.decadeBreakdown)}/>
                <BasicPanel primary={"Breakdown by Day"} data={<DecadePieChart data={props.weekDayBreakdown}/>}/>
                <BasicPanel primary={"Breakdown by Hour"} data={<HourChart data={props.hourBreakdown}/>} last={true}/>
            </div>
            <div className={'disclaimer'}>
                Disclaimer: Spotify does not provide time zone data, so breakdowns by hour and day are based on Mountain Daylight Time (UTC -6:00).
            </div>
        </>

    )
}

function DecadePieChart(props) {

    // noinspection JSValidateTypes
    return (
        <div className={'decade-wrapper'}>
            <Chart
                chartType="PieChart"
                data={props.data}
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

        // noinspection JSValidateTypes
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
                            titleTextStyle: {
                                color: '#cce2e6',
                            },
                            textStyle: {
                                color: '#cce2e6',
                            }
                        },
                        vAxis: {
                            title: 'Count',
                            titleTextStyle: {
                                color: '#cce2e6',
                            },
                            textStyle: {
                                color: '#cce2e6',
                            },
                            viewWindow: {
                                min: 0,
                                max: findMax(props.data)
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
                {/*<div className={'panel-commentary'}>{props.commentary}</div>*/}
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
function getCommentary(minutes, metric, data) {
    switch (metric) {
        case 'totalMinutes':
            return "Rookie numbers."
        case 'averageYear':
            return "That was a good year."
        case 'averageLength':
            return "That's not very long."
        case 'medianYear':
            return "That was a better year."
        case 'percentExplicit':
            return "That's too high."
        case 'totalTracks':
            return "Looks like you spend too much time on Spotify."
        case 'uniqueAlbums':
            return "Wow, not a whole lot of diversity there."
        case 'uniqueArtists':
            return "Nice!"
        case 'uniqueTracks':
            return "That's pretty ok."
        case 'decadeBreakdown':
            return "Looks like you need more diversity."
        case 'hourBreakdown':
            return "You're a night owl."
        default:
            return "That's pretty ok."
    }
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

function makeIntDoubleDigit(num) {
    if (num < 10) return '0' + num
    return num
}

function validateTimes(startTime, endTime) {
    startTime = +startTime
    endTime = +endTime
    if (isNaN(startTime) || isNaN(endTime)) return false
    if (startTime < 0) return false
    if (endTime < startTime) return false
    if (endTime < 1145746800000) return false // The day Spotify was released
    return true
}

function dateToMillis(date) {
    const beginningOfDay = new Date(date.getFullYear(), date.getMonth(), date.getDate());
    return beginningOfDay.getTime();
}

function addCommaToNumber(num) {
    return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

function addWeekDayToData(data) {
    // assumes data is formatted correctly
    let result = []
    result.push(["Day", "Count"])
    result.push(["Sunday", data[0]])
    result.push(["Monday", data[1]])
    result.push(["Tuesday", data[2]])
    result.push(["Wednesday", data[3]])
    result.push(["Thursday", data[4]])
    result.push(["Friday", data[5]])
    result.push(["Saturday", data[6]])
    return result
}

function findMax(data) {
    let max = 0
    data.forEach(item => {
        if (item > max) max = item
    })
    return max
}
