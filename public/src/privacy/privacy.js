import './privacy.css';
import {PrimaryInfo} from "../util/util";
import {Link} from "react-router-dom";

export function Privacy() {
    return (
        <div>
            <PrimaryInfo text="Privacy Policy"/>
            <div className={'terms-list default-text-color'}>
                <p><b>We respect your privacy and are committed to protecting your personal information. This Privacy Policy
                    explains how we collect, use, and disclose information when you use our web app that tracks your Spotify
                    listening history and reports interesting statistics for it. By using our web app, you consent to our
                    collection, use, and disclosure of your information in accordance with this Privacy Policy.</b><br/></p>
                <div className='terms-list-item'>
                    <b>1. Information we collect:</b> We collect information from you when you authorize access to your
                    Spotify account. This information includes your Spotify listening history, playlists, and other
                    information related to your Spotify account.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>2. How we use your information:</b> We use your information to generate statistics and reports about
                    your Spotify listening history. We may also use your information to improve our web app and to provide
                    you with personalized content and recommendations.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>3. How we share your information:</b> We do not sell, rent, or lease your personal information to
                    third parties. However, we may share your information with third-party service providers who help us
                    operate our web app and provide you with our services. We require these service providers to protect
                    your personal information and to use it only for the purposes for which we disclose it to them.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>4. Security of your information:</b> We take reasonable measures to protect your personal information
                    from unauthorized access, use, or disclosure. However, no data transmission over the Internet or electronic
                    storage can be guaranteed to be 100% secure.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>5. Retention of your information:</b> We retain your personal information for as long as necessary
                    to provide you with our services and to comply with our legal obligations. We may also retain your information
                    for research or statistical purposes, but we will anonymize or delete your information when it is no
                    longer necessary for these purposes.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>6. Your rights and choices:</b> You have the right to access and correct your personal information
                    that we hold. You may also request that we delete your personal information or restrict our processing
                    of it. However, we may not be able to provide you with our services if you request that we delete or
                    restrict the processing of your personal information.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>7. Children's privacy:</b> Our web app is not intended for children under the age of 13. We do not
                    knowingly collect personal information from children under the age of 13. If you are under the age of
                    13, please do not use our web app or provide us with any personal information.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>8. Changes to our Privacy Policy:</b> We reserve the right to modify or revise this Privacy Policy
                    at any time, without notice or liability. Your continued use of our web app following the posting of
                    any changes to this Privacy Policy constitutes acceptance of those changes.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>9. Contact Us:</b> If you have any questions about this Privacy Policy or our web app, please contact
                    us <Link className={'custom-link'} to={'/contact'}><u>here</u></Link>.<br/><br/>
                </div>
            </div>
        </div>
    )
}
