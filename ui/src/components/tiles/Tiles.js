import React from 'react';
import './Tiles.css'
import SeeDetails from "../seeDetails/SeeDetails";
import {FcLike, FcFlashOff, FcFlashOn, FcLikePlaceholder, FcDislike} from "react-icons/fc";
import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en.json'
import ReactTimeAgo from "react-time-ago";

TimeAgo.addDefaultLocale(en)

const Tiles = ({entries}) => {

    const getDate = (date) => {
        return new Date(date)
    }
    return (
        <div className="tiles-container">
            {
                entries === null ? <div>No data found</div> : entries.map((entry, key) => {
                    return (
                        <div key={key} className={(() => {
                            if (entry.status == 'success') {
                                return ("tile colored success")
                            } else if (entry.status == 'skipped') {
                                return ("tile colored skipped")
                            } else if (entry.status == 'failed') {
                                return ("tile colored failure")
                            } else if (entry.status == 'running') {
                                return ("tile colored running")
                            } else {
                                return ("tile colored")
                            }
                        })()}>
                            {(() => {
                                if (entry.status == 'success') {
                                    return (
                                        <>
                                            <div className="tile-top">
                                                <span><FcLike size="25px" color="#16a085"/></span>
                                                <span className="tag">{entry.origin}</span>
                                            </div>
                                            <div className="tile-inner">
                                                <p className=""><b className="tag success">Success</b>
                                                    <span className="tag">
                                                        <ReactTimeAgo date={getDate(entry.event_date)}
                                                                  locale="en-US"/>
                                                    </span>
                                                </p>
                                                <span
                                                    className="font-bold font-weight-1-7em">{entry.project_name}</span>
                                            </div>
                                        </>
                                    )
                                } else if (entry.status == 'skipped') {
                                    return (
                                        <>
                                            <div className="tile-top">
                                                <span><FcDislike size="25px" color="#16a085"/></span>
                                                <span className="tag">{entry.origin}</span>
                                            </div>
                                            <div className="tile-inner">
                                                <p className=""><b className="tag skipped">Skipped</b>
                                                    <span className="tag">
                                                        <ReactTimeAgo date={getDate(entry.event_date)}
                                                                  locale="en-US"/>
                                                    </span>
                                                </p>
                                                <span
                                                    className="font-bold font-weight-1-7em">{entry.project_name}</span>
                                            </div>
                                        </>
                                    )
                                } else if (entry.status == 'failed') {
                                    return (
                                        <>
                                            <div className="tile-top">
                                                <span><FcFlashOff size="25px" color="#16a085"/></span>
                                                <span className="tag">{entry.origin}</span>
                                            </div>
                                            <div className="tile-inner">
                                                <p className=""><b className="tag red">Failed</b>
                                                    <span className="tag">
                                                        <ReactTimeAgo date={getDate(entry.event_date)}
                                                                  locale="en-US"/>
                                                    </span>
                                                </p>
                                                <span
                                                    className="font-bold font-weight-1-7em">{entry.project_name}</span>
                                            </div>
                                        </>
                                    )
                                } else if (entry.status == 'running') {
                                    return (
                                        <>
                                            <div className="tile-top">
                                                <span><FcFlashOn size="25px" color="#16a085"/></span>
                                                <span className="tag">{entry.origin}</span>
                                            </div>
                                            <div className="tile-inner">
                                                <p className=""><b className="tag running">Running</b>
                                                    <span className="tag">
                                                        <ReactTimeAgo date={getDate(entry.event_date)}
                                                                      locale="en-US"/>
                                                    </span>
                                                </p>
                                                <span
                                                    className="font-bold font-weight-1-7em">{entry.project_name}</span>
                                            </div>
                                        </>
                                    )
                                } else {
                                    return (
                                        <>
                                            <i className="fa fa-question fa-2x"
                                               aria-hidden="true"></i>
                                            <div className="tile-inner">
                                                <p className="">Unknown</p>
                                                <span
                                                    className="font-bold font-weight-1-7em">{entry.project_name}</span>
                                            </div>
                                        </>
                                    )
                                }
                            })()}
                            <div className="tile-footer">
                                {/*<p>{entry.event_date}</p>*/}
                                <SeeDetails caption={"Details"} target={entry.origin_url}/>
                                {/*<a href={"https://gitlab.com"} target="_blank" rel="noreferrer"*/}
                                {/*   style={{marginLeft: '10px'}}>History</a>*/}
                            </div>
                        </div>
                    )
                })
            }
        </div>
    );
};

export default Tiles;
