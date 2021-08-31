export interface ApiError {
  statusCode: number
  error: boolean
}

export const ApiErrorFromResponse = (response: Response): ApiError => (
  {
    statusCode: response.status,
    error: true,
  }
);

export const isApiError = (
  payload: any | ApiError,
): payload is ApiError => (payload as ApiError).error !== undefined;
