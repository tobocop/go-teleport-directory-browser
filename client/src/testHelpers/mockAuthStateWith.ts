import * as AuthContext from '../session/AuthContextProvider';
import { AuthState } from '../session/AuthContextProvider';

export const mockAuthStateWith = (state: AuthState): jest.SpyInstance => {
  const hookSpy = jest.spyOn(AuthContext, 'useAuthState');
  hookSpy.mockReturnValue(state);
  return hookSpy;
};
