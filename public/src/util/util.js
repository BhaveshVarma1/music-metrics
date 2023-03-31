import './util.css';
import logo from './logo.png';
import {Link} from 'react-router-dom';
import React, {useEffect, useState} from "react";
import {GoogleLogin} from '@react-oauth/google';

// GLOBAL CONSTANTS ----------------------------------------------------------------------------------------------------

export const BASE_URL_API = 'https://dev.musicmetrics.app';
export const BASE_URL_WEB = 'https://dev.musicmetrics.app';
export const USERNAME_MIN_LENGTH = 6;
export const USERNAME_MAX_LENGTH = 30;
export const PASSWORD_MIN_LENGTH = 8;
export const PASSWORD_MAX_LENGTH = 30;
export const NAME_MAX_LENGTH = 60;
export const EMAIL_MAX_LENGTH = 254;
const HTTP_METHODS = {
    '/login': 'POST',
    '/register': 'POST',
    '/updateCode': 'POST',
    '/averageYear': 'GET',
    '/songCounts': 'GET',
}

// ELEMENTS COMMON TO EVERY PAGE ---------------------------------------------------------------------------------------

export function Header() {

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
                if (isOpen && !event.target.classList.contains('hamburger')) {
                    console.log("here")
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

    const [navButtons, setNavButtons] = useState(navButtonsPortrait);

    // Dynamically change the navigation bar based on the orientation of the device.
    useEffect(() => {
        const mediaQuery = window.matchMedia('(orientation: portrait)');
        setNavButtons(mediaQuery.matches ? navButtonsPortrait : navButtonsLandscape);

        const handleOrientationChange = (event) => {
            setNavButtons(event.matches ? navButtonsPortrait : navButtonsLandscape);
        };

        mediaQuery.addEventListener('change', handleOrientationChange);
        return () => {
            mediaQuery.removeEventListener('change', handleOrientationChange);
        };
    }, [])

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
    return (
        <footer className="footer default-text-color">
            <p>
                <span onClick={() => clearStorage()}>&copy;</span> (pending) <span onClick={() => logStorage()}>2023</span> Noah Pratt <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/privacy"} className='custom-link'> Privacy Policy</Link> <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/terms"} className='custom-link'> Terms of Service</Link> <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/about"} className='custom-link'> About</Link> <span className={'text-color-white'}>&#8226;</span>
                <Link to={"/contact"} className='custom-link'> Contact Us</Link> <span className={'text-color-white'}>&#8226;</span>
                <a href="https://github.com/prattnj/music-metrics" target='_blank' className='custom-link'> GitHub</a>
            </p>
        </footer>
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

export function LoginForm() {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorVisible, setErrorVisible] = useState(false);
    const [errorText, setErrorText] = useState('');

    function handleUsernameChange(event) {
        setUsername(event.target.value);
    }

    function handlePasswordChange(event) {
        setPassword(event.target.value);
    }

    function handleLogin() {
        if (username.length > USERNAME_MAX_LENGTH || username.length < USERNAME_MIN_LENGTH
            || password.length > PASSWORD_MAX_LENGTH || password.length < PASSWORD_MIN_LENGTH) {
            setErrorVisible(true);
            setErrorText('Invalid username or password.');
        } else {
            /*fetchData('/login', {username, password})
                .then(data => {
                    if (data.success) {
                        setErrorVisible(false);
                        setErrorText('');
                        localStorage.token = data.token;
                        location.reload();
                    } else {
                        setErrorVisible(true);
                        setErrorText(data.message);
                    }
                })*/
        }
    }

    return (
        <div className='login-input-wrapper'>
            <input type="text" placeholder="Username" className="login-input" onChange={handleUsernameChange}/>
            <input type="password" placeholder="Password" className="login-input" onChange={handlePasswordChange}/>
            <LoginError isVisible={errorVisible} text={errorText}/>
            <LoginButton text={'LOGIN'} click={() => handleLogin()}/>
            <p className={'default-text-color'}>Don't have an account? <Link to={'/register'} className={'custom-link'}><u>Create one</u></Link> or</p>
            <LoginWithGoogle/>
            <RegisterMessage/>
        </div>
    )
}

export function LoginButton(props) {
    return (
        <div className='login-button-wrapper'>
            <div className='login-button' onClick={() => authenticate()}>
                <b>{props.text}</b>
            </div>
        </div>
    )
}

export function LoginWithGoogle() {

    const responseMessage = (response) => {
        console.log('google success: ' + response);
    };
    const errorMessage = (error) => {
        console.log('google failure: ' + error);
    };

    return (
        <GoogleLogin onSuccess={() => responseMessage} onError={() => errorMessage} />
    )
}

export function LoginError(props) {
    return (
        <div>
            {props.isVisible && <div className={'login-error'}>{props.text}</div>}
        </div>
    )
}

export function getToken() {
    return localStorage.getItem('token');
}

export function RegisterMessage() {
    return (
        <div className={'register-message default-text-color'}>By registering, you agree to our
            <Link className={'custom-link'} to={'/privacy'}> Privacy Policy</Link> and
            <Link className={'custom-link'} to={'/terms'}> Terms of Service</Link>.
        </div>
    )
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

export class fetchData {
}