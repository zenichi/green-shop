import axios from 'axios';
import { getHeaders } from './api-helpers';
import config from './api.config';

export const getProducts = currency => {
  return axios({
    method: 'get',
    url: config.apiHost() + `products/${currency}`,
    headers: getHeaders()
  });
};

export const getProductById = (id, currency) => {
  return axios({
    method: 'get',
    url: config.apiHost() + `products/${id}/${currency}`,
    headers: getHeaders()
  });
};
