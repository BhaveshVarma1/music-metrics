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
    const [errorMessage, setErrorMessage] = useState('')
    const [hoveredIndex, setHoveredIndex] = useState(null)

    function onDrop(acceptedFiles, rejectedFiles) {
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

    function handleHover(index) {
        setHoveredIndex(index)
    }

    function removeItem(item) {
        if (errorMessage === 'Too many files') setErrorMessage('')
        setFiles(files.filter(file => file.path !== item.path))
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
            {errorMessage !== '' && <p className={'dropzone-error'}>{errorMessage}</p>}
            <ul>
                {files.map((file, index) => (
                    <li key={file.name} style={{position: "relative"}}>
                        <div
                            className={'dropzone-item'}
                            onMouseEnter={() => handleHover(index)}
                            onMouseLeave={() => handleHover(null)}
                        >
                            {file.name}
                        </div>
                        {hoveredIndex === index && <div className={'dropzone-item-remove'} onClick={() => removeItem(file)} onMouseEnter={() => handleHover(index)}></div>}
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
