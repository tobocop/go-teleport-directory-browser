import { setCookie } from '../SetCookie';

export interface ApiClient {
  authenticate(username: string, password: string): Promise<boolean>
}

export class ApiClientImpl {
  private baseRoute = '/api';

  private readonly csrf: string;

  constructor() {
    const randomNumbers = new Uint32Array(8);
    crypto.getRandomValues(randomNumbers);
    this.csrf = encodeURIComponent(btoa(randomNumbers.join('')));
    setCookie('csrf-token', this.csrf, 6);
  }

  authenticate(username: string, password: string): Promise<boolean> {
    return fetch(`${this.baseRoute}/authenticate`, {
      method: 'POST',
      cache: 'no-cache',
      headers: {
        'X-CSRF-Token': this.csrf,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username,
        password,
      }),
    }).then((r) => {
      if (r.status === 204) {
        return true;
      }
      // TODO: Should include better error handling and return status code
      throw new Error(`${r.status} Failed code`);
    });
  }
}
