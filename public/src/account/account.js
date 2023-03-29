import './account.css';
import {getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";

export function Account() {
    if (getToken() == null) {
        sessionStorage.setItem('route', 'account')
        return (
            <div>
                <PrimaryInfo text="Log in to see account info..."/>
                <LoginButton text="LOGIN TO SPOTIFY"/>
            </div>
        )
    } else {
        return (
            <div>
                <PrimaryInfo text="Account Information"/>
                <SecondaryInfo text={"Username: " + localStorage.getItem('username')}/>
                <SecondaryInfo text={"Display Name: " + localStorage.getItem('display_name')}/>
                <SecondaryInfo text={"Email: " + localStorage.getItem('email')}/>
                <SecondaryInfo text={"Account Created: " + unixMillisToString(localStorage.getItem('timestamp'))}/>
            </div>
        )
    }

}

function unixMillisToString(unixMillis) {
    const date = new Date(+unixMillis)
    return date.toLocaleString()
}
