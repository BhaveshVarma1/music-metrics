import './account.css';
import {getToken, LoginButton, PrimaryInfo, SecondaryInfo} from "../util/util";
import React, {useCallback, useState} from "react";
import {useDropzone} from "react-dropzone";

//const maxFileSize = 10000000

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

    /*function UploadFiles() {

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
            reader.readAsText(file1)
        }

        return (
            <div className={'upload-files'}>
                <input type={"file"} onChange={handleFileChange} />
                <div className={'upload-button'} onClick={doUpload}>UPLOAD</div>
            </div>
        )
    }

    function handleDrop(files, rejectedFiles) {
        console.log(files)
        console.log(rejectedFiles)
    }*/

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
        if (acceptedFiles?.length) setFiles(previousFiles => [...previousFiles, ...acceptedFiles])
    }, [])

    const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop })

    return (
        <form>
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
                    <li key={file.name}>{file.name}</li>
                ))}
            </ul>
        </form>
    )
}

function unixMillisToString(unixMillis) {
    const date = new Date(+unixMillis)
    return date.toLocaleString()
}
