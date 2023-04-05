// noinspection JSUnresolvedVariable

import './util.css';
import logo from './logo.png';
import {Link} from 'react-router-dom';
import React, {useEffect, useState} from "react";
import {websocket} from "../index";

// GLOBAL CONSTANTS ----------------------------------------------------------------------------------------------------

export const BASE_URL_API = 'https://dev.musicmetrics.app';
export const BASE_URL_WEB = 'https://dev.musicmetrics.app';
const HTTP_METHODS = {
    '/api/v1/updateCode': 'POST',
    '/api/v1/averageYear': 'GET',
    '/api/v1/topSongs': 'GET',
    '/api/v1/topArtists': 'GET',
    '/api/v1/topAlbums': 'GET',
    '/api/v1/decadeBreakdown': 'GET',
}

// ELEMENTS COMMON TO EVERY PAGE ---------------------------------------------------------------------------------------

export function Header() {
    const [navButtons, setNavButtons] = useState(null);

    function NavButton(props) {
        return (
            <Link to={props.url} className='custom-link'>
                <div className="nav-button">
                    {props.text}
                </div>
            </Link>
        )
    }

    function HamburgerMenu() {
        const [isOpen, setIsOpen] = useState(false);

        function toggle() {
            setIsOpen(!isOpen);
        }

        useEffect(() => {
            document.addEventListener('click', (event) => {
                if (!event.target.classList.toString().includes('hamburger')) {
                    setIsOpen(false);
                }
            })
        }, [])

        return (
            <div className='hamburger'>
                <div className='hamburger-button' onClick={toggle}>
                    <div className='hamburger-button-line'/>
                    <div className='hamburger-button-line'/>
                    <div className='hamburger-button-line'/>
                </div>
                {isOpen && (
                    <div className='hamburger-menu'>
                        <ul>
                            <Link to="/" className='custom-link'><li onClick={toggle}>HOME</li></Link>
                            <Link to="/stats" className='custom-link'><li onClick={toggle}>STATS</li></Link>
                            <Link to="/account" className='custom-link'><li onClick={toggle}>ACCOUNT</li></Link>
                        </ul>
                    </div>
                )}
            </div>
        )
    }

    // Dynamically change the navigation bar based on the orientation of the device.
    useEffect(() => {
        // The default navigation bar in landscape mode
        const navButtonsLandscape = (
            <div className='nav-buttons default-text-color'>
                <NavButton url="/" text="HOME"/>
                <NavButton url="/stats" text="STATS"/>
                <NavButton url="/account" text="ACCOUNT"/>
            </div>
        )
        // The navigation bar in portrait mode, defined below.
        const navButtonsPortrait = (
            <HamburgerMenu/>
        )
        const mediaQuery = window.matchMedia('(orientation: portrait)');
        setNavButtons(mediaQuery.matches ? navButtonsPortrait : navButtonsLandscape);

        const handleOrientationChange = (event) => {
            setNavButtons(event.matches ? navButtonsPortrait : navButtonsLandscape);
        };

        mediaQuery.addEventListener('change', handleOrientationChange);
        return () => {
            mediaQuery.removeEventListener('change', handleOrientationChange);
        };
    }, []);

    return (
        <header className="header-all">
            <div>
                <Link to="/">
                    <img
                        className="logo-primary"
                        src = {logo}
                        alt = "Unavailable."
                    />
                </Link>
            </div>
            {navButtons}
        </header>
    )
}

export function Footer() {

    function onClickTemp() {
        websocket.send("hello there from the client")
    }

    return (
        <footer className="footer default-text-color">
            <p>
                <span onClick={() => clearStorage()}>&copy;</span> (pending) <span onClick={() => logStorage()}>2023</span> Noah <span onClick={() => onClickTemp()}>Pratt</span> <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/privacy"} className='custom-link'> Privacy Policy</Link> <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/terms"} className='custom-link'> Terms of Service</Link> <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/about"} className='custom-link'> About</Link> <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/contact"} className='custom-link'> Contact Us</Link> <span className={'text-color-white'}>&#8226;</span>
                <a href="https://github.com/prattnj/music-metrics" target='_blank' rel='noreferrer' className='custom-link'> GitHub</a>
            </p>
        </footer>
    )
}

export function ConnectionModal() {
    const [circleStyle, setCircleStyle] = useState('connection-modal-disconnected');
    const [message, setMessage] = useState('Offline');

    useEffect(() => {
        websocket.onopen = () => {
            setCircleStyle('connection-modal-connected');
        }
        websocket.onclose = () => {
            setCircleStyle('connection-modal-disconnected');
            setMessage('Offline');
        }
        websocket.onmessage = (event) => {
            let numConn = findFirstInteger(event.data);
            if (numConn === 1) {
                setMessage(numConn + ' user online');
            } else {
                setMessage(numConn + ' users online');
            }
        }
    }, [])

    return (
        <div className={'connection-modal'}>
            <div className={'connection-modal-circle ' + circleStyle}> </div>
            <div>{message}</div>
        </div>
    )
}

export function PrimaryInfo(props) {

    return (
        <div className="primary-info">
            <b>{props.text}</b>
        </div>
    )
}

export function SecondaryInfo(props) {
    return (
        <div className="secondary-info">
            {props.text}
        </div>
    )
}

// LOGIN ELEMENTS ------------------------------------------------------------------------------------------------------

export function LoginButton(props) {
    return (
        <div className='login-button-wrapper'>
            <div className='login-button' onClick={() => authenticate()}>
                <b>{props.text}</b>
            </div>
        </div>
    )
}

export function getToken() {
    return localStorage.getItem('token');
}

// USEFUL METHODS ------------------------------------------------------------------------------------------------------

export function fetchInit(endpoint, requestBody, token) {
    if (HTTP_METHODS[endpoint] === 'GET') {
        return {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            }
        }
    } else {
        return {
            method: HTTP_METHODS[endpoint],
            headers: {
                'Content-Type': 'application/json',
                'Authorization': token
            },
            body: JSON.stringify(requestBody)
        }
    }
}

export function authenticate() {
    let client_id = '8b99139c99794d4b9e89b8367b0ac3f4'
    let redirect_uri = BASE_URL_WEB + '/spotify-landing'
    let state = Math.floor(Math.random() * 100000000) // random 8 digit number
    sessionStorage.setItem('state', state.toString())
    let scope = 'user-read-playback-state ' +
        'playlist-read-private ' +
        'playlist-read-collaborative ' +
        'user-follow-read ' +
        'user-read-currently-playing ' +
        'user-read-playback-position ' +
        'user-read-email ' +
        'user-top-read ' +
        'user-read-recently-played ' +
        'user-read-private ' +
        'user-library-read'

    let url = 'https://accounts.spotify.com/authorize'
    url += '?response_type=code'
    url += '&client_id=' + encodeURIComponent(client_id)
    url += '&scope=' + encodeURIComponent(scope)
    url += '&redirect_uri=' + encodeURIComponent(redirect_uri)
    url += '&state=' + encodeURIComponent(state)

    window.location = url;
}

function clearStorage() {
    localStorage.clear();
    sessionStorage.clear();
}

function logStorage() {
    console.log(localStorage);
    console.log(sessionStorage);
}

function findFirstInteger(str) {
    const match = str.match(/\d+/);
    if (match) {
        return parseInt(match[0], 10);
    }
    return null;
}
