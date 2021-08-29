export interface ApiClient {
  authenticate(username: string, password: string): Promise<boolean>
}

export class ApiClientImpl {
  private baseRoute = '/api';

  authenticate(username: string, password: string): Promise<boolean> {
    return fetch(`${this.baseRoute}/authenticate`, {
      method: 'POST',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: encodeURIComponent(username),
        password: encodeURIComponent(password),
      }),
    }).then((r) => {
      if (r.status === 204) {
        return true;
      }
      throw new Error(`${r.status} Failed code`);
    });
  }
}
