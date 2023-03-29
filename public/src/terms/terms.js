import './terms.css';
import {PrimaryInfo} from "../util/util";
import {Link} from "react-router-dom";

export function Terms() {
    return (
        <div>
            <PrimaryInfo text="Terms of Service"/>
            <div className='terms-list default-text-color'>
                <p><b>Welcome to Music Metrics. We strive to offer a service that is fun and engaging while also protecting the privacy and
                    security of our users. Please read these Terms of Service carefully, as they govern your use of our service.</b><br/></p>
                <div className='terms-list-item'>
                    <b>1. Acceptance of Terms:</b> By accessing or using our web app, you agree to these Terms of Service, our Privacy Policy,
                    and any other guidelines, rules, or additional terms that we may provide to you. If you do not agree to these terms,
                    you may not use our web app.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>2. Use of Service:</b> Our web app is designed to track your Spotify listening history and report interesting
                    statistics for it. You may use our web app only for lawful purposes and in accordance with these Terms of Service.
                    You are solely responsible for any content that you post or share through our web app.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>3. User Accounts:</b> To access our web app, you may be required to create a user account. You are responsible
                    for maintaining the confidentiality of your account information and for all activities that occur under your account.
                    You agree to notify us immediately of any unauthorized use of your account or any other breach of security.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>4. Spotify Authorization:</b> In order to track your Spotify listening history, our web app requires you
                    to authorize access to your Spotify account. By authorizing access, you acknowledge that we will have access
                    to your Spotify listening history and that we will use this information to generate statistics and reports.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>5. Intellectual Property Rights:</b> All content and materials available through our web app, including
                    but not limited to text, graphics, logos, images, and software, are owned by or licensed to us and are protected
                    by copyright, trademark, and other intellectual property laws. You may not reproduce, modify, distribute, or
                    otherwise use any of our content or materials without our express written permission.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>6. Privacy:</b> We take the privacy of our users seriously and are committed to protecting your personal information.
                    Please review our <Link className={'custom-link'} to={'/privacy'}><u>Privacy Policy</u></Link> for information
                    on how we collect, use, and disclose your personal information.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>7. Disclaimer of Warranties:</b> Our web app is provided on an "as is" and "as available" basis, without
                    warranties of any kind, either express or implied, including but not limited to warranties of merchantability,
                    fitness for a particular purpose, or non-infringement. We do not guarantee the accuracy, timeliness, or completeness
                    of any information provided through our web app.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>8. Limitation of Liability:</b> To the fullest extent permitted by applicable law, we shall not be liable
                    for any damages arising out of or in connection with your use of our web app, including but not limited to direct,
                    indirect, incidental, consequential, or punitive damages.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>9. Indemnification:</b> You agree to indemnify and hold us harmless from any claims, losses, liabilities, damages,
                    costs, or expenses (including reasonable attorneys' fees) arising out of or in connection with your use of our web
                    app or your violation of these Terms of Service.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>10. Termination:</b> We reserve the right to terminate or suspend your access to our web app at any time,
                    without notice or liability, for any reason or no reason, including but not limited to your violation
                    of these Terms of Service.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>11. Governing Law:</b> These Terms of Service shall be governed by and construed in accordance with
                    the laws of the jurisdiction in which we are located, without giving effect to any principles of conflicts of law.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>12. Changes to Terms of Service:</b> We reserve the right to modify or revise these Terms of Service at any time,
                    without notice or liability. Your continued use of our web app following the posting of any changes to these
                    Terms of Service constitutes acceptance of those changes.<br/><br/>
                </div>
                <div className='terms-list-item'>
                    <b>13. Contact Us:</b> If you have any questions about these Terms of Service or our web app, please contact
                    us <Link className={'custom-link'} to={'/contact'}><u>here</u></Link>.<br/><br/>
                </div>
            </div>
        </div>
    )
}
