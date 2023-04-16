import {BASE_URL_API, BASE_URL_WEB, fetchInit, PrimaryInfo, SecondaryInfo} from "./util";

export function SpotifyLanding() {
    onload()
    return (
        <div>
            <PrimaryInfo text={"Redirecting..."}/>
            <SecondaryInfo text={"This will take slightly longer if this is your first time logging in."}/>
        </div>
    )
}

function onload() {
    const landing_url = window.location.href
    if (!validateState(landing_url)) {
        // ERROR: INVALID STATE (CSRF ATTACK)
        window.location = BASE_URL_WEB
    }
    storeAuthInfo(landing_url)
    if (sessionStorage.getItem('error') != null) {
        // FAILURE (user did not log in)
        localStorage.removeItem('token')
        window.location = BASE_URL_WEB
    } else {
        // SUCCESS (user logged in)
        let code = sessionStorage.getItem('code')
        if (code == null) {
            console.log("ERROR: code is null")
            window.location = BASE_URL_WEB
        }
        fetch(BASE_URL_API + '/api/v1/updateCode', fetchInit('/api/v1/updateCode', {code: code}, null))
            .then(response => response.json())
            .then(data => {
                localStorage.setItem('token', data.token)
                localStorage.setItem('username', data.username)
                localStorage.setItem('display_name', data.displayName)
                localStorage.setItem('email', data.email)
                localStorage.setItem('timestamp', data.timestamp)
                window.location = BASE_URL_WEB + '/' + sessionStorage.getItem('route')
            }).catch(error => {
                // Internal Server Error
                console.log("ERROR: " + error)
                localStorage.removeItem('token')
                // todo remove this line?
                localStorage.clear()
                sessionStorage.setItem('error', 'Internal Server Error')
                window.location = BASE_URL_WEB
            })
    }
}

function storeAuthInfo(url) {
    url = url.replace('#', '?')

    const urlParams = new URLSearchParams(new URL(url).search)

    if (!urlParams.has('state')) return;
    if (!urlParams.has('error')) {
        const code = urlParams.get('code')
        sessionStorage.setItem('code', code)
    } else {
        const error = urlParams.get('error')
        sessionStorage.setItem('error', error)
    }
}

function validateState(url) {
    const urlParams = new URLSearchParams(url)
    const stored_state = sessionStorage.getItem('state')
    if (!urlParams.has('state') || stored_state == null) return false;
    return stored_state === urlParams.get('state')
}