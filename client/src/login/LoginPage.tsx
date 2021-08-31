import React, { FormEvent, useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { useApi } from '../api/ApiContextProvider';
import { Routes } from '../routing/Routes';
import { useAuthState } from '../session/AuthContextProvider';
import { isApiError } from '../api/ApiError';
import { loginErrorFromStatusCode } from './loginErrorFromStatusCode';

export const LoginPage = () => {
  const history = useHistory();
  const api = useApi();
  const { authenticated, setAuthenticated } = useAuthState();

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  useEffect(() => {
    if (authenticated) {
      history.push(Routes.ROOT);
    }
  }, [authenticated]);

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    setError('');
    setIsLoading(true);
    api.authenticate(username, password)
      .finally(() => setIsLoading(false))
      .then((r) => {
        if (isApiError(r)) {
          setError(loginErrorFromStatusCode(r.statusCode));
        } else {
          setAuthenticated(true);
        }
      });
  };

  const disableLogin = username === '' || password === '' || isLoading;

  return (
    <div>
      <p>
        Please provide a username and password to login.
      </p>
      {error !== '' && error}
      <form onSubmit={handleSubmit}>
        <label htmlFor="username">
          Username
          <input
            name="username"
            id="username"
            type="text"
            onChange={(event) => setUsername(event.target.value)}
          />
        </label>
        <label htmlFor="password">
          Password
          <input
            name="password"
            id="password"
            type="password"
            onChange={(event) => setPassword(event.target.value)}
          />
        </label>
        <input
          type="submit"
          disabled={disableLogin}
          value={isLoading ? 'Loading...' : 'Login'}
        />
      </form>
    </div>
  );
};
