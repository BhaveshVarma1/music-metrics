import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {Home} from './home/home';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter, Route, Routes} from "react-router-dom"
import {Footer, Header} from "./util/util";
import {NotFound} from "./404/404";
import {Stats} from "./stats/stats";
import {Privacy} from "./privacy/privacy";
import {Terms} from "./terms/terms";
import {Account} from "./account/account";
import {About} from "./about/about";
import {Contact} from "./contact/contact";
import {Playlist} from "./playlist/playlist";
import {Social} from "./social/social";
import {SpotifyLanding} from "./util/spotify-landing";

const root = ReactDOM.createRoot(document.getElementById('root'));

export const websocket = new WebSocket('wss://dev.musicmetrics.app/ws');
websocket.onopen = () => {
    console.log('Websocket connected')
}
websocket.onmessage = (event) => {
    console.log('Websocket message received: ', event.data)
}

root.render(
    <BrowserRouter>
        <Header />
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/stats" element={<Stats />} />
            <Route path="/account" element={<Account />} />
            <Route path="/privacy" element={<Privacy />} />
            <Route path="/terms" element={<Terms />} />
            <Route path="/about" element={<About />} />
            <Route path="/contact" element={<Contact />} />
            <Route path="/playlist" element={<Playlist />} />
            <Route path="/social" element={<Social />} />
            <Route path="/spotify-landing" element={<SpotifyLanding />} />
            <Route path="*" element={<NotFound />} />
        </Routes>
        <Footer />
    </BrowserRouter>

);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

