import React from 'react';
import './SeeDetails.css'

const SeeDetails = ({caption, target}) => {
    return (
        <>
            <button className="btn btn-6 btn-6a" onClick={()=> window.open(target, "_blank")}>{caption}</button>
        </>
    );
};

export default SeeDetails;
