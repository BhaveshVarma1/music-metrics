import './playlist.css';
import {getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";

export function Playlist() {
    if (getToken() == null) {
        sessionStorage.setItem('route', 'playlist')
        return (
            <div>
                <PrimaryInfo text={"Log in to use playlist builder..."}/>
                <LoginButton text={"LOGIN TO SPOTIFY"}/>
            </div>
        )
    } else {
        return (
            <div>
                <PrimaryInfo text={"🚧Playlist Builder🚧"}/>
                <SecondaryInfo text={"Page coming soon..."}/>
            </div>
        )
    }
}