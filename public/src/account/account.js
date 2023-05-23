import './account.css';
import {getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import React, {useState} from "react";

export function Account() {

    if (getToken() == null || getToken() === 'undefined') {
        sessionStorage.setItem('route', 'stats')
        return (
            <div>
                <PrimaryInfo text="Log in to continue to stats..."/>
                <LoginButton text="LOGIN TO SPOTIFY"/>
            </div>
        )
    }

    function UploadFiles() {

        const [file1, setFile1] = useState(null)

        function handleFileChange(event) {
            setFile1(event.target.files[0])
        }

        function doUpload() {
            const reader = new FileReader()
            console.log(file1)
            reader.onload = function (event) {
                const data = event.target.result
                console.log("string: " + JSON.stringify(data))
            }
        }

        return (
            <div className={'upload-files'}>
                <input type={"file"} onChange={handleFileChange} />
                <div className={'upload-button'} onClick={doUpload}>UPLOAD</div>
            </div>
        )
    }

    return (
        <div>
            <PrimaryInfo text="Account Information"/>
            <SecondaryInfo text={"Username: " + localStorage.getItem('username')}/>
            <SecondaryInfo text={"Display Name: " + localStorage.getItem('display_name')}/>
            <SecondaryInfo text={"Email: " + localStorage.getItem('email')}/>
            <SecondaryInfo text={"Account Created: " + unixMillisToString(localStorage.getItem('timestamp'))}/>
            <UploadFiles/>
        </div>
    )
}

function unixMillisToString(unixMillis) {
    const date = new Date(+unixMillis)
    return date.toLocaleString()
}
