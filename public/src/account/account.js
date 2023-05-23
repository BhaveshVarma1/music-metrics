import './account.css';
import {getToken, LoginButton, PrimaryInfo} from "../util/util";
import React, {useState} from "react";
import {useDropzone} from "react-dropzone";

const maxFileSize = 20000000
const maxFiles = 15

export function Account() {

    // LOGIN SCREEN
    if (getToken() == null || getToken() === 'undefined') {
        sessionStorage.setItem('route', 'account')
        return (
            <div>
                <PrimaryInfo text="Log in see account info..."/>
                <LoginButton text="LOGIN TO SPOTIFY"/>
            </div>
        )
    }

    return (
        <div>
            <PrimaryInfo text="Account Information"/>
            <div className={'table-acct'}>
                <div className={'table-row-acct'}>
                    <div>Username</div>
                    <div>{localStorage.getItem('username')}</div>
                </div>
                <div className={'table-row-acct'}>
                    <div>Display Name</div>
                    <div>{localStorage.getItem('display_name')}</div>
                </div>
                <div className={'table-row-acct'}>
                    <div>Email</div>
                    <div>{localStorage.getItem('email')}</div>
                </div>
                <div className={'table-row-acct'}>
                    <div>Account Created</div>
                    <div>{unixMillisToString(localStorage.getItem('timestamp'))}</div>
                </div>
            </div>
            <Dropzone/>
        </div>
    )
}

function Dropzone() {

    const [files, setFiles] = useState([])
    const [errorMessage, setErrorMessage] = useState('')
    const [hoveredIndex, setHoveredIndex] = useState(-1)

    function onDrop(acceptedFiles, rejectedFiles) {
        setHoveredIndex(-1)
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

    function submit() {
        //console.log(files)
        files.forEach(file => {
            const reader = new FileReader()
            reader.onload = (event) => {
                const fileContent = event.target.result;
                const jsonString = JSON.stringify(fileContent);
                console.log(file.name, jsonString.length);
            }
            reader.readAsText(file)
        })
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
            <div className={'dropzone-info'}>Upload your extended streaming history here. What's this?</div>
            <div {...getRootProps({
                className: 'dropzone'
            })}>
                <input {...getInputProps()} />
                {isDragActive ? (
                    <p>Drop the files here ...</p>
                ) : (
                    <p>Drag and drop .json files here, or click to select files</p>
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
            {files.length !== 0 && (
                <div className={'dropdown-submit-wrapper'}>
                    <div className={'login-button dropzone-submit'} onClick={submit}><b>SUBMIT</b></div>
                </div>
            )}
        </div>
    )
}

function unixMillisToString(unixMillis) {
    const date = new Date(+unixMillis)
    return date.toLocaleString()
}
