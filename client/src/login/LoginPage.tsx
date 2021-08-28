import React, { FormEvent, useState } from 'react';
import { useApi } from '../api/ApiContextProvider';

export const LoginPage = () => {
  const api = useApi();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    setIsLoading(true);
    api.authenticate(username, password)
      .catch((e) => setError(e.message))
      .finally(() => setIsLoading(false));
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
