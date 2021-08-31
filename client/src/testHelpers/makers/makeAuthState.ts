import { AuthState } from '../../session/AuthContextProvider';

export const makeAuthState = (overrides: Partial<AuthState> = {}): AuthState => ({
  authenticated: false,
  setAuthenticated: () => {},
  ...overrides,
});
