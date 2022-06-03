import './App.css';
import axios from "axios";
import Dashboard from "./pages/dashboard/Dashboard";
import {ToastContainer} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

let apiUrl = window.__RUNTIME_CONFIG__.API_URL;
let wsUrl = window.__RUNTIME_CONFIG__.WS_URL;
let environment = window.__RUNTIME_CONFIG__.NODE_ENV;

const api = axios.create({
    baseURL: apiUrl + `/api/v1`
})

function App() {

    return (
        <div className="container">
            <main>
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
