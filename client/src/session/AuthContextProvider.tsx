import React, {
  createContext, ReactNode, useContext, useEffect, useState,
} from 'react';
import { useApi } from '../api/ApiContextProvider';
import { isApiError } from '../api/ApiError';

export interface AuthState {
  authenticated: boolean
  setAuthenticated: (newValue: boolean) => void
}

const AuthContext = createContext<AuthState | null>(null);

export const AuthContextProvider = ({ children }: { children: ReactNode }) => {
  const api = useApi();
  const [authenticated, setAuthenticated] = useState<boolean | null>(null);
  const [error, setError] = useState('');

  useEffect(() => {
    api.authenticated()
      .then((r) => {
        if (isApiError(r)) {
          if (r.statusCode !== 401) {
            setError('Server error, please try to use this app later');
          }
          setAuthenticated(false);
        } else {
          setAuthenticated(true);
        }
      });
  }, [api, setAuthenticated]);

  if (authenticated === null) {
    return <div>Loading...</div>;
  }

  if (error !== '') {
    return <div>{error}</div>;
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
