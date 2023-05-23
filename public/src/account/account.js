import './account.css';
import {getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import React, {useCallback, useState} from "react";
import {useDropzone} from "react-dropzone";

const maxFileSize = 10000000

export function Account() {

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

    return (
        <div>
            <PrimaryInfo text="Account Information"/>
            <SecondaryInfo text={"Username: " + localStorage.getItem('username')}/>
            <SecondaryInfo text={"Display Name: " + localStorage.getItem('display_name')}/>
            <SecondaryInfo text={"Email: " + localStorage.getItem('email')}/>
            <SecondaryInfo text={"Account Created: " + unixMillisToString(localStorage.getItem('timestamp'))}/>
            {/*<Dropzone onDrop={handleDrop} multiple={true} maxSize={maxFileSize}>{() => { return (<div> <p>Drop file here</p></div> );}}</Dropzone>*/}
            <Dropzone/>
        </div>
    )
}

function Dropzone() {

    const [files, setFiles] = useState([])

    const onDrop = useCallback((acceptedFiles, rejectedFiles) => {
        if (acceptedFiles?.length) setFiles(previousFiles => [...previousFiles, ...acceptedFiles.filter(file => !previousFiles.includes(file))])
        if (rejectedFiles?.length) console.log(rejectedFiles)
    }, [])

    const { getRootProps, getInputProps, isDragActive } = useDropzone({
        onDrop,
        accept: {
            'text/json': ['.json'],
        },
        maxFiles: 10,
        maxSize: maxFileSize,
    })

    return (
        <div className={'dropzone-all'}>
            <div {...getRootProps({
                className: 'dropzone'
            })}>
                <input {...getInputProps()} />
                {isDragActive ? (
                    <p>Drop the files here ...</p>
                ) : (
                    <p>Drag and drop the files here, or click to select files</p>
                )}
            </div>

            <ul>
                {files.map(file => (
                    <li key={file.name}>
                        <div className={'dropzone-item'}>{file.name}</div>
                    </li>
                ))}
            </ul>
        </div>
    )
}

function unixMillisToString(unixMillis) {
    const date = new Date(+unixMillis)
    return date.toLocaleString()
}
