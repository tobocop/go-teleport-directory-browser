import { Route } from 'react-router-dom';
import React from 'react';
import { Routes } from './Routes';
import { LoginPage } from '../login/LoginPage';
import { AuthenticatedRoute } from './AuthenticatedRoute';

const AuthSuccess = () => <div>Auth Successful</div>;

export const AppRoutes = () => (
  <>
    <Route exact strict path={Routes.LOGIN} component={LoginPage} />
    <AuthenticatedRoute exact path={`${Routes.ROOT}*`}>
      <AuthSuccess />
    </AuthenticatedRoute>
  </>
);
