import { setCookie } from '../SetCookie';
import { ApiError, ApiErrorFromResponse } from './ApiError';

export interface ApiClient {
  authenticate(username: string, password: string): Promise<boolean | ApiError>
  authenticated(): Promise<boolean | ApiError>
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

  authenticated(): Promise<boolean | ApiError> {
    return this.makeCall(
      '/me',
      { method: 'GET' },
    ).then((r) => {
      if (r.status === 204) {
        return true;
      }
      return ApiErrorFromResponse(r);
    });
  }

  authenticate(username: string, password: string): Promise<boolean | ApiError> {
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
      return ApiErrorFromResponse(r);
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
