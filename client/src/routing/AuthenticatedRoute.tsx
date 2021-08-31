import React from 'react';
import { Redirect, Route, RouteProps } from 'react-router-dom';
import { useAuthState } from '../session/AuthContextProvider';
import { Routes } from './Routes';

export const AuthenticatedRoute = ({ children, ...rest }: Omit<Omit<RouteProps, 'render'>, 'component'>) => {
  const { authenticated } = useAuthState();
  return (
    <Route
      /* eslint-disable-next-line react/jsx-props-no-spreading */
      {...rest}
      render={() => (authenticated ? (children) : (<Redirect to={Routes.LOGIN} />))}
    />
  );
};
