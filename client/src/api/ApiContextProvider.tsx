import {ApiClient, ApiClientImpl} from "./ApiClient";
import {createContext, ReactNode, useContext} from "react";

const ApiContext = createContext<ApiClient | undefined>(undefined)

export const ApiContextProvider = ({children}: {children: ReactNode}) => {
    return <ApiContext.Provider value={new ApiClientImpl()}>
        {children}
    </ApiContext.Provider>
}

export const useApi = (): ApiClient => {
    const apiClient = useContext<ApiClient | undefined>(ApiContext)

    if(apiClient === undefined) {
        throw Error("Attempted to use api client when outside of the provider context")
    }

    return apiClient
}
