import { setCookie } from '../SetCookie';

export interface ApiClient {
  authenticate(username: string, password: string): Promise<boolean>
  authenticated(): Promise<boolean>
}

export class ApiClientImpl implements ApiClient {
  private baseRoute = '/api';

  private readonly csrf: string;

  constructor() {
    const randomNumbers = new Uint32Array(8);
    crypto.getRandomValues(randomNumbers);
    this.csrf = encodeURIComponent(btoa(randomNumbers.join('')));
    setCookie('csrf-token', this.csrf, 6);
  }

  authenticated(): Promise<boolean> {
    return this.makeCall(
      '/me',
      { method: 'GET' },
    ).then((r) => {
      if (r.status === 204) {
        return true;
      }
      throw new Error(`${r.status} Failed code`);
    });
  }

  authenticate(username: string, password: string): Promise<boolean> {
    return this.makeCall(
      '/authenticate',
      { method: 'POST' },
      {
        username,
        password,
      },
    ).then((r) => {
      if (r.status === 204) {
        return true;
      }
      // TODO: Should include better error handling and return status code
      throw new Error(`${r.status} Failed code`);
    });
  }

  private makeCall(
    url: string,
    fetchParams: Partial<RequestInit>,
    body: any | null = null,
  ): Promise<Response> {
    const init = {
      cache: 'no-cache' as RequestCache,
      headers: {
        'X-CSRF-Token': this.csrf,
        'Content-Type': 'application/json',
      },
      ...fetchParams,
    };

    if (body !== null) {
      init.body = JSON.stringify(body);
    }
    return fetch(`${this.baseRoute}${url}`, init);
  }
}
