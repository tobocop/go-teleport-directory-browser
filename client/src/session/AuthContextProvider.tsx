import React, {
  createContext, ReactNode, useContext, useEffect, useState,
} from 'react';
import { useApi } from '../api/ApiContextProvider';

export interface AuthState {
  authenticated: boolean
  setAuthenticated: (newValue: boolean) => void
}

const AuthContext = createContext<AuthState | null>(null);

export const AuthContextProvider = ({ children }: { children: ReactNode }) => {
  const api = useApi();
  const [authenticated, setAuthenticated] = useState<boolean | null>(null);

  useEffect(() => {
    api.authenticated()
      .then(() => setAuthenticated(true))
      .catch(() => setAuthenticated(false));
  }, [api, setAuthenticated]);

  if (authenticated === null) {
    return <div>Loading...</div>;
  }
  return (
    <AuthContext.Provider value={{ authenticated, setAuthenticated }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuthState = (): AuthState => {
  const auth = useContext<AuthState | null>(AuthContext);

  if (auth === null) {
    throw Error('Attempted to use api client when outside of the provider context');
  }

  return auth;
};
