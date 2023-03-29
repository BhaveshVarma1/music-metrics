import './home.css';
import social from './social.png';
import playlist from './playlist.png';
import stats from './stats.png';
import {PrimaryInfo} from '../util/util';
import {Link} from "react-router-dom";

export function Home() {
    return (
        <div>
            <PrimaryInfo text="Unleash the power of your playlist."/>
            <div className='landing-info'>
                <Link to={"/stats"} className={'custom-link'}><LandingTile image={stats} text="See in-depth stats about all your listening habits."/></Link>
                <Link to={"/playlist"} className={'custom-link'}><LandingTile image={playlist} text="Get custom playlists recommended based on your stats."/></Link>
                <Link to={"/social"} className={'custom-link'}><LandingTile image={social} text="Share your fun facts with all your friends any time."/></Link>
            </div>
        </div>
    )
}

function LandingTile(props) {
    return (
        <div className="landing-tile">
            <img src={props.image} alt="Image not found." className="landing-image"/>
            <div className="landing-tile-text">
              {props.text}
            </div>
        </div>
    )
}

