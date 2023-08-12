import './contact.css'
import {PrimaryInfo, SecondaryInfo} from "../util/util";

export function Contact() {
    return (
        <div>
            <PrimaryInfo text={"Contact Us"}/>
            <SecondaryInfo text={"Email: musicmetricsapp@gmail.com"}/>
            <SecondaryInfo text={<div>Instagram: <a className={'custom-link'} target={"_blank"} href={"https://instagram.com/_musicmetrics"}>@_musicmetrics</a></div>}/>
        </div>
    )
}
