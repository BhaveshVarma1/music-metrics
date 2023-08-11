import './social.css';
import {getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";

export function Social() {
    if (getToken() == null) {
        sessionStorage.setItem('route', 'social')
        return (
            <div>
                <PrimaryInfo text={"Log in to share stats..."}/>
                <LoginButton text={"LOGIN TO SPOTIFY"}/>
            </div>
        )
    } else {
        return (
            <div>
                <PrimaryInfo text={"🚧Share with Friends🚧"}/>
                <SecondaryInfo text={"Page coming soon..."}/>
            </div>
        )
    }

}