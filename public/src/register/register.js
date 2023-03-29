import React, {useState} from 'react';
import './register.css';
import {
    EMAIL_MAX_LENGTH, LoginError,
    LoginWithGoogle, NAME_MAX_LENGTH,
    PASSWORD_MAX_LENGTH, PASSWORD_MIN_LENGTH,
    PrimaryInfo,
    RegisterMessage,
    USERNAME_MAX_LENGTH, USERNAME_MIN_LENGTH
} from "../util/util";

export function RegisterForm () {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [password2, setPassword2] = useState('');
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [errorVisible, setErrorVisible] = useState(false);
    const [errorText, setErrorText] = useState('');

    function handleUsernameChange(event) {
        setUsername(event.target.value);
    }

    function handlePasswordChange(event) {
        setPassword(event.target.value);
    }

    function handlePassword2Change(event) {
        setPassword2(event.target.value);
    }

    function handleFirstNameChange(event) {
        setFirstName(event.target.value);
    }

    function handleLastNameChange(event) {
        setLastName(event.target.value);
    }

    function handleEmailChange(event) {
        setEmail(event.target.value);
    }

    function handleRegister() {
        const valid = sanitize();
        if (valid !== '') {
            setErrorVisible(true);
            setErrorText(valid);
        } else {
            /*fetchData('/register', {username, password, firstName, lastName, email})
                .then(data => {
                    if (data.success) {
                        setErrorVisible(false);
                        setErrorText('');
                        localStorage.token = data.token;
                        window.location.href = 'https://dev.musicmetrics.app';
                    } else {
                        setErrorVisible(true);
                        setErrorText(data.message);
                    }
                })*/
        }
    }

    function sanitize() {
        if (firstName === '' || lastName === '' || username === '' || password === '' || password2 === '' || email === '') {
            return 'Please fill out all fields.'
        } else if (firstName.length > NAME_MAX_LENGTH || lastName.length > NAME_MAX_LENGTH) {
            return 'Name too long.'
        } else if (username.length > USERNAME_MAX_LENGTH) {
            return 'Username too long.'
        } else if (username.length < USERNAME_MIN_LENGTH) {
            return 'Username too short.'
        } else if (password.length > PASSWORD_MAX_LENGTH) {
            return 'Password too long.'
        } else if (password.length < PASSWORD_MIN_LENGTH) {
            return 'Password too short.'
        } else if (password !== password2) {
            return 'Passwords do not match.'
        } else if (!validateEmail()) {
            return 'Invalid email.'
        } else {
            return ''
        }
    }

    function validateEmail() {
        if (email.length > EMAIL_MAX_LENGTH) return false
        const re = /\S+@\S+\.\S+/;
        return re.test(email);
    }

    return (
        <div>
            <PrimaryInfo text="Create a Music Metrics account"/>
            <div className='login-input-wrapper'>
                <input type="text" placeholder="First name" className="login-input" onChange={handleFirstNameChange}/>
                <input type="text" placeholder="Last name" className="login-input" onChange={handleLastNameChange}/>
                <input type="text" placeholder="Username" className="login-input" onChange={handleUsernameChange}/>
                <input type="password" placeholder="Password" className="login-input" onChange={handlePasswordChange}/>
                <input type="password" placeholder="Confirm password" className="login-input" onChange={handlePassword2Change}/>
                <input type="text" placeholder="Email" className="login-input" onChange={handleEmailChange}/>
                <LoginError isVisible={errorVisible} text={errorText}/>
                <div className='login-button bottom-margin' onClick={() => handleRegister()}>Create your free account</div>
                <LoginWithGoogle />
                <RegisterMessage />
            </div>
        </div>
    )

}
