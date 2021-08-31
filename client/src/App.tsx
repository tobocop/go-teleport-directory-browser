import React from 'react';
import { BrowserRouter, Switch } from 'react-router-dom';
import { ApiContextProvider } from './api/ApiContextProvider';
import { AppRoutes } from './routing/AppRoutes';
import { AuthContextProvider } from './session/AuthContextProvider';

function App() {
  return (
    <ApiContextProvider>
      <AuthContextProvider>
        <BrowserRouter>
          <Switch>
            <AppRoutes />
          </Switch>
        </BrowserRouter>
      </AuthContextProvider>
    </ApiContextProvider>
  );
}

export default App;
