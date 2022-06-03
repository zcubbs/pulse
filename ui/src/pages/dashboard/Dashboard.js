import React, {useCallback, useEffect, useState} from 'react';
// import './Overwatch.css';
import {toast} from "react-toastify";
import Tiles from "../../components/tiles/Tiles";
import ReconnectingWebSocket from "reconnecting-websocket";

const Dashboard = ({api, wsUrl}) => {

    const [entries, setEntries] = useState([])
    const [websocket, setWebsocket] = useState()

    useEffect(() => {
        setWebsocket(new ReconnectingWebSocket(wsUrl + "/ws", null, {debug: true, reconectInterval: 3000}));
    }, [])


    useEffect(() => {
        if (websocket) {
            websocket.onopen = () => {
                console.log('⚡ z/Rocket websocket connected')
            };

            websocket.onerror = (e) => {
                console.log(e)
            };

            websocket.onclose = (e) => {
                console.log(e)
            };

            websocket.onmessage = (e) => {
                wsReceive(e)
            };
        }
    }, [websocket])

    const wsSend = useCallback((type, msg) => {
        websocket.send(JSON.stringify({
            message_type: type,
            message: msg
        }))
    }, [])

    const wsReceive = useCallback((e) => {
        console.log('// ** ws:received ** //')
        if (e.data) {
            let message = JSON.parse(e.data)

            if (message.message_type === 'watch_refresh') {
                toast.dark(message.message, {
                    position: "bottom-right",
                    autoClose: 3000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                });
                getWatchEntries()
            }
        }
    }, [])

    useEffect(async () => {
            await getWatchEntries()
        }, []
    )

    function getWatchEntries() {
        return api.get('/watch/entries')
            .then(res => {
                setEntries(res.data)
                console.log('watch entries', res.data)
            }).catch(function (error) {
                // handle error
                console.log(error)
                toast.error('Unable to connect to backend api', {
                    position: "bottom-right",
                    autoClose: 5000,
                    hideProgressBar: false,
                    closeOnClick: true,
                    pauseOnHover: true,
                    draggable: true,
                    progress: undefined,
                });
            });
    }

    return (
        <Tiles entries={entries}/>
    );
};

export default Dashboard;
