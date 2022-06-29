import React, {useEffect, useState} from 'react';
import Tiles from "../../components/tiles/Tiles";
import {PipelineStatusPromiseClient} from "../../proto/pipelines_grpc_web_pb";
import {GetStatusRequest} from "../../proto/pipelines_pb";
import {toast} from "react-toastify";


const client = new PipelineStatusPromiseClient("http://localhost:10000");

const Dashboard = ({api, wsUrl}) => {

    const [entries, setEntries] = useState([])

    const callGrpcService = () => {
        console.log("Calling grpc service");
        const request = new GetStatusRequest();
        request.setGroup('all');


        let stream = client.getStatus(request, {});

        let _entries = [];
        stream.on('data', (response) => {
            let entry = {
                name: response.getName(),
                status: response.getStatus(),
                branch: response.getBranch(),
                commit: response.getCommit(),
                author: response.getAuthor(),
                message: response.getMessage(),
                date: response.getDate(),
                url: response.getUrl(),
                platform: response.getPlatform(),
            }
            _entries.push(entry)
        });

        return _entries
    }

    const getWatchEntries = () => {
        const _entries = callGrpcService()
        setEntries(_entries)
    }

    useEffect(() => {
        setInterval(function run() {
            getWatchEntries();
            return run;
        }(), 3000);
    }, [])

    return (
            <Tiles entries={entries}/>
    );
};

export default Dashboard;
