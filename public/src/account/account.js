import './account.css';
import {BASE_URL_API, fetchInit, getToken, LoginButton, PrimaryInfo} from "../util/util";
import React, {useState} from "react";
import {useDropzone} from "react-dropzone";

const maxFileSize = 20000000
const maxFiles = 15

export function Account() {

    // LOGIN SCREEN
    if (getToken() == null || getToken() === 'undefined') {
        sessionStorage.setItem('route', 'account')
        return (
            <div>
                <PrimaryInfo text="Log in see account info..."/>
                <LoginButton text="LOGIN TO SPOTIFY"/>
            </div>
        )
    }

    return (
        <div>
            <PrimaryInfo text="Account Information"/>
            <div className={'table-acct'}>
                <div className={'table-row-acct'}>
                    <div>Username</div>
                    <div>{truncateStr(localStorage.getItem('username'))}</div>
                </div>
                <div className={'table-row-acct'}>
                    <div>Display Name</div>
                    <div>{truncateStr(localStorage.getItem('display_name'))}</div>
                </div>
                <div className={'table-row-acct'}>
                    <div>Email</div>
                    <div>{localStorage.getItem('email')}</div>
                </div>
                <div className={'table-row-acct'}>
                    <div>Account Created</div>
                    <div>{unixMillisToString(localStorage.getItem('timestamp'))}</div>
                </div>
            </div>
            <Dropzone/>
        </div>
    )
}

function Dropzone() {

    const [files, setFiles] = useState([])
    const [errorMessage, setErrorMessage] = useState('')
    const [hoveredIndex, setHoveredIndex] = useState(-1)
    const [popupVisible, setPopupVisible] = useState(false)
    const [loadingMessage, setLoadingMessage] = useState('')

    function onDrop(acceptedFiles, rejectedFiles) {
        setLoadingMessage('')
        setHoveredIndex(-1)
        if (acceptedFiles?.length) {
            if (files.length + acceptedFiles.length > maxFiles) {
                setErrorMessage('Too many files')
                return
            }
            setErrorMessage('')
            const uniqueFiles = acceptedFiles.filter(file => !files.some(f => f.path === file.path))
            setFiles(previousFiles => [...previousFiles, ...uniqueFiles])
        }
        if (rejectedFiles?.length) setErrorMessage('Only .json files under 20MB are accepted')
    }

    function removeItem(item) {
        if (errorMessage === 'Too many files') setErrorMessage('')
        setFiles(files.filter(file => file.path !== item.path))
    }

    /**
     * Processes each file, and if all are formatted correctly, submits them to the server.
     */
    function submit() {
        setLoadingMessage('Loading...')
        setErrorMessage('')
        console.log(files)

        const uploadPromises = []
        const bodies = []
        files.forEach(file => {
            const reader = new FileReader()

            const promise = new Promise((resolve, reject) => {
                console.log("here3")
                reader.onload = (event) => {
                    if (!isFormattedCorrectly(event.target.result)) {
                        reject('Incorrect format')
                        return
                    }
                    bodies.push(event.target.result)
                }
                reader.readAsText(file)
            })
            uploadPromises.push(promise)
        })

        console.log("here")
        // Wait for all promises to resolve
        Promise.all(uploadPromises).then(() => {
            const fetchPromises = []
            console.log("here2")
            bodies.forEach(body => {
                const fetchPromise = new Promise((resolve, reject) => {
                    fetch(BASE_URL_API + '/api/v1/load/' + localStorage.getItem('username'), fetchInit('/api/v1/load', body, getToken()))
                        .then(response => response.json())
                        .then(data => {
                            console.log(data)
                            resolve(data)
                        }).catch(error => {
                            console.error(error)
                            reject(error)
                        })
                })
                fetchPromises.push(fetchPromise)
            })
            Promise.all(fetchPromises).then(() => {
                setFiles([])
                setLoadingMessage('Success! You will be able to view your updated stats within 24 hours.')
            }).catch(error => {
                console.error(error)
                setLoadingMessage('')
                setErrorMessage('There was an error on our part uploading your files. Please try again later.')
            })
        }).catch(error => {
            console.error(error)
            setLoadingMessage('')
            setErrorMessage('There seems to be a problem with the files you uploaded. Make sure they are the correct files and try again.')
        })
    }

    const { getRootProps, getInputProps, isDragActive } = useDropzone({
        onDrop,
        accept: {
            'text/json': ['.json'],
        },
        maxSize: maxFileSize,
    })

    return (
        <div className={'dropzone-all'}>
            <div className={'dropzone-info'}>
                Upload your extended streaming history here.
                <span className={'custom-link whats-this'} onClick={() => setPopupVisible(true)}> <u>What's this?</u></span>
            </div>
            <div {...getRootProps({
                className: 'dropzone'
            })}>
                <input {...getInputProps()} />
                {isDragActive ? (
                    <p>Drop the files here ...</p>
                ) : (
                    <p>Drag and drop .json files here, or click to select files</p>
                )}
            </div>
            {loadingMessage !== '' && <div className={'loading-acct'}>{loadingMessage}</div>}
            {errorMessage !== '' && <p className={'dropzone-error'}>{errorMessage}</p>}
            <ul>
                {files.map((file, index) => (
                    <li key={file.name} style={{position: "relative"}}>
                        <div
                            className={'dropzone-item'}
                            onMouseEnter={() => setHoveredIndex(index)}
                            onMouseLeave={() => setHoveredIndex(null)}
                        >
                            {file.name}
                        </div>
                        {hoveredIndex === index && <div className={'dropzone-item-remove'} onClick={() => removeItem(file)} onMouseEnter={() => setHoveredIndex(index)}></div>}
                    </li>
                ))}
            </ul>
            {files.length !== 0 && (
                <div className={'dropdown-submit-wrapper'}>
                    <div className={'login-button dropzone-submit'} onClick={submit}><b>SUBMIT</b></div>
                </div>
            )}
            {popupVisible && (
                <div className={'popup-container'}>
                    <div className={'popup-content'}>
                        <div style={{width: "100%"}}>
                            <h1>Extended Streaming History</h1>
                            <div>To obtain your extended streaming history, visit your
                                <a href={"https://www.spotify.com/us/account/privacy/"} target={"_blank"} rel={"noreferrer"} className={'custom-link'}> <u>Privacy Settings</u></a> on Spotify.
                             Select 'Extended streaming history' and click 'Request data'. It will be sent to you within 30 days.
                             Once you receive the data as a .zip, download and extract it. Then upload the endsong_x.json files here.
                            </div>
                        </div>
                        <div className={'login-button popup-ok'} onClick={() => setPopupVisible(false)}><b>OK</b></div>
                    </div>
                </div>
            )}
        </div>
    )
}

function unixMillisToString(unixMillis) {
    const date = new Date(+unixMillis)
    return date.toLocaleDateString()
}

function truncateStr(str) {
    let num = 25
    if (str.length <= num) return str
    return str.slice(0, num) + '...'
}

function isFormattedCorrectly(file) {
    file = JSON.parse(file)
    if (!Array.isArray(file)) {
        console.log('Not an array')
        return false;
    }

    for (let i = 0; i < file.length; i++) {
        const item = file[i];

        if (
            typeof item !== 'object' ||
            !item.hasOwnProperty('ts') ||
            !item.hasOwnProperty('username') ||
            !item.hasOwnProperty('platform') ||
            !item.hasOwnProperty('ms_played') ||
            !item.hasOwnProperty('conn_country') ||
            !item.hasOwnProperty('ip_addr_decrypted') ||
            !item.hasOwnProperty('user_agent_decrypted') ||
            !item.hasOwnProperty('master_metadata_track_name') ||
            !item.hasOwnProperty('master_metadata_album_artist_name') ||
            !item.hasOwnProperty('master_metadata_album_album_name') ||
            !item.hasOwnProperty('spotify_track_uri') ||
            !item.hasOwnProperty('episode_name') ||
            !item.hasOwnProperty('episode_show_name') ||
            !item.hasOwnProperty('spotify_episode_uri') ||
            !item.hasOwnProperty('reason_start') ||
            !item.hasOwnProperty('reason_end') ||
            !item.hasOwnProperty('shuffle') ||
            !item.hasOwnProperty('skipped') ||
            !item.hasOwnProperty('offline') ||
            !item.hasOwnProperty('offline_timestamp') ||
            !item.hasOwnProperty('incognito_mode')
        ) {
            return false;
        }
    }

    return true;
}
