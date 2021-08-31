export const loginErrorFromStatusCode = (code: number): string => {
  switch (code) {
    case 401:
      return 'Provided credentials were invalid';
    case 400:
    case 500:
      return 'Server error, please try and login later';
    default:
      return 'Login request failed, please try again later';
  }
};
