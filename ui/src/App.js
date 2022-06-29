import './App.css';
import axios from "axios";
import Dashboard from "./pages/dashboard-v2/Dashboard";
import {ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import {useEffect} from "react";

let apiUrl = window.__RUNTIME_CONFIG__.API_URL;
let wsUrl = window.__RUNTIME_CONFIG__.WS_URL;
let environment = window.__RUNTIME_CONFIG__.NODE_ENV;

const api = axios.create({
    baseURL: apiUrl + `/api/v1`
})

function App() {
    function refreshPage() {
        window.location.reload();
    }
    function goToGithubRepo() {
        window.open("https://github.com/zcubbs/pulse", '_blank').focus();
    }

    useEffect(() => {
        console.log("Environment: " + environment);
    },[environment]);

    return (
        <div className="container">
            <main>
                <div className="top-nav">
                    <nav className="nav-items">
                        <div className="title" onClick={refreshPage}><span><img className="logo" src="/pulse_logo.png" alt=""/></span>v1
                        </div>
                        <div className="repo-link" onClick={goToGithubRepo}>
                            <span><img className="logo" width="24px" height="24px" src="/github.png"
                                       alt=""/></span>
                        </div>
                    </nav>
                </div>
                <div className="main__container">
                    {/*<p>ENVIRONMENT = {environment}</p>*/}
                    {/*<p>API_URL = {api}</p>*/}
                    <ToastContainer/>
                    <Dashboard api={api} wsUrl={wsUrl}/>
                </div>
            </main>
        </div>
    );
}

export default App;
