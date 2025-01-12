import type { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios';
import axios, { Axios } from 'axios';

import { ApiError } from '@/helpers/api_error';
import { clearToken, getToken } from '@/helpers/cookies';
import type { RequestType } from '@/helpers/types';

const API_URL = import.meta.env.VITE_API_URL;

/**
 * A singleton class responsible for making HTTP requests using Axios.
 */
class ApiService {
  /**
   * The single instance of the ApiService class.
   */
  private static instance: ApiService;

  /**
   * The Axios instance used for making HTTP requests.
   */
  private axios: Axios;

  /**
   * Private constructor to prevent direct instantiation.
   */
  private constructor() {
    /**
     * Initialize the Axios instance with default settings.
     */
    this.axios = axios;
    this.axios.defaults.timeout = 10000;
    this.axios.defaults.baseURL = API_URL;
  }

  /**
   * Returns the single instance of the ApiService class.
   *
   * @returns {ApiService} The instance of the ApiService class.
   */
  public static getInstance(): ApiService {
    if (!ApiService.instance) {
      ApiService.instance = new ApiService();
    }

    return ApiService.instance;
  }

  /**
   * Sends an HTTP request using the specified method and URL.
   *
   * This method is a wrapper around the Axios library, providing a convenient way to send HTTP requests.
   *
   * @param {RequestType} type The type of HTTP request (e.g., "get", "post", etc.).
   * @param {string} url The URL of the request.
   * @param {T} [data] The data to be sent in the request body.
   * @param {Record<string, string>} [headers] Additional headers to be sent with the request.
   * @returns {Promise<AxiosResponse<R>} A promise resolving to the response data.
   * @throws {ApiError} If the request fails.
   */
  public async send<T, R>(
    type: RequestType,
    url: string,
    data?: T,
    headers: Record<string, string> = {},
  ): Promise<AxiosResponse<R>> {
    try {
      const request: AxiosRequestConfig = {
        method: type,
        url,
        headers: {
          'Content-Type': 'application/json',
          ...headers,
        },
      };

      if (!url.includes('login')) {
        request.headers = {
          accessToken: getToken() || '',
        };
      }

      if (data) {
        request.data = JSON.stringify(data);
      } else {
        request.data = {};
      }

      return await this.axios.request(request);
    } catch (e) {
      const error = e as AxiosError;
      if (error.status === 401) {
        clearToken();
        localStorage.clear();
        window.location.href = '/login';
      }
      throw new ApiError(error);
    }
  }
}

export const api = ApiService.getInstance();