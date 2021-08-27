import {useEffect, useState} from "react";

export const TestHttps = () => {
    const [ret, setRet] = useState("")

    useEffect( () => {
        fetch("/api/hello")
            .then((res) => res.text())
            .then(data => setRet(data))
    })

    return <div>ret: {ret}</div>
}