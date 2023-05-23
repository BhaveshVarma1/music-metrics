import './account.css';
import {getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import React, {useState} from "react";
import {useDropzone} from "react-dropzone";

const maxFileSize = 20000000
const maxFiles = 5

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
            <Dropzone/>
        </div>
    )
}

function Dropzone() {

    const [files, setFiles] = useState([])

    function onDrop(acceptedFiles, rejectedFiles) {
        if (acceptedFiles?.length) {
            //if (files.length + acceptedFiles.length > maxFiles) return
            setFiles(previousFiles => [...previousFiles, ...acceptedFiles.filter(file => !previousFiles.some(previousFile => previousFile.name === file.name))])
        }
        if (rejectedFiles?.length) console.log(rejectedFiles)
    }

    const { getRootProps, getInputProps, isDragActive } = useDropzone({
        onDrop,
        accept: {
            'text/json': ['.json'],
        },
        maxFiles: maxFiles,
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
