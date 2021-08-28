import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Routes } from './Routes';
import { LoginPage } from './login/LoginPage';
import { ApiContextProvider } from './api/ApiContextProvider';

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <ApiContextProvider>
          <Route path={Routes.LOGIN} component={LoginPage} />
        </ApiContextProvider>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
